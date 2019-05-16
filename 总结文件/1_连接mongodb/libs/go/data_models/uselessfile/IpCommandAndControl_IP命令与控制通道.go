package uselessfile

//IpCommandAndControl_IP命令与控制通道

type IpCommandAndControl struct {
	Ip      string                 `json:"ip" bson:"ip"` 
	Details map[string]interface{} `json:"details" bson:"details"` 
}
