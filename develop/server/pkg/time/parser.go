package time

import "time"

func Parse(t string) (time.Time, error) {
	return time.Parse(
		"2006-01-02",
		t,
	)
}
