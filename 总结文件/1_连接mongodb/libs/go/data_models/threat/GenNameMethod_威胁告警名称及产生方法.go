package threat

//威胁告警名称及产生方法

type GenNameMethod struct {
	Name             string          `json:"name" bson:"name"` 
	NamingRuleId     string          `json:"naming_rule_id" bson:"naming_rule_id"` 
	NamingRuleDetail GenMethodDetail `json:"naming_rule_detail" bson:"naming_rule_detail"` 
}

type GenMethodDetail struct {
}

type AttackerAssert struct {
	GenMethodDetail
	AttackerId  string   `json:"attacker_id" bson:"attacker_id"` 
	AssetGroups []string `json:"asset_groups" bson:"asset_groups"` 
}
