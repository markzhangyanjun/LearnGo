package IOCs

type NetworkFileTransfrerEnvionDetailsHTTP struct {
	URL string `json:"url" bson:"url"`
}

type IoAEmail interface{}

type NetworkFileTransfrerEnvionDetailsMails struct {
	SenderIP   string   `json:"mail_sender_ip" bson:"mail_sender_ip"`
	SenderPort int      `json:"mail_sender_port" bson:"mail_sender_port"`
	Emails     IoAEmail `json:"email_addrs" bson:"email_addrs"`
}

type EnvionDetails interface{}

const (
	NetworkFileTransferEnvionDirectionClientToServer = "c2s"
	NetworkFileTransferEnvionDirectionServerToClient = "s2c"
)

type NetworkFileTransferEnvion struct {
	Type      string        `json:"type" bson:"type"`
	Filename  string        `json:"filename" bson:"filename"`
	SrcIP     string        `json:"srcip" bson:"srcip"`
	SrcPort   int           `json:"srcport" bson:"srcport`
	DstIP     string        `json:"dstip" bson:"dstip"`
	Direction string        `json:"direction" bson:"direction"`
	DstPort   int           `json:"dstport" bson:"dstport`
	Proto     []string      `json:"proto" bson:"proto"`
	Details   EnvionDetails `json:"details,omitempty" bson:"details,omitempty"`
}
