package uselessfile

//时间监控点

type TimeWatched struct {
	Birth   int `json:"birth" bson:"birth"` 
	Execute int `json:"execute" bson:"execute"` 
}
