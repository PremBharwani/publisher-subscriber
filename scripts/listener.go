package main

import (
	"context"
	"log"
	"math/big"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)


func main() {
	
	pubAddress := "0xACfdB16DBdD5c43d97cDCC3Fdb6eAC1620Ae98aD"
	subAddress := "0x53Aebfbae8532B6E6568349fe4EAD593a543fab9"
	eventAddress := "0xf89C54DffF0e777b936ba82e7F7b3c3eBcd1292c"
	
	// LastBlockId := big.NewInt(0)

	// Set the LastBlockId to the latest block
	client, err := ethclient.Dial("HTTP://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}
	latestBlock, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	LastBlockId := latestBlock.Number()

	for {

		client, err := ethclient.Dial("HTTP://127.0.0.1:7545")
		if err != nil {
			log.Fatal(err)
		}
		query := ethereum.FilterQuery{
			Addresses: []common.Address{
				common.HexToAddress(pubAddress),
				common.HexToAddress(subAddress),
				common.HexToAddress(eventAddress),
			},
			FromBlock: LastBlockId,
		}
		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			fmt.Printf("Error In Filtering Logs \n")
		}
		


		LastBlockId = listen_pub_logs(logs, LastBlockId) 
		LastBlockId = listen_event_logs(logs, LastBlockId) 
		LastBlockId = listen_sub_logs(logs, LastBlockId)


	}
	


}