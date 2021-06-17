package main

import (
	"fmt"
	"github.com/shaofeinus/gomq"
	"github.com/shaofeinus/gomq/example/rpc/rpcfuncs"
	"log"
	"os"
)

// Execute RPC functions on some queues
func main() {
	queues := os.Args[1:]
	// AMQP URL is taking from the env variable GOMQ_AMQP_URL. See config.go
	gomq.SetupMQ("")
	defer gomq.TeardownMQ()
	forever := make(chan bool)
	log.Printf(fmt.Sprintf(" [*] Waiting for messages on queues %v. To exit press CTRL+C", queues))
	rpcfuncs.RegisterRPC()
	gomq.WorkOnRPC(queues)
	<-forever
}
