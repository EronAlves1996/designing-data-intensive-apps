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
	if a.height != 1 {
		t.Fatalf("Expect %d found %d", 1, a.height)
	}

	a = Insert(a, 2)

	if a.height != 2 {
		t.Fatalf("Expect %d found %d", 2, a.height)
	}

	a = Insert(a, 3)

	if a.height != 2 {
		t.Fatalf("Expect %d found %d", 2, a.height)
	}

	if a.Value != 2 {
		t.Fatalf("Expect %d found %d", 2, a.Value)
	}

	if a.Right.Value != 3 {
		t.Fatalf("Expect %d found %d", 3, a.Right.Value)
	}

	if a.Left.Value != 1 {
		t.Fatalf("Expect %d found %d", 1, a.Left.Value)
	}
}

func TestInsertSequentialBackwards(t *testing.T) {
	n := createBaseCase()
	a := &n

	if a.height != 1 {
		t.Fatalf("Expect %d found %d", 1, a.height)
	}

	a = Insert(a, 0)

	if a.height != 2 {
		t.Fatalf("Expect %d found %d", 2, a.height)
	}

	a = Insert(a, -1)

	if a.height != 2 {
		t.Fatalf("Expect %d found %d", 2, a.height)
	}

	expect := 0
	if a.Value != expect {
		t.Fatalf("Expect %d found %d", expect, a.Value)
	}

	expect = 1
	if a.Right.Value != expect {
		t.Fatalf("Expect %d found %d", expect, a.Right.Value)
	}

	expect = -1
	if a.Left.Value != expect {
		t.Fatalf("Expect %d found %d", expect, a.Left.Value)
	}
}

func TestInsertSequentialTwoLevels(t *testing.T) {
	n := createBaseCase()
	a := &n
	if a.height != 1 {
		t.Fatalf("Expect %d found %d", 1, a.height)
	}

	a = Insert(a, 2)

	if a.height != 2 {
		t.Fatalf("Expect %d found %d", 2, a.height)
	}

	a = Insert(a, 3)

	if a.height != 2 {
		t.Fatalf("Expect %d found %d", 2, a.height)
	}

	a = Insert(a, 4)

	if a.height != 3 {
		t.Fatalf("Expect %d found %d", 3, a.height)
	}

	if a.Value != 2 {
		t.Fatalf("Expect %d found %d", 2, a.Value)
	}

	if a.Right.Value != 3 {
		t.Fatalf("Expect %d found %d", 3, a.Right.Value)
	}

	if a.Right.Right.Value != 4 {
		t.Fatalf("Expect %d found %d", 4, a.Right.Value)
	}

	if a.Left.Value != 1 {
		t.Fatalf("Expect %d found %d", 1, a.Left.Value)
	}
}

func TestInsertNotSequential(t *testing.T) {
	n := createBaseCase()
	a := &n
	a = Insert(a, 2)

	if a.Value != 1 {
		t.Fatalf("Expect %d found %d", 1, a.Value)
	}

	if a.Right.Value != 2 {
		t.Fatalf("Expect %d found %d", 2, a.Right.Value)
	}

	a = Insert(a, -2)

	if a.Value != 1 {
		t.Fatalf("Expect %d found %d", 1, a.Value)
	}

	if a.Right.Value != 2 {
		t.Fatalf("Expect %d found %d", 2, a.Right.Value)
	}

	if a.Left.Value != -2 {
		t.Fatalf("Expect %d found %d", -2, a.Right.Value)
	}

	a = Insert(a, -5)

	if a.Value != 1 {
		t.Fatalf("Expect %d found %d", 1, a.Value)
	}

	if a.Right.Value != 2 {
		t.Fatalf("Expect %d found %d", 2, a.Right.Value)
	}

	if a.Left.Value != -2 {
		t.Fatalf("Expect %d found %d", -2, a.Right.Value)
	}

	if a.Left.Left.Value != -5 {
		t.Fatalf("Expect %d found %d", -5, a.Right.Value)
	}

	a = Insert(a, -6)

	if a.Value != 1 {
		t.Fatalf("Expect %d found %d", 1, a.Value)
	}

	if a.Left.Value != -5 {
		t.Fatalf("Expect %d found %d", -5, a.Left.Value)
	}

	if a.Left.Left.Value != -6 {
		t.Fatalf("Expect %d found %d", -6, a.Left.Left.Value)
	}

	if a.Left.Right.Value != -2 {
		t.Fatalf("Expect %d found %d", -2, a.Left.Right.Value)
	}

	if a.Right.Value != 2 {
		t.Fatalf("Expect %d found %d", 2, a.Right.Value)
	}
}

func TestInsertNotSequentialWithHardDisbalancing(t *testing.T) {
	n := createBaseCase()
	a := &n
	a = Insert(a, 2)
	a = Insert(a, -2)
	a = Insert(a, -5)
	a = Insert(a, -6)
	a = Insert(a, -1)

	if a.Value != -2 {
		t.Fatalf("Expect %d found %d", -2, a.Value)
	}

	if a.Right.Value != 1 {
		t.Fatalf("Expect %d found %d", 1, a.Value)
	}

	if a.Right.Right.Value != 2 {
		t.Fatalf("Expect %d found %d", 2, a.Value)
	}

	if a.Right.Left.Value != -1 {
		t.Fatalf("Expect %d found %d", -1, a.Value)
	}

	if a.Left.Value != -5 {
		t.Fatalf("Expect %d found %d", -5, a.Value)
	}

	if a.Left.Left.Value != -6 {
		t.Fatalf("Expect %d found %d", -6, a.Value)
	}
}

func TestInsertNotSequentialWithHardDisbalancing2(t *testing.T) {
	n := createBaseCase()
	a := &n
	a = Insert(a, 2)
	a = Insert(a, -2)
	a = Insert(a, -5)
	a = Insert(a, -6)
	a = Insert(a, -7)

	if a.Value != -5 {
		t.Fatalf("Expect %d found %d", -5, a.Value)
	}

	if a.Right.Value != 1 {
		t.Fatalf("Expect %d found %d", 1, a.Value)
	}

	if a.Right.Right.Value != 2 {
		t.Fatalf("Expect %d found %d", 2, a.Value)
	}

	if a.Right.Left.Value != -2 {
		t.Fatalf("Expect %d found %d", -2, a.Value)
	}

	if a.Left.Value != -6 {
		t.Fatalf("Expect %d found %d", -6, a.Value)
	}

	if a.Left.Left.Value != -7 {
		t.Fatalf("Expect %d found %d", -7, a.Value)
	}
}

// func TestInsertBalanced(t *testing.T) {
// 	n := createBaseCase()
// 	r := &n
// 	AssertEquals(r.balanceFactor, 0, t)

// 	r = Insert(r, -1)
// 	AssertEquals(r.balanceFactor, -1, t)

// 	r = Insert(r, 2)
// 	AssertEquals(r.balanceFactor, 0, t)

// 	AssertEquals(r.Value, 1, t)
// 	AssertEquals(r.Left.Value, -1, t)
// 	AssertEquals(r.Right.Value, 2, t)
// }
