package main

import (
	"fmt"
	"html"
	"net/http"
)

func main() {
	// 1. Create my root handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Log the call
		newLogMsg(fmt.Sprintf("Hey from Root, %q", html.EscapeString(r.URL.Path)))
	})

	// 2. Create my feature handler
	http.HandleFunc("/feature", func(w http.ResponseWriter, r *http.Request) {
		// Log the call
		newLogMsg(fmt.Sprintf("Hey from /feature, %q", html.EscapeString(r.URL.Path)))

		// Get the etcd endpoint from Vault
		etcdURL, err := getEtcdEndpointFromVault()
		if err != nil {
			newLogMsg(fmt.Sprintf("failed to get etcd endpoint from Vault: %v", err))
		}

		// Create an etcd client
		client, err := newEtcdClient(etcdURL)
		if err != nil {
			newLogMsg(fmt.Sprintf("failed to create etcd client: %v", err))
		}

		// Get a key
		_, _ = client.Get("key", false, false)

		// ... do something else
	})

	// 3. Start the server
	newLogMsg(http.ListenAndServe(":8080", nil).Error())
}
