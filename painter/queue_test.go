package painter

import (
	"golang.org/x/exp/shiny/screen"
	"testing"
)

func TestMQueuePush(t *testing.T) {
	mq := MessageQueue{}

	if !mq.Empty() {
		t.Error("Expected empty queue, got non-empty")
	}

	op := OperationFunc(func(screen.Texture) {})
	mq.Push(op)

	if mq.Empty() {
		t.Error("Expected non-empty queue, got empty")
	}
}

func TestMQueuePull(t *testing.T) {
	mq := &MessageQueue{}

	expectedOp := new(Mock)
	mq.Push(expectedOp)

	op := mq.Pull()
	if op != expectedOp {
		t.Errorf("pull() from non-empty queue failed: expected %v, got %v", expectedOp, op)
	}

	if !mq.Empty() {
		t.Errorf("empty() after pull() failed: expected true, got false")
	}
}

func TestMQueueEmpty(t *testing.T) {
	mq := MessageQueue{}

	if !mq.Empty() {
		t.Errorf("empty() on empty queue failed: expected true, got false")
	}

	mq.Push(new(Mock))

	if mq.Empty() {
		t.Errorf("empty() on non-empty queue failed: expected false, got true")
	}
}
