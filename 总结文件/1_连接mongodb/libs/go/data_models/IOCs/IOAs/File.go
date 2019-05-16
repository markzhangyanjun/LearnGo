package IOAs

import "LearnGo/总结文件/1_连接mongodb/libs/go/data_models/IOCs"

type File struct {
	MD5      string        `json:"md5" bson:"md5"` 
	FileType string        `json:"filetype" bson:"filetype"` 
	Size     uint64        `json:"size" bson:"size"` 
	Malname  string        `json:"malname" bson:"malname"` 
	Envion   []IOCs.Envion `json:"envion" bson:"envion"` 
}
