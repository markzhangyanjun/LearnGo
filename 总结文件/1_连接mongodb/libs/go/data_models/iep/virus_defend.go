package iep

// same file info, but use different type comparing to virus_detect.
// i.e. virus_defend ct type is int.. virus_defend ct type is string.
type FileInfo struct {
	Malname     string   `json:"malname"`
	Filename    string   `json:"filename"`
	Status      int      `json:"status"`
	MD5         string   `json:"md5"`
	CT          int64    `json:"ct"`
	MT          int64    `json:"mt"`
	Ver         string   `json:"ver"`
	DigSig      []string `json:"digsig"`
	Size        uint64   `json:"size"`
	Path        string   `json:"path"`
	Instruction string   `json:"instruction"`
	Option      string   `json:"option"`
}

type VirusDefend struct {
	Type     string     `json:"type"`
	Ts       int64      `json:"ts"`
	Client   ClientInfo `json:"client"`
	FileInfo []FileInfo `json:"fileinfo"`
}
