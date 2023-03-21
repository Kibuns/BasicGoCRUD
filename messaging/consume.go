package messaging

import (
	"fmt"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func consumeMessage(message string) {
	conn, err := amqp.Dial(os.Getenv("CLOUDAMQP_URL"))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declare a queue
	q, err := ch.QueueDeclare(
		"strat_queue", // queue name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare queue")

	// Consume from the queue
	msgs, err := ch.Consume(
		q.Name, // Queue name
		"",     // Consumer tag
		true,   // Auto-acknowledge
		false,  // Exclusive
		false,  // No-local
		false,  // No-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	// Receive messages
	for msg := range msgs {
		fmt.Println(string(msg.Body))
	}
}
