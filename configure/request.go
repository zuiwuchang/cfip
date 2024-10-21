package configure

type Request struct {
	// 請求超時時間
	Timeout string `json:"timeout"`

	// 對同一 ip 要進行測試多少次
	Count int `json:"count"`

	// 請求 url
	URL string `json:"url"`

	// 只要請求返回的 http status 與此相等時才認爲請求成功，
	// 如果 code <1 則 將 [200,299] 的 htttp status 都認爲成功
	Code int `json:"code"`

	// 防止 tls 阻斷，如果設置則每次以不同的 sni/host 請求 url
	SNI []string `json:"sni"`

	// 如果設置每次都以不同的 UserAgent 發送請求
	UserAgent []string `json:"userAgent"`
}
