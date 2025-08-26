package hashable

import (
	"hash"

	"github.com/samber/lo"
)

// Map is a Hashable which is used to describe map-like structures (maps themselves,
// structs, etc).
//
// It is hashed as an unordered list of key value pairs.
type Map map[string]Hashable

// Hash ...
func (m Map) Hash(hf func() hash.Hash) []byte {
	return Tagged("Map", unorderedList(
		lo.Map(lo.Entries(m), func(attr lo.Entry[string, Hashable], _ int) Hashable {
			return Tagged(attr.Key, attr.Value)
		}),
	)).Hash(hf)
}
