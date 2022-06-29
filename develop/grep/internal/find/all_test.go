package find

import (
	"github.com/stretchr/testify/require"
	flags2 "grep/internal/flags"
	"strings"
	"testing"
)

var dataAll = `Isle of white heat,
Ovations of palms, clapping their fronds, yea!
But then, “Hey,
Make way, deadbeat!”
They crumple and stomp creativity.
And on again off am I’m,
bound for oasis, my locomotivity,
weaving in sandiness footsteps of time.`

func TestAll(t *testing.T) {
	req := require.New(t)

	flags := flags2.Flags{
		Ignore: true,
	}

	text := strings.Split(dataAll, "\n")
	t.Run("All matches", func(t *testing.T) {
		res, ok := All(text, "ey", &flags)
		req.True(ok)
		req.Equal([]string{
			"But then, “Hey,",
			"They crumple and stomp creativity.",
		}, res)
	})
}
