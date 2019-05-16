package threat

//Envion_环境及传输信息

type Envion struct {
	Type     string `json:"type" bson:"type"` 
	Filename string `json:"filename" bson:"filename"` 
}
