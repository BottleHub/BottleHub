package client

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

// Connects to the EVM client
func ConnectClient() *ethclient.Client {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Client!")

	return client
}
