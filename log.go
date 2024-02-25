package main

import (
	"github.com/go-kit/log"
	"os"
)

// Trusty logging examples
// github.com/go-kit/log // lower scored in Trusty
// github.com/sirupsen/logrus // higher scored in Trusty

// logMessage logs a message
func newLogMsg(msg string) {
	// Create a new logger
	logger := log.NewLogfmtLogger(os.Stdout)

	// Log the message
	err := logger.Log(msg)
	if err != nil {
		os.Exit(1)
	}
}
