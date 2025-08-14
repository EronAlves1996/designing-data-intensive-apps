package tree

import "fmt"

type Node[T any] struct {
	Value         T
	balanceFactor int
	Left          *Node[T]
	Right         *Node[T]
	compare       func(T, T) int
}

func NewNode[T any](value T, compare func(T, T) int) Node[T] {
	return Node[T]{
		Value:   value,
		compare: compare,
	}
}

func (n *Node[T]) InsertAux(v T, h int) int {
	if n.compare(v, n.Value) < 0 {
		if n.Left == nil {
			nn := NewNode(v, n.compare)
			n.Left = &nn
			n.balanceFactor -= 1
			return n.balanceFactor
		}

		n.balanceFactor += n.Left.InsertAux(v, h-1)
		return n.balanceFactor
	}

	if n.Right == nil {
		nn := NewNode(v, n.compare)
		n.Right = &nn
		n.balanceFactor += 1
		return n.balanceFactor
	}

	n.balanceFactor += n.Right.InsertAux(v, h+1)
	return n.balanceFactor
}

func (n *Node[T]) DebugBalance() {
	fmt.Println(n.balanceFactor)
}

func (n *Node[T]) Insert(v T) {
	n.InsertAux(v, 0)
}

func (n *Node[T]) Remove(v T) *T {
	if n.compare(n.Value, v) == 0 {
		return &n.Value
	}

	if n.compare(v, n.Value) < 0 {
		if n.Left == nil {
			return nil
		}

		r := n.Left.Remove(v)
		if r != nil {
			n.Left = nil
			return nil
		}
	}

	if n.Right == nil {
		return nil
	}

	r := n.Right.Remove(v)
	if r != nil {
		n.Right = nil
	}

	return nil
}

func (n *Node[T]) Ordered() []T {
	r := make([]T, 0)

	if n.Left != nil {
		r = append(r, n.Left.Ordered()...)
	}

	r = append(r, n.Value)

	if n.Right != nil {
		r = append(r, n.Right.Ordered()...)
	}

	return r
}
