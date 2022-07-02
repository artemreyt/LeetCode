package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch2DMatrix(t *testing.T) {
	matr := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}

	tests := []struct {
		m   [][]int
		t   int
		exp bool
	}{
		{
			m:   matr,
			t:   13,
			exp: false,
		},
		{
			m:   matr,
			t:   11,
			exp: true,
		},
		{
			m:   matr,
			t:   -1,
			exp: false,
		},
		{
			m:   matr,
			t:   1,
			exp: true,
		},
		{
			m:   matr,
			t:   2,
			exp: false,
		},
		{
			m:   matr,
			t:   60,
			exp: true,
		},
		{
			m:   matr,
			t:   59,
			exp: false,
		},
		{
			m:   matr,
			t:   100,
			exp: false,
		},
	}

	for _, test := range tests {
		assert.Equalf(t, test.exp, searchMatrix(test.m, test.t), "searchMatrix(%d) binarySearch approach failed", test.t)
		assert.Equalf(t, test.exp, searchMatrixTwoPointers(test.m, test.t), "searchMatrix(%d) 2 pointers approach failed", test.t)
	}
}
