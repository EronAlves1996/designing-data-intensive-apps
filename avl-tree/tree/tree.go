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

func rotateLeft[T any](n *Node[T]) *Node[T] {
	actual := n
	child := actual.Right
	actual.Right = nil
	child.Left = actual
	return child
}

func Insert[T any](n *Node[T], v T) *Node[T] {
	if n.compare(v, n.Value) < 0 {
		if n.Left == nil {
			nn := NewNode(v, n.compare)
			n.Left = &nn
			n.balanceFactor = -1
			return n
		}

		Insert(n.Left, v)
		n.balanceFactor += n.Left.balanceFactor
		return n
	}

	if n.Right == nil {
		nn := NewNode(v, n.compare)
		n.Right = &nn
		n.balanceFactor += 1
		return n
	}

	Insert(n.Right, v)
	n.balanceFactor += n.Right.balanceFactor
	if n.balanceFactor > 1 {
		r := rotateLeft(n)
		r.balanceFactor = reviewBalance(r)
		return r
	}
	return n
}

func reviewBalance[T any](r *Node[T]) int {
	le := 0
	ri := 0
	if r.Right != nil {
		ri = reviewBalance(r.Right) + 1
	}

	if r.Left != nil {
		le = reviewBalance(r.Left) + 1
	}

	return ri - le
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
