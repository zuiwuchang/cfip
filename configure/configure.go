package configure

import (
	"encoding/json"
	"log"

	"github.com/google/go-jsonnet"
)

type Configure struct {
	// 間隔多久執行一次任務，如果不設置則只執行一次
	Interval string `json:"interval"`
	// 並行工作數量，建議不要設置太高，太高的並行存在諸多問題
	// 1. 對測試的網站造成太大壓力，並且可能會被加入黑名單禁止訪問
	// 2. 容易引發防火牆的 sni 阻斷，一段時間內都禁止你和 cloudflare 的任何連接
	// 3. 對運行測試的計算機造成過大的壓力
	// 默認 5 就很不錯，無論對運行 測試的計算機 cloudflare 測試網站造成的影響都很微弱，也不容易引發 sni 阻斷
	Worker int `json:"worker"`

	// 要測試的 ip 段
	IP []string `json:"ip"`
	// 請求設定
	Request Request `json:"request"`
	// 查找目標設定
	Found Found `json:"found"`
}

func Load(filepath string) (conf *Configure, e error) {
	vm := jsonnet.MakeVM()
	jsonStr, e := vm.EvaluateFile(filepath)
	if e != nil {
		log.Fatalln(e)
		return
	}
	var c Configure
	e = json.Unmarshal([]byte(jsonStr), &c)
	if e != nil {
		log.Fatalln(e)
	}

	conf = &c
	return
}
