package tree

import "testing"

func createBaseCase() Node[int] {
	return NewNode(1, func(a, b int) int {
		return a - b
	})
}

func TestInsertSequential(t *testing.T) {
	n := createBaseCase()

	n.Insert(2)
	n.Insert(3)

	AssertEquals(n.Value, 1, t)
	AssertEquals(n.Right.Value, 2, t)
	AssertEquals(n.Right.Right.Value, 3, t)
}

func AssertEquals(a, b int, t *testing.T) {
	if a != b {
		t.Errorf("Want %d, found %d", a, b)
	}
}

func TestInsertBalanced(t *testing.T) {
	n := createBaseCase()
	n.Insert(-1)
	n.Insert(2)

	AssertEquals(n.Value, 1, t)
	AssertEquals(n.Left.Value, -1, t)
	AssertEquals(n.Right.Value, 2, t)
}
