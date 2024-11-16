package week4

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxArrCalculatedCorrectly(t *testing.T) {
	testData := []struct {
		data           []int32
		expectedMaxLen int32
	}{
		{[]int32{1, 1, 2, 2, 4, 4, 5, 5, 5}, 5},
		{[]int32{4, 6, 5, 3, 3, 1}, 3},
		{[]int32{1, 2, 2, 3, 1, 2}, 5},
		{[]int32{4, 2, 3, 4, 4, 9, 98, 98, 3, 3, 3, 4, 2, 98, 1, 98, 98, 1, 1, 4, 98, 2, 98, 3, 9, 9, 3, 1, 4, 1, 98, 9, 9, 2, 9, 4, 2, 2, 9, 98, 4, 98, 1, 3, 4, 9, 1, 98, 98, 4, 2, 3, 98, 98, 1, 99, 9, 98, 98, 3, 98, 98, 4, 98, 2, 98, 4, 2, 1, 1, 9, 2, 4}, 22},
		{[]int32{1, 1, 1, 1}, 4},
	}

	for _, datum := range testData {
		t.Run(fmt.Sprintf("%v", datum.data), func(t *testing.T) {
			calculatedMax := pickingNumbers(datum.data)

			assert.Equal(t, datum.expectedMaxLen, calculatedMax)
		})
	}
}
