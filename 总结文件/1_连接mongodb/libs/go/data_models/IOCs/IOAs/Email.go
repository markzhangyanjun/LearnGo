package IOAs

type Email struct {
	Sender  string   `json:"sender" bson:"sender"` 
	From    string   `json:"from" bson:"from"` 
	To      []string `json:"to" bson:"to"` 
	Cc      []string `json:"cc" bson:"cc"` 
	Bcc     []string `json:"bcc" bson:"bcc"` 
	Subject string   `json:"subject" bson:"subject"` 
}
