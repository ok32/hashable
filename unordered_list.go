package hashable

import (
	"bytes"
	"hash"
	"slices"

	"github.com/samber/lo"
)

// UnorderedList creates a Hashable from a list of other hashables.
//
// The hashing algorithms is:
// 1. hash each list item
// 2. sort the item hashes lexicographically
// 3. write each item's hash to the resulting hash
//
// Items order is not respected. This means that hashing two unordered lists
// with the same items but with different item orders produces the same
// hashes.
func UnorderedList(items ...Hashable) Hashable {
	return Tagged("UnorderedList", unorderedList(items))
}

type unorderedList []Hashable

func (l unorderedList) Hash(hf func() hash.Hash) []byte {
	itemHashes := lo.Map(l, func(item Hashable, _ int) []byte {
		return item.Hash(hf)
	})

	slices.SortFunc(itemHashes, func(a, b []byte) int {
		return bytes.Compare(a, b)
	})

	h := hf()

	for _, itemHash := range itemHashes {
		h.Write(itemHash)
	}

	return h.Sum(nil)
}
