package rabbitmq

import (
	"context"
	"errors"
	"fmt"
	"github.com/blazee5/film-api/films/internal/config"
	sl "github.com/blazee5/film-api/films/lib/logger/slog"
	amqp "github.com/rabbitmq/amqp091-go"
	"log/slog"
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

func NewChannelConn(conn *amqp.Connection, log *slog.Logger) (*amqp.Channel, error) {
	ch, err := conn.Channel()

	if err != nil {
		log.Info("failed to create a channel in rabbitmq:", sl.Err(err))
		return nil, err
	}

	return ch, nil
}

func NewQueueConn(ch *amqp.Channel, log *slog.Logger) (*amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Info("failed to declare a queue:", sl.Err(err))
		return nil, err
	}

	return &q, nil
}

func PublishMessage(ctx context.Context, message string, ch *amqp.Channel, q *amqp.Queue) error {
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

func NewConsumer(ctx context.Context, ch *amqp.Channel, q *amqp.Queue, consumeName string) (<-chan amqp.Delivery, error) {
	msgs, err := ch.Consume(q.Name, consumeName, false, false, false, false, nil)

	if err != nil {
		return nil, err
	}

	return msgs, errors.New("")
}
