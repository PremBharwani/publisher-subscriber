// ========================================================================
//* Contains api call methods to manage topics
//* These mutate the local map as well as the kafka topics
// ========================================================================
package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
)

// =================
// = Config values =
// =================
var brokersList string 			= "localhost:9092" // Should be something like : "host1:xyz,host2:xyz"
var producerBatchSize int 		= 1
var consumerGroupId string 		= "myGroup"
var numPartitions int 			= 1
var replicationFactor int 		= 1
// =================

var mProducerAgent *kafka.Producer
var mConsumerAgent *kafka.Consumer
var mAdminAgent *kafka.AdminClient

func initializeKafkaAgents(){
	mProducerAgent =  create_publisher()
	mConsumerAgent = create_consumer()
	mAdminAgent = create_admin_client()
}

func createTopic(c *gin.Context){//*Creates a topic on local map & kafka
	var json ApiRequestJson
	if err:=c.BindJSON(&json); err!=nil{
		fmt.Println("Error binding JSON!")
		c.JSON(400, gin.H{"error": err.Error()})
	}else{
		isTopicAdded := create_kafka_topic(mAdminAgent, json.EventQueueID)
		if isTopicAdded {
			add_queue_to_map(json.EventQueueID)
			c.JSON(200, gin.H{"message": fmt.Sprintf("Topic %s added", json.EventQueueID)})
		}else{
			c.JSON(400, gin.H{"message": fmt.Sprintf("Topic %s not added", json.EventQueueID)})
		}
	}
}

func deleteTopic(c *gin.Context){//*Deletes a topic from local map & kafka
	var json ApiRequestJson
	if err:=c.BindJSON(&json); err!=nil{
		fmt.Println("Error binding JSON!")
		c.JSON(400, gin.H{"error": err.Error()})
	}else{
		isTopicDeleted := delete_kafka_topic(mAdminAgent, json.EventQueueID)
		if isTopicDeleted {
			remove_queue_from_map(json.EventQueueID)
			c.JSON(200, gin.H{"message": fmt.Sprintf("Topic %s deleted", json.EventQueueID)})
		}else{
			c.JSON(400, gin.H{"message": fmt.Sprintf("Topic %s not deleted", json.EventQueueID)})
		}
	}
}

func listTopics(c *gin.Context){//*Lists all topics on kafka
	topicList := list_kafka_topics(mAdminAgent)
	c.JSON(200, gin.H{"kafka_topics": topicList, "local_topics_map": uniqEventQueueMap})
}

// func testMain(){//TODO: Delete this later.

// 	initializeKafkaAgents()

// 	// Testing topic creation/deletion
// 	topicList := list_kafka_topics(a); fmt.Println("Topic list before adding : ", topicList)
// 	isTopicAdded := create_topic(a, "prem-eq")
// 	if(isTopicAdded){fmt.Printf("Topic added\n")}
// 	topicList = list_kafka_topics(a)
// 	fmt.Println("Topic list after adding : ", topicList)
// 	isTopicDeleted := delete_topic(a, "prem-eq")
// 	if(isTopicDeleted) { fmt.Println("Topic deleted") }
// 	topicList = list_kafka_topics(a)
// 	fmt.Println("Topic list after adding : ", topicList)
// 	fmt.Println("=======================")
// 	// Testing publishing of events.
// 	_= publish_event(p, "prem-eq", "hello there!");	_ = publish_event(p, "prem-eq", "hello there 2!")
// 	_= publish_event(p, "prem-eq", "hello there!");	_ = publish_event(p, "prem-eq", "hello there 2!")
// 	_= publish_event(p, "prem-eq", "hello there!");	_ = publish_event(p, "prem-eq", "hello there 2!")
// 	_= publish_event(p, "prem-eq", "hello there!");	_ = publish_event(p, "prem-eq", "hello there 2!")
// 	_= publish_event(p, "prem-eq", "hello there!");	_ = publish_event(p, "prem-eq", "hello there 2!")
// 	_= publish_event(p, "prem-eq", "hello there!");	_ = publish_event(p, "prem-eq", "hello there 2!")
// 	_= publish_event(p, "prem-eq", "hello there!");	_ = publish_event(p, "prem-eq", "hello there 2!")
// 	_= publish_event(p, "prem-eq", "hello there!");	_ = publish_event(p, "prem-eq", "hello there 2!")
// 	_= publish_event(p, "prem-eq", "hello there!");	_ = publish_event(p, "prem-eq", "hello there 2!")
// 	_= publish_event(p, "prem-eq", "hello there!");	_ = publish_event(p, "prem-eq", "hello there 2!")

// 	eventMessages := get_events(c, "prem-eq")
// 	fmt.Println("Event messages : ", eventMessages)


// 	p.Close()
// 	c.Close()
// 	a.Close()
// }


