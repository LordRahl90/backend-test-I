package service

import (
	"testing"
	"twitterbot/dataobject"
)

func TestWriteToSpreadSheet(t *testing.T) {
	data := dataobject.SpreadSheetData{
		ProfileName: "Adewale Salau",
		Followers:   5000,
		Description: "Hello to the new world",
		HashTags:    []string{"hello", "no"},
	}
	WriteToSpreadSheet(data)
}
