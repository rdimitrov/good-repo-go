package main

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"os"
)

// Example of creating a Hashicorp Vault client and reading a secret
func getEtcdEndpointFromVault() (*string, error) {
	var mySecretValue string
	// Set up the Vault client
	vaultAddr := os.Getenv("VAULT_ADDR") // e.g., "http://127.0.0.1:8200"
	if vaultAddr == "" {
		err := fmt.Errorf("VAULT_ADDR environment variable is not set")
		newLogMsg(err.Error())
		return nil, err
	}

	// Set up the Vault configuration
	config := &api.Config{
		Address: vaultAddr,
	}

	// Create a new Vault client
	client, err := api.NewClient(config)
	if err != nil {
		newLogMsg(fmt.Sprintf("failed to create Vault client: %v", err))
		return nil, err
	}

	// Set the authentication token
	vaultToken := os.Getenv("VAULT_TOKEN")
	if vaultToken == "" {
		err := fmt.Errorf("VAULT_TOKEN environment variable is not set")
		newLogMsg(err.Error())
		return nil, err
	}
	client.SetToken(vaultToken)

	// Read the secret
	etcdSecret, err := client.Logical().Read("secret/data/my-etcd")
	if err != nil {
		newLogMsg(fmt.Sprintf("failed to read secret: %v", err))
	}

	// Extract the desired data from the secret
	if etcdSecret != nil && etcdSecret.Data != nil {
		data, ok := etcdSecret.Data["data"].(map[string]interface{})
		if !ok {
			newLogMsg("data type assertion failed")
		}

		// Get the value
		mySecretValue, ok = data["etcd-host-urls"].(string)
		if !ok {
			newLogMsg("etcd-host-urls type assertion failed or not found")
		}
	} else {
		newLogMsg("No data found at the specified path")
	}

	// All good
	return &mySecretValue, nil
}
