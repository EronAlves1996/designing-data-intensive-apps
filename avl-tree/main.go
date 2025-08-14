package main

import (
	"fmt"

	"github.com/EronAlves1996/designing-data-intensive-apps/avl-tree/tree"
)

func main() {
	t := tree.NewNode(10, func(a, b int) int {
		return a - b
	})

	t.Insert(15)
	t.Insert(40)
	t.Insert(5)
	t.Insert(-5)
	t.Insert(100)

	fmt.Println(t.Ordered())
}
