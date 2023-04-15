package main

import (
	// "bytes"
	"context"
	"fmt"
	"log"
    // "strconv"
	"math/big"
	"strings"
    // "encoding/hex"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	// "github.com/ethereum/go-ethereum/core/types"
	// "golang.org/x/crypto/sha3"
	"github.com/miguelmota/go-solidity-sha3"

)

func convertUint64ToBigInt(num uint64) *big.Int {
    bigInt := new(big.Int)
    bigInt.SetUint64(num)
    return bigInt
}
func convertUint64ToBigInt1(num uint) *big.Int {
    bigInt := new(big.Int)
    bigInt.SetUint64(uint64(num))
    return bigInt
}
func publisher_created(pub_id common.Address){
	fmt.Printf("Publisher Created : %s \n", pub_id.Hex())
}
func publisher_added(pub_id common.Address, stream_id *big.Int){
	fmt.Printf("Publisher Added : %s %v \n", pub_id.Hex(), stream_id)
}
func publisher_removed(pub_id common.Address, stream_id *big.Int){
	fmt.Printf("Publisher Removed : %s %v \n", pub_id.Hex(), stream_id)
}
func publisher_deleted(pub_id common.Address){
	fmt.Printf("Publisher Deleted : %s \n", pub_id.Hex())
}
func published_data(pub_id common.Address, stream_id *big.Int, data string){
	fmt.Printf("Published Data : %s %v %s \n", pub_id.Hex(), stream_id, data)
}

func hashing(str string) common.Hash {
    hashBytes := solsha3.SoliditySHA3([]string{"string"}, []interface{}{str})
    var hashResult common.Hash
    copy(hashResult[:], hashBytes[:])
    return hashResult
}
// func hashing(str string) common.Hash {
// 	hash := sha3.NewLegacyKeccak256()
// 	hash.Write([]byte(str))
// 	hashBytes := hash.Sum(nil)
// 	var hashResult common.Hash
// 	copy(hashResult[:], hashBytes[:])
// 	return hashResult
// }

func main() {

	latestblock:= big.NewInt(163)
	hash_pub_created := hashing("publisher_created(address)")
	hash_pub_added := hashing("publisher_added(address,uint256)")
	hash_pub_removed := hashing("publisher_removed(address,uint256)")
	hash_pub_deleted := hashing("publisher_deleted(address)")
	hash_pub_data := hashing("published_data(address,uint256,string)")

	// fmt.Printf("Hashes : %s\n %s \n%s\n %s\n %s \n", hash_pub_created.Hex(), hash_pub_added.Hex(), hash_pub_removed.Hex(), hash_pub_deleted.Hex(), hash_pub_data.Hex())

	pubabistring:="[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pub_id\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stream_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"published_data\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pub_id\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stream_id\",\"type\":\"uint256\"}],\"name\":\"publisher_added\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pub_id\",\"type\":\"address\"}],\"name\":\"publisher_created\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pub_id\",\"type\":\"address\"}],\"name\":\"publisher_deleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pub_id\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stream_id\",\"type\":\"uint256\"}],\"name\":\"publisher_removed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stream_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"pub_id\",\"type\":\"address\"}],\"name\":\"add_publisher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_address_publisher\",\"type\":\"address\"}],\"name\":\"create_publisher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_id\",\"type\":\"address\"}],\"name\":\"delete_publisher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_id\",\"type\":\"address\"}],\"name\":\"get_publisher\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stream_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"pub_id\",\"type\":\"address\"}],\"name\":\"pub_to_event\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"publisher\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"address_publisher\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stream_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"pub_id\",\"type\":\"address\"}],\"name\":\"remove_publisher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	pubAbi, err := abi.JSON(strings.NewReader(pubabistring))
    if err != nil {
        log.Fatal(err)
    }


	for{

	client, err := ethclient.Dial("HTTP://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress("0xa14D52f7AE4855Fc1f4A61cDBe53f51405443616")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		FromBlock: latestblock,
	}
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		fmt.Printf("Error In Connecting 2 \n")
	}
	


	for _, vLog := range logs {
		// if(convertUint64ToBigInt(vLog.BlockNumber)==latestblock){
		// 	if(convertUint64ToBigInt1(vLog.Index)<latestlogindex){
		// 		continue
		// 	}
		// }
        switch vLog.Topics[0] {
			case hash_pub_created:
				event := make(map[string]interface{})
				err := pubAbi.UnpackIntoMap(event, "publisher_created", vLog.Data)
				if err != nil {
					fmt.Printf("Error In Unpacking \n")
				}
				pub_id := event["pub_id"].(common.Address)
				publisher_created(pub_id)
			case hash_pub_added:
				// fmt.Printf("xx")
				event := make(map[string]interface{})
				err := pubAbi.UnpackIntoMap(event, "publisher_added", vLog.Data)
				if err != nil {
					fmt.Printf("Error In Unpacking \n")
				}
				pub_id := event["pub_id"].(common.Address)
				stream_id := event["stream_id"].(*big.Int)
				// fmt.Printf("Publisher Added : %s \n", pub_id.Hex())
				publisher_added(pub_id, stream_id)

			case hash_pub_removed:
				event := make(map[string]interface{})
				err := pubAbi.UnpackIntoMap(event, "publisher_removed", vLog.Data)
				if err != nil {
					fmt.Printf("Error In Unpacking \n")
				}
				pub_id := event["pub_id"].(common.Address)
				stream_id := event["stream_id"].(*big.Int)
				// fmt.Printf("Publisher Removed : %s \n", pub_id.Hex())
				publisher_removed(pub_id, stream_id)
				
			
			case hash_pub_deleted:
				event := make(map[string]interface{})
				err := pubAbi.UnpackIntoMap(event, "publisher_deleted", vLog.Data)
				if err != nil {
					fmt.Printf("Error In Unpacking \n")
				}
				pub_id := event["pub_id"].(common.Address)
				// fmt.Printf("Publisher Deleted : %s \n", pub_id.Hex())
				publisher_deleted(pub_id)
			
			case hash_pub_data:
				fmt.Print("Publisher Data \n")
				event := make(map[string]interface{})
				err := pubAbi.UnpackIntoMap(event, "published_data", vLog.Data)
				if err != nil {
					fmt.Printf("Error In Unpacking \n")
				}
				pub_id := event["pub_id"].(common.Address)
				stream_id := event["stream_id"].(*big.Int)
				data := event["data"].(string)
				// fmt.Printf("Publisher Data : %s \n", pub_id.Hex())
				published_data(pub_id, stream_id, data)
			default:
				fmt.Printf("Error In Switch \n")
		}
		sum := big.NewInt(0)
		sum = sum.Add(latestblock,big.NewInt(1))
		latestblock = sum
		// latestlogindex = convertUint64ToBigInt1(vLog.Index)
    }
}
}