package stack_test

import (
	"testing"

	"github.com/droxer/algo/stack"
)

func TestStack(t *testing.T) {
	s := stack.New()

	// Test empty stack
	if !s.IsEmpty() {
		t.Error("Expected empty stack")
	}
	if s.Size() != 0 {
		t.Errorf("Expected size 0, got %d", s.Size())
	}

	// Test Push and Size
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Size() != 3 {
		t.Errorf("Expected size 3, got %d", s.Size())
	}

	// Test Peek
	if val, ok := s.Peek(); !ok || val != 3 {
		t.Errorf("Expected peek value 3, got %v", val)
	}

	// Test Pop
	if val, ok := s.Pop(); !ok || val != 3 {
		t.Errorf("Expected pop value 3, got %v", val)
	}
	if s.Size() != 2 {
		t.Errorf("Expected size 2 after pop, got %d", s.Size())
	}

	// Test multiple Pops
	if val, ok := s.Pop(); !ok || val != 2 {
		t.Errorf("Expected pop value 2, got %v", val)
	}
	if val, ok := s.Pop(); !ok || val != 1 {
		t.Errorf("Expected pop value 1, got %v", val)
	}

	// Test empty stack after pops
	if !s.IsEmpty() {
		t.Error("Expected empty stack after all pops")
	}
	if val, ok := s.Pop(); ok {
		t.Errorf("Expected pop to fail on empty stack, got %v", val)
	}
	if val, ok := s.Peek(); ok {
		t.Errorf("Expected peek to fail on empty stack, got %v", val)
	}
}

func TestStackWithDifferentTypes(t *testing.T) {
	s := stack.New()

	// Test with different types
	s.Push("string")
	s.Push(42)
	s.Push(true)

	if val, ok := s.Pop(); !ok || val != true {
		t.Errorf("Expected pop value true, got %v", val)
	}
	if val, ok := s.Pop(); !ok || val != 42 {
		t.Errorf("Expected pop value 42, got %v", val)
	}
	if val, ok := s.Pop(); !ok || val != "string" {
		t.Errorf("Expected pop value 'string', got %v", val)
	}
}

	