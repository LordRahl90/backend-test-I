package utility

import (
	"fmt"
	"log"
	"testing"
)

func TestSplitString(t *testing.T) {
	result := SplitString("hello,world moving create, dontStop")
	strLength := len(result)

	if strLength != 5 {
		log.Fatal("Length of the string should be 5 actualt:", strLength)
	}
}

func TestReadTwitterConfigurations(t *testing.T) {
	result := ReadTwitterConfigurations()
	if result == nil {
		log.Fatal("Invalid Config File Detected")
	}
}

func TestGetGoogleClient(t *testing.T) {
	service := GetGoogleClient()
	if service == nil {
		log.Fatal("Invalid Service Detected")
	}
}

func TestReadSpreadSheetConfiguration(t *testing.T) {
	config := ReadSpreadSheetConfiguration()
	fmt.Println(config)
	if config.SpreadSheetID == "" {
		log.Fatal("Config File is empty.... ")
	}
}

func TestFetchSpreadSheetID(t *testing.T) {
	spreadSheetID := FetchSpreadSheetID()
	log.Fatal(spreadSheetID)
}

func TestCreateNewSpreadSheet(t *testing.T) {
	result := CreateNewSpreadSheet()
	log.Fatal(result)
}

func TestSaveSpreadSheetID(t *testing.T) {
	spreadSheetID := "1A-m9VxUoYa_Tu96bmNMlYYpM-ubp8aFls8CkvW0225k"
	SaveSpreadSheetID(spreadSheetID)
}
