package utils

import (
	"log"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalln(err, msg)
	}
}