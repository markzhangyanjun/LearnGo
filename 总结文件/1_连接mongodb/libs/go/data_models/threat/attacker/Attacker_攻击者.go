package attacker

import "LearnGo/总结文件/1_连接mongodb/libs/go/data_models/IOCs"

//攻击者

type Attacker struct {
	ID          string               `json:"id" bson:"id"` 
	Name        string               `json:"name" bson:"name"` 
	Type        string               `json:"type" bson:"type"` 
	Country     string               `json:"country" bson:"country"` 
	Description string               `json:"description" bson:"description"` 
	Details     IAttackerFindDetails `json:"details" bson:"details"` 
}

//攻击者寻找详情
type IAttackerFindDetails interface{}

type AttackerFindMethod string

const AttackerFindMethodByIntelligence AttackerFindMethod = "/pointer/attacker/intelligence/ioc"
const AttackerFindMethodByIocDeduce AttackerFindMethod = "/pointer/attacker/deduce/ioc"

//知识与情报直接指出
type PointedByIoc struct {
	FindMethod AttackerFindMethod `json:"find_method" bson:"find_method"` 
	IoaList    []IOCs.IIoA        `json:"ioc_list" bson:"ioc_list"` 
}
