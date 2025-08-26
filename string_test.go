package hashable

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	r := require.New(t)

	r.Equal(
		defHash(Tagged("String", stubHashable(hashBytes([]byte("abc"))))),
		defHash(String("abc")),
	)
}
