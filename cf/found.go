package cf

import (
	"context"
	"net"
	"sync"
)

type Found struct {
	r     IPRange
	ip    int
	valid int
	test  int

	sync.Mutex

	ctx    context.Context
	cancel context.CancelFunc

	keys map[string]bool
	ch   chan string

	usedTest int
}

func newFound(r IPRange, ip, valid, test int) *Found {
	ctx, cancel := context.WithCancel(context.Background())
	return &Found{
		ctx:    ctx,
		cancel: cancel,
		r:      r,
		ip:     ip,
		valid:  valid,
		test:   test,
		keys:   make(map[string]bool),
		ch:     make(chan string),
	}
}
func (f *Found) serve() {
	ch := make(chan []net.IP)
	go func() {
		for {
			ips := make([]net.IP, 0, 50)
			for i := 0; i < 50; i++ {
				ip := f.r.Random()
				ips = append(ips, ip)
			}
			select {
			case ch <- ips:
			case <-f.ctx.Done():
				return
			}
		}
	}()

	var (
		ips  []net.IP
		ip   net.IP
		strs []string
		s    string
	)
	for {
		select {
		case ips = <-ch:
		case <-f.ctx.Done():
			return
		}
		f.Lock()
		// check end

		// range new ip
		for _, ip = range ips {
			s = ip.String()
			if f.keys[s] {
				continue
			}

			strs = append(strs, s)
			f.keys[s] = true
			f.usedTest++
		}
		for _, s = range strs {
			delete(f.keys, s)
		}
		f.Unlock()

		for _, s = range strs {
			select {
			case f.ch <- s:
			case <-f.ctx.Done():
				return
			}
		}
	}
}
func (f *Found) Get() (ctx context.Context, ip string, e error) {
	select {
	case <-f.ctx.Done():
		e = f.ctx.Err()
		return
	case ip = <-f.ch:
		ctx = f.ctx
	}
	return
}
