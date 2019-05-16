package intelligences

type DomainWhoIsContact struct {
	Name         string `json:"name" bson:"name"` 
	Organization string `json:"organization" bson:"organization"` 
	Street1      string `json:"street1" bson:"street1"` 
	City         string `json:"city" bson:"city"` 
	State        string `json:"state" bson:"state"` 
	PostalCode   string `json:"postalCode" bson:"postalCode"` 
	Country      string `json:"country" bson:"country"` 
	CountryCode  string `json:"countryCode" bson:"countryCode"` 
	Email        string `json:"email" bson:"email"` 
	Telephone    string `json:"telephone" bson:"telephone"` 
	Fax          string `json:"fax" bson:"fax"` 
}

type DomainWhoIs struct {
	CreateDate            string             `json:"created_date" bson:"created_date"` 
	UpdatedDate           string             `json:"updated_date" bson:"updated_date"` 
	ExpiresDate           string             `json:"expires_date" bson:"expires_date"` 
	Registrant            DomainWhoIsContact `json:"current_whois" bson:"current_whois"` 
	AdministrativeContact DomainWhoIsContact `json:"administrative_contact" bson:"administrative_contact"` 
}

type ResolutionHistory struct {
	//TODO:resolution_history
}

type DomainIntelligencesDetails struct {
	CurrentWhois      DomainWhoIs         `json:"current_whois" bson:"current_whois"` 
	SubDomains        []string            `json:"sub_domains" bson:"sub_domains"` 
	ResolutionHistory []ResolutionHistory `json:"resolution_history" bson:"resolution_history"` 
}
