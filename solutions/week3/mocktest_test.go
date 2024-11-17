package week3

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetTotalX(t *testing.T) {
	testData := []struct {
		a     []int32
		b     []int32
		count int32
	}{
		{
			[]int32{2, 6},
			[]int32{24, 36},
			2,
		},
		{
			[]int32{2, 4},
			[]int32{16, 32, 96},
			3,
		},
	}

	for _, datum := range testData {
		t.Run(fmt.Sprintf("%v+%v", datum.a, datum.a), func(t *testing.T) {
			calculated := getTotalX(datum.a, datum.b)

			require.Equal(t, datum.count, calculated)
		})
	}
}
