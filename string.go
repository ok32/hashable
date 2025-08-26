package hashable

import (
	"hash"
)

// String creates a Hashable from a string.
//
// The hashing is trivial: just write the string's bytes to the hash.
func String[T ~string](s T) Hashable {
	return Tagged(
		"String",
		&str{
			inner: string(s),
		},
	)
}

type str struct {
	inner string
}

func (s str) Hash(hf func() hash.Hash) []byte {
	return bytez{inner: []byte(s.inner)}.Hash(hf)
}
