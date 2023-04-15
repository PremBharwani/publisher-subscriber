package main

import (
	"fmt"
	"log"
	"math/big"
	"time"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
)
func listen_pub_logs(logs []types.Log, LastBlockId *big.Int ) *big.Int{

	hash_pub_created := CalculateHash("publisher_created(address)")
	hash_pub_added := CalculateHash("publisher_added(address,uint256)")
	hash_pub_removed := CalculateHash("publisher_removed(address,uint256)")
	hash_pub_deleted := CalculateHash("publisher_deleted(address)")
	hash_pub_data := CalculateHash("published_data(address,uint256,string)")


	
	pubabistring := "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pub_id\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stream_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"published_data\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pub_id\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stream_id\",\"type\":\"uint256\"}],\"name\":\"publisher_added\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pub_id\",\"type\":\"address\"}],\"name\":\"publisher_created\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pub_id\",\"type\":\"address\"}],\"name\":\"publisher_deleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pub_id\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stream_id\",\"type\":\"uint256\"}],\"name\":\"publisher_removed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stream_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"pub_id\",\"type\":\"address\"}],\"name\":\"add_publisher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address_publisher\",\"type\":\"address\"}],\"name\":\"create_publisher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_id\",\"type\":\"address\"}],\"name\":\"delete_publisher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"event_publish_limit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pub_id\",\"type\":\"address\"}],\"name\":\"get_publisher\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stream_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"pub_id\",\"type\":\"address\"}],\"name\":\"publish_to_eventstream\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"publisher_list\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stream_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"pub_id\",\"type\":\"address\"}],\"name\":\"remove_publisher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"set_limit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	pubAbi, err := abi.JSON(strings.NewReader(pubabistring))
    if err != nil {
        log.Fatal(err)
    }



	for _, vLog := range logs {

		// eventArgs := make(map[string]interface{})
		// err := pubAbi.UnpackIntoMap(eventArgs, "publisher_created", vLog.Data)
		// if err!=nil{
		// 	fmt.Printf("Error In Unpacking \n")
		// }
	
		switch vLog.Topics[0] {
			
			case hash_pub_created:
				eventArgs := make(map[string]interface{})
				err := pubAbi.UnpackIntoMap(eventArgs, "publisher_created", vLog.Data)
				if err!=nil{
					fmt.Printf("Error In Unpacking \n")
				}
				// pub_id := eventArgs["pub_id"].(common.Address)
				val := make_dynamic_api_call("POST", "http://localhost:8080/create-user", fmt.Sprintf("{\"userWalletAddress\": \"%s\"}", eventArgs["pub_id"] ) )
				fmt.Printf("%s: %s\n", time.Now().Format("2006-01-02 15:04:05"), val)
			case hash_pub_added:
				eventArgs := make(map[string]interface{})
				err := pubAbi.UnpackIntoMap(eventArgs, "publisher_added", vLog.Data)
				if err!=nil{
					fmt.Printf("Error In Unpacking \n")
				}
				// fmt.Printf("pub_id: %s | stream_id: %s\n", eventArgs["pub_id"], eventArgs["stream_id"])
				val := make_dynamic_api_call("POST", "http://localhost:8080/add-user-access", fmt.Sprintf("{\"userWalletAddress\": \"%s\",\"eventQueueId\": \"%s\",\"action\": \"publish\"}", eventArgs["pub_id"], eventArgs["stream_id"]) )
				fmt.Printf("%s: %s\n", time.Now().Format("2006-01-02 15:04:05"), val)
				// pub_id := eventArgs["pub_id"].(common.Address)
				// stream_id := eventArgs["stream_id"].(*big.Int)
				// fmt.Printf("Publisher Added : %s \n", pub_id.Hex())

			case hash_pub_removed:
				eventArgs := make(map[string]interface{})
				err := pubAbi.UnpackIntoMap(eventArgs, "publisher_removed", vLog.Data)
				if err!=nil{
					fmt.Printf("Error In Unpacking \n")
				}
				val := make_dynamic_api_call("POST", "http://localhost:8080/remove-user-access", fmt.Sprintf("{\"userWalletAddress\": \"%s\",\"eventQueueId\": \"%s\",\"action\": \"publish\"}", eventArgs["pub_id"], eventArgs["stream_id"]) )
				fmt.Printf("%s: %s\n", time.Now().Format("2006-01-02 15:04:05"), val)

				// pub_id := eventArgs["pub_id"].(common.Address)
				// stream_id := eventArgs["stream_id"].(*big.Int)
				// fmt.Printf("Publisher Removed : %s \n", pub_id.Hex())
				
			
			case hash_pub_deleted:
				eventArgs := make(map[string]interface{})
				err := pubAbi.UnpackIntoMap(eventArgs, "publisher_deleted", vLog.Data)
				if err!=nil{
					fmt.Printf("Error In Unpacking \n")
				}
				val := make_dynamic_api_call("POST", "http://localhost:8080/remove-subscriber-access", fmt.Sprintf("{\"userWalletAddress\": \"%s\"}", eventArgs["pub_id"]) )
				fmt.Printf("%s: %s\n", time.Now().Format("2006-01-02 15:04:05"), val)
				// pub_id := eventArgs["pub_id"].(common.Address)
				// fmt.Printf("Publisher Deleted : %s \n", pub_id.Hex())
			
			case hash_pub_data:
				eventArgs := make(map[string]interface{})
				err := pubAbi.UnpackIntoMap(eventArgs, "published_data", vLog.Data)
				if err!=nil{
					fmt.Printf("Error In Unpacking \n")
				}
				val := make_dynamic_api_call("POST", "http://localhost:8080/publish-event", fmt.Sprintf("{\"userWalletAddress\": \"%s\",\"eventQueueId\": \"%s\",\"message\": \"%s\"}", eventArgs["pub_id"], eventArgs["stream_id"], eventArgs["data"]) )
				fmt.Printf("%s: %s\n", time.Now().Format("2006-01-02 15:04:05"), val)
				// pub_id := eventArgs["pub_id"].(common.Address)
				// stream_id := eventArgs["stream_id"].(*big.Int)
				// data := eventArgs["data"].(string)
				// fmt.Printf("Publisher Data : %s \n", pub_id.Hex())
			default:
		}
		LastBlockId.Add(LastBlockId, big.NewInt(int64(1)))
	}

	return LastBlockId

}