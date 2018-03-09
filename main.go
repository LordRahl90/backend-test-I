package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"twitterbot/service"
	"twitterbot/utility"
)

func main() {

	accessConfig := utility.ReadTwitterConfigurations()

FETCHPARAMS:
	searchTerms := startUp()

	if len(searchTerms) <= 0 || len(searchTerms[0]) <= 0 {
		fmt.Println("Please provide at lease a search term...")
		goto FETCHPARAMS
	}

	for _, v := range searchTerms {
		if v != "" || v == " " {
			go service.FetchService(v, accessConfig)
		}
	}

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	fmt.Println("Stopping Service")
}

//Accepts the user input and splits it based on the delimeter
func startUp() (searchTerms []string) {
	var input string
	fmt.Println("Enter Each Hashtag Separated by a comma or space and press enter to proceed:")
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')

	searchTerms = utility.SplitString(input)
	fmt.Println("Filtering streams based on the following hashtags: ", searchTerms)
	return
}
