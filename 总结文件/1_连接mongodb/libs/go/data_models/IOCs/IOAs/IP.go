package IOAs

type IP struct {
	IPStr   string `json:"ip" bson:"ip"` 
	Details struct {
		IPInfo struct {
			IPStr                    string `json:"ip" bson:"ip"` 
			Country                  string `json:"country" bson:"country"` 
			Region                   string `json:"region" bson:"region"` 
			City                     string `json:"city" bson:"city"` 
			Unit                     string `json:"unit" bson:"unit"` 
			Isp                      string `json:"isp" bson:"isp"` 
			Time_zones1              string `json:"time_zones1" bson:"time_zones1"` 
			Time_zones2              string `json:"time_zones2" bson:"time_zones2"` 
			Area_code                string `json:"area_code" bson:"area_code"` 
			International_phone_code string `json:"international_phone_code" bson:"international_phone_code"` 
			Country_codes            string `json:"country_codes" bson:"country_codes"` 
			World_code_string        string `json:"world_code_string" bson:"world_code_string"` 

			Location struct {
				Latitude  string `json:"latitude" bson:"latitude"` 
				Longitude string `json:"longitude" bson:"longitude"` 
			} `json:"location" bson:"location"` 
		} `json:"ipinfo" bson:"ipinfo"` 
	} `json:"details" bson:"details"` 
}

type IPPort struct {
	IP
	Port  int      `json:"port" bson:"port"` 
	Proto []string `json:"proto" bson:"proto"` 
}
