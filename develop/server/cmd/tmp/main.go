package main

import (
	"fmt"
	"time"
)

func main() {
	times := []time.Time{
		NewDate(2020, 1, 1),
		NewDate(2020, 1, 2),
		NewDate(2020, 1, 3),
		NewDate(2020, 1, 4),
	}

	start := NewDate(2020, 1, 2)
	end := start.AddDate(0,0,1)
	for _, v  := range times {
		if (v == start || v.After(start)) && v.Before(end) {
			fmt.Println(v)
		}
	}

	fmt.Println()
}

func NewDate(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0,0,0, 0, time.UTC)
}