package hashable

// Unit creates a Hashable which doesn't depend on any value.
// Useful when you need to hash some nil's.
//
// It's hashed as a tagged 0-lenght byte slice.
func Unit() Hashable {
	return Tagged("Unit", bytez{})
}
