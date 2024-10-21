package cf

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
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

	request *Request
}

func New(conf *configure.Configure) (c *Context, e error) {
	var interval time.Duration
	if conf.Interval != `` {
		interval, e = time.ParseDuration(conf.Interval)
		if e != nil {
			return
		}
	}
	req, e := newRequest(&conf.Request)
	if e != nil {
		return
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
		request:  req,
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
	var (
		ctx  context.Context
		ip   *Valid
		e    error
		last time.Time
		used time.Duration
		code int
	)
	for {
		ctx, ip, e = found.Get()
		if e != nil {
			break
		}

		last = time.Now()
		code, e = c.doReq(ctx, found, ip)
		if e != nil {
			continue
		} else if code != 0 {
			continue
		}

		used = time.Since(last)
		log.Println(`-------------used`, used, ip, code)
		ip.used = append(ip.used, used)
		os.Exit(1)
	}
}

var netdialder net.Dialer

func (c *Context) doReq(ctx context.Context, found *Found, ip *Valid) (status int, e error) {
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				i := strings.LastIndex(addr, ":")
				if i < 0 {
					return nil, errors.New(`unknow addr:` + addr)
				}
				if strings.Index(ip.ip, ":") > 0 {
					addr = "[" + ip.ip + "]" + addr[i:]
				} else {
					addr = ip.ip + addr[i:]
				}

				return netdialder.DialContext(ctx, network, addr)
			},
		},
	}
	defer client.CloseIdleConnections()

	req, cancel, e := c.request.Request(ctx)
	if e != nil {
		log.Println(e)
		return
	}
	defer cancel()
	userAgent := c.request.UserAgent()
	if userAgent != `` {
		req.Header.Set(`user-agent`, userAgent)
	}

	resp, e := client.Do(req)
	if e != nil {
		if !errors.Is(e, context.DeadlineExceeded) &&
			!errors.Is(e, context.Canceled) {
			log.Println(e, req.URL)
		}
		return
	}
	code := resp.StatusCode
	if (c.request.code == 0 && code < 300 && code >= 200) || c.request.code == code {
		_, e = io.ReadAll(resp.Body)
		if e != nil {
			log.Println(e, req.URL)
			return
		}
		e = resp.Body.Close()
		if e != nil {
			log.Println(e, req.URL)
			return
		}
		status = code
	} else {
		log.Println(`code not match`, code, req.URL)
		io.ReadAll(resp.Body)
		resp.Body.Close()
	}
	return
}
