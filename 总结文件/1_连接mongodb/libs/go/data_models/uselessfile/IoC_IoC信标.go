package uselessfile

//IoC信标

type IoC struct {
	Type string `json:"type" bson:"type"` 
}

//URL类型IoC
type URLTypeAndIoC struct {
	IoC
	Domain string `josn:"domain"`
}

//IP类型及Ioc

type IpTypeAndIoc struct {
	IoC
	Domain string `json:"domain" bson:"domain"` 
}

//域名累成Ioc

type DomainTypeIoC struct {
	IoC
	Domain string `json:"domain" bson:"domain"` 
}
