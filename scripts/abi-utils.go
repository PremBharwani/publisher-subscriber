package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Function to extract the key value from a json file
func extractKeyValue(jsonFilePath string, key string) string {
	// Open our jsonFile using jsonFilePath which is relative path
	jsonFile, err := os.Open(jsonFilePath)

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened " + jsonFilePath)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var result map[string]interface{}

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &result)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	fmt.Printf("Converted to value : %s\n", result[key].(string))
	return result[key].(string)
}