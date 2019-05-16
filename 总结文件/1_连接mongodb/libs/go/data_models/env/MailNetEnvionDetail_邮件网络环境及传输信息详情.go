package threat

//MailNetEnvionDetail_邮件网络环境及传输信息详情

type MailNetEnvionDetail struct {
	NetEnvionAndTransInfoDetail
	Mail_sender_ip   string `json:"mail_sender_ip" bson:"mail_sender_ip"` 
	Mail_sender_port int    `json:"mail_sender_port" bson:"mail_sender_port"` 

	//TODO：其他
}
