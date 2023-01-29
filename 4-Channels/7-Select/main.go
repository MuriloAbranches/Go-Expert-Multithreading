package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id  int64
	msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)

	var i int64 = 0

	go func() {
		for {
			time.Sleep(time.Second * 2)
			atomic.AddInt64(&i, 1)
			msg := Message{i, "Hello from RabbitMQ"}
			c1 <- msg
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			atomic.AddInt64(&i, 1)
			msg := Message{i, "Hello from Kafka"}
			c2 <- msg
		}
	}()

	for {
		select {
		case msg := <-c1:
			fmt.Printf("Received from RabbitMQ - ID: %d, %s\n", msg.id, msg.msg)

		case msg := <-c2:
			fmt.Printf("Received from RabbitMQ - ID: %d, %s\n", msg.id, msg.msg)

		case <-time.After(time.Second * 3):
			println("timeout")

			// default:
			// 	println("default")
		}
	}
}
