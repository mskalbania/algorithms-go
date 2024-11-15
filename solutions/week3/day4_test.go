package week3

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSockPairsCountedCorrectly(t *testing.T) {
	testCases := []struct {
		data          []int32
		expectedPairs int32
	}{
		{
			[]int32{1, 2, 1, 2, 1, 3, 2},
			int32(2),
		},
		{
			[]int32{},
			int32(0),
		},
		{
			[]int32{1, 2, 3, 4},
			int32(0),
		},
		{
			[]int32{1, 1, 1, 1, 1},
			int32(2),
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCase.data), func(t *testing.T) {
			calculated := sockMerchant(int32(len(testCase.data)), testCase.data)

			require.Equal(t, testCase.expectedPairs, calculated)
		})
	}
}
