package bloomfilter

import (
	"fmt"
	"hash/fnv"
)

type BloomFilter struct {
	bitset []bool
	k      int
	size   int
}

func NewBloomFilter(size, k int) *BloomFilter {
	return &BloomFilter{
		bitset: make([]bool, size),
		k:      k,
		size:   size,
	}
}

func hash(s string, seed int) int {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprintf("%d-%s", seed, s)))
	return int(h.Sum32())
}

func (bf *BloomFilter) Add(s string) {
	for i := 0; i < bf.k; i++ {
		pos := hash(s, i) % bf.size
		bf.bitset[pos] = true
	}
}

func (bf *BloomFilter) Exists(s string) bool {
	for i := 0; i < bf.k; i++ {
		pos := hash(s, i) % bf.size
		if !bf.bitset[pos] {
			return false
		}
	}
	return true
}
