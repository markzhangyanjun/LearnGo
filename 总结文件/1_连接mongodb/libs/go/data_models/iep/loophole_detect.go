package iep

type CVEStatusInfo struct {
	ID     string `json:"id"`
	Status int    `json:"status"`
}
type PatchInfo struct {
	Status      int           `json:"status"`
	Name        string        `json:"name"`
	Desc        string        `json:"desc"`
	Level       int           `json:"level"`
	Kbid        string        `json:"kbid"`
	DownloadUrl string        `json:"download_url"`
	Type        int           `json:"type"`
	CVE         []string      `json:"cve"`
	CVEStatus   []interface{} `json:"cve_status"`
}
type LoopholeDetect struct {
	Type   string      `json:"type"`
	Ts     int64       `json:"ts"`
	Client ClientInfo  `json:"client"`
	Patch  []PatchInfo `json:"patch"`
}
