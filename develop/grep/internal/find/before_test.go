package find

import (
	"github.com/stretchr/testify/require"
	flags2 "grep/internal/flags"
	"strings"
	"testing"
)

var dataBefore = `Isle of white heat,
Ovations of palms, clapping their fronds, yea!
But then, “Hey,
Make way, deadbeat!”
They crumple and stomp creativity.
And on again off am I’m,
bound for oasis, my locomotivity,
weaving in sandiness footsteps of time.`

func TestBefore(t *testing.T) {
	req := require.New(t)
	flags := flags2.Flags{
		Before: 1,
		Ignore: true,
	}

	text := strings.Split(dataBefore, "\n")
	t.Run("Before, 1 line", func(t *testing.T) {
		res, ok := Before(text, "but", &flags)
		req.True(ok)
		req.Equal([]string{
			"Ovations of palms, clapping their fronds, yea!",
			"But then, “Hey,",
		}, res)
	})

	flags.Before = 4
	t.Run("Before with range start of text", func(t *testing.T) {
		res, ok := Before(text, "WAY", &flags)
		req.True(ok)
		req.Equal([]string{
			"Isle of white heat,",
			"Ovations of palms, clapping their fronds, yea!",
			"But then, “Hey,",
			"Make way, deadbeat!”",
		}, res)
	})
}
