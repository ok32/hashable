package hashable

import (
	"encoding/binary"
	"fmt"
	"hash"

	"golang.org/x/exp/constraints"
)

// Number creates a Hashable from a numerically typed value.
//
// The hashing is implemented as writing the number's binary representation to the hash.
func Number[T primitiveNumber](x T) Hashable {
	return Tagged(
		"Number",
		number[T]{
			inner: x,
		},
	)
}

type primitiveNumber interface {
	constraints.Integer | constraints.Float
}

type number[T primitiveNumber] struct {
	inner T
}

func (n number[T]) Hash(hf func() hash.Hash) []byte {
	correctedVal := any(n.inner)
	switch v := correctedVal.(type) {
	case int:
		correctedVal = int64(v)
	case uint:
		correctedVal = uint64(v)
	}

	h := hf()

	err := binary.Write(h, binary.LittleEndian, correctedVal)
	if err != nil {
		panic(fmt.Sprintf("BUG: binary.Write (writing to a hash.Hash) returns an error only if a non-supported type has been provided, but they are all filtered out with type constraints: %v", err))
	}

	return h.Sum(nil)
}
