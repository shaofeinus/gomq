package main

import (
	"github.com/shaofeinus/gomq"
	"github.com/shaofeinus/gomq/examples/rpc/funcs"
	"github.com/shaofeinus/gomq/rpc"
	"log"
	"os"
)

// Calls an RPC function by its name and passing a message data
func main() {
	funcName := os.Args[1]
	message := os.Args[2]
	// AMQP URL is taking from the env variable GOMQ_AMQP_URL. See config.go
	gomq.SetupMQ("")
	defer gomq.TeardownMQ()
	funcs.Setup()
	err := rpc.InvokeRPC(funcName, map[string]interface{}{"message": message})
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
