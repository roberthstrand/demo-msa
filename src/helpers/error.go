package helpers

import (
	"log"
	"os"
)

func ErrorCheck(e error) {
	if e != nil {
		log.Fatalln(e)
		os.Exit(1)
	}
}
