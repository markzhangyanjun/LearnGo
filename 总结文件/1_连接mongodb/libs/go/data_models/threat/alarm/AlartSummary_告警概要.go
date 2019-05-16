package alarm

//告警概要

type AlarmState string

const AlarmStateOpen AlarmState = "open"
const AlarmStateClosed AlarmState = "closed"

type AlarmSummary struct {
	State          AlarmState `json:"state" bson:"state"` 
	GeneratedAt    int64      `json:"generated_at" bson:"generated_at"` 
	FirstSeenAt    int64      `json:"first_seen_at" bson:"first_seen_at"` 
	LastSeenAt     int64      `json:"last_seen_at" bson:"last_seen_at"` 
	EventCount     int        `json:"event_count" bson:"event_count"` 
	AttackDuration int64      `json:"attack_duration" bson:"attack_duration"` 
}
