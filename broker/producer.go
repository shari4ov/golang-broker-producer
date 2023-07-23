package broker

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"notification-parser/config"
)

func ConnectRabbit(config config.RabbitConnectionConfig) (*amqp.Connection, error) {
	conn, err := amqp.Dial(fmt.Sprintf("%s://%s:%s@%s:%s/", config.Protocol, config.User, config.Password, config.Host, config.Port))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

type Queue struct {
	Conn      *amqp.Connection
	channel   *amqp.Channel
	QueueName string
}

func New(amqpConnection *amqp.Connection, amqpChannel *amqp.Channel, queueName string) *Queue {
	return &Queue{
		Conn:      amqpConnection,
		channel:   amqpChannel,
		QueueName: queueName,
	}
}
func (q *Queue) DeclareQueue() (*amqp.Channel, error) {
	ch, err := q.Conn.Channel()
	if err != nil {
		return nil, err
	}
	_, err = ch.QueueDeclare(
		q.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	err = ch.Qos(
		1,
		0,
		false,
	)
	if err != nil {
		return nil, err
	}
	return ch, nil
}

func (q *Queue) SendMessage(message []byte) error {
	ctx := context.Background()
	err := q.channel.PublishWithContext(ctx,
		"",          // exchange
		q.QueueName, // routing key
		false,       // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         message,
		})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
