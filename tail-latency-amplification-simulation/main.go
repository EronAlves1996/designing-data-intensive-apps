package main

import (
	"fmt"
	"math/rand"
	"slices"
	"sync"
	"time"
)

func queryNode(nodeID int) string {
	response := "OK"
	prop := rand.Float32()
	ms := 10
	if prop <= 0.95 {
		ms = rand.Intn(40) + 10
	} else {
		ms = rand.Intn(800) + 200
	}

	<-time.After(time.Duration(ms) * time.Millisecond)
	fmt.Printf("Node %d responding\n", nodeID)
	return response
}

func clientRequest() {
	var wg sync.WaitGroup
	for i := range 10 {
		n := i
		wg.Go(func() {
			queryNode(n)
		})
	}
	wg.Wait()
	fmt.Println("Client request returned")
}

func clientRequestImproved(successThreshold int) {
	var wg sync.WaitGroup
	wg.Add(successThreshold)
	for i := range 10 {
		go func(nodeID int) {
			queryNode(nodeID)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("Client request returned")
}

func main() {
	times := make([]int, 0, 100)
	for i := range 100 {
		start := time.Now()
		clientRequest()
		elapsed := time.Since(start)
		fmt.Printf("Request %d took %d ms to complete\n", i, elapsed.Milliseconds())
		times = append(times, int(elapsed.Milliseconds()))
	}
	slices.Sort(times)
	total := 0
	p95 := times[94]
	for _, a := range times {
		total += a
	}
	avg := int(total) / 100
	fmt.Printf("The avg response time is %d ms and the p95 is %d ms\n", avg, p95)
}
