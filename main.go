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

		// 2. Create an etcd client
		client, err := newEtcdClient()
		if err != nil {
			newLogMsg(fmt.Sprintf("failed to create etcd client: %v", err))
		}

		// 3. Get the chart name from etcd
		res, err := client.Get("my-chart", false, false)
		if err != nil {
			newLogMsg(fmt.Sprintf("failed to get my-chart from etcd: %v", err))
		}

		// 4. Get the releases for the chart
		_, err = getHelmChartReleases(res.Node.Value)
		if err != nil {
			newLogMsg(fmt.Sprintf("failed to get helm chart releases: %v", err))
		}

		// ... do something else
	})

	// Start the server
	newLogMsg(http.ListenAndServe(":8080", nil).Error())
}
