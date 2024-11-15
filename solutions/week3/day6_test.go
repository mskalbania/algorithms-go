package week3

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTriangleSelectedCorrectly(t *testing.T) {
	testCases := []struct {
		sticks           []int32
		expectedTriangle []int32
	}{
		{
			[]int32{1, 2, 3, 4, 5, 10},
			[]int32{3, 4, 5},
		},
		{
			[]int32{1, 1, 1, 2, 3, 5},
			[]int32{1, 1, 1},
		},
		{
			[]int32{1, 2, 3},
			[]int32{-1},
		},
		{
			[]int32{5, 5, 5, 5, 5},
			[]int32{5, 5, 5},
		},
		{
			[]int32{3, 4, 5},
			[]int32{3, 4, 5},
		},
		{
			[]int32{3, 4, 5},
			[]int32{3, 4, 5},
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCase.sticks), func(t *testing.T) {
			selected := maximumPerimeterTriangle(testCase.sticks)

			require.Equal(t, testCase.expectedTriangle, selected)
		})
	}
}
