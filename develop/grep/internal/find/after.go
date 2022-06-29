package find

import (
	"grep/internal"
	"grep/internal/flags"
	"strconv"
)

func After(text []string, pattern string, flags *flags.Flags) ([]string, bool) {
	l := flags.After
	window := internal.NewQueue(l)

	flag := false
	for i, line := range text {
		Contains(line, pattern, flags.Ignore, flags.Fixed, &flag)
		// before match
		if flag {
			// if less than context -> add
			if window.Len() <= l {
				if flags.LineNum {
					window.Enqueue(strconv.Itoa(i) + " " + line)
				} else {
					window.Enqueue(line)
				}
			} else {
				break
			}
		}
	}

	if !flag {
		return nil, false
	}

	return window.GetData(), true
}
