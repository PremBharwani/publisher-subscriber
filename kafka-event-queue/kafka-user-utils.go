// ========================================================================
//*Contains the functions that create publisher, consumer and admin client.
// TODO: Add remove functionality later.
// ========================================================================
package main

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)


func create_publisher() *kafka.Producer{
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": brokersList,
		"batch.size": producerBatchSize,
	})
	
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}else { fmt.Println("Created Producer") }
	return producer
}

func create_consumer() *kafka.Consumer{	
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":    brokersList,
		"group.id":             consumerGroupId,
		// "fetch.max.wait.ms":    100,
		"auto.offset.reset":    "smallest"})
	
	if err != nil {
		fmt.Printf("Failed to create consumer: %s\n", err)
		os.Exit(1)
	}else { fmt.Println("Created Consumer") }

	return consumer
}

func create_admin_client() *kafka.AdminClient{
	adminClient, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": brokersList})
	if err != nil {
		fmt.Printf("Failed to create Admin client: %s\n", err)
		os.Exit(1)
	}else{ fmt.Println("Created Admin Client") }
	return adminClient	
}

