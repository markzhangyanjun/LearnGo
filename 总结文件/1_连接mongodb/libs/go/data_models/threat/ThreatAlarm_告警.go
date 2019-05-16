package threat

type AlarmType string

const AlarmTypeThreat AlarmType = "threat"

type AlarmHdr struct {
	Type     AlarmType `json:"type" bson:"type"` 
	ThreatID int       `json:"threat_id" bson:"threat_id"` 
}
