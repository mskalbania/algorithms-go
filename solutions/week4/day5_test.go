package week4

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClosestNumberWorksCorrectly(t *testing.T) {
	testCases := []struct {
		arr      []int32
		expected []int32
	}{
		{
			[]int32{5, 2, 3, 4, 1},
			[]int32{1, 2, 2, 3, 3, 4, 4, 5},
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCase.arr), func(t *testing.T) {
			result := closestNumbers(testCase.arr)
			require.Equal(t, testCase.expected, result)
		})
	}
}
