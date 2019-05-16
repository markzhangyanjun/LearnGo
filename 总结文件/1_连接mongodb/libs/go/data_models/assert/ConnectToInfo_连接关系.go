package assert

//连接关系

type ConnectToInfo struct {
	BasicAssertInfo
	ConnectType string `json:"connect_type" bson:"connect_type"` 
}
