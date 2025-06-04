package btree_test

import (
	"testing"

	"github.com/droxer/algo/btree"
)

type item int

func (i item) Less(other btree.Item) bool {
	return i < other.(item)
}

func TestGet(t *testing.T) {
	bt := btree.New(2)
	bt.Insert(item(4))
	bt.Insert(item(6))
	bt.Insert(item(8))
	bt.Insert(item(18))
	bt.Insert(item(20))
	bt.Insert(item(22))
	bt.Insert(item(24))
	bt.Insert(item(26))
	bt.Insert(item(28))
	bt.Insert(item(30))

	if bt.Get(item(11)) != nil {
		t.Fatalf("expected is nil, actual is %v", bt.Get(item(7)))
	}

	if bt.Get(item(8)) != item(8) {
		t.Fatalf("expected is 5, actual is %v", bt.Get(item(7)))
	}
}
