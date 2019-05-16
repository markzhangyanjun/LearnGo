package alarm

import "LearnGo/总结文件/1_连接mongodb/libs/go/data_models/IOCs"



//告警相关信息

type RelatedDetails interface{}

type AlarmRelatedInfo struct {
	IntelligenceEvents []IntelligencesEvents `json:"intelligence_events" bson:"intelligence_events"`
	Cases              []Cases               `json:"cases" bson:"cases"`
}

//告警相关外部事件
type IntelligencesEvents struct {
	Type    string         `json:"type" bson:"type"`
	Name    string         `json:"name" bson:"name"`
	Items   []KeyValue     `json:"items" bson:"items"`
	Details RelatedDetails `json:"details,omit_empty" bson:"details"`
}

//事件相关内容
type KeyValue struct {
	Key   string `json:"key" bson:"key"`
	Value string `json:"value" bson:"value"`
}

//相关Case
type Cases struct {
	Type        string         `json:"type" bson:"type"`
	Name        string         `json:"name" bson:"name"`
	Level       string         `json:"level" bson:"level"`
	Description string         `json:"description" bson:"description"`
	Details     RelatedDetails `json:"details,omit_empty" bson:"details,omit_empty"`
}

//案例详情

type RelatedFindMethod string

const RelatedFindMethodByVulnerability RelatedFindMethod = "/related/cases/vulnerability"
const RelatedIntelligencesEventsCaseFindMethodByIoc RelatedFindMethod = "/related/intelligences_events/ioc"

type RelatedDetailsByAttacker struct {
	FindMethod RelatedFindMethod `json:"find_method" bson:"find_method"`
	IoCList    []IOCs.IIoA       `json:"ioc_list" bson:"ioc_list"`
}

type RelatedDetailsByVulnerability struct {
	FindMethod         RelatedFindMethod `json:"find_method" bson:"find_method"`
	VulnerabilityLists []IOCs.IIoE       `json:"vulnerability_lists" bson:"vulnerability_lists"`
}
