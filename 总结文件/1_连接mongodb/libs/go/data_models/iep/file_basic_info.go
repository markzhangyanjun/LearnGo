package iep

type FileBasicInfo struct {
	Malname  string `json:"malname"`
	Filename string `json:"filename"`
	Status   int    `json:"status"`
	MD5      string `json:"md5"`
	CT       string `json:"ct"`
	MT       string `json:"mt"`
	Ver      string `json:"ver"`
	DigSig   string `json:"digsig"`
	Size     uint64 `json:"size"`
	Path     string `json:"path"`
}
