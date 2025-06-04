package linkedlist_test

import (
	"testing"

	"github.com/droxer/algo/linkedlist"
)

func TestLinkedList(t *testing.T) {
	l := linkedlist.New()

	// Test empty list
	if !l.IsEmpty() {
		t.Error("Expected empty list")
	}
	if l.Size() != 0 {
		t.Errorf("Expected size 0, got %d", l.Size())
	}

	// Test Insert
	l.Insert(1)
	l.Insert(2)
	l.Insert(3)

	if l.Size() != 3 {
		t.Errorf("Expected size 3, got %d", l.Size())
	}

	// Test Search
	if !l.Search(1) {
		t.Error("Expected to find value 1")
	}
	if !l.Search(2) {
		t.Error("Expected to find value 2")
	}
	if !l.Search(3) {
		t.Error("Expected to find value 3")
	}
	if l.Search(4) {
		t.Error("Expected to not find value 4")
	}

	// Test Delete
	if !l.Delete(2) {
		t.Error("Expected to delete value 2")
	}
	if l.Size() != 2 {
		t.Errorf("Expected size 2 after delete, got %d", l.Size())
	}
	if l.Search(2) {
		t.Error("Expected value 2 to be deleted")
	}

	// Test Delete non-existent value
	if l.Delete(4) {
		t.Error("Expected delete of non-existent value to return false")
	}

	// Test Delete head
	if !l.Delete(1) {
		t.Error("Expected to delete head value 1")
	}
	if l.Size() != 1 {
		t.Errorf("Expected size 1 after deleting head, got %d", l.Size())
	}

	// Test Delete last element
	if !l.Delete(3) {
		t.Error("Expected to delete last value 3")
	}
	if !l.IsEmpty() {
		t.Error("Expected empty list after deleting all elements")
	}
}

func TestLinkedListWithDifferentTypes(t *testing.T) {
	l := linkedlist.New()

	// Test with different types
	l.Insert("string")
	l.Insert(42)
	l.Insert(true)

	if !l.Search("string") {
		t.Error("Expected to find string value")
	}
	if !l.Search(42) {
		t.Error("Expected to find integer value")
	}
	if !l.Search(true) {
		t.Error("Expected to find boolean value")
	}

	// Test Delete with different types
	if !l.Delete("string") {
		t.Error("Expected to delete string value")
	}
	if !l.Delete(42) {
		t.Error("Expected to delete integer value")
	}
	if !l.Delete(true) {
		t.Error("Expected to delete boolean value")
	}

	if !l.IsEmpty() {
		t.Error("Expected empty list after deleting all elements")
	}
}
