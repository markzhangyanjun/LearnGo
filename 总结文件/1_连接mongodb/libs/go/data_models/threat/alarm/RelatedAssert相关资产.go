package alarm

import "LearnGo/总结文件/1_连接mongodb/libs/go/data_models/assert"

//相关资产

type RelatedAssert struct {
	ImpactsDirectly  assert.AssertList `json:"impacts_directly" bson:"impacts_directly"` 
	ImpactsPotential assert.AssertList `json:"impacts_potential" bson:"impacts_potential"` 
}
