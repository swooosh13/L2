package find

import (
	"github.com/stretchr/testify/require"
	flags2 "grep/internal/flags"
	"strings"
	"testing"
)

var data2 = `Isle of white heat,
Ovations of palms, clapping their fronds, yea!
But then, “Hey,
Make way, deadbeat!”
They crumple and stomp creativity.
And on again off am I’m,
bound for oasis, my locomotivity,
weaving in sandiness footsteps of time.`

func TestAfter(t *testing.T) {
	req := require.New(t)
	flags := flags2.Flags{
		After: 1,
		Ignore: true,
	}

	text := strings.Split(data2, "\n")
	t.Run("After 1 line", func(t *testing.T) {
		res, ok := After(text, "but", &flags)
		req.True(ok)
		req.Equal([]string{
			"But then, “Hey,",
			"Make way, deadbeat!”",
		}, res)
	})

	flags.After = 3
	t.Run("After end of ", func(t *testing.T) {
		res, ok := After(text, "MY", &flags)
		req.True(ok)
		req.Equal([]string{
			"bound for oasis, my locomotivity,",
			"weaving in sandiness footsteps of time.",
		}, res)
	})
}
