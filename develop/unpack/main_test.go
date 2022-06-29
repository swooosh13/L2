package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	req := require.New(t)

	t.Run("Basic realization", func(t *testing.T) {
		req.Equal("", unpack(""))
		req.Equal("aaaabccddddde", unpack("a4bc2d5e"))
		req.Equal("abcd", unpack("abcd"))
		req.Equal("", unpack("45"))
	})

	t.Run("advanced usage", func(t *testing.T) {
		req.Equal("qwe45", unpack(`qwe\4\5`))
		req.Equal("qwe44444", unpack(`qwe\45`))
		req.Equal(`qwe\\\\\`, unpack(`qwe\\5`))
		// 8 штук = 4 + 4
		req.Equal(`\\\\\\\\`, unpack(`\\4\\4`))
		req.Equal(`\\фрр\`, unpack(`\\2фр2\\`))
		req.Equal(`\`, unpack(`\\`))
	})
}
