package tree

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

func (n *Node[T]) Insert(v T) {
	if n.compare(v, n.Value) < 0 {
		if n.Left == nil {
			nn := NewNode(v, n.compare)
			n.Left = &nn
			n.balanceFactor = -1
			return
		}

		n.Left.Insert(v)
		n.balanceFactor += n.Left.balanceFactor
		return
	}

	if n.Right == nil {
		nn := NewNode(v, n.compare)
		n.Right = &nn
		n.balanceFactor += 1
		return
	}

	n.Right.Insert(v)
	n.balanceFactor += n.Right.balanceFactor
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
