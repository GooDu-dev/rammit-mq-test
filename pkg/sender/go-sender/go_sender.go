package gosender

import (
	"context"
	"log"
	"time"

	"github.com/GooDu-Dev/function-parser-go/pkg/inventory"
	"github.com/GooDu-Dev/function-parser-go/pkg/message-queue/types"
	"github.com/GooDu-Dev/function-parser-go/pkg/sender"
	amqp "github.com/rabbitmq/amqp091-go"
)

type GoSender struct {
}

func NewGoSender() sender.Sender {

	return &GoSender{}
}

func (gs *GoSender) Send(ctx context.Context, channel *amqp.Channel, msg types.Message, pkg inventory.Package) error {

	ctxTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := channel.PublishWithContext(ctxTimeout,
		msg.Metadata().Exchange,
		msg.Metadata().RouteingKey,
		msg.Metadata().Mandatory,
		msg.Metadata().Immediate,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        pkg.Process().JSON(),
		})
	if err != nil {
		return err
	}

	log.Printf("%s [x] Sent %s with %s\n", pkg.ID(), pkg.Process().Function(), pkg.Process().Params())
	return nil
}
