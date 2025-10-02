package main

import (
	"fmt"
	"math/rand"
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
