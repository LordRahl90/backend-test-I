package utility

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"twitterbot/dataobject"

	"golang.org/x/oauth2/google"
	spreadsheet "gopkg.in/Iwark/spreadsheet.v2"
)

//FailOnError function
func FailOnError(err error, msg string) {
	if err != nil {
		log.Println(msg)
		log.Panic(err)
	}
}

//SplitString based on space or comma
func SplitString(input string) (result []string) {
	result = strings.FieldsFunc(input, Split)
	return
}

//Split function based on rune
func Split(r rune) bool {
	return r == ',' || r == ' '
}

//GetGoogleClient - function to retrieve googles client API
func GetGoogleClient() (service *spreadsheet.Service) {
	data, err := ioutil.ReadFile("client_secret.json")
	fmt.Println(os.Getwd())
	FailOnError(err, "Cannot Load the Service configuration file")

	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	FailOnError(err, "Cannot Load the proper configurations")

	client := conf.Client(context.TODO())
	service = spreadsheet.NewServiceWithClient(client)
	return
}

//ReadTwitterConfigurations - function that Reads the configuration file
func ReadTwitterConfigurations() (config *dataobject.TwitterAccessConfig) {
	//reading the configuration file
	fileContent, err := ioutil.ReadFile("config.json")
	FailOnError(err, "Cannot Read Config File")

	//convert the read file to an object
	err = json.Unmarshal(fileContent, &config)
	FailOnError(err, "Invalid JSON format detected")
	return
}

//ReadSpreadSheetConfiguration - function to read the spreadsheet configuration
func ReadSpreadSheetConfiguration() (config *dataobject.SpreadSheetConfig) {
	data, err := ioutil.ReadFile("spreadsheet.json")
	FailOnError(err, "Cannot Load Spreadsheet Configuration file")

	err = json.Unmarshal(data, &config)
	FailOnError(err, "Invalid JSON Format Detected")

	fmt.Println("Config file loaded successfully...")
	return
}

//FetchSpreadSheetID - function to retrieve the spreadsheet ID
func FetchSpreadSheetID() (spreadSheetID string) {
	spreadsheetInfo := ReadSpreadSheetConfiguration()

	spreadSheetID = spreadsheetInfo.SpreadSheetID

	if spreadSheetID == "" {
		fmt.Println("Creating New Spreadsheet")
		spreadSheetID = CreateNewSpreadSheet()
	} else {
		fmt.Println("Spreadsheet ID exists")
	}

	return
}

//CreateNewSpreadSheet - function to create a new spreadsheet
func CreateNewSpreadSheet() (spreadsheetID string) {
	service := GetGoogleClient()
	newSpreadSheet, err := service.CreateSpreadsheet(spreadsheet.Spreadsheet{
		Properties: spreadsheet.Properties{
			Title: "Hello World Twitter Bot",
		},
	})

	FailOnError(err, "Cannot Create The Defined Spreadsheet")

	fmt.Println(newSpreadSheet)
	fmt.Println(newSpreadSheet.ID)
	go SaveSpreadSheetID(newSpreadSheet.ID)
	return newSpreadSheet.ID
}

//SaveSpreadSheetID - function to update The spreadsheetID
func SaveSpreadSheetID(spreadSheetID string) {
	var spreadSheetConfig dataobject.SpreadSheetConfig
	spreadSheetConfig.SpreadSheetID = spreadSheetID

	b, err := json.Marshal(spreadSheetConfig)
	FailOnError(err, "Cannot Convert to JSON data")
	err = ioutil.WriteFile("spreadsheet.json", b, os.ModePerm)
	FailOnError(err, "Cannot Write to file spreadsheet.json")
}
