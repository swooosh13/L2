package anagrams

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindAnagrams(t *testing.T) {
	req := require.New(t)
	words := []string{"ПятКа", "пятка", "пяТак", "пятак", "слиток", "стОлик", "тяпка", "Листок", "дОМ", "МОД", "код"}
	t.Run("Test anagram", func(t *testing.T) {
		res := FindAnagrams(&words)
		req.Equal([]string{"тяпка", "пятка", "пятак"}, *((*res)["пятка"]))
		req.Equal([]string{"столик", "слиток", "листок"}, *((*res)["слиток"]))
		req.Equal([]string{"мод", "дом"}, *((*res)["дом"]))
	})

	t.Run("Nil test", func(t *testing.T) {
		res := FindAnagrams(&words)
		req.Nil((*res)["код"])
	})
}
