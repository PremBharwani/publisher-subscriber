// util methods to make api calls
package main

import (
	"net/http"
	"strings"
	"io/ioutil"
)

func make_dynamic_api_call(actionString string, url string, inputJsonBody string) string{
	//? Check if you want to give back the status code?
	req, err := http.NewRequest(actionString, url, strings.NewReader(inputJsonBody) )
	if err != nil {
		// handle error
	}
	req.Header.Set("User-Agent", "MyUserAgent/1.0")

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