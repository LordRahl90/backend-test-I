package dataobject

// TwitterAccessConfig - struct to save the access details
type TwitterAccessConfig struct {
	ConsumerKey    string `json:"consumerKey"`
	ConsumerSecret string `json:"consumerSecret"`
	AccessToken    string `json:"accessToken"`
	AccessSecret   string `json:"accessSecret"`
}

//UserProfile Object
type UserProfile struct {
	ProfileName string   `json:"profileName"`
	Followers   int      `json:"noOfFollowers"`
	Description int      `json:"userDescription"`
	HashTags    []string `json:"hasTags"`
}
