package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	client, err := ethclient.Dial(os.Getenv("URL"))
	if err != nil {
		log.Fatal(err)
	}

	// Get the latest block number
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	latestBlockNumber := header.Number.Uint64()

	// Get the block with the given block number
	block, err := client.BlockByNumber(context.Background(), new(big.Int).SetUint64(latestBlockNumber))
	if err != nil {
		log.Fatal(err)
	}

	// Iterate over the transactions in the block
	for _, tx := range block.Transactions() {
		// Print the transaction hash and the transaction value
		fmt.Println(tx.Hash().Hex(), tx.Value().String())
	}
}
