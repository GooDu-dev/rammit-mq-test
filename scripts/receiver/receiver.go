package main

import (
	"fmt"
	"log"
	"os"

	messagequeue "github.com/GooDu-Dev/function-parser-go/pkg/message-queue"
	msqTypes "github.com/GooDu-Dev/function-parser-go/pkg/message-queue/types"
)

func main() {

	// ctx, cancel := signal.NotifyContext(
	// 	context.Background(),
	// 	syscall.SIGINT,
	// 	syscall.SIGTERM,
	// 	syscall.SIGQUIT,
	// )
	// defer cancel()

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

	msgs := msgQ.ConsumeQueue(ch, &msqTypes.ConsumeConfig{
		QueueName: "test_queue",
		Consumer:  "",
		AutoAck:   false,
		Exclusive: false,
		NoLocal:   false,
		NoWait:    false,
	})

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a function %s", d.Body)
		}
	}()

	log.Printf("[*] Waitiong for message. to exit press CTRL+C")
	<-forever
}
