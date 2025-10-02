package main

import (
	"math/rand"
	"time"
)

func queryNode(_ int) string {
	response := "OK"
	prop := rand.Float32()
	ms := 10
	if prop <= 0.95 {
		ms = rand.Intn(40) + 10
	} else {
		ms = rand.Intn(800) + 200
	}

	<-time.After(time.Duration(ms) * time.Millisecond)
	return response
}
