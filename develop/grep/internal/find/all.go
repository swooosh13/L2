package find

import (
	"grep/internal/flags"
	"strconv"
)

func All(text []string, pattern string, flags *flags.Flags) ([]string, bool) {
	res := make([]string, 0)
	flag := false
	for i, line := range text {
		if exists := IsContains(line, pattern, flags.Ignore, flags.Fixed); exists {
			flag = true
			if flags.LineNum {
				res = append(res, strconv.Itoa(i)+" "+line)
			} else {
				res = append(res, line)
			}
		}
	}

	return res, flag
}
