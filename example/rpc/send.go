package main

import (
	"github.com/shaofeinus/gomq"
	"github.com/shaofeinus/gomq/example/rpc/rpcfuncs"
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
	rpcfuncs.RegisterRPC()
	err := gomq.InvokeRPC(funcName, map[string]interface{}{"message": message})
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
