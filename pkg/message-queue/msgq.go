package messagequeue

import (
	"fmt"
	"log"

	"github.com/GooDu-Dev/function-parser-go/pkg/message-queue/types"
	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageQueue interface {
	Connect() error
	OpenChannel() (*amqp.Channel, error)
	CreateQueue(channel *amqp.Channel, cfg *types.QueueConfig) amqp.Queue
	ConsumeQueue(channel *amqp.Channel, cfg *types.ConsumeConfig) <-chan amqp.Delivery
	Disconnect() error
}

type messageQueue struct {
	cfg        *types.MessageQueueConfig
	connection *amqp.Connection
}

func NewMessageQueue(cfg *types.MessageQueueConfig) MessageQueue {
	return &messageQueue{
		cfg: cfg,
	}
}

func (mq *messageQueue) Connect() error {

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", mq.cfg.Username, mq.cfg.Password, mq.cfg.Host, mq.cfg.Port))
	if err != nil {
		return err
	}

	mq.connection = conn

	return nil
}

func (mq *messageQueue) OpenChannel() (*amqp.Channel, error) {

	ch, err := mq.connection.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %s", err.Error())
	}

	return ch, nil
}

func (mq *messageQueue) CreateQueue(channel *amqp.Channel, cfg *types.QueueConfig) amqp.Queue {
	q, err := channel.QueueDeclare(
		cfg.QueueName,
		cfg.Durable,
		cfg.DeleteWhenUnused,
		cfg.Exclusive,
		cfg.NoWait,
		nil,
	)

	if err != nil {
		log.Fatalf("cannot create queue: %s", err.Error())
		return amqp.Queue{}
	}

	return q
}

func (mq *messageQueue) ConsumeQueue(channel *amqp.Channel, cfg *types.ConsumeConfig) <-chan amqp.Delivery {
	msg, err := channel.Consume(
		cfg.QueueName,
		cfg.Consumer,
		cfg.AutoAck,
		cfg.Exclusive,
		cfg.NoLocal,
		cfg.NoWait,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to consume message : %s", err.Error())
		return nil
	}

	return msg
}

func (mq *messageQueue) Disconnect() error {
	if mq.connection == nil {
		return fmt.Errorf("connection does not exist")
	}
	err := mq.connection.Close()
	if err != nil {
		return err
	}
	return nil
}
