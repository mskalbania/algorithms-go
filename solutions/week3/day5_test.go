package week3

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBirdsCountedCorrectly(t *testing.T) {
	testCases := []struct {
		data           []int32
		expectedResult int32
	}{
		{
			[]int32{1, 1, 2, 2, 3},
			int32(1),
		},
		{
			[]int32{1, 1, 2, 2, 3, 3},
			int32(1),
		},
		{
			[]int32{5, 5, 4, 4, 3, 3, 2, 2, 1, 1},
			int32(1),
		},
		{
			[]int32{1, 2, 2, 3, 3, 3, 4, 4, 4, 4},
			int32(4),
		},
		{
			[]int32{1, 1, 1, 1, 1},
			int32(1),
		},
		{
			[]int32{},
			int32(0),
		},
		{
			[]int32{1, 2, 3, 4, 5},
			int32(1),
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCase.data), func(t *testing.T) {
			calculated := migratoryBirds(testCase.data)

			require.Equal(t, testCase.expectedResult, calculated)
		})
	}
}
