package main

import (
	"fmt"
	"html"
	"net/http"
)

func main() {
	// Root handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Log the call
		newLogMsg(fmt.Sprintf("Hey from Root, %q", html.EscapeString(r.URL.Path)))
	})

	// Feature handler
	http.HandleFunc("/feature", func(w http.ResponseWriter, r *http.Request) {
		// 1. Log the call
		newLogMsg(fmt.Sprintf("Hey from /feature, %q", html.EscapeString(r.URL.Path)))

		// 2. Get chart name from etcd
		chart, err := etcdGet("my-helm-chart")
		if err != nil {
			newLogMsg(fmt.Sprintf("failed to create etcd client: %v", err))
		}

		// 3. Get releases for this chart
		_, err = getHelmChartReleases(chart)
		if err != nil {
			newLogMsg(fmt.Sprintf("failed to get helm chart releases: %v", err))
		}

		// ... do something else
	})

	// Start the server
	newLogMsg(http.ListenAndServe(":8080", nil).Error())
}
