package threat_event

import (
	"LearnGo/总结文件/1_连接mongodb/libs/go/data_models/IOCs/IOAs"
	"LearnGo/总结文件/1_连接mongodb/libs/go/data_models/IOCs/IOEs"
	"LearnGo/总结文件/1_连接mongodb/libs/go/data_models/assert"
	"LearnGo/总结文件/1_连接mongodb/libs/go/data_models/tags"
	"gopkg.in/mgo.v2/bson"
)

//威胁事件

type Event struct {
	ID                bson.ObjectId        `json:"-" bson:"_id"`
	EventID           string               `json:"event_id" bson:"event_id"`
	GeneratedAt       int64                `json:"generated_at" bson:"generated_at"`
	FirstSeenAt       int64                `json:"first_seen_at" bson:"first_seen_at"`
	LastSeenAt        int64                `json:"last_seen_at" bson:"last_seen_at"`
	LogCount          int                  `json:"log_count" bson:"log_count"`
	ThreatSource      ThreatSource         `json:"threat_source" bson:"threat_source"`
	VulnerabilityUsed []IOEs.Vulnerability `json:"vulnerability_used" bson:"vulnerability_used"`
	PayloadUsed       []IOAs.File          `json:"payload_used" bson:"payload_used"`
	Victim            assert.Assert        `json:"victim" bson:"victim"`
	Tags              []tags.Tags          `json:"tags" bson:"tags"`
	Action            ThreatAction         `json:"action" bson:"action"`
	Evidences         []Evidences          `json:"evidences" bson:"evidences"`
	Details           []Evidences          `json:"details,omitempty" bson:"details,omitempty"`
}
