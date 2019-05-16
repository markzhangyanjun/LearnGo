package IOCs

type HostFileEnvionDetails struct {
	Type string `json:"type" bson:"type"`

	Filename    string `json:"filename" bson:"filename"`
	TimeWatched struct {
		Birth   int64 `json:"birth" bson:"birth"`
		Execute int64 `json:"execute" bson:"execute"`
	} `json:"time_watched" bson:"time_watched"`
	CTime int64 `json:"ctime" bson:"ctime"`
	MTime int64 `json:"mtime" bson:"mtime"`
	ATime int64 `json:"atime" bson:"atime"`
}
