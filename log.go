package main

import (
	"github.com/go-kit/log"
	"os"
)

// Low-score example - logging - github.com/go-kit/log
func newLogMsg(msg string) {
	// Create a new logger
	logger := log.NewLogfmtLogger(os.Stdout)

	// Log the message
	err := logger.Log(msg)
	if err != nil {
		os.Exit(1)
	}
}

//// High-score example - logging - github.com/sirupsen/logrus
//func newLogMsg(msg string) {
//	// Create a new logger
//	logger := logrus.New()
//
//	// Log the message
//	logger.Log(logrus.InfoLevel, msg)
//}
