package cf

import (
	"container/heap"
	"context"
	"fmt"
	"log"
	"net"
	"sort"
	"sync"
	"time"
)

type Valid struct {
	ip   string
	used []time.Duration
	avg  time.Duration
	last time.Time
}
type Found struct {
	r     IPRange
	ip    int
	valid int
	test  int

	sync.Mutex

	ctx    context.Context
	cancel context.CancelFunc

	keys map[string]bool
	ch   chan *Valid

	tests  int
	valids maxValid

	next chan *Valid
}
type maxValid []*Valid

func (a maxValid) Len() int      { return len(a) }
func (a maxValid) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a maxValid) Less(i, j int) bool { return a[i].avg > a[j].avg }
func (h *maxValid) Push(x interface{}) {
	*h = append(*h, x.(*Valid))
}
func (h *maxValid) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
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
		ch:     make(chan *Valid),
		next:   make(chan *Valid, 100),
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
	go func() {
		var (
			ip   *Valid
			wait time.Duration
		)
		for {
			select {
			case ip = <-f.next:
			case <-f.ctx.Done():
				return
			}
			wait = time.Since(ip.last)
			if wait > time.Second*5 {
				time.Sleep(wait - time.Second*5)
			}

			select {
			case f.ch <- ip:
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
		if f.checkEnd() {
			f.Unlock()
			return
		}

		// range new ip
		if strs != nil {
			strs = strs[:0]
		}
		for _, ip = range ips {
			s = ip.String()
			if f.keys[s] {
				continue
			}

			strs = append(strs, s)
			f.keys[s] = true
		}
		for _, s = range strs {
			delete(f.keys, s)
		}
		f.Unlock()

		for _, s = range strs {
			select {
			case f.ch <- &Valid{
				ip: s,
			}:
			case <-f.ctx.Done():
				return
			}
		}
	}
}
func (f *Found) checkEnd() bool {
	if len(f.valids) >= f.valid && f.tests >= f.test {
		f.cancel()
		return true
	}
	return false
}
func (f *Found) Get() (ctx context.Context, ip *Valid, e error) {
	select {
	case <-f.ctx.Done():
		e = f.ctx.Err()
		return
	case ip = <-f.ch:
		ctx = f.ctx

		f.Lock()
		defer f.Unlock()
		e = f.ctx.Err()
		if e != nil {
			return
		}
		f.tests++
		log.Printf("tests[%d]=%v valids=%d\n", f.tests, ip, len(f.valids))

	}
	return
}
func (f *Found) SetOk(ip *Valid) {
	f.Lock()
	defer f.Unlock()

	ip.avg = 0
	for _, v := range ip.used {
		ip.avg += v
	}

	heap.Push(&f.valids, ip)
	if len(f.valids) > f.valid {
		heap.Pop(&f.valids)
	}
	f.keys[ip.ip] = true

	if !f.checkEnd() {
		return
	}

	sort.Sort(f.valids)
	f.valids = f.valids[len(f.valids)-f.ip:]
	for _, ip := range f.valids {
		fmt.Println(ip)
	}
}
func (f *Found) Next(ip *Valid) {
	select {
	case f.next <- ip:
		return
	case <-f.ctx.Done():
		return
	default:
	}

	go func() {
		select {
		case f.next <- ip:
		case <-f.ctx.Done():
		}
	}()
}
