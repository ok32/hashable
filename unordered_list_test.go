package hashable

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnorderedList(t *testing.T) {
	r := require.New(t)

	h := defHasher()
	h.Write([]byte("item1"))
	h.Write([]byte("item2"))
	expHash := h.Sum(nil)

	r.Equal(
		defHash(Tagged("UnorderedList", stubHashable(expHash))),
		defHash(UnorderedList(
			stubHashable([]byte("item1")),
			stubHashable([]byte("item2")),
		)),
	)

	r.Equal(
		defHash(Tagged("UnorderedList", stubHashable(expHash))),
		defHash(UnorderedList(
			stubHashable([]byte("item2")),
			stubHashable([]byte("item1")),
		)),
	)
}
