package main

import (
	"fmt"
	"log"
	"time"
	"math/big"
	"strings"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi"

)


func listen_event_logs(logs []types.Log, LastBlockId *big.Int) *big.Int {

	const abi_string_queue = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stream_id\",\"type\":\"uint256\"}],\"name\":\"topic_added\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stream_id\",\"type\":\"uint256\"}],\"name\":\"topic_deleted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"add_topic\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_topic_id\",\"type\":\"uint256\"}],\"name\":\"delete_topic\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"num_queues\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"


	contractQueueABI, err := abi.JSON(strings.NewReader(abi_string_queue))
	if err != nil {
		fmt.Printf("Error In Reading Event Queue ABI \n")
	}


	for _, vLog := range logs {

		eventArgs := make(map[string]interface{})

		err1 := contractQueueABI.UnpackIntoMap(eventArgs, "topic_added", vLog.Data)
		if err1 != nil {
			log.Fatal(err1)
		}

		if vLog.Topics[0] == CalculateHash("topic_added(uint256)") {
			val := make_dynamic_api_call("POST", "http://localhost:8080/create-topic", fmt.Sprintf("{\"userWalletAddress\": \"%s\", \"eventQueueId\":\"%s\"}", "lmao", eventArgs["stream_id"]))
			fmt.Printf("%s: %s\n", time.Now().Format("2006-01-02 15:04:05"), val)
			
		} else if vLog.Topics[0] == CalculateHash("topic_deleted(uint256))") {
			val := make_dynamic_api_call("POST", "http://localhost:8080/delete-topic", fmt.Sprintf("{\"userWalletAddress\": \"%s\", \"eventQueueId\":\"%s\"}", "lmao", eventArgs["stream_id"]))
			fmt.Printf("%s: %s\n", time.Now().Format("2006-01-02 15:04:05"), val)
		}

	}

	return LastBlockId
}

