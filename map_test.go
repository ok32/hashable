package hashable

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	r := require.New(t)

	// just check the internal structure is fine;
	// the unorderedList and Tagged themselves are tested separately
	r.Equal(
		defHash(Tagged("Map", unorderedList{
			Tagged("a", String("abc")),
			Tagged("b", Number(147)),
		})),
		defHash(Map{
			"a": String("abc"),
			"b": Number(147),
		}),
	)
}
