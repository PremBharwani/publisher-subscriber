// ========================================================================
//*Contains the functions that handle kafka topics like create, list and delete.
// TODO: Add remove functionality later.
// ========================================================================
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func list_kafka_topics(a *kafka.AdminClient) []string { // Method to list kafka topic ids

	maxDur, err := time.ParseDuration("60s")
	if err != nil {
		panic("ParseDuration(60s)")
	}

	m, err := a.GetMetadata(nil, true, int(maxDur.Seconds()))
	if err != nil {
		fmt.Printf("Failed to get metadata: %v\n", err)
		os.Exit(1)
	}

	topics := make([]string, 0)

	for _, topic := range m.Topics {topics = append(topics, topic.Topic )}

	return topics
}

func create_kafka_topic(adminClient *kafka.AdminClient, topicId string) bool{// Method to create a kafka topic

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	maxDur, err := time.ParseDuration("60s")
	if err != nil {
		panic("ParseDuration(60s)")
	}
	_, err = adminClient.CreateTopics(
		ctx,
		[]kafka.TopicSpecification{{
			Topic:             topicId,
			NumPartitions:     numPartitions,
			ReplicationFactor: replicationFactor}},

		kafka.SetAdminOperationTimeout(maxDur))
	if err != nil {
		fmt.Printf("Failed to create topic: %v\n", err)
		os.Exit(1)
	}else{
		fmt.Printf("Created the topic %s\n", topicId)
	}

	return true
}

func delete_kafka_topic(adminClient *kafka.AdminClient, topicId string) bool{// Method to delete a kafka topic
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Delete topics on cluster
	// Set Admin options to wait for the operation to finish (or at most 60s)
	maxDur, err := time.ParseDuration("60s")
	if err != nil {
		panic("ParseDuration(60s)")
	}

	results, err := adminClient.DeleteTopics(ctx, []string{topicId}, kafka.SetAdminOperationTimeout(maxDur))
	if err != nil {
		fmt.Printf("Failed to delete topics: %v\n", err)
		os.Exit(1)
	}

	// Print results
	for _, result := range results {
		fmt.Printf("%s\n", result)
	}
	return true
}

