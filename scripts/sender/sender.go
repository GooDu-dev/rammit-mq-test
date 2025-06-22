package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	ivtTypes "github.com/GooDu-Dev/function-parser-go/pkg/inventory/types"
	messagequeue "github.com/GooDu-Dev/function-parser-go/pkg/message-queue"
	msqTypes "github.com/GooDu-Dev/function-parser-go/pkg/message-queue/types"
	gosender "github.com/GooDu-Dev/function-parser-go/pkg/sender/go-sender"
)

func main() {

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer cancel()

	msgQ := messagequeue.NewMessageQueue(&msqTypes.MessageQueueConfig{
		Username: "guest",
		Password: "guest",
		Host:     "localhost",
		Port:     "5672",
	})

	msgQ.Connect()
	defer msgQ.Disconnect()

	ch, err := msgQ.OpenChannel()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	q := msgQ.CreateQueue(ch, &msqTypes.QueueConfig{
		QueueName: "test_queue",
	})

	sender := gosender.NewGoSender()

	msg := msqTypes.NewMessage(msqTypes.MessageMetadata{
		Exchange:    "",
		RouteingKey: q.Name,
		Mandatory:   false,
		Immediate:   false,
	})

	pc := ivtTypes.NewProcess("PrintHelloWorldNTimes", map[string]interface{}{
		"n": 10,
	})

	pkg := ivtTypes.NewPackage(string("this-is-id"), pc, 10, 1, 10, false)

	err = sender.Send(ctx, ch, msg, pkg)
	if err != nil {
		log.Fatal("error sending message: ", err.Error())
	}
}
