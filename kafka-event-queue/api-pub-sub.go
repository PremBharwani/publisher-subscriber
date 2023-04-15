package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type EventRequestJson struct {
	UserWalletAddress 	string `json:"userWalletAddress"`// binding:"required"`
	EventQueueID      	string `json:"eventQueueId"`// binding:"required"`
	Message 		 	string `json:"message"`// binding:"required"`
}

func publishEvent(c *gin.Context){
	var json EventRequestJson
	if err:=c.BindJSON(&json); err!=nil{
		fmt.Println("Error binding JSON!")
		c.JSON(400, gin.H{"error": err.Error()})
	}else{
		isEventPublished := publish_event(mProducerAgent, json.EventQueueID, json.Message)
		if isEventPublished {
			c.JSON(200, gin.H{"message": "Event published successfully!"})
		}else{
			c.JSON(400, gin.H{"message": "Event not published!"})
		}
	}
}

func consumeEvent(c *gin.Context){
	var json EventRequestJson
	if err:=c.BindJSON(&json); err!=nil{
		fmt.Println("Error binding JSON!")
		c.JSON(400, gin.H{"error": err.Error()})
	}else{
		eventMessages := get_events(mConsumerAgent, json.EventQueueID)
		if len(eventMessages) > 0 {
			c.JSON(200, gin.H{"message": eventMessages})
		}else{
			c.JSON(400, gin.H{"message": "No events found!"})
		}
	}
}