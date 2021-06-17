package main

import (
	"fmt"
	"github.com/shaofeinus/gomq"
	"github.com/shaofeinus/gomq/example/pubsub/subs"
	"log"
	"os"
)

// Act on events for some subscribers
func main() {
	subscribers := os.Args[1:]
	// AMQP URL is taking from the env variable GOMQ_AMQP_URL. See config.go
	gomq.SetupMQ("")
	defer gomq.TeardownMQ()

	forever := make(chan bool)
	log.Printf(fmt.Sprintf(" [*] Waiting for events for subs %v. To exit press CTRL+C", subscribers))
	subs.RegisterPubSub()
	for _, subscriber := range subscribers {
		err := gomq.WorkOnSubscriber("GREET", subscriber)
		if err != nil {
			log.Fatalf(fmt.Sprintf("Failed to work on subscriber \"%s\": %s", subscriber, err.Error()))
		}
	}
	<-forever
}
