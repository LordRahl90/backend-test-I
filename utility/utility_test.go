package utility

import (
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
