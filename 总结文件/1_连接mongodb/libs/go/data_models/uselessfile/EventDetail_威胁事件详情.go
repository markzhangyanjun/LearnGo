package uselessfile

//EventDetail_威胁事件详情

type EventDetail struct {
	EventEvidencesSummary
	Content map[string]interface{} `json:"content" bson:"content"` 
}
