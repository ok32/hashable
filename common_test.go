package hashable

import (
	"crypto/sha256"
	"hash"
)

var defHasher = sha256.New

func hashBytes(bs []byte) []byte {
	h := defHasher()
	h.Write(bs)
	return h.Sum(nil)
}

func defHash(h Hashable) []byte {
	return h.Hash(defHasher)
}

type stubHashable []byte

func (sh stubHashable) Hash(_ func() hash.Hash) []byte {
	return sh
}
