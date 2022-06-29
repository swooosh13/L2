package find

import (
	"github.com/stretchr/testify/require"
	flags2 "grep/internal/flags"
	"strings"
	"testing"
)

var dataCount = `Isle of white heat,
Ovations of palms, clapping their fronds, yea!
But then, “Hey,
Make way, deadbeat!”
They crumple and stomp creativity.
And on again off am I’m,
bound for oasis, my locomotivity,
weaving in sandiness footsteps of time.`

func TestCount(t *testing.T) {
	req := require.New(t)

	flags := flags2.Flags{
		Ignore:  true,

	}

	text := strings.Split(dataContext, "\n")
	t.Run("Count word with ignore case", func(t *testing.T) {
		res := Count(text, "and", &flags)
		req.Equal(3, res)
	})

	flags.Ignore = false
	t.Run("Count word without ignore case", func(t *testing.T) {
		res := Count(text, "and", &flags)
		req.Equal(2, res)
	})
}
