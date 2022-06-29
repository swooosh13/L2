package lines

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUniqLines(t *testing.T) {
	req := require.New(t)
	t.Run("Print only uniq lines", func(t *testing.T) {
		lines := []string{
			"a",
			"b",
			"a ",
			"b",
			"c",
			"c",
		}
		req.Equal([]string{"a", "b", "a ", "c"}, UniqLines(lines))
	})
}

func TestIsSorted(t *testing.T) {
	req := require.New(t)
	t.Run("Not sorted strings", func(t *testing.T) {
		lines := []string{
			"a",
			"b",
			"a ",
			"b",
			"c",
			"c",
		}
		req.False(IsSorted(lines))
	})

	t.Run("Sorted strings", func(t *testing.T) {
		lines := []string{
			"a",
			"abc",
			"b",
			"b",
			"c",
			"c",
		}
		req.True(IsSorted(lines))
	})
}

func TestDefaultSort(t *testing.T) {
	req := require.New(t)

	t.Run("Simple sort", func(t *testing.T) {
		lines := []string{
			"g",
			"b",
			"a ",
			"b",
			"h",
			"c",
		}

		Sort(&lines)
		req.Equal([]string{"a ", "b", "b", "c", "g", "h"}, lines)
	})

	t.Run("Reverse sort", func(t *testing.T) {
		lines := []string{
			"g",
			"b",
			"a ",
			"b",
			"h",
			"c",
		}

		ReverseSort(&lines)
		req.Equal([]string{"h", "g", "c", "b", "b", "a "}, lines)
	})
}
