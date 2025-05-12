package _type

import "sort"

type Int64Slice []int64

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type order bool

const (
	ASC  order = true
	DESC order = false
)

// SortInt64 orders the int64 array according to ascending or descending.
func SortInt64(a []int64, o order) {
	if o {
		sort.Sort(Int64Slice(a))
	} else {
		sort.Sort(sort.Reverse(Int64Slice(a)))
	}
}
