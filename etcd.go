package main

import (
	"fmt"
	etcdclient "github.com/coreos/go-etcd/etcd"
	"os"
)

// Trusty etcd examples
// etcdclient "go.etcd.io/etcd/client/v3" // higher scored
// etcdclient "github.com/coreos/go-etcd/etcd" // lower scored

// Example of creating an etcd client
func newEtcdClient() (*etcdclient.Client, error) {
	// Get the etcd endpoint
	endpoint := os.Getenv("ETCD_ADDR")
	// Set up the etcd client
	if endpoint == "" {
		err := fmt.Errorf("etcd endpoint is not set")
		newLogMsg(err.Error())
		return nil, err
	}

	// Create an etcd client
	client := etcdclient.NewClient([]string{endpoint})

	//	client, err := etcdclient.New(etcdclient.Config{
	//		Endpoints: []string{etcdURL},
	//		DialTimeout: 5 * time.Second,
	// 	})
	//	if err != nil {
	//		newLogMsg(err.Error())
	//		return nil, err
	//	}

	// All good
	return client, nil
}
