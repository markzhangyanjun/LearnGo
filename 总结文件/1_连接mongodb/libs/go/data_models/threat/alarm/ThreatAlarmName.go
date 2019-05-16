package alarm

type RuleIDType string

const ThreatAlarmNameRuleIDCodeStatic RuleIDType = "/acd_access/code/20190414t1"

type ThreatAlarmName struct {
	Name    string                     `json:"name" bson:"name"`
	RuleID  RuleIDType                 `json:"naming_rule_id" bson:"naming_rule_id"`
	Details ThreatAlarmNameRuleDetails `json:"naming_rule_detail" bson:"naming_rule_detail"`
}

type ThreatAlarmNameRuleDetails interface{}

type ThreatAlarmNameRuleByAttackerAndVictimDetails struct {
	AttackerID        string   `json:"attacker_id" bson:"attacker_id"`
	VictimAssertGroup []string `json:"asset_groups" bson:"asset_groups"`
}
