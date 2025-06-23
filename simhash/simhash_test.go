package simhash_test

import (
	"math/bits"
	"testing"

	"github.com/droxer/algo/simhash"
)

func hammingDistance(x, y uint64) int {
	return bits.OnesCount64(x ^ y)
}

func TestSimHash(t *testing.T) {
	text1 := "hello world simhash algorithm"
	text2 := "hello simhash algorithm world"
	text3 := "completely different text"

	hash1 := simhash.SimHash(text1)
	hash2 := simhash.SimHash(text2)
	hash3 := simhash.SimHash(text3)

	if hammingDistance(hash1, hash2) > 3 {
		t.Errorf("Hamming(text1, text2): %d", hammingDistance(hash1, hash2))
	}

	if hammingDistance(hash1, hash3) < 10 {
		t.Errorf("Hamming(text1, text3): %d", hammingDistance(hash1, hash3))
	}
}
