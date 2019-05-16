package threat_event

type ThreatSource struct {
	AttackerID        string            `json:"attacker_id" bson:"attacker_id"` 
	MalwareScene      MalwareScene      `json:"malware" bson:"malware"` 
	CommandAndControl CommandAndControl `json:"command_and_control" bson:"command_and_control"` 
}
