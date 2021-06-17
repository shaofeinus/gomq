package subs

import (
	"fmt"
	"github.com/shaofeinus/gomq"
)

// RegisterPubSub Define events and register subscribers. Should be called during initialization for both sender and receiver
func RegisterPubSub() {
	greetEvent := gomq.RegisterEvent("GREET")
	gomq.Subscribe(greetEvent, "a", SubA)
	gomq.Subscribe(greetEvent, "b", SubB)
}

func SubA(args map[string]interface{}) {
	fmt.Printf("Subscriber A: %v\n", args)
}

func SubB(args map[string]interface{}) {
	fmt.Printf("Subscriber B: %v\n", args)
}
