package service

import (
	"fmt"
	"strconv"
	"twitterbot/dataobject"
	"twitterbot/utility"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

//FetchService - function to fetch The tweets
func FetchService(searchTerm string, accessConfig *dataobject.TwitterAccessConfig) {
	//TODO: Refactor the config process into an utility function
	config := oauth1.NewConfig(accessConfig.ConsumerKey, accessConfig.ConsumerSecret)
	token := oauth1.NewToken(accessConfig.AccessToken, accessConfig.AccessSecret)

	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	//use stream to search, so as to make it continous
	params := &twitter.StreamFilterParams{
		Track:         []string{searchTerm},
		StallWarnings: twitter.Bool(true),
	}
	demux := twitter.NewSwitchDemux()

	stream, err := client.Streams.Filter(params)
	utility.FailOnError(err, "Cannot Start Stream")

	demux.Tweet = func(tweet *twitter.Tweet) {
		profileName := tweet.User.Name
		userDescription := tweet.User.Description
		followersCount := tweet.User.FollowersCount
		hashTags := []string{searchTerm}

		if followersCount >= 1000 && followersCount <= 50000 {
			//we consider these users
			fmt.Printf("Profile Name: %s\n Number of Followers: %d\n User description: %s\n HashTags: %s\n\n", profileName, followersCount, userDescription, hashTags)
			fmt.Println("-----------------------------------------------------------------------------------------------------------")
			//create an instance of the user
			object := dataobject.SpreadSheetData{
				ProfileName: profileName,
				Followers:   followersCount,
			}
			go WriteToSpreadSheet(object)
		}
	}

	//this continues till we r done
	go demux.HandleChan(stream.Messages)
}

//WriteToSpreadSheet - This writes the array values into the spreadsheet
func WriteToSpreadSheet(input dataobject.SpreadSheetData) {
	service := utility.GetGoogleClient()
	spreadSheetID := utility.FetchSpreadSheetID()

	spreadsheet, err := service.FetchSpreadsheet(spreadSheetID)
	utility.FailOnError(err, "We cannot fetch the specified spreadsheet")

	//get the first sheet
	sheet, err := spreadsheet.SheetByIndex(0)
	utility.FailOnError(err, "Cannot Find the specified Sheet")

	//Sets the headers of the spreadsheet
	sheet.Update(0, 0, "Profile Name")
	sheet.Update(0, 1, "Followers Count")

	//Iterate through the rows to see if the first column is empty.
	//if its empty, we keep the value in that row.
	//This is buggy as many go-routines can try to access this point at the same time, thereby leading to overwritten records.
	for r := range sheet.Rows {
		if sheet.Rows[r][0].Value == "" {
			sheet.Update(r, 0, input.ProfileName)
			sheet.Update(r, 1, strconv.Itoa(input.Followers))
			break
		}
	}

	err = sheet.Synchronize()
	utility.FailOnError(err, "Cannot Synchronize the spreadsheet")
	fmt.Println("Sheet Updated Successfully..")
}
