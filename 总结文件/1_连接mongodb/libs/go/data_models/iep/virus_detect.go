package iep


type VirusDetect struct {
	Type     string     `json:"type"`
	Ts       int64     `json:"ts"`
	Client   ClientInfo `json:"client"`
	FileInfo []FileBasicInfo `json:"fileinfo"`
}

