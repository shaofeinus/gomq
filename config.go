package gomq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
)

var AMQP_URL = os.Getenv("GOMQ_AMQP_URL")

var CONN *amqp.Connection
var CH *amqp.Channel
var QUEUES = make(map[string]amqp.Queue)

func SetupMQ(amqpUrl string) {
	if amqpUrl == "" {
		amqpUrl = AMQP_URL
	}
	conn, err := amqp.Dial(amqpUrl)
	failOnError(err, "Failed to connect to RabbitMQ")
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	CONN = conn
	CH = ch
}

func TeardownMQ() {
	_ = CONN.Close()
	_ = CH.Close()
}
