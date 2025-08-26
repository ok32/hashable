package hashable

import (
	"hash"
)

// Bytes creates a Hashable from a bytes slice.
//
// The hashing is trivial: just write the bytes to the hash.
func Bytes[T ~[]byte](bs T) Hashable {
	return Tagged(
		"Bytes",
		&bytez{
			inner: []byte(bs),
		},
	)
}

type bytez struct {
	inner []byte
}

func (bs bytez) Hash(hf func() hash.Hash) []byte {
	h := hf()

	h.Write(bs.inner)

	return h.Sum(nil)
}
