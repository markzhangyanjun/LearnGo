package threat_event

import "LearnGo/总结文件/1_连接mongodb/libs/go/data_models/IOCs/IOAs"

type CommandAndControl struct {
	IP     []IOAs.IP     `json:"ip" bson:"ip"` 
	IPPort []IOAs.IPPort `json:"ipport" bson:"ipport"` 
	Domain []IOAs.Domain `json:"domain" bson:"domain"` 
	Url    []IOAs.Url    `json:"url" bson:"url"` 
}
