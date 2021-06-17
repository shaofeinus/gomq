package gomq

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
)

func AddQueue(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) error {
	q, err := CH.QueueDeclare(
		name,       // name
		durable,    // durable
		autoDelete, // delete when unused
		exclusive,  // exclusive
		noWait,     // no-wait
		args,       // arguments
	)
	QUEUES[name] = q
	return err
}

func GetQueue(queue string) (*amqp.Queue, error) {
	_, ok := QUEUES[queue]
	if !ok {
		err := AddQueue(
			queue, // name
			true,  // durable
			false, // delete when unused
			false, // exclusive
			false, // no-wait
			nil,   // arguments
		)
		if err != nil {
			return nil, err
		}
	}
	q := QUEUES[queue]
	return &q, nil
}

func BindQueue(queue string, exchange string) error {
	q, err := GetQueue(queue)
	if err != nil {
		return err
	}
	err = CH.QueueBind(
		q.Name,   // queue name
		"",       // routing key
		exchange, // exchange
		false,
		nil,
	)
	return err
}

func SendMessageToQueue(queue string, message []byte) error {
	q, err := GetQueue(queue)
	if err != nil {
		return err
	}
	err = CH.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	if err != nil {
		return err
	}
	return nil
}

func SendMessageToExchange(exchange string, message []byte) error {
	err := CH.ExchangeDeclare(
		exchange, // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return err
	}
	err = CH.Publish(
		exchange, // exchange
		"",       // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	if err != nil {
		return err
	}
	return nil
}

func SendJSONToQueue(queue string, data interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return SendMessageToQueue(queue, b)
}

func SendJSONToExchange(exchange string, data interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return SendMessageToExchange(exchange, b)
}

func ConsumeMessage(queue string, handle func(message []byte)) {
	q, err := GetQueue(queue)
	failOnError(err, "Failed to get queue")

	err = CH.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := CH.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			handle(d.Body)
			_ = d.Ack(false)
		}
	}()
}

func ConsumeJSON(queue string, handle func(map[string]interface{})) {
	ConsumeMessage(queue, func(message []byte) {
		var d map[string]interface{}
		_ = json.Unmarshal(message, &d)
		handle(d)
	})
}
