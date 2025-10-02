package main

import (
	"context"
	"fmt"
	"math/rand"
	"slices"
	"sync"
	"time"
)

func queryNode(ctx context.Context, nodeID int) string {
	response := "OK"
	prop := rand.Float32()
	ms := 10
	if prop <= 0.95 {
		ms = rand.Intn(40) + 10
	} else {
		ms = rand.Intn(800) + 200
	}

	select {
	case <-ctx.Done():
		fmt.Printf("Request for node %d discarded: threshold achieved\n", nodeID)
	case <-time.After(time.Duration(ms) * time.Millisecond):
		fmt.Printf("Node %d responding\n", nodeID)
	}
	return response
}

func clientRequest() {
	var wg sync.WaitGroup
	ctx := context.Background()
	for i := range 10 {
		n := i
		wg.Go(func() {
			queryNode(ctx, n)
		})
	}
	wg.Wait()
	fmt.Println("Client request returned")
}

func clientRequestImproved(successThreshold int) {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	responses := make(chan string, 10)
	for i := range 10 {
		wg.Go(func() {
			select {
			case <-ctx.Done():
				return
			case responses <- queryNode(ctx, i):
			}
		})
	}

	for range successThreshold {
		<-responses
	}
	cancel()
	fmt.Println("Client request returned")

	wg.Wait()
}

func runSimulation(successThreshold int) string {
	times := make([]int, 0, 100)
	for i := range 100 {
		start := time.Now()
		clientRequestImproved(successThreshold)
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
	return fmt.Sprintf("For a successThreshold %d, The avg response time is %d ms and the p95 is %d ms\n", successThreshold, avg, p95)
}

func main() {
	stats9 := runSimulation(9)
	stats7 := runSimulation(7)
	stats8 := runSimulation(8)
	fmt.Print(stats9)
	fmt.Print(stats7)
	fmt.Print(stats8)
}
