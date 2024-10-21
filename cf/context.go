package cf

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/zuiwuchang/cfip/configure"
)

type Context struct {
	r        IPRange
	interval time.Duration
	worker   int

	ip    int
	valid int
	test  int
}

func New(conf *configure.Configure) (c *Context, e error) {
	var interval time.Duration
	if conf.Interval != `` {
		interval, e = time.ParseDuration(conf.Interval)
	}

	worker := conf.Worker
	if worker < 1 {
		worker = 5
	}
	var r IPRange
	for _, s := range conf.IP {
		e = r.AddCIDR(s)
		if e != nil {
			return
		}
	}
	ip := conf.Found.IP
	if ip < 1 {
		ip = 5
	}
	valid := conf.Found.Valid
	if valid < ip {
		valid = ip * 5
	}
	test := conf.Found.Test
	if test < valid {
		test = valid * 20
	}
	c = &Context{
		interval: interval,
		r:        r,
		worker:   worker,
		ip:       ip,
		valid:    valid,
		test:     test,
	}
	return
}
func (c *Context) Serve() (e error) {
	if c.interval > 0 {
		for {
			c.serve()

			log.Println(`next after`, c.interval)
			time.Sleep(c.interval)
		}
	} else {
		e = c.serve()
	}

	return
}
func (c *Context) serve() (e error) {
	var wait sync.WaitGroup
	wait.Add(c.worker + 1)

	found := newFound(c.r, c.ip, c.valid, c.test)
	for i := 0; i < c.worker; i++ {
		go func() {
			c.do(found)
			defer wait.Done()
		}()
	}
	go found.serve()

	wait.Wait()
	return
}
func (c *Context) do(found *Found) {
	for {
		ctx, ip, e := found.Get()
		if e != nil {
			break
		}

		fmt.Println(ctx, ip)
	}
}
