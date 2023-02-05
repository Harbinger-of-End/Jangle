package pkg

import (
	"log"
	"os"
)

var logger *log.Logger = log.New(os.Stderr, "", log.LUTC)

func CheckError(err error) {
	if err != nil {
		logger.Fatal(err.Error())

	}
}
