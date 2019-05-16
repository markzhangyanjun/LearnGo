package threat

//HttpDownlodaAndTransInfoDetail_Http网络下载及传输详情

type HttpDownlodaAndTransInfoDetail struct {
	NetEnvionAndTransInfoDetail
	URL string `json:"url" bson:"url"` 

	//TODO:其他
}
