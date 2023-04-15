package main

import (
	// "fmt"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	// =============================================
	// ============= API Endpoints =================
	// =============================================

	// Calls through pub.sol 
	r.POST("/create-user", createUser)
	r.DELETE("/delete-user", deleteUser)

	// List userAccessMap & topicList
	r.GET("/user-access-list", listUserAccess)
	r.GET("/topic-list", listTopics)

	// Calls through topics.sol 
	r.POST("/create-topic", createTopic)
	r.DELETE("/delete-topic", deleteTopic)

	// Calls to remove publisher or subscriber access of a user compeletely
	r.POST("/remove-publisher-access", removePublisherAccess)
	r.POST("/remove-subscriber-access", removeSubscriberAccess)

	// Calls to control access of users to various topics
	r.POST("/add-user-access", addUserAccess)
	r.DELETE("/remove-user-access", removeUserAccess)

	// Calls to publish events to a topic/event queue
	r.POST("publish-event", publishEvent)
	r.GET("consume-event", consumeEvent)


	initializeLists() 		//*Init lists to maintain user access & topic list
	initializeKafkaAgents() //*Init pub, sub & admin kafka agents to perform operations 

	r.Run() // listen and serve on 0.0.0.0:8080
}