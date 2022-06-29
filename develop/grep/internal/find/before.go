package find

import (
	"grep/internal"
	"grep/internal/flags"
	"strconv"
)

func Before(text []string, pattern string, flags *flags.Flags) ([]string, bool) {
	l := flags.Before
	window := internal.NewQueue(l)

	flag := false
	for i, line := range text {
		// before match
		if !flag {
			// if less than context -> add
			if window.Len() <= l {
				if flags.LineNum {
					window.Enqueue(strconv.Itoa(i) + " " + line)
				} else {
					window.Enqueue(line)
				}
			} else {
				// dequeue last && enqueue new
				_ = window.Dequeue()
				if flags.LineNum {
					window.Enqueue(strconv.Itoa(i) + " " + line)
				} else {
					window.Enqueue(line)
				}
			}
			// мы нашли
		} else {
			break
		}

		Contains(line, pattern, flags.Ignore, flags.Fixed, &flag)
	}

	if !flag {
		return nil, false
	}

	return window.GetData(), true
}
