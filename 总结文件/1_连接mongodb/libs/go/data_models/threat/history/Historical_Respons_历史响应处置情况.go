package history

import "LearnGo/总结文件/1_连接mongodb/libs/go/data_models/assert"

//历史响应处理情况
type HistoricalResponse struct {
	Asserts        assert.BasicAssertInfo  `json:"assert" bson:"assert"` 
	ResponseStatus string                  `json:"response_status" bson:"response_status"` 
	TimeLine       HistoryResponseTimeline `json:"time_line" bson:"time_line"` 
}

//历史响应处置时间轴

type HistoryResponseTimeline struct {
	Potential_impact_found int `json:"potential_impact_found" bson:"potential_impact_found"` 
	Potential_impact_fixed int `json:"potential_impact_fixed" bson:"potential_impact_fixed"` 
	Direct_impact_found    int `json:"direct_impact_found" bson:"direct_impact_found"` 
	Direct_impact_fixed    int `json:"direct_impact_fixed" bson:"direct_impact_fixed"` 
}
