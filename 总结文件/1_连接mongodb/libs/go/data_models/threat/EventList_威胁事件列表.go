package threat

import "LearnGo/总结文件/1_连接mongodb/libs/go/data_models/threat/threat_event"

//威胁事件列表

type EventList struct {
	Count     int                   `json:"count" bson:"count"` 
	EventList []*threat_event.Event `json:"eventlist" bson:"eventlist"` 
}
