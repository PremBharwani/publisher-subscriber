package main

import (
	"context"
	"fmt"
	"log"

	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	// "github.com/ethereum/go-ethereum/event"
	"golang.org/x/crypto/sha3"
)

func CalculateHash(str string) common.Hash {
	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(str))
	hashBytes := hash.Sum(nil)
	var hashResult common.Hash
	copy(hashResult[:], hashBytes[:])
	return hashResult
}

func main() {

	const abi_string_sub = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"subscriber_id\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"event_stream_id\",\"type\":\"uint256\"}],\"name\":\"requested_for_events\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"subscriber_id\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"event_stream_id\",\"type\":\"address\"}],\"name\":\"subscribed_to_event\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"subscriber_id\",\"type\":\"address\"}],\"name\":\"subscriber_created\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"subscriber_limit_set\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"subscriber_id\",\"type\":\"address\"}],\"name\":\"subscriber_removed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"subscriber_id\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"event_stream_id\",\"type\":\"address\"}],\"name\":\"unsubscribed_to_event\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subscriber_id\",\"type\":\"address\"}],\"name\":\"create_subscriber\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"subscriber_id\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"event_streams_subscribed\",\"type\":\"address[]\"}],\"internalType\":\"struct Sub.Subscriber\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"}],\"name\":\"delete_subscriber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"event_subscribe_limit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stream_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sub_id\",\"type\":\"address\"}],\"name\":\"get_events\",\"outputs\":[{\"components\":[{\"internalType\":\"string[100]\",\"name\":\"events\",\"type\":\"string[100]\"},{\"internalType\":\"uint256\",\"name\":\"last_index\",\"type\":\"uint256\"}],\"internalType\":\"struct Sub.events_data\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"events\",\"type\":\"string[]\"}],\"name\":\"relay_events\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"set_limit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"event_stream_id\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"}],\"name\":\"subscribe_to_event\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"event_stream_id\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"}],\"name\":\"unsubscribe_to_event\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	const abi_string_queue = "[{\"inputs\":[],\"name\":\"create_event_stream\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_event_id\",\"type\":\"uint256\"}],\"name\":\"delete_event_stream\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_event_id\",\"type\":\"uint256\"}],\"name\":\"get_all_queue_events\",\"outputs\":[{\"internalType\":\"bytes32[200]\",\"name\":\"\",\"type\":\"bytes32[200]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_event_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_last_seen_index\",\"type\":\"uint256\"}],\"name\":\"get_data_from_event_stream\",\"outputs\":[{\"internalType\":\"bytes32[200]\",\"name\":\"\",\"type\":\"bytes32[200]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_event_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_last_seen_index\",\"type\":\"uint256\"}],\"name\":\"get_next_event\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_next_hash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"max_queue_size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"num_queues\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_event_id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"publish_to_event_stream\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"queue_front_index\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"queue_next_index\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"queues\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

	contractSubABI, err := abi.JSON(strings.NewReader(abi_string_sub))
	if err != nil {
		fmt.Printf("Error In Reading Subscriber ABI \n")
	}

	// contractQueueABI, err := abi.JSON(strings.NewReader(abi_string_queue))
	// if err != nil {
	// 	fmt.Printf("Error In Reading Event Queue ABI \n")
	// }

	sub_contract_address := common.HexToAddress("0xF0Dd6d22b59DB5325Dae833d87b1a683b27D702D")
	// event_contract_address := "0x0f10024fb7415240861C78535166Df909272A37C"

	LastBlockId := big.NewInt(0)

	for {

		client, err := ethclient.Dial("HTTP://127.0.0.1:7545")
		if err != nil {
			log.Fatal(err)
		}

		query := ethereum.FilterQuery{
			Addresses: []common.Address{sub_contract_address},
			// {
			// common.HexToAddress(sub_contract_address),
			// common.HexToAddress(event_contract_address),
			// },
			FromBlock: LastBlockId,
		}

		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			log.Fatal(err)
		}

		for _, vLog := range logs {

			if vLog.Topics[0] == CalculateHash("subscriber_created(address)") {

				eventArgs := make(map[string]interface{})

				err1 := contractSubABI.UnpackIntoMap(eventArgs, "subscriber_created", vLog.Data)
				if err1 != nil {
					log.Fatal(err1)
				} else {
					// call api here
					// subscriber_id
				}
			}

			if vLog.Topics[0] == CalculateHash("subscriber_created(address)") {

				eventArgs := make(map[string]interface{})

				err1 := contractSubABI.UnpackIntoMap(eventArgs, "subscriber_created", vLog.Data)
				if err1 != nil {
					log.Fatal(err1)
				} else {
					// call api here
					// subscriber_id
				}
			}

			if vLog.Topics[0] == CalculateHash("subscribed_to_event(address,address)") {

				eventArgs := make(map[string]interface{})

				err1 := contractSubABI.UnpackIntoMap(eventArgs, "subscribed_to_event", vLog.Data)
				if err1 != nil {
					log.Fatal(err1)
				} else {
					// call api here
					// subscriber_id
					// event_stream_id
					fmt.Printf("its working")
				}
			}

			if vLog.Topics[0] == CalculateHash("unsubscribed_to_event(address,address)") {

				eventArgs := make(map[string]interface{})

				err1 := contractSubABI.UnpackIntoMap(eventArgs, "unsubscribed_to_event", vLog.Data)
				if err1 != nil {
					log.Fatal(err1)
				} else {
					// call api here
					// subscriber_id
					// event_stream_id
				}
			}

			if vLog.Topics[0] == CalculateHash("requested_for_events(address,address)") {

				eventArgs := make(map[string]interface{})

				err1 := contractSubABI.UnpackIntoMap(eventArgs, "requested_for_events", vLog.Data)
				if err1 != nil {
					log.Fatal(err1)
				} else {
					// call api here
					// subscriber_id
					// event_stream_id
				}
			}

			LastBlockId.Add(LastBlockId, big.NewInt(int64(1)))
		}
	}

}
