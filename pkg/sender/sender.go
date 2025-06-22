package sender

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/GooDu-Dev/function-parser-go/pkg/inventory"
	"github.com/GooDu-Dev/function-parser-go/pkg/message-queue/types"
)

type Sender interface {
	Send(ctx context.Context, channel *amqp.Channel, msg types.Message, pkg inventory.Package) error
}
