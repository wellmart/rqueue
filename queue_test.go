package rqueue

import "testing"

func TestQueue(t *testing.T) {
	queue := New()

	queue.Enqueue(1)

	if queue.Length() != 1 {
		t.Error("Invalid length of queue")
	}

	queue.TryDequeue()

	if queue.Length() != 0 {
		t.Error("Invalid length of queue")
	}

	queue.TryDequeue()
}

func TestQueueResize(t *testing.T) {
	queue := New()

	queue.Capacity()

	for i := 0; i < 65; i++ {
		queue.Enqueue(1)
	}

	for i := 0; i < 64; i++ {
		queue.TryDequeue()
	}
}
