package find

import (
	"github.com/stretchr/testify/require"
	flags2 "grep/internal/flags"
	"strings"
	"testing"
)

var dataContext = `Isle of white heat,
Ovations of palms, clapping their fronds, yea!
But then, “Hey,
Make way, deadbeat!”
They crumple and stomp creativity.
And on again off am I’m,
bound for oasis, my locomotivity,
weaving in sandiness footsteps of time.`

func TestWithContext(t *testing.T) {
	req := require.New(t)

	flags := flags2.Flags{
		Ignore:  true,
		Context: 2,
	}

	text := strings.Split(dataContext, "\n")
	t.Run("Context in center", func(t *testing.T) {
		res, ok := WithContext(text, "and", &flags)
		req.True(ok)
		req.Equal(`But then, “Hey,
Make way, deadbeat!”
They crumple and stomp creativity.
And on again off am I’m,
bound for oasis, my locomotivity,`, strings.Join(res, "\n"))
	})

	flags.Context = 1
	t.Run("Context in center with 1 el around", func(t *testing.T) {
		res, ok := WithContext(text, "and", &flags)
		req.True(ok)
		req.Equal(`Make way, deadbeat!”
They crumple and stomp creativity.
And on again off am I’m,`, strings.Join(res, "\n"))
	})

	flags.Context = 2
	t.Run("Context in start", func(t *testing.T) {
		res, ok := WithContext(text, "palms", &flags)
		req.True(ok)
		req.Equal(`Isle of white heat,
Ovations of palms, clapping their fronds, yea!
But then, “Hey,
Make way, deadbeat!”`, strings.Join(res, "\n"))
	})
}
