package find

import (
	"grep/internal"
	"grep/internal/flags"
	"strconv"
)

func WithContext(text []string, pattern string, flags *flags.Flags) ([]string, bool) {
	context := flags.Context
	window := internal.NewQueue(context)

	flag := false
	for i, line := range text {
		// before match
		if !flag {
			// if less than context -> add
			if window.Len() <= context {
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
		} else if context > 0 {
			if flags.LineNum {
				window.Enqueue(strconv.Itoa(i) + " " + line)
			} else {
				window.Enqueue(line)
			}
			context--
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
