package find

import (
	"grep/internal/flags"
)

func Count(text []string, pattern string, flags *flags.Flags) int {
	count := 0
	for _, line := range text {
		if ok := IsContains(line, pattern, flags.Ignore, flags.Fixed); ok {
			count++
		}
	}

	return count
}

func InvertCount(text []string, pattern string, flags *flags.Flags) int {
	return len(text) - Count(text, pattern, flags)
}