package types

import (
	"testing"
)

func TestInt64(t *testing.T) {
	s := []int64{-98, -2, -1024, -100, 0, -3, -1, 1, 3, 9, 2, 8, 6, 56, 7, 63, 89, 1024, 25000}
	sASC := []int64{-1024, -100, -98, -3, -2, -1, 0, 1, 2, 3, 6, 7, 8, 9, 56, 63, 89, 1024, 25000}
	sDESC := []int64{25000, 1024, 89, 63, 56, 9, 8, 7, 6, 3, 2, 1, 0, -1, -2, -3, -98, -100, -1024}
	SortInt64(s, ASC)
	for i, v := range s {
		if v != sASC[i] {
			t.Fatal("function [SortInt64] got an unexpected result.")
		}
	}
	SortInt64(s, DESC)
	for i, v := range s {
		if v != sDESC[i] {
			t.Fatal("function [SortInt64] got an unexpected result.")
		}
	}
}
