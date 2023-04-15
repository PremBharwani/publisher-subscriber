package main

import (
	"fmt"
	"log"
	"time"
	"encoding/json"
	// "net/http"
	// "io/ioutil"
	// "bytes"
	"math/big"
	"strings"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi"


)

type apiRespStruct struct {
	Message []string `json:"message"`
}


var greetM bool = false
func listen_sub_logs(logs []types.Log, LastBlockId *big.Int) *big.Int {

	if !greetM {fmt.Println("Listening to Sub Logs"); greetM = true}

	var abi_string_sub = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"subscriber_id\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"event_stream_id\",\"type\":\"uint256\"}],\"name\":\"requested_for_events\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"subscriber_id\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"event_stream_id\",\"type\":\"uint256\"}],\"name\":\"subscribed_to_event\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"subscriber_id\",\"type\":\"address\"}],\"name\":\"subscriber_created\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"subscriber_id\",\"type\":\"address\"}],\"name\":\"subscriber_removed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"subscriber_id\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"event_stream_id\",\"type\":\"uint256\"}],\"name\":\"unsubscribed_to_event\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stream_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sub_id\",\"type\":\"address\"}],\"name\":\"call_for_events\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkEvent\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address_subscriber\",\"type\":\"address\"}],\"name\":\"create_subscriber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_id\",\"type\":\"address\"}],\"name\":\"delete_subscriber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"event_subscribe_limit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sub_id\",\"type\":\"address\"}],\"name\":\"get_back_events\",\"outputs\":[{\"components\":[{\"internalType\":\"string[50]\",\"name\":\"events\",\"type\":\"string[50]\"},{\"internalType\":\"uint8\",\"name\":\"last_index\",\"type\":\"uint8\"}],\"internalType\":\"struct Sub.events_data\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"events\",\"type\":\"string[]\"},{\"internalType\":\"address\",\"name\":\"sub_id\",\"type\":\"address\"}],\"name\":\"relay_events\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relay_events_called\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"set_limit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stream_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sub_id\",\"type\":\"address\"}],\"name\":\"subscribe_to_event\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"subscriber_list\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stream_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sub_id\",\"type\":\"address\"}],\"name\":\"unsubscribe_to_event\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"


	// fmt.Println("ABI String : ", abi_string_sub)

	contractSubABI, err := abi.JSON(strings.NewReader(abi_string_sub))
	if err != nil {
		fmt.Printf("Error In Reading Subscriber ABI \n")
	}


	

	for _, vLog := range logs {
		// eventArgs := make(map[string]interface{})
		// err1 := contractSubABI.UnpackIntoMap(eventArgs, "subscriber_created", vLog.Data)
		// if err1 != nil {
		// 	log.Fatal(err1)
		// } 
		// fmt.Printf("Log : %s", vLog.Topics[0])
		switch vLog.Topics[0] {

			case CalculateHash("subscriber_created(address)") :
				eventArgs := make(map[string]interface{})
				err1 := contractSubABI.UnpackIntoMap(eventArgs, "subscriber_created", vLog.Data)
				if err1 != nil {
					log.Fatal(err1)
				} 
				val := make_dynamic_api_call("POST", "http://localhost:8080/create-user", fmt.Sprintf("{\"userWalletAddress\": \"%s\"}", eventArgs["subscriber_id"]) )
				fmt.Printf("%s: %s\n", time.Now().Format("2006-01-02 15:04:05"), val)

			case CalculateHash("subscriber_removed(address)") :
				eventArgs := make(map[string]interface{})
				err1 := contractSubABI.UnpackIntoMap(eventArgs, "subscriber_removed", vLog.Data)
				if err1 != nil {
					log.Fatal(err1)
				} 
				val := make_dynamic_api_call("POST", "http://localhost:8080/remove-subscriber-access", fmt.Sprintf("{\"userWalletAddress\": \"%s\"}", eventArgs["subscriber_id"]) )
				fmt.Printf("%s: %s\n", time.Now().Format("2006-01-02 15:04:05"), val)
				
			case CalculateHash("subscribed_to_event(address,uint256)") :
				eventArgs := make(map[string]interface{})
				err1 := contractSubABI.UnpackIntoMap(eventArgs, "subscribed_to_event", vLog.Data)
				if err1 != nil {
					log.Fatal(err1)
				} 
				val := make_dynamic_api_call("POST", "http://localhost:8080/add-user-access", fmt.Sprintf("{\"userWalletAddress\": \"%s\",\"eventQueueId\": \"%s\", \"action\": \"subscribe\"}", eventArgs["subscriber_id"], eventArgs["event_stream_id"]) )
				fmt.Printf("%s: %s\n", time.Now().Format("2006-01-02 15:04:05"), val)
				// call api here
				// subscriber_id
				// event_stream_id

			case CalculateHash("unsubscribed_to_event(address,uint256)") :
				eventArgs := make(map[string]interface{})
				err1 := contractSubABI.UnpackIntoMap(eventArgs, "unsubscribed_to_event", vLog.Data)
				if err1 != nil {
					log.Fatal(err1)
				} 
				val := make_dynamic_api_call("DELETE", "http://localhost:8080/remove-user-access", fmt.Sprintf("{\"userWalletAddress\": \"%s\",\"eventQueueId\": \"%s\", \"action\": \"subscribe\"}", eventArgs["subscriber_id"], eventArgs["event_stream_id"]) )
				fmt.Printf("%s: %s\n", time.Now().Format("2006-01-02 15:04:05"), val)

			case CalculateHash("requested_for_events(address,uint256)") :
				eventArgs := make(map[string]interface{})
				err1 := contractSubABI.UnpackIntoMap(eventArgs, "requested_for_events", vLog.Data)
				if err1 != nil {
					log.Fatal(err1)
				} 
				eventMessages := make_dynamic_api_call("GET", "http://localhost:8080/consume-event", 
					fmt.Sprintf("{\"userWalletAddress\":\"%s\",\"eventQueueId\": \"%s\"}", eventArgs["subscriber_id"], eventArgs["event_stream_id"]) )
				// fmt.Printf("Messages recieved on the listener server : %s\n", eventMessages)
				// var1 : = make_api_call("http://localhost:8080/create-user", fmt.Sprintf("{\"userWalletAddress\": \"%s\"}", eventArgs["subscriber_id"]) )
				// subscriber_id
				// event_stream_id

				var jsonMessage apiRespStruct
				json.Unmarshal([]byte(eventMessages), &jsonMessage)
				messageEventsArr := jsonMessage.Message
				fmt.Printf("over hrere %s\n", messageEventsArr[0])
				
				//Iterate over messageEventsArr and make a json string
				var jsonString string
				jsonString += "["
				for i := 0; i < len(messageEventsArr); i++ {
					jsonString += fmt.Sprintf("\"%s\"", messageEventsArr[i])
					if i!= len(messageEventsArr)-1 {
						jsonString += ", "
					}
				}
				//Remove last comma from the json string

				jsonString += "]"
				// fmt.Printf("check this out %s\n", jsonString)
				// dummyJsonString := fmt.Sprintf("{\"type\":\"relay_events\",\"sub_id\": \"%s\",\"data\": [\"1\", \"2\"]}", eventArgs["subscriber_id"])
				
				fmt.Printf("Messages recieved on the listener server : %s\n", messageEventsArr)
				// fmt.Printf("JSon obj : %s\n", dummyJsonString)//, messageEventsArr))
				apiResp := make_dynamic_api_call("POST", "http://localhost:3000/send-events", 
					fmt.Sprintf("{\"type\":\"relay_events\",\"sub_id\": \"%s\",\"data\": %s}", eventArgs["subscriber_id"], jsonString) )
				if(apiResp!=""){
					fmt.Printf("API Response to send it to the sol: %s", apiResp)
				}
				fmt.Printf("Messages recieved on the listener server : %s\n", eventMessages)
			default :
				continue
		}
	}

	return LastBlockId
}

