package main

import (
	"strconv"

	stackdriver "github.com/axiomzen/logrus-stackdriver-formatter"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.Formatter = stackdriver.NewFormatter(
		stackdriver.WithService("test-service"),
		stackdriver.WithVersion("v0.1.0"),
	)

	logger.Debug("starting application")
	logger.Info("application up and running")

	_, err := strconv.ParseInt("text", 10, 64)
	if err != nil {
		logger.WithError(err).Errorln("unable to parse integer")
	}

	logger.Debug("exiting")
}
