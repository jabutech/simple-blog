package helper

import "log"

func FatalError(message string, err error) {
	if err != nil {
		log.Fatal(message, err.Error())
	}
}
