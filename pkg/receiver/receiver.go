package receiver

import (
	"github.com/GooDu-Dev/function-parser-go/pkg/inventory"
	"github.com/GooDu-Dev/function-parser-go/pkg/message-queue/types"
)

type Receiver interface {
	Receive(msg types.Message) (inventory.Package, error)
}
