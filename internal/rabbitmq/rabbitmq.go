package rabbitmq

import (
	"context"
	"fmt"
	"github.com/blazee5/film-api/internal/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

func NewRabbitMQConn(cfg *config.Config) (*amqp.Connection, error) {
	connAddr := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		cfg.RabbitMQUser,
		cfg.RabbitMQPassword,
		cfg.RabbitMQHost,
		cfg.RabbitMQPort,
	)
	return amqp.Dial(connAddr)
}

func SendMessage(ctx context.Context, message string, ch *amqp.Channel, q *amqp.Queue) error {
	err := ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

	if err != nil {
		return err
	}

	return nil
}
