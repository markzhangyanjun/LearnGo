package iep

type ClientInfo struct {
	Name      string `json:"name"`
	IP        string `json:"ip"`
	UUID      string `json:"uuid"`
	GroupID   int    `json:"group_id"`
	CompanyID int    `json:"company_id"`
	ServerIP  string `json:"server_ip"`
	ServerUID string `json:"server_uid"`
}

