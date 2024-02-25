package main

import (
	"fmt"
	"github.com/go-kit/log"
	uuid "github.com/satori/go.uuid"
	"html"
	stdlog "log"
	"net/http"
	"os"
)

func main() {
	// Hey handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logMessage(fmt.Sprintf("Hey, %q", html.EscapeString(r.URL.Path)))
		logMessage(fmt.Sprintf("request_id:, %s", uuid.NewV4()))
	})

	// Bye handler
	http.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		logMessage(fmt.Sprintf("Bye, %q", html.EscapeString(r.URL.Path)))
		logMessage(fmt.Sprintf("request_id:, %s", uuid.NewV4()))
	})

	// Listen and serve
	stdlog.Fatal(http.ListenAndServe(":8080", nil))
}

// logMessage logs a message
func logMessage(msg string) {
	logger := log.NewLogfmtLogger(os.Stdout)
	logger.Log(msg)
}
