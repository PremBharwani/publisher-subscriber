package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var bootstrapServers string = "localhost:9092"

func create_kafka_topic(topicName string, partitionCount int, replicationFactor int) bool{
	
	// Create a new AdminClient.
	// AdminClient can also be instantiated using an existing
	// Producer or Consumer instance, see NewAdminClientFromProducer and
	// NewAdminClientFromConsumer.
	a, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": bootstrapServers})
	if err != nil {
		fmt.Printf("Failed to create Admin client: %s\n", err)
		os.Exit(1)
	}

	// Contexts are used to abort or limit the amount of time
	// the Admin call blocks waiting for a result.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create topics on cluster.
	// Set Admin options to wait for the operation to finish (or at most 60s)
	maxDur, err := time.ParseDuration("60s")
	if err != nil {
		panic("ParseDuration(60s)")
	}
	results, err := a.CreateTopics(
		ctx,
		// Multiple topics can be created simultaneously
		// by providing more TopicSpecification structs here.
		[]kafka.TopicSpecification{{
			Topic:             topicName,
			NumPartitions:     partitionCount,
			ReplicationFactor: replicationFactor}},
		// Admin options
		kafka.SetAdminOperationTimeout(maxDur))
	if err != nil {
		fmt.Printf("Failed to create topicName: %v\n", err)
		// os.Exit(1)
		return false
	}

	// Print results
	for _, result := range results {
		fmt.Printf("%s\n", result)
	}

	a.Close()
	return true
}

func list_kafka_topics() []string {
	a, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": bootstrapServers})
	if err != nil {
		fmt.Printf("Failed to create Admin client: %s\n", err)
		os.Exit(1)
	}

	// Contexts are used to abort or limit the amount of time
	// the Admin call blocks waiting for a result.
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// Get metadata for all topics
	// Set Admin options to wait for the operation to finish (or at most 60s)
	maxDur, err := time.ParseDuration("60s")
	if err != nil {
		panic("ParseDuration(60s)")
	}

	m, err := a.GetMetadata(nil, true, int(maxDur.Seconds()))
	if err != nil {
		fmt.Printf("Failed to get metadata: %v\n", err)
		os.Exit(1)
	}

	// Print metadata for all topics
	// fmt.Printf("%d brokers, %d topics:\n", len(m.Brokers), len(m.Topics))
	topics := make([]string, 0)

	for _, topic := range m.Topics {
		// fmt.Println(topic.Topic)
		topics = append(topics, topic.Topic )
	}

	a.Close()
	return topics
}

func delete_kafka_topic(topicName string) bool{

	a, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": bootstrapServers})
	if err != nil {
		fmt.Printf("Failed to create Admin client: %s\n", err)
		os.Exit(1)
	}

	// Contexts are used to abort or limit the amount of time
	// the Admin call blocks waiting for a result.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Delete topics on cluster
	// Set Admin options to wait for the operation to finish (or at most 60s)
	maxDur, err := time.ParseDuration("60s")
	if err != nil {
		panic("ParseDuration(60s)")
	}

	results, err := a.DeleteTopics(ctx, []string{topicName}, kafka.SetAdminOperationTimeout(maxDur))
	if err != nil {
		fmt.Printf("Failed to delete topics: %v\n", err)
		// os.Exit(1)
		return false
	}

	// Print results
	for _, result := range results {
		fmt.Printf("%s\n", result)
	}

	a.Close()
	return true
}

func main() {

	if len(os.Args) != 5 {
		fmt.Fprintf(os.Stderr,
			"Usage: %s <bootstrap-servers> <topic> <partition-count> <replication-factor>\n",
			os.Args[0])
		os.Exit(1)
	}

	bootstrapServers := os.Args[1]
	bootstrapServers += "oklol"

	topic := os.Args[2]
	numParts, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("Invalid partition count: %s: %v\n", os.Args[3], err)
		os.Exit(1)
	}
	replicationFactor, err := strconv.Atoi(os.Args[4])
	if err != nil {
		fmt.Printf("Invalid replication factor: %s: %v\n", os.Args[4], err)
		os.Exit(1)
	}

	
	res := create_kafka_topic(topic, numParts, replicationFactor)
	fmt.Println(res)
	topicList := list_kafka_topics()
	fmt.Println("Topics list:", topicList)
	res1 := delete_kafka_topic(topic)
	fmt.Println(res1)
	topicList = list_kafka_topics()
	fmt.Println("Topics list after deletion:", topicList)

}