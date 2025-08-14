package tree

import "testing"

func createBaseCase() Node[int] {
	return NewNode(1, func(a, b int) int {
		return a - b
	})
}

func TestInsertSequential(t *testing.T) {
	n := createBaseCase()
	a := &n
	AssertEquals(a.balanceFactor, 0, t)

	a = Insert(a, 2)
	AssertEquals(a.balanceFactor, 1, t)

	a = Insert(a, 3)
	AssertEquals(a.balanceFactor, 0, t)

	AssertEquals(a.Value, 2, t)
	AssertEquals(a.Right.Value, 3, t)
	AssertEquals(a.Left.Value, 1, t)
}

func AssertEquals(a, b int, t *testing.T) {
	if a != b {
		t.Fatalf("Want %d, found %d", b, a)
	}
}

// func TestInsertBalanced(t *testing.T) {
// 	n := createBaseCase()
// 	AssertEquals(n.balanceFactor, 0, t)

// 	n.Insert(-1)
// 	AssertEquals(n.balanceFactor, -1, t)

// 	n.Insert(2)
// 	AssertEquals(n.balanceFactor, 0, t)

// 	AssertEquals(n.Value, 1, t)
// 	AssertEquals(n.Left.Value, -1, t)
// 	AssertEquals(n.Right.Value, 2, t)
// }
