package main

import (
	"github.com/shaofeinus/gomq"
	"github.com/shaofeinus/gomq/example/pubsub/subs"
	"log"
	"os"
)

// Trigger an event by its name and passing a message data
func main() {
	event := os.Args[1]
	message := os.Args[2]
	// AMQP URL is taking from the env variable GOMQ_AMQP_URL. See config.go
	gomq.SetupMQ("")
	defer gomq.TeardownMQ()
	subs.RegisterPubSub()
	err := gomq.Publish(event, map[string]interface{}{"message": message})
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
