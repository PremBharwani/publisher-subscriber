// ========================================================================
//*Contains the functions to subscribe and get events from a topic.
// TODO: Add remove functionality later.
// ========================================================================
package main

import (
	"fmt"
    "os"
    "os/signal"
    "syscall"
    "time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func get_events(consumer *kafka.Consumer, topicId string) []string {
	var eventMessages []string = make([]string, 0)
	
	err := consumer.SubscribeTopics([]string{topicId}, nil)
	if err != nil {fmt.Printf("Failed to subscribe to topic: %v\n", err)}
    // Set up a channel for handling Ctrl-C, etc
    sigchan := make(chan os.Signal, 1)
    signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	recievedMessage := false

    // Process messages
    run := true
    for run {
        select {
        case sig := <-sigchan:
            fmt.Printf("Caught signal %v: terminating\n", sig)
            run = false
        default:
            ev, err := consumer.ReadMessage(100 * time.Millisecond)
            if err != nil {
                // Errors are informational and automatically handled by the consumer
				if (recievedMessage==true) {run = false}
                continue
            }
            // fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n", *ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
			eventMessages = append(eventMessages, string(ev.Value))
			recievedMessage = true
        }
    }

	return eventMessages

}