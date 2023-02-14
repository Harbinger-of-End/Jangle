package pkg

import (
	"log"
	"os"

	"github.com/getsentry/sentry-go"
)

var logger *log.Logger = log.New(os.Stderr, "", log.LUTC)

func CheckError(err error) {
	if err != nil {
		logger.Println(err.Error())
		sentry.CaptureMessage(err.Error())
	}
}
