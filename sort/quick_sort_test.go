package sort

import (
	"testing"

	ds "github.com/DavidCai1993/datastructures.go"
	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	assert := assert.New(t)

	cs := QuickSort(ds.Comparables{
		ds.IntComparable(14),
		ds.IntComparable(33),
		ds.IntComparable(27),
		ds.IntComparable(10),
		ds.IntComparable(35),
		ds.IntComparable(19),
		ds.IntComparable(42),
		ds.IntComparable(44),
	})

	assert.Equal(ds.Comparables{
		ds.IntComparable(10),
		ds.IntComparable(14),
		ds.IntComparable(19),
		ds.IntComparable(27),
		ds.IntComparable(33),
		ds.IntComparable(35),
		ds.IntComparable(42),
		ds.IntComparable(44),
	}, cs)
}
