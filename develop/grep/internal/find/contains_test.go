package find

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

var (
	data1 = `
The moon is emerging.
It going to be here
soon.
Now, it hangs in the air, full and stark.
That is probably God,
with a divine
silver spoon,
groping in the fish-soup of stars. 
`
)

func TestContains(t *testing.T) {
	req := require.New(t)
	text := strings.Split(data1, "\n")
	t.Run("Contains without ignore case", func(t *testing.T) {
		contains := false
		Contains(text[4], "it", false, false, &contains)
		req.True(contains)

		contains = false
		Contains(text[4], "It", false, false, &contains)
		req.True(!contains)
	})

	t.Run("Contains with ignore case", func(t *testing.T) {
		contains := false
		Contains(text[4], "now", true, false, &contains)
		req.True(contains)

		contains = false
		Contains(text[4], "Now", true, false, &contains)
		req.True(contains)
	})
	
	t.Run("Contains with fixed case", func(t *testing.T) {
		contains := false
		Contains(text[5], "That is probably God,", false, true, &contains)
		req.True(contains)

		contains = false
		Contains(text[5], "That is probably God", false, true, &contains)
		req.True(!contains)
	})
}
