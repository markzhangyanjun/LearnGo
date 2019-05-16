package cases

//案例详情

type CaseDetail struct {
	FindMethod         string              `json:"find_method" bson:"find_method"` 
	VulnerabilityLists []*VulnerabilityIoE `json:"vulnerability_lists" bson:"vulnerability_lists"` 
}
