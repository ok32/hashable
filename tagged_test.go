package hashable

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTagged(t *testing.T) {
	r := require.New(t)

	h := defHasher()
	h.Write(hashBytes([]byte("SomeTag")))
	h.Write([]byte("fake_hash"))
	expHash := h.Sum(nil)

	r.Equal(
		expHash,
		defHash(Tagged("SomeTag", stubHashable([]byte("fake_hash")))),
	)
}
