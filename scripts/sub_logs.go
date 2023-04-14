package main

import (
	// "bytes"
	"context"
	"fmt"
	"log"
	// "github.com/ethereum/go-ethereum/core/types"

	// "encoding/hex"
	"math/big"
	"strings"
	"golang.org/x/crypto/sha3"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)




type trial_emits3 struct {
	s3 string
	addr common.Address
}



func CalculateHash(str string) common.Hash {
	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(str))
	hashBytes := hash.Sum(nil)
	var hashResult common.Hash
	copy(hashResult[:], hashBytes[:])
	return hashResult
}



func main(){
	
	const abi_string = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"s1\",\"type\":\"string\"}],\"name\":\"trial_emits1\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"s2\",\"type\":\"string\"}],\"name\":\"trial_emits2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"s3\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"trial_emits3\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"s1\",\"type\":\"string\"}],\"name\":\"f1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"s2\",\"type\":\"string\"}],\"name\":\"f2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"s3\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"f3\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	contractABI, err := abi.JSON(strings.NewReader(abi_string))
	if err != nil {
		fmt.Printf("Error In Reading ABI \n")
	}
	
	LastBlockId := big.NewInt(0)

	
	for{
		client, err := ethclient.Dial("HTTP://127.0.0.1:7545")
		if err != nil {
			log.Fatal(err)
		}

		contractAddress := common.HexToAddress("0x4e8F6811b519dEAA45B031f145401D54109Fa6B1")

		query := ethereum.FilterQuery{
			Addresses: []common.Address{contractAddress},
			FromBlock: LastBlockId,		
		}

		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			log.Fatal(err)
		}
		

		for _, vLog := range logs {

			if vLog.Topics[0] == CalculateHash("trial_emits3(string,address)"){
				
				eventArgs := make(map[string]interface{})

				err1 := contractABI.UnpackIntoMap(eventArgs, "trial_emits3", vLog.Data)
				if err1 != nil {
					log.Fatal(err1)
				} else {
				fmt.Printf("%s\n" , eventArgs["s3"])
				}
			}
			LastBlockId.Add(LastBlockId, big.NewInt(int64(1)))			
		}
	}

}