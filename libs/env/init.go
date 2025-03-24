package env

import (
	"fmt"
	"log"

	"github.com/hashicorp/vault/api"
)

const vaultAddr = "http://127.0.0.1:8200"
const vaultToken = "root"

func CreateClient() *api.Client {
	config := &api.Config{
		Address: vaultAddr,
	}

	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("Error creating Vault client: %v", err)
	}

	client.SetToken(vaultToken)
	status, err := client.Sys().Health()
	if err != nil {
		log.Fatalf("Error checking Vault health: %v", err)
	}
	fmt.Println("Vault Health Status:", status.Initialized)
	return client
}
