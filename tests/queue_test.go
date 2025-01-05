package tests

import (
	"testing"
	"time"

	"github.com/Dnreikronos/transactions/queue"
	"github.com/stretchr/testify/assert"
)

func TestConsumeQueue(t *testing.T) {
	message := []byte(`{"description": "Test transaction", "value": 100.00, "date": "2025-01-02", "currency": "USD"}`)

	processedMessages := make(chan []byte, 1)

	process := func(msg []byte) {
		processedMessages <- msg
	}

	go queue.ConsumeQueue(process)

	err := queue.PublishToQueue(message)
	assert.NoError(t, err, "Expected message to be published to the queue")

	select {
	case msg := <-processedMessages:
		assert.Equal(t, message, msg, "Expected processed message to match the original")
	case <-time.After(3 * time.Second):
		t.Fatal("Expected message to be processed within the time limit")
	}
}
