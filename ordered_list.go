package hashable

import (
	"hash"

	"github.com/samber/lo"
)

// OrderedList creates a Hashable from a list of other hashables.
// The hashing algorithms is:
// 1. hash each list item
// 2. write each item's hash to the resulting hash
//
// Items order is respected. This means that hashing two ordered lists
// with the same items but with different item orders produces different
// hashes.
func OrderedList(items ...Hashable) Hashable {
	return Tagged("OrderedList", orderedList(items))
}

type orderedList []Hashable

func (l orderedList) Hash(hf func() hash.Hash) []byte {
	itemHashes := lo.Map(l, func(item Hashable, _ int) []byte {
		return item.Hash(hf)
	})

	h := hf()

	for _, itemHash := range itemHashes {
		h.Write(itemHash)
	}

	return h.Sum(nil)
}
