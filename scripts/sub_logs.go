package main

import (
	// "bytes"
	"context"
	"fmt"
	"log"

	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type subscriber_limit_set struct {
	limit uint
}

type subscriber_created struct {
	SubscriberAddress common.Address
}

type subscriber_removed struct {
	SubscriberAddress common.Address
}

type subscribed_to_event struct {
	SubscriberAddress common.Address
	EventStreamId     common.Address
}

type unsubscribe_to_event struct {
	SubscriberAddress common.Address
	EventStreamId     common.Address
}

func main() {

	client, err := ethclient.Dial("HTTP://127.0.0.1:7545")
	if err != nil {
		log.Printf("Error in connecting to client \n")
	}

	const abi_string = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"subscriber_id\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"event_stream_id\",\"type\":\"address\"}],\"name\":\"subscribed_to_event\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"subscriber_id\",\"type\":\"address\"}],\"name\":\"subscriber_created\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"subscriber_limit_set\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"subscriber_id\",\"type\":\"address\"}],\"name\":\"subscriber_removed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"subscriber_id\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"event_stream_id\",\"type\":\"address\"}],\"name\":\"unsubscribed_to_event\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subscriber_id\",\"type\":\"address\"}],\"name\":\"create_subscriber\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"subscriber_id\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"event_streams_subscribed\",\"type\":\"address[]\"}],\"internalType\":\"struct Sub.Subscriber\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"subscriber_id\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"event_streams_subscribed\",\"type\":\"address[]\"}],\"internalType\":\"struct Sub.Subscriber\",\"name\":\"s\",\"type\":\"tuple\"}],\"name\":\"delete_subscriber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"event_subscribe_limit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"set_limit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"event_stream_id\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"subscriber_id\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"event_streams_subscribed\",\"type\":\"address[]\"}],\"internalType\":\"struct Sub.Subscriber\",\"name\":\"s\",\"type\":\"tuple\"}],\"name\":\"unsubscribe_to_event\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

	contractABI, err := abi.JSON(strings.NewReader(abi_string))
	if err != nil {
		fmt.Printf("Error In Reading ABI \n")
	}

	// eventSignature := []byte("subscriber_created(address)")
	contractAddress := common.HexToAddress("0xAcaf4C3a4e50a3ef7FCCfA4dA0EcC01CF1bF0BC5")
	LastBlockId := big.NewInt(22)
	i := 1

	for i > 0 {

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		FromBlock: LastBlockId,
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		fmt.Printf("Error In Getting Logs")
	}


	for _, vLog := range logs {

		var event1 subscriber_limit_set

		err1 := contractABI.UnpackIntoInterface(&event1, "subscriber_limit_set", vLog.Data)
		if err1 == nil {
			fmt.Printf("%v", event1.limit)
		}

		var event2 subscriber_created

		err2 := contractABI.UnpackIntoInterface(&event2, "subscriber_created", vLog.Data)
		if err2 == nil {
			fmt.Printf(event2.SubscriberAddress.Hex())
			fmt.Printf("\n")
		}

		var event3 subscriber_removed

		err3 := contractABI.UnpackIntoInterface(&event3, "subscriber_removed", vLog.Data)
		if err3 == nil {
			fmt.Printf("%v", event3.SubscriberAddress)
			fmt.Printf("\n")
		}

		var event4 subscribed_to_event

		err4 := contractABI.UnpackIntoInterface(&event4, "subscribed_to_event", vLog.Data)
		if err4 == nil {
			fmt.Printf("%v", event4.SubscriberAddress)
			fmt.Printf("%v", event4.EventStreamId)
			fmt.Printf("\n")
		}

		var event5 unsubscribe_to_event

		err5 := contractABI.UnpackIntoInterface(&event5, "unsubscribe_to_event", vLog.Data)
		if err5 == nil {
			fmt.Printf("%v", event5.SubscriberAddress)
			fmt.Printf("%v", event5.EventStreamId)
			fmt.Printf("\n")
		}

		// fmt.Printf("%v\n", vLog)

		// var event6 string
		// err6 := contractABI.UnpackIntoInterface(&event6, "emit_a_string", vLog.Data)
		// if err6 == nil {
		// 	// log.Fatal(err6)
		// }
		// fmt.Printf(event6)
		// fmt.Printf("\n")

		//   address := common.BytesToAddress(event.subscriber_id)
		//   fmt.Println(string(event.Value[:])) // bar

	}

	LastBlockId.Add(LastBlockId, big.NewInt(int64(1)))

	}
}
