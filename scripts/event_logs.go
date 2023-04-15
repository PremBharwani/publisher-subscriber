package scripts

import (
	"fmt"
	"log"

	"math/big"
	"strings"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi"

)


func listen_event_logs(logs []types.Log, LastBlockId *big.Int) *big.Int {

	const abi_string_queue = "[{\"inputs\":[],\"name\":\"create_event_stream\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_event_id\",\"type\":\"uint256\"}],\"name\":\"delete_event_stream\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_event_id\",\"type\":\"uint256\"}],\"name\":\"get_all_queue_events\",\"outputs\":[{\"internalType\":\"bytes32[200]\",\"name\":\"\",\"type\":\"bytes32[200]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_event_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_last_seen_index\",\"type\":\"uint256\"}],\"name\":\"get_data_from_event_stream\",\"outputs\":[{\"internalType\":\"bytes32[200]\",\"name\":\"\",\"type\":\"bytes32[200]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_event_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_last_seen_index\",\"type\":\"uint256\"}],\"name\":\"get_next_event\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_next_hash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"max_queue_size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"num_queues\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_event_id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"publish_to_event_stream\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"queue_front_index\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"queue_next_index\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"queues\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"


	contractQueueABI, err := abi.JSON(strings.NewReader(abi_string_queue))
	if err != nil {
		fmt.Printf("Error In Reading Event Queue ABI \n")
	}

	for _, vLog := range logs {

		if vLog.Topics[0] == CalculateHash("topic_added(uint)") {

			eventArgs := make(map[string]interface{})

			err1 := contractQueueABI.UnpackIntoMap(eventArgs, "topic_added", vLog.Data)
			if err1 != nil {
				log.Fatal(err1)
			} else {
				// call api here
				// subscriber_id
			}
		}

		if vLog.Topics[0] == CalculateHash("topic_deleted(uint))") {

			eventArgs := make(map[string]interface{})

			err1 := contractQueueABI.UnpackIntoMap(eventArgs, "subscriber_created", vLog.Data)
			if err1 != nil {
				log.Fatal(err1)
			} else {
				// call api here
				// subscriber_id
			}
		}

	}

	return LastBlockId
}

