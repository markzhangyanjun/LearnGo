package iep

type ProcProtect struct {
	Type              string     `json:"type"`
	Ts                int64     `json:"ts"`
	Client            ClientInfo `json:"client"`
	OpType            string     `json:"op_type"`
	Status            string     `json:"status"`
	ParentProcessPath string     `json:"parent_process_path"`
	ParentProcessMD5  string     `json:"parent_process_md5"`
	ChildProcessMD5   string     `json:"child_process_md5"`
	Parameters        string     `json:"parameters"`
	Malname           string     `json:"malname"`
}
