package assert

//资产

type Assert interface{}

type BasicAssertInfo struct {
	AssertID      string `json:"assert_id" bson:"assert_id"` 
	AssertVersion int    `json:"assert_version" bson:"assert_version"` 
	Type          string `json:"type" bson:"type"` 
	Domain        string `json:"domain" bson:"domain"` 
	AssertGroup   string `json:"assert_group" bson:"assert_group"` 
	User          string `json:"user" bson:"user"` 
	Maintainer    string `json:"maintainer" bson:"maintainer"` 
}

//资产列表
type AssertList struct {
	Count      int      `json:"count" bson:"count"` 
	AssertList []Assert `json:"assert_list" bson:"assert_list"` 
}
