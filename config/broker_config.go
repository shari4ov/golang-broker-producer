package config

import amqp "github.com/rabbitmq/amqp091-go"

type RabbitConnectionConfig struct {
	Protocol string
	User     string
	Password string
	Host     string
	Port     string
}
type AmqpChannels struct {
	RemindersQueue *amqp.Channel
}
