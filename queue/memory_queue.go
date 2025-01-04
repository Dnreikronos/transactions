package queue

import (
	"errors"
	"log"
)

var transactionQueue = make(chan []byte, 100)

func PublishToQueue(message []byte) error {
	select {
	case transactionQueue <- message:
		log.Println("Message published to in-memory queue")
		return nil
	default:
		log.Println("Queue is full, failed to enqueue message")
		return errors.New("queue is full")
	}
}
