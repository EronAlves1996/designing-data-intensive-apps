package tree

type Node[T comparable] struct {
	value  T
	height int
	left   *Node[T]
	right  *Node[T]
}
