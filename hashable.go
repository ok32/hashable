// Package hashable provides a collection of hashable primitives
// which can be composed together to describe how to hash an arbitrary go
// structure.
package hashable

import "hash"

// Hashable represents a thing which is, being given a hash algorithm,
// knows how to hash itself.
type Hashable interface {
	Hash(func() hash.Hash) []byte
}
