package other

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddCorrectlyAddsElementsToSet(t *testing.T) {
	testData := []struct {
		input          []string
		add            []string
		expectedOutput []string
	}{
		{
			[]string{"a", "b", "c"},
			[]string{"d"},
			[]string{"a", "b", "c", "d"},
		},
		{
			[]string{},
			[]string{"a", "b", "c"},
			[]string{"a", "b", "c"},
		},
	}
	for i, datum := range testData {
		t.Run(fmt.Sprintf("TestCase-%v", i+1), func(t *testing.T) {
			//given
			set := NewSet[string](datum.input...)

			//when
			set.Add(datum.add...)

			//then
			values := set.Values()
			require.ElementsMatch(t, datum.expectedOutput, values)
		})
	}
}

func TestRemoveCorrectlyRemovesElementsFromSet(t *testing.T) {
	testData := []struct {
		input          []string
		remove         string
		expectedOutput bool
	}{
		{
			[]string{},
			"a",
			false,
		},
		{
			[]string{"a"},
			"b",
			false,
		},
		{
			[]string{"a"},
			"a",
			true,
		},
	}
	for i, datum := range testData {
		t.Run(fmt.Sprintf("TestCase-%v", i), func(t *testing.T) {
			//given
			set := NewSet(datum.input...)

			//when
			result := set.Remove(datum.remove)

			//then
			require.Equal(t, datum.expectedOutput, result)
		})
	}
}
