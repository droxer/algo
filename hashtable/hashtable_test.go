package hashtable_test

import (
	"testing"

	"github.com/droxer/algo/hashtable"
)

func TestHashTable(t *testing.T) {
	ht := hashtable.New(10)

	// Test Put and Get
	ht.Put("key1", "value1")
	ht.Put("key2", "value2")

	if val, exists := ht.Get("key1"); !exists || val != "value1" {
		t.Errorf("Expected value1 for key1, got %v", val)
	}
	if val, exists := ht.Get("key2"); !exists || val != "value2" {
		t.Errorf("Expected value2 for key2, got %v", val)
	}

	// Test non-existent key
	if _, exists := ht.Get("key3"); exists {
		t.Error("Expected key3 to not exist")
	}

	// Test Update
	ht.Put("key1", "newvalue1")
	if val, exists := ht.Get("key1"); !exists || val != "newvalue1" {
		t.Errorf("Expected newvalue1 for key1, got %v", val)
	}

	// Test Delete
	ht.Delete("key1")
	if _, exists := ht.Get("key1"); exists {
		t.Error("Expected key1 to be deleted")
	}

	// Test Size
	if ht.Size() != 1 {
		t.Errorf("Expected size 1, got %d", ht.Size())
	}
}

func TestHashTableCollision(t *testing.T) {
	ht := hashtable.New(10)

	// Test collision handling
	ht.Put("key1", "value1")
	ht.Put("key2", "value2")
	ht.Put("key3", "value3")

	if val, exists := ht.Get("key1"); !exists || val != "value1" {
		t.Errorf("Expected value1 for key1, got %v", val)
	}
	if val, exists := ht.Get("key2"); !exists || val != "value2" {
		t.Errorf("Expected value2 for key2, got %v", val)
	}
	if val, exists := ht.Get("key3"); !exists || val != "value3" {
		t.Errorf("Expected value3 for key3, got %v", val)
	}
}

func TestHashTableResize(t *testing.T) {
	ht := hashtable.New(10)

	// Fill the table to trigger resize
	for i := 0; i < 100; i++ {
		ht.Put(string(rune(i)), i)
	}

	// Verify all values are still accessible after resize
	for i := 0; i < 100; i++ {
		if val, exists := ht.Get(string(rune(i))); !exists || val != i {
			t.Errorf("Expected %d for key %c, got %v", i, rune(i), val)
		}
	}
}
