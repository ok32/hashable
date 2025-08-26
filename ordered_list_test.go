package hashable

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOrderedList(t *testing.T) {
	r := require.New(t)

	h := defHasher()
	h.Write([]byte("item1"))
	h.Write([]byte("item2"))
	expHash := h.Sum(nil)

	r.Equal(
		defHash(Tagged("OrderedList", stubHashable(expHash))),
		defHash(OrderedList(
			stubHashable([]byte("item1")),
			stubHashable([]byte("item2")),
		)),
	)

	r.NotEqual(
		defHash(Tagged("OrderedList", stubHashable(expHash))),
		defHash(OrderedList(
			stubHashable([]byte("item2")),
			stubHashable([]byte("item1")),
		)),
	)
}
