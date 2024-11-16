package week3

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPageCountCalculatedCorrectly(t *testing.T) {
	testData := []struct {
		bookPages     int32
		targetPage    int32
		expectedSwaps int32
	}{
		{5, 3, 1},
		{6, 2, 1},
		{5, 4, 0},
		{7, 3, 1},
		{6, 4, 1},
		{6, 5, 1},
	}

	for _, datum := range testData {
		t.Run(fmt.Sprintf("N=%d;P=%d", datum.bookPages, datum.targetPage), func(t *testing.T) {
			calculated := pageCount(datum.bookPages, datum.targetPage)

			require.Equal(t, datum.expectedSwaps, calculated)
		})
	}
}
