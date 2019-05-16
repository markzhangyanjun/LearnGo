package threat_event

type EvidenceType string

const PTDSourceAny EvidenceType = "/source/ptd/<any>"
const PTDSourceHttp EvidenceType = "/source/ptd/http"
const PTDSourceDns EvidenceType = "/source/ptd/dns"
const PTDSourceMail EvidenceType = "/source/ptd/mail"
const PTDSourceMySql EvidenceType = "/source/ptd/mysql"
const PTDSourceFtp EvidenceType = "/source/ptd/ftp"
const PTDSourceGenericFileTransmit EvidenceType = "/source/ptd/generic_file"

type Evidences struct {
	Type EvidenceType `json:"type" bson:"type"` 

	ID string `json:"evidences_id" bson:"evidences_id"` 

	Content interface{} `json:"content,omitempty" bson:"content,omitempty"` 
}
