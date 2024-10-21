package cf

import (
	"context"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/zuiwuchang/cfip/configure"
)

type Request struct {
	// 請求超時時間
	timeout time.Duration

	// 對同一 ip 要進行測試多少次
	count int

	// 請求 url
	url *url.URL

	// 只要請求返回的 http status 與此相等時才認爲請求成功，
	// 如果 code <1 則 將 [200,299] 的 htttp status 都認爲成功
	// code:200,
	// 防止 tls 阻斷
	sni []string
	// 如果設置每次都以不同的 UserAgent 發送請求
	userAgent []string

	code int
	i    int
	sync.Mutex
}

func newRequest(conf *configure.Request) (r *Request, e error) {
	var timeout time.Duration
	if conf.Timeout == `` {
		timeout = time.Second
	} else {
		timeout, e = time.ParseDuration(conf.Timeout)
		if e != nil {
			return
		}
	}
	count := conf.Count
	if count < 1 {
		count = 3
	}
	u, e := url.ParseRequestURI(conf.URL)
	if e != nil {
		return
	}

	r = &Request{
		timeout:   timeout,
		count:     count,
		url:       u,
		sni:       conf.SNI,
		userAgent: conf.UserAgent,
		code:      conf.Code,
	}
	return
}
func (r *Request) Request(ctx context.Context) (req *http.Request, cancel context.CancelFunc, e error) {
	u := *r.url

	n := len(r.sni)
	if n > 0 {
		host := u.Host
		i := strings.LastIndex(host, ":")
		if i >= 0 {
			host = host[i:]
		} else {
			host = ``
		}
		r.Lock()
		u.Host = r.sni[r.i] + host
		r.i++
		if r.i == n {
			r.i = 0
		}
		r.Unlock()
	}
	ctx, cancel = context.WithTimeout(ctx, r.timeout)
	req, e = http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	return
}
func (r *Request) UserAgent() string {
	n := len(r.userAgent)
	if n > 0 {
		return r.userAgent[rand.Intn(n)]
	}
	return ``
}
