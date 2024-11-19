package week4

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKangarooReturnsCorrectResult(t *testing.T) {
	testCases := []struct {
		x1, v1, x2, v2 int32
		expectedResult string
	}{
		{
			2, 1, 1, 2,
			"YES",
		},
		{
			0, 3, 4, 2,
			"YES",
		},
		{
			43, 2, 70, 2,
			"NO",
		},
	}
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%v,%v,%v,%v", testCase.x1, testCase.v1, testCase.x2, testCase.v2), func(t *testing.T) {
			result := kangaroo(testCase.x1, testCase.v1, testCase.x2, testCase.v2)
			require.Equal(t, testCase.expectedResult, result)
		})
	}
}
