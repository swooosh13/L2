package lines

import (
	"github.com/stretchr/testify/require"
	"sortutil/flags"
	"testing"
)

func TestSortByColumn(t *testing.T) {
	req := require.New(t)
	flags := flags.Flags{}

	data1 := []string{
		"a b c d",
		"b c d a",
		"d e f g",
		"f x y z o ",
	}

	t.Run("Sort", func(t *testing.T) {
		d := data1
		flags.K = 1

		expected := []string{
			"a b c d",
			"b c d a",
			"d e f g",
			"f x y z o ",
		}

		req.Equal(expected, SortByColumn(d, &flags))
	})

	data2 := []string{
		"b c d a",
		"a b c d",
		"f x y z o ",
		"d e f g",
	}

	t.Run("Sort2", func(t *testing.T) {
		flags.K = 1

		expected := []string{
			"a b c d",
			"b c d a",
			"d e f g",
			"f x y z o ",
		}

		req.Equal(expected, SortByColumn(data2, &flags))
	})

	data3 := []string {
		"a a z y",
		"a b y g",
		"f g t .",
	}

	t.Run("Sort by third col", func(t *testing.T) {
		flags.K = 3
		expected := []string{
			"f g t .",
			"a b y g",
			"a a z y",
		}
		req.Equal(expected, SortByColumn(data3, &flags))
	})
}
