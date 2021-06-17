package main

import (
	"fmt"
	"github.com/shaofeinus/gomq"
	"github.com/shaofeinus/gomq/examples/pubsub/funcs"
	"github.com/shaofeinus/gomq/pubsub"
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
	funcs.Setup()
	for _, subscriber := range subscribers {
		err := pubsub.WorkOnSubscriber("GREET", subscriber)
		if err != nil {
			log.Fatalf(fmt.Sprintf("Failed to work on subscriber \"%s\": %s", subscriber, err.Error()))
		}
	}
	<-forever
}
