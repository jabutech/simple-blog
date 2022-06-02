package helper

import "log"

// FatalError handle error with message
func FatalError(message string, err error) {
	if err != nil {
		log.Fatal(message, err.Error())
	}
}
