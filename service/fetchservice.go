package service

import (
	"fmt"
	"twitterbot/dataobject"
	"twitterbot/utility"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

//FetchService - function to fetch The tweets
func FetchService(searchTerm string, accessConfig *dataobject.TwitterAccessConfig) {
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
		}
	}

	//this continues till we r done
	go demux.HandleChan(stream.Messages)
}
