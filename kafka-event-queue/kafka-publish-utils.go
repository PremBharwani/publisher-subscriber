// ========================================================================
//*Contains the functions that publish events to a topic.
// TODO: Add remove functionality later.
// ========================================================================
package main

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func publish_event(publisher *kafka.Producer, topicId string, message string) bool{
	
	quitGoRoutine := make(chan bool)

	// Listen to all the events on the default events channel
	go func() {
		for e := range publisher.Events() {

			doIQuit := <-quitGoRoutine
			if doIQuit { break }
			
			switch ev := e.(type) {
			case *kafka.Message:
				m := ev
				if m.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
				} else {
					fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
						*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
				}
			case kafka.Error:
				fmt.Printf("Error: %v\n", ev)
			default:
				fmt.Printf("Ignored event: %s\n", ev)
			}
		}
		fmt.Println("Done listening to events! Exiting go routine")
	}()

	fmt.Println("Trying to publish now!")
	err := publisher.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topicId, Partition: kafka.PartitionAny},
		Value:          []byte(message),
		// Headers:        []kafka.Header{{Key: "myTestHeader", Value: []byte("header values are binary")}},
	}, nil)
	fmt.Println("Done publishing now!")
	if err != nil {
		if err.(kafka.Error).Code() == kafka.ErrQueueFull {
			// Producer queue is full, wait 1s for messages
			// to be delivered then try again.
			fmt.Print("Producer queue is full, waiting 1s for messages to be delivered\n")
			time.Sleep(time.Second)
		}
		fmt.Printf("Failed to produce message: %v\n", err)
		return false
	}

	for publisher.Flush(10000) > 0 {
		fmt.Print("Still waiting to flush outstanding messages\n")
	}

	quitGoRoutine <- true
	fmt.Println("End of the main publish function")
	
	return true
}