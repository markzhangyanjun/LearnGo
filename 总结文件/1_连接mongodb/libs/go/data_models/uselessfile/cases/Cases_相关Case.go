package cases

//相关Case
type Cases struct {
	Type        string      `json:"type" bson:"type"` 
	Name        string      `json:"name" bson:"name"` 
	Level       string      `json:"level" bson:"level"` 
	Description string      `json:"description" bson:"description"` 
	Details     *CaseDetail `json:"details" bson:"details"` 
}
