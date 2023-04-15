package scripts

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
	
	pubAddress := "0x123456789123456789123456789"
	subAddress := "0x123456789123456789123456789"
	eventAddress := "0x123456789123456789123456789"
	
	LastBlockId := big.NewInt(0)


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