package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ApiRequestJson struct {
	UserWalletAddress string `json:"userWalletAddress"`// binding:"required"`
	EventQueueID      string `json:"eventQueueId"`// binding:"required"`
	Action            string `json:"action"`// binding:"required"`
}

func main() {
	
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
	
		c.JSON(200, gin.H{
			"message": "pong",
		})

	})

	r.GET("/user-access-list", func(c *gin.Context) {
		userAccesListMap := get_user_access_list()
		c.JSON(200, gin.H{
			"message": userAccesListMap,
		})
	})

	r.GET("/verify-user", func(c *gin.Context) {
		var json ApiRequestJson
		var res bool
		if err:=c.BindJSON(&json); err == nil {
			res = verify_user(json.UserWalletAddress, json.EventQueueID, json.Action) // Attempt to perform the action with the user
			
		}else{
			fmt.Println("Error in binding the json: ", err)
			return
		}

		if res {
			c.JSON(200, gin.H{
				"message": fmt.Sprintf("User wallet address [%s] has access to [%s] the event queue id [%s]", json.UserWalletAddress, json.Action, json.EventQueueID),
			})
			return
		}
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("User wallet address [%s] does not have access to [%s] the event queue id [%s]", json.UserWalletAddress, json.Action, json.EventQueueID),
		})

	})

	r.POST("/add-user-access", func(c *gin.Context) {
		
		var json ApiRequestJson
		res := false
		if err:=c.BindJSON(&json); err == nil {
			res = add_access_for_user(json.UserWalletAddress, json.EventQueueID, json.Action) // Attempt to perform the action with the user 
		}else{
			fmt.Println("Error in binding the json: ", err)
		}
		
		if(res){ // Success 
			c.JSON(200, gin.H{
				"message": fmt.Sprintf("Added user wallet address [%s] to the [%s] access list for event queue id [%s]", json.UserWalletAddress, json.Action, json.EventQueueID),
				"list_after_modif": get_user_access_list(),
			})
		}else{
			c.JSON(400, gin.H{
				"message": fmt.Sprintf("Failed to add user wallet address [%s] to the [%s] access list for event queue id [%s]", json.UserWalletAddress, json.Action, json.EventQueueID),
			})
		}
	})

	r.POST("/add-user", func(c *gin.Context) {
		
		var json ApiRequestJson
		if err:=c.BindJSON(&json); err == nil {
			add_user_to_map(json.UserWalletAddress) // Adding the user to the access list
			c.JSON(200, gin.H{
				"message": fmt.Sprintf("Added user wallet address [%s] to the access list", json.UserWalletAddress),
				"list_after_modif": get_user_access_list(),
				}) 
		}else{
			fmt.Println("Error in binding the json: ", err)
		}
		
	})

	r.DELETE("/delete-user-access", func(c *gin.Context) {
		var json ApiRequestJson
		res := false
		if err:=c.BindJSON(&json); err == nil {
			res = remove_access_for_user(json.UserWalletAddress, json.EventQueueID, json.Action) // Attempt to perform the action with the user 
		}else{
			fmt.Println("Error in binding the json: ", err)
		}
		
		if(res){ // Success 
			c.JSON(200, gin.H{
				"message": fmt.Sprintf("Removed user wallet address [%s] from the [%s] access list for event queue id [%s]", json.UserWalletAddress, json.Action, json.EventQueueID),
				"list_after_modif": get_user_access_list(),
			})
		}else{
			c.JSON(400, gin.H{
				"message": fmt.Sprintf("Failed to remove user wallet address [%s] from the [%s] access list for event queue id [%s]", json.UserWalletAddress, json.Action, json.EventQueueID),
			})
		}
	})

	r.DELETE("/delete-user", func(c *gin.Context) {
		var json ApiRequestJson
		if err:=c.BindJSON(&json); err == nil {
			remove_user_from_map(json.UserWalletAddress) // Remove the user to the access list
			c.JSON(200, gin.H{
				"message": fmt.Sprintf("Removed user wallet address [%s] from the access list", json.UserWalletAddress),
				"list_after_modif": get_user_access_list(),
				})
		}else{
			fmt.Println("Error in binding the json: ", err)
		}
	})

	r.GET("/event-queue-list", func(c *gin.Context) {
		uniqueEventQueueMap := get_queue_access_list()
		c.JSON(200, gin.H{
			"message": uniqueEventQueueMap,
		})
	})

	r.POST("/add-event-queue", func(c *gin.Context) {
		var json ApiRequestJson
		if err:=c.BindJSON(&json); err == nil {
			add_queue_to_map(json.EventQueueID) // Adding the user to the access list
			c.JSON(200, gin.H{
				"message": fmt.Sprintf("Added event queue id [%s] to the unique event queue list", json.EventQueueID),
				"list_after_modif": get_queue_access_list(),
				})
		}else{
			fmt.Println("Error in binding the json: ", err)
		}
	})

	r.DELETE("/delete-event-queue", func(c *gin.Context) {
		var json ApiRequestJson
		if err:=c.BindJSON(&json); err == nil {
			remove_queue_from_map(json.EventQueueID) // Remove the user to the access list
			c.JSON(200, gin.H{
				"message": fmt.Sprintf("Removed event queue id [%s] from the unique event queue list", json.EventQueueID),
				"list_after_modif": get_queue_access_list(),
				})
		}else{
			fmt.Println("Error in binding the json: ", err)
		}
	})

	initializeLists() // Initialize the access list & the unique event queue MAPS ( Because GOLANG map implementation needs us to initialize the map before using it)
	
	r.Run() // listen and serve on 0.0.0.0:8080

}
