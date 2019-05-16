package IOAs

import "LearnGo/总结文件/1_连接mongodb/libs/go/data_models/intelligences"

type Domain struct {
	Domain  string                                    `json:"domain" bson:"domain"`
	Details *intelligences.DomainIntelligencesDetails `json:"details,omitempty" bson:"details,omitempty"`
}
