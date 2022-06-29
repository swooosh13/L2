package main

import "fmt"

type PublicTransportStrategy struct {}

func (r *PublicTransportStrategy) buildRoute(from int, to int) {
	avgSpeed := 30
	trafficJam := 1
	total := to - from

	totalTime := total/avgSpeed + total/(avgSpeed*trafficJam)
	fmt.Printf("Road From:[%d] to B:[%d] Avg speed:[%d] Traffic jam:[%d] Total:[%d]km Total time:[%d]m\n",
		from, to, avgSpeed, trafficJam, total, totalTime)
}

