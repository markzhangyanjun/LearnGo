package threat

//EndEnvion_终端环境及传递信息

type EndEnvion struct {
	Envion
	Ctime int `json:"ctime" bson:"ctime"` 
	Mtime int `json:"mtime" bson:"mtime"` 
	Atime int `json:"atime" bson:"atime"` 
}
