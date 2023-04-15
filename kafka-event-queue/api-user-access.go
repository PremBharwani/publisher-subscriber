// ========================================================================
// *Contains the api call methods for the user access management
// ========================================================================
package main

import (

	"fmt"
	"github.com/gin-gonic/gin"
)

type ApiRequestJson struct { //* Struct to take in requests & bind them to extract json key-values
	UserWalletAddress string `json:"userWalletAddress"`
	EventQueueID      string `json:"eventQueueId"`
	Action            string `json:"action"`
}

func createUser(c *gin.Context){ //*Creates a user with the wallet address as its unique identifier
	var json ApiRequestJson
	if err:=c.BindJSON(&json); err!=nil{
		fmt.Println("Error binding JSON!")
		c.JSON(400, gin.H{"error": err.Error()})
	}else{
		add_user_to_map(json.UserWalletAddress)
		c.JSON(200, gin.H{"message": fmt.Sprintf("User with wallet address [%s] created!", json.UserWalletAddress)})
	}
}

func deleteUser(c *gin.Context){ //*Deletes a user with the wallet address as its unique identifier
	var json ApiRequestJson
	if err:=c.BindJSON(&json); err!=nil{
		fmt.Println("Error binding JSON!")
		c.JSON(400, gin.H{"error": err.Error()})
	}else{
		remove_user_from_map(json.UserWalletAddress)
		c.JSON(200, gin.H{"message": fmt.Sprintf("User with wallet address [%s] deleted!", json.UserWalletAddress)})
	}
}

func addUserAccess(c *gin.Context){//*Adds access for the user to the topic
	var json ApiRequestJson
	if err:=c.BindJSON(&json); err!=nil{
		fmt.Println("Error binding JSON!")
		c.JSON(400, gin.H{"error": err.Error()})
	}else{
		res := add_access_for_user(json.UserWalletAddress, json.EventQueueID, json.Action)
		if res {
			c.JSON(200, gin.H{"message": fmt.Sprintf("User with wallet address [%s] granted access to [%s] the event queue id [%s]", json.UserWalletAddress, json.Action, json.EventQueueID)})
		}else{
			c.JSON(400, gin.H{"message": fmt.Sprintf("User with wallet address [%s] FAILED to grant access to [%s] the event queue id [%s]", json.UserWalletAddress, json.Action, json.EventQueueID)})
		}
	}
}

func removeUserAccess(c *gin.Context){//* Removes access for the user to the topic
	var json ApiRequestJson
	if err:=c.BindJSON(&json); err!=nil{
		fmt.Println("Error binding JSON!")
		c.JSON(400, gin.H{"error": err.Error()})
	}else{
		res := remove_access_for_user(json.UserWalletAddress, json.EventQueueID, json.Action)
		if res {
			c.JSON(200, gin.H{"message": fmt.Sprintf("User with wallet address [%s] removed access to [%s] the event queue id [%s]", json.UserWalletAddress, json.Action, json.EventQueueID)})
		}else{
			c.JSON(400, gin.H{"message": fmt.Sprintf("User with wallet address [%s] FAILED to remove access to [%s] the event queue id [%s]", json.UserWalletAddress, json.Action, json.EventQueueID)})
		}
	}	
}

func removePublisherAccess(c *gin.Context){//* Removes the publisher access for the user to all topics
	var json ApiRequestJson
	if err:=c.BindJSON(&json); err!=nil{
		fmt.Println("Error binding JSON!")
		c.JSON(400, gin.H{"error": err.Error()})
	}else{
		accessListMap[json.UserWalletAddress].PublisherAccess = []string{}
		c.JSON(200, gin.H{"message": fmt.Sprintf("User with wallet address [%s] removed access to all topics as publisher", json.UserWalletAddress)})
	}
}

func removeSubscriberAccess(c *gin.Context){//* Removes the subscriber access for the user to all topics
	var json ApiRequestJson
	if err:=c.BindJSON(&json); err!=nil{
		fmt.Println("Error binding JSON!")
		c.JSON(400, gin.H{"error": err.Error()})
	}else{
		accessListMap[json.UserWalletAddress].SubscriberAccess = []string{}
		c.JSON(200, gin.H{"message": fmt.Sprintf("User with wallet address [%s] removed access to all topics as subscriber", json.UserWalletAddress)})
	}
}

func listUserAccess(c *gin.Context){
	// Return accessListMap as JSON
	c.JSON(200, gin.H{"userAccessListMap": accessListMap})
}