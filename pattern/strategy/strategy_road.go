package main

import "fmt"

// конкретная стратегия
type RoadStrategy struct {}

func (r *RoadStrategy) buildRoute(from int, to int) {
	avgSpeed := 40
	trafficJam := 2
	total := to - from

	totalTime := total/avgSpeed + total/(avgSpeed*trafficJam)
	fmt.Printf("Road From:[%d] to B:[%d] Avg speed:[%d] Traffic jam:[%d] Total:[%d]km Total time:[%d]m\n",
		from, to, avgSpeed, trafficJam, total, totalTime)
}
