package funcs

import (
	"fmt"
	"github.com/shaofeinus/gomq/pubsub"
)

// Setup Define events and register subscribers. Should be called during initialization for both sender and receiver
func Setup() {
	greetEvent := pubsub.RegisterEvent("GREET")
	pubsub.Subscribe(greetEvent, "a", SubA)
	pubsub.Subscribe(greetEvent, "b", SubB)
}

func SubA(args map[string]interface{}) {
	fmt.Printf("Subscriber A: %v\n", args)
}

func SubB(args map[string]interface{}) {
	fmt.Printf("Subscriber B: %v\n", args)
}
