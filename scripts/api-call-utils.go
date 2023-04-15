// util methods to make api calls
package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"fmt"
)

func make_dynamic_api_call(actionString string, url string, inputJsonBody string) string{
	//? Check if you want to give back the status code?
	req, err := http.NewRequest(actionString, url, strings.NewReader(inputJsonBody) )
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "MyUserAgent/1.0")
	req.Header.Set("Content-Type", "application/json")
	
	// Print out the req to verify 
	fmt.Println("=====================================")
	fmt.Println("Request inside the make-dyn-api-call ", req)

	
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return (string(body))
}