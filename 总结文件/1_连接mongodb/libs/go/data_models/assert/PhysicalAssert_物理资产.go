package assert

//物理资产
type PhysicalAssert struct {
	BasicAssertInfo
	Detail PhysicalAssertDetail `json:"detail" bson:"detail"` 
}

//物理资产详情
type PhysicalAssertDetail struct {
	PhysicalAddress PhysicalAddress  `json:"physical_address" bson:"physical_address"` 
	NetInfo         NetInfo          `json:"network" bson:"network"` 
	ConnectTo       []*ConnectToInfo `json:"connect_to" bson:"connect_to"` 
	UsedBy          PersonnelDetails `json:"used_by" bson:"used_by"` 
	MaintenanceBy   PersonnelDetails `json:"maintenance_by" bson:"maintenance_by"` 
}

//PhysicalAddress_物理地址

type PhysicalAddress struct {
	Position PhysicalAddressPosition `json:"position" bson:"position"` 
}

//网络信息

type NetInfo struct {
	ManagementIp []NetworkConnInfoDetails `json:"management_ip" bson:"management_ip"` 
	InUse        []NetworkConnInfoDetails `json:"in_use" bson:"in_use"` 
}

//PersonnelDetails_人员详情

type PersonnelDetails struct {
	HumanId   string `json:"human_id" bson:"human_id"` 
	Name      string `json:"name" bson:"name"` 
	Group     string `json:"group" bson:"group"` 
	RankName  string `json:"rank_name" bson:"rank_name"` 
	RankLevel int    `json:"rank_level" bson:"rank_level"` 
}

//PhysicalAddressPosition_物理地理位置
type PhysicalAddressPosition struct {
	City    string `json:"city" bson:"city"` 
	Region  string `json:"region" bson:"region"` 
	Unit    string `json:"unit" bson:"unit"` 
	Address string `json:"address" bson:"address"` 
}

type ConnInfoType string

const NetworkConnInfoDetailsTypeIPV4 ConnInfoType = "ipv4"
const NetworkConnInfoDetailsTypeIPV6 ConnInfoType = "ipv6"

//IP访问地点
type NetworkConnInfoDetails struct {
	Type    ConnInfoType `json:"type" bson:"type"` 
	Ip      string       `json:"ip" bson:"ip"` 
	Network string       `json:"network" bson:"network"` 
	Gateway string       `json:"gateway" bson:"gateway"` 
	Mac     string       `json:"mac" bson:"mac"` 
	AclId   int          `json:"acl_id" bson:"acl_id"` 
}
