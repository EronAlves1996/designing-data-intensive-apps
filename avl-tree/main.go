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
	t.DebugBalance()
	t.Insert(40)
	t.DebugBalance()
	t.Insert(5)
	t.DebugBalance()
	t.Insert(-5)
	t.DebugBalance()
	t.Insert(100)
	t.DebugBalance()

	fmt.Println(t.Ordered())
}
