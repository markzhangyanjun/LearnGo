package uselessfile

//CommandAndControl_威胁事件利用命令与控制通道

type CommandAndControl struct {
	Ip     []map[string]interface{} `json:"ip" bson:"ip"` 
	Ipport []map[string]interface{} `json:"ipport" bson:"ipport"` 
	Domain []map[string]interface{} `json:"domain" bson:"domain"` 
	Url    []string                 `json:"url" bson:"url"` 
}
