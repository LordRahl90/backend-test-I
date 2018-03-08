package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
	"twitterbot/dataobject"
	"twitterbot/service"
	"twitterbot/utility"
)

func main() {

	var input string
	fmt.Println("Enter Each Hashtag Separated by a comma or space and press enter to proceed:")
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')

	//reading the configuration file
	fileContent, err := ioutil.ReadFile("config.json")
	utility.FailOnError(err, "Cannot Read Config File")

	//convert the read file to an object
	var accessConfig dataobject.TwitterAccessConfig
	err = json.Unmarshal(fileContent, &accessConfig)

	searchTerms := utility.SplitString(input)

	for _, v := range searchTerms {
		if v != "" || v == " " {
			go service.FetchService(v, &accessConfig)
		}
	}

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	fmt.Println("Stopping Service")
}
