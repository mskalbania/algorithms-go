package softavail

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetOsDetailsWorksCorrectly(t *testing.T) {
	testData := []struct {
		name            string
		expectedDetails []string
	}{
		{
			"Windows 8",
			[]string{"Windows", "false", "2009-02-12 15:00:00 +0000 UTC"},
		},
		{
			"Ubuntu 11",
			[]string{"Linux", "true", "true", "false"},
		},
	}

	for _, datum := range testData {
		t.Run(datum.name, func(t *testing.T) {
			details, err := getOsDetails(datum.name)

			require.NoError(t, err)
			require.Equal(t, datum.expectedDetails, details)
		})
	}
}
