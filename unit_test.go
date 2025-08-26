package hashable

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnit(t *testing.T) {
	r := require.New(t)

	h := defHasher()
	expHash := h.Sum(nil)

	r.Equal(
		defHash(Tagged("Unit", stubHashable(expHash))),
		defHash(Unit()),
	)
}
