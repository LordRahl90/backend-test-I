package dataobject

// TwitterAccessConfig - struct to save the access details
type TwitterAccessConfig struct {
	ConsumerKey    string `json:"consumerKey"`
	ConsumerSecret string `json:"consumerSecret"`
	AccessToken    string `json:"accessToken"`
	AccessSecret   string `json:"accessSecret"`
}

//SpreadSheetConfig - Spreadsheet configuration struct
type SpreadSheetConfig struct {
	SpreadSheetID string `json:"spreadsheetID"`
}

//SpreadSheetData Object
type SpreadSheetData struct {
	ProfileName string   `json:"profileName"`
	Followers   int      `json:"noOfFollowers"`
	Description string   `json:"userDescription"`
	HashTags    []string `json:"hasTags"`
}
