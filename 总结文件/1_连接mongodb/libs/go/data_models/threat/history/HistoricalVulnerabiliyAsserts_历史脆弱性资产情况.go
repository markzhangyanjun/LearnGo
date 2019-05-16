package history

import "LearnGo/总结文件/1_连接mongodb/libs/go/data_models/assert"

//历史脆弱性资产情况

type HistoricalVulnerabiliyAsserts struct {
	Asserts             assert.Assert         `json:"assert" bson:"assert"` 
	VulnerabiliyHistory []VulnerabiliyHistory `json:"vulnerabiliy_history" bson:"vulnerabiliy_history"` 
}

//脆弱性性存在历史项
type VulnerabiliyHistory struct {
	VulnerabilityId string                           `json:"vulnerability_id" bson:"vulnerability_id"` 
	Status          assert.AssertVulnerabilityStatus `json:"status" bson:"status"` 
	TimeLine        VulnerabiliyHistoryTimeline      `json:"time_line" bson:"time_line"` 
}

//脆弱性存在历史项时间轴
type VulnerabiliyHistoryTimeline struct {
	//脆弱性引入硬件、软件或固件的时间。通常是最早受影响的软件的发布时间
	VulnerabilityIntroducedTime int64 `json:"vulnerability_introduced_time" bson:"vulnerability_introduced_time"` 
	//安全通告发布时间（对公众公布时间）
	SecurityAdvisoryFound int64 `json:"security_advisory_found" bson:"security_advisory_found"` 
	//POC发现事件（在此时间后，可能受到攻击）
	WildlyPoCSeen int64 `json:"wildly_poc_seen" bson:"wildly_poc_seen"` 
	//补丁发布时间（在此时间后，可以根除此脆弱性）
	PatchReleased int64 `json:"patch_released" bson:"patch_released"` 

	//此硬件、软件资产在网内首次出现时间
	AssertFirstSeen int64 `json:"assert_first_seen" bson:"assert_first_seen"` 
	//脆弱性在此资产上发现时间（被扫描器等，首次发现时间）
	VulnerabilityFound int64 `json:"vulnerability_found" bson:"vulnerability_found"` 
	//检测到的利用此脆弱性攻击的时间
	VulnerabilityAttack []int64 `json:"vulnerability_attack_detected" bson:"vulnerability_attack_detected"` 

	//对此脆弱性响应的开始
	VulnerabilityResponseStart int64 `json:"vulnerability_response_start" bson:"vulnerability_response_start"` 
	//对此脆弱性影响的结束时间
	VulnerabilityFixedTime int64 `json:"vulnerability_fixed_time" bson:"vulnerability_fixed_time"` 
	//修复确认时间
	FixVerificationTime int64 `json:"fix_verification_time" bson:"fix_verification_time"` 
}
