package week4

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRotateLeft(t *testing.T) {
	testData := []struct {
		d        int32
		arr      []int32
		expected []int32
	}{
		{
			2,
			[]int32{1, 2, 3, 4, 5},
			[]int32{3, 4, 5, 1, 2},
		},
		{
			5,
			[]int32{1, 2, 3, 4, 5},
			[]int32{1, 2, 3, 4, 5},
		},
		{
			1,
			[]int32{1},
			[]int32{1},
		},
		{
			100,
			[]int32{1},
			[]int32{1},
		},
	}

	for _, datum := range testData {
		t.Run(fmt.Sprintf("%v", datum.arr), func(t *testing.T) {
			calculated := rotateLeft(datum.d, datum.arr)

			require.Equal(t, datum.expected, calculated)
		})
	}
}
