package threat

//NetEnvion_网络环境及传输信息

type NetEnvion struct {
	Envion
	Proto  []string               `json:"Proto" bson:"Proto"` 
	Detail map[string]interface{} `json:"details" bson:"details"` 
}
