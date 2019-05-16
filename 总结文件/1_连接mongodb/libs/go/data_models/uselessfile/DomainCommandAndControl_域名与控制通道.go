package uselessfile

//DomainCommandAndControl_域名与控制通道

type DomainCommandAndControl struct {
	Domain  string                 `json:"domain" bson:"domain"` 
	Details map[string]interface{} `json:"details" bson:"details"` 
}
