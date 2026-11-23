package sorter

import (
	"math/rand"
	"testing"

	"push-swap/stack"
)

func check(t *testing.T, nums []int) int {
	ops := Sort(nums)

	a := make([]int, len(nums))
	copy(a, nums)
	s := &stack.State{A: a}
	for _, op := range ops {
		if !stack.Apply(s, op) {
			t.Fatalf("invalid instruction: %s", op)
		}
	}
	if !stack.Solved(s) {
		t.Fatalf("not sorted for %v: a=%v b=%v", nums, s.A, s.B)
	}
	return len(ops)
}

func TestFixedCases(t *testing.T) {
	cases := [][]int{
		{},
		{1},
		{1, 2},
		{2, 1},
		{3, 2, 1},
		{2, 3, 1},
		{4, 67, 3, 87, 23},
		{2, 1, 3, 6, 5, 8},
		{-5, 10, 0, -1},
	}
	for _, c := range cases {
		check(t, c)
	}
}

func TestRandom(t *testing.T) {
	for n := 0; n <= 50; n++ {
		for k := 0; k < 20; k++ {
			check(t, rand.Perm(n))
		}
	}
}

func TestAlreadySorted(t *testing.T) {
	if ops := Sort([]int{1, 2, 3, 4}); len(ops) != 0 {
		t.Errorf("expected no instructions, got %v", ops)
	}
}

func TestCounts(t *testing.T) {
	t.Logf("100 numbers: %d instructions", check(t, rand.Perm(100)))
}
