package tree

type Node[T any] struct {
	Value   T
	height  int
	Left    *Node[T]
	Right   *Node[T]
	compare func(T, T) int
}

func NewNode[T any](value T, compare func(T, T) int) Node[T] {
	return Node[T]{
		Value:   value,
		compare: compare,
		height:  1,
	}
}

func rotateLeft[T any](n *Node[T]) *Node[T] {
	actual := n
	child := actual.Right
	actual.Right = nil
	if child.Left == nil {
		child.Left = actual
	} else {
		if child.Left.height > 1 {
			child = rotateRight(child)
		}
		child.Left.Left = actual
		child.Left = rotateRight(child.Left)
	}
	child.Left = actual
	actual.height = height(actual)
	child.height = height(child)
	return child
}

func rotateRight[T any](n *Node[T]) *Node[T] {
	actual := n
	child := actual.Left
	actual.Left = nil
	if child.Right == nil {
		child.Right = actual
	} else {
		if child.Right.height > 1 {
			child = rotateLeft(child)
		}
		child.Right.Right = actual
		child.Right = rotateLeft(child.Right)
	}
	actual.height = height(actual)
	child.height = height(child)
	return child
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rotate[T any](n *Node[T]) *Node[T] {
	if n == nil {
		return n
	}
	diff := (height(n.Right) - height(n.Left))
	if diff > 1 {
		r := rotateLeft(n)
		return r
	}
	if diff < -1 {
		r := rotateRight(n)
		return r
	}
	return n
}

func Insert[T any](n *Node[T], v T) *Node[T] {
	if n.compare(v, n.Value) < 0 {
		if n.Left == nil {
			nn := NewNode(v, n.compare)
			n.Left = &nn
			n.height = height(n)
			return n
		}

		n.Left = Insert(n.Left, v)
		n.height = height(n)
		return rotate(n)
	}

	if n.Right == nil {
		nn := NewNode(v, n.compare)
		n.Right = &nn
		n.height = height(n)
		return n
	}

	n.Right = Insert(n.Right, v)
	n.height = height(n)

	return rotate(n)
}

func height[T any](n *Node[T]) int {
	if n == nil {
		return 0
	}

	if n.Left == nil && n.Right == nil {
		return 1
	}

	if n.Right == nil {
		return n.Left.height + 1
	}

	if n.Left == nil {
		return n.Right.height + 1
	}

	return max(n.Right.height, n.Left.height) + 1
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
