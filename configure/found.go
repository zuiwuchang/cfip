package configure

type Found struct {
	// 每次任務的目標是要尋找多少個 ip
	IP int `json:"ip"`
	// 至少要尋找多少個有效 ip，從中選擇最優(平均延遲最低) ip
	Valid int `json:"valid"`
	// 至少要測試多少個 ip
	Test int `json:"test"`
	// 將結果以 post json 發送到此 url
	URL string `json:"url"`
	// 最多要測試多少 ip
	Max int `json:"max"`
}
