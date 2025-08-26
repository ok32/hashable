package hashable

import (
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumber(t *testing.T) {
	r := require.New(t)

	h := defHasher()
	err := binary.Write(h, binary.LittleEndian, uint64(147))
	r.NoError(err)
	expHash := h.Sum(nil)

	r.Equal(
		defHash(Tagged("Number", stubHashable(expHash))),
		defHash(Number(147)),
	)
}
