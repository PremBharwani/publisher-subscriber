package scripts

import (
	"fmt"
	"log"
	"math/big"
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


	pubabistring:="[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pub_id\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stream_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"published_data\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pub_id\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stream_id\",\"type\":\"uint256\"}],\"name\":\"publisher_added\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pub_id\",\"type\":\"address\"}],\"name\":\"publisher_created\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pub_id\",\"type\":\"address\"}],\"name\":\"publisher_deleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pub_id\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stream_id\",\"type\":\"uint256\"}],\"name\":\"publisher_removed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stream_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"pub_id\",\"type\":\"address\"}],\"name\":\"add_publisher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_address_publisher\",\"type\":\"address\"}],\"name\":\"create_publisher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_id\",\"type\":\"address\"}],\"name\":\"delete_publisher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_id\",\"type\":\"address\"}],\"name\":\"get_publisher\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stream_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"pub_id\",\"type\":\"address\"}],\"name\":\"pub_to_event\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"publisher\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"address_publisher\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stream_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"pub_id\",\"type\":\"address\"}],\"name\":\"remove_publisher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	pubAbi, err := abi.JSON(strings.NewReader(pubabistring))
    if err != nil {
        log.Fatal(err)
    }



	for _, vLog := range logs {
	
		switch vLog.Topics[0] {
			
			case hash_pub_created:
				event := make(map[string]interface{})
				err := pubAbi.UnpackIntoMap(event, "publisher_created", vLog.Data)
				if err != nil {
					fmt.Printf("Error In Unpacking \n")
				}
				// pub_id := event["pub_id"].(common.Address)
			
			case hash_pub_added:
				// fmt.Printf("xx")
				event := make(map[string]interface{})
				err := pubAbi.UnpackIntoMap(event, "publisher_added", vLog.Data)
				if err != nil {
					fmt.Printf("Error In Unpacking \n")
				}
				// pub_id := event["pub_id"].(common.Address)
				// stream_id := event["stream_id"].(*big.Int)
				// fmt.Printf("Publisher Added : %s \n", pub_id.Hex())

			case hash_pub_removed:
				event := make(map[string]interface{})
				err := pubAbi.UnpackIntoMap(event, "publisher_removed", vLog.Data)
				if err != nil {
					fmt.Printf("Error In Unpacking \n")
				}
				// pub_id := event["pub_id"].(common.Address)
				// stream_id := event["stream_id"].(*big.Int)
				// fmt.Printf("Publisher Removed : %s \n", pub_id.Hex())
				
			
			case hash_pub_deleted:
				event := make(map[string]interface{})
				err := pubAbi.UnpackIntoMap(event, "publisher_deleted", vLog.Data)
				if err != nil {
					fmt.Printf("Error In Unpacking \n")
				}
				// pub_id := event["pub_id"].(common.Address)
				// fmt.Printf("Publisher Deleted : %s \n", pub_id.Hex())
			
			case hash_pub_data:
				fmt.Print("Publisher Data \n")
				event := make(map[string]interface{})
				err := pubAbi.UnpackIntoMap(event, "published_data", vLog.Data)
				if err != nil {
					fmt.Printf("Error In Unpacking \n")
				}
				// pub_id := event["pub_id"].(common.Address)
				// stream_id := event["stream_id"].(*big.Int)
				// data := event["data"].(string)
				// fmt.Printf("Publisher Data : %s \n", pub_id.Hex())
			default:
		}
		LastBlockId.Add(LastBlockId, big.NewInt(int64(1)))
	}

	return LastBlockId

}