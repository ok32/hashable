package hashable

import (
	"hash"
)

// Tagged creates a Hashable from a string tag and another Hashable.
// Useful for distinguishing different Hashables which otherwise (if not wrapped in Tagged)
// would have the same internal structure (and would produce the same hash too).
//
// It's hashed as an ordered list consisting of the tag and inner Hashable.
func Tagged(tag string, inner Hashable) Hashable {
	return &tagged{
		tag:   tag,
		inner: inner,
	}
}

type tagged struct {
	tag   string
	inner Hashable
}

func (t *tagged) Hash(hf func() hash.Hash) []byte {
	return orderedList{
		&str{inner: t.tag},
		t.inner,
	}.Hash(hf)
}
