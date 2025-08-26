package hashable

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBytes(t *testing.T) {
	r := require.New(t)

	r.Equal(
		defHash(Tagged("Bytes", stubHashable(hashBytes([]byte{1, 2, 3})))),
		defHash(Bytes([]byte{1, 2, 3})),
	)
}
