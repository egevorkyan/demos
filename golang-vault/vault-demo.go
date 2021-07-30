package main

import (
	"fmt"
	vault "github.com/mch1307/vaultlib"
	"log"
)

func main() {
	//Reads environment variable, Like:
	//VAULT_ADDR and VAULT_TOKEN, fallbacks to default values if needed
	vConf := vault.NewConfig()

	//Create new client
	client, err := vault.NewClient(vConf)
	if err != nil {
		log.Fatal(err)
	}

	//Get data from Vault
	auth, err := client.GetSecret("demo1/auth")
	if err != nil {
		log.Fatal(err)
	}

	//Key:Value data reading
	for k, v := range auth.KV {
		fmt.Printf("Key: %v, Value: %v\n", k, v)
	}

	//Second method, assigning with variables preparing config
	vConfDemo := vault.NewConfig()
	vConfDemo.Address = "http://localhost:8200"
	vConfDemo.Token = "s.u6ZGJSYVR7r6Z7P8Tqgiv8E4"

	//Get data from vault using user demo
	vclient, err := vault.NewClient(vConfDemo)
	if err != nil {
		log.Fatal(err)
	}
	data, err := vclient.GetSecret("demo/vault")
	if err != nil {
		log.Fatal(err)
	}

	//Key:Value data reading
	for k, v := range data.KV {
		fmt.Printf("Key: %v, Value: %v\n", k, v)
	}
}
