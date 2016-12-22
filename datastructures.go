package datastructures

// Comparable represents a type which is comparable.
type Comparable interface {
	Compare(Comparable) int
}

// IntComparable represents a comparable int.
type IntComparable int

// Compare implements the Comparable interface.
func (ic IntComparable) Compare(i Comparable) int {
	other, ok := i.(IntComparable)

	if !ok {
		panic("param should be an intComparable")
	}

	if int(ic) == int(other) {
		return 0
	} else if int(ic) < int(other) {
		return -1
	} else {
		return 1
	}
}
