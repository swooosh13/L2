package main

import "fmt"

type WalkStrategy struct {}

func (r *WalkStrategy) buildRoute(from int, to int) {
	avgSpeed := 5
	total := to - from

	totalTime := total/avgSpeed
	fmt.Printf("Road From:[%d] to B:[%d] Avg speed:[%d] Total:[%d]km Total time:[%d]m\n",
		from, to, avgSpeed, total, totalTime)
}

