package ntctf

//NtctfDescription_威胁行为描述
type NTCTFDescription struct {
	Stage      string `json:"stage" bson:"stage"` 
	Objective  string `json:"objective" bson:"objective"` 
	Action     string `json:"action" bson:"action"` 
	KeyPhrases string `json:"key_phrases" bson:"key_phrases"` 
}

func NewUnknownNTCTFDescription() NTCTFDescription {
	return NTCTFDescription{
		"未知",
		"未知",
		"未知",
		"未知",
	}
}

func NewNTCTFDescription(stage string, objective string, action string, keyphrases string) NTCTFDescription {
	return NTCTFDescription{
		stage,
		objective,
		action,
		keyphrases,
	}
}
