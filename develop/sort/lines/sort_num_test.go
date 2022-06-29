package lines

import (
	"github.com/stretchr/testify/require"
	"sortutil/flags"
	"testing"
)

func TestSortByN(t *testing.T) {
	req := require.New(t)

	data := []string{
		"12 0 91 -1 a",
		"56 a s 9 12 3",
		"1 975 sad  9332 23 f",
		"2 -50 40  -24 23 f",
		"-2 50 -40  24 -23 c",
	}

	t.Run("sort by number column", func(t *testing.T) {
		req.Equal([]string{
			"-2 50 -40  24 -23 c",
			"1 975 sad  9332 23 f",
			"2 -50 40  -24 23 f",
			"12 0 91 -1 a",
			"56 a s 9 12 3",
		},
		SortByN(data, &flags.Flags{K: 1}))
	})

	t.Run("Sort with column without number and reverse", func(t *testing.T) {
		req.Equal([]string{
			"1 975 sad  9332 23 f",
			"-2 50 -40  24 -23 c",
			"12 0 91 -1 a",
			"2 -50 40  -24 23 f",
			"56 a s 9 12 3",
		},
			SortByN(data, &flags.Flags{K: 2, R: true}))
	})
}
