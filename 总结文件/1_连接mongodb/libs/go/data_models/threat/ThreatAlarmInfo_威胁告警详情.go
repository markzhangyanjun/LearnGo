package threat

import (
	"LearnGo/总结文件/1_连接mongodb/libs/go/data_models/IOCs"
	"LearnGo/总结文件/1_连接mongodb/libs/go/data_models/tags"
	"LearnGo/总结文件/1_连接mongodb/libs/go/data_models/threat/alarm"
	"LearnGo/总结文件/1_连接mongodb/libs/go/data_models/threat/attacker"
	"LearnGo/总结文件/1_连接mongodb/libs/go/data_models/threat/evaluate"
	"LearnGo/总结文件/1_连接mongodb/libs/go/data_models/threat/history"
)

//威胁告警详情
type ThreatAlarm struct {
	Type                          AlarmType                               `json:"type" bson:"type"`
	ThreatID                      string                                  `json:"threat_id" bson:"threat_id"`
	ThreatName                    alarm.ThreatAlarmName                   `json:"name" bson:"name"`
	AlarmSummary                  alarm.AlarmSummary                      `json:"summary" bson:"summary"`
	ConcernedBy                   []string                                `json:"concerned_by" bson:"concerned_by"`
	AlartTag                      []tags.Tags                             `json:"alert_tags" bson:"alert_tags"`
	Evaluations                   evaluate.ThreatEvaluations              `json:"evaluations" bson:"evaluations"`
	Attacker                      attacker.Attacker                       `json:"attacker" bson:"attacker"`
	Related                       alarm.AlarmRelatedInfo                  `json:"related" bson:"related"`
	Asserts                       alarm.RelatedAssert                     `json:"asserts" bson:"asserts"`
	ThreatEvents                  EventList                               `json:"threat_events" bson:"threat_events"`
	HistoricalVulnerabiliyAsserts []history.HistoricalVulnerabiliyAsserts `json:"historical_vulnerabiliy_asserts" bson:"historical_vulnerabiliy_asserts"`
	HistoricalResponseAsserts     []history.HistoricalResponse            `json:"historical_response_asserts" bson:"historical_response_asserts"`
	IoAs                          []IOCs.IIoA                             `json:"IoAs" bson:"IoAs"`
	IoEs                          []IOCs.IIoE                             `json:"IoEs" bson:"IoEs"`
}
