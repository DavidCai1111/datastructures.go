package search

import (
	"testing"

	ds "github.com/DavidCai1993/datastructures.go"
	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	assert := assert.New(t)

	t.Run("found", func(t *testing.T) {
		i := BinarySearch(ds.Comparables{
			ds.IntComparable(10),
			ds.IntComparable(14),
			ds.IntComparable(19),
			ds.IntComparable(27),
			ds.IntComparable(33),
			ds.IntComparable(35),
			ds.IntComparable(42),
			ds.IntComparable(44),
			ds.IntComparable(50),
		}, ds.IntComparable(42))

		assert.Equal(6, i)
	})

	t.Run("not found", func(t *testing.T) {
		i := BinarySearch(ds.Comparables{
			ds.IntComparable(10),
			ds.IntComparable(14),
			ds.IntComparable(19),
			ds.IntComparable(27),
			ds.IntComparable(33),
			ds.IntComparable(35),
			ds.IntComparable(42),
			ds.IntComparable(44),
			ds.IntComparable(50),
		}, ds.IntComparable(99))

		assert.Equal(-1, i)

		i = BinarySearch(ds.Comparables{}, ds.IntComparable(14))

		assert.Equal(-1, i)
	})
}
