package utility

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

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

func getGoogleClient() {
	data, err := ioutil.ReadFile("client_secret.json")
	FailOnError(err, "Cannot Load the Service configuration file")

	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	FailOnError(err, "Cannot Load the proper configurations")

	client := conf.Client(context.TODO())
	service := spreadsheet.NewServiceWithClient(client)

	fmt.Printf("%T", service)
}
