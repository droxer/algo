package simhash
//
// SimHash is a function that calculates the similarity between two strings.
// It uses the SimHash algorithm to calculate the similarity between two strings.
// The algorithm is based on the concept of similarity hashing, which is a technique
// for calculating the similarity between two strings.
// The algorithm is based on the concept of similarity hashing, which is a technique
//
import (
	"hash/fnv"
	"strings"
)

func tokenize(text string) []string {
	return strings.Fields(text)
}

func hash64(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}

func SimHash(text string) uint64 {
	tokens := tokenize(text)
	v := [64]int{}

	for _, token := range tokens {
		h := hash64(token)
		for i := 0; i < 64; i++ {
			if (h>>i)&1 == 1 {
				v[i] += 1
			} else {
				v[i] -= 1
			}
		}
	}

	var fingerprint uint64 = 0
	for i := 0; i < 64; i++ {
		if v[i] >= 0 {
			fingerprint |= (1 << i)
		}
	}
	return fingerprint
}
