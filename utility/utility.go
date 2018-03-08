package utility

import (
	"log"
	"strings"
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

//GetTwitterClient - function to get the http twitter client
func GetTwitterClient() {

}
