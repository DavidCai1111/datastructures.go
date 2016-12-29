package sort

import (
	"math"

	ds "github.com/DavidCai1993/datastructures.go"
)

// ShellSort sorts the given []Comparable by shell sort algorithm.
// Time complexity O(n^1.3)
// Space complexity O(1)
func ShellSort(c ds.Comparables, interval int) ds.Comparables {
	if len(c) <= 1 {
		return c
	}

	if interval == 1 {
		return InsertionSort(c)
	}

	sub := make([][]int, 0)

	for i := 0; i <= int(math.Ceil(float64(len(c)/2))); i++ {
		if i+interval < len(c) {
			sub = append(sub, []int{i, i + interval})
		} else {
			sub = append(sub, []int{i})
		}
	}

	for _, s := range sub {
		if len(s) == 1 {
			continue
		}

		if c[s[1]].Compare(c[s[0]]) < 0 {
			c[s[1]], c[s[0]] = c[s[0]], c[s[1]]
		}
	}

	return ShellSort(c, int(math.Ceil(float64(interval/2))))
}
