package threat_event

import "LearnGo/总结文件/1_连接mongodb/libs/go/data_models/threat/ntctf"

//威胁行为

type ThreatAction struct {
	NTCTF ntctf.NTCTFDescription `json:"ntctf" bson:"ntctf"` 
}
