package sorter

import (
	"testing"

	"push-swap/stack"
)

func shuffled(size, seed int) []int {
	nums := make([]int, size)
	for i := range nums {
		nums[i] = i
	}
	x := seed
	for i := size - 1; i > 0; i-- {
		x = (x*1103515245 + 12345) % 2147483648
		j := (x / 65536) % (i + 1)
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}

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

func equalOps(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
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

func TestShuffled(t *testing.T) {
	for n := 0; n <= 80; n++ {
		for k := 0; k < 20; k++ {
			nums := shuffled(n, n*100+k+1)
			for i := range nums {
				nums[i] -= n / 2
			}
			check(t, nums)
		}
	}
}

func TestAlreadySorted(t *testing.T) {
	if ops := Sort([]int{1, 2, 3, 4}); len(ops) != 0 {
		t.Errorf("expected no instructions, got %v", ops)
	}
}

func TestSimplify(t *testing.T) {
	cases := []struct {
		in   []string
		want []string
	}{
		{[]string{"ra", "rb"}, []string{"rr"}},
		{[]string{"rb", "ra"}, []string{"rr"}},
		{[]string{"rrb", "rra"}, []string{"rrr"}},
		{[]string{"ra", "rra"}, nil},
		{[]string{"pb", "ra", "pa"}, []string{"pb", "ra", "pa"}},
	}
	for _, c := range cases {
		if got := simplify(c.in); !equalOps(got, c.want) {
			t.Errorf("simplify(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestLimits(t *testing.T) {
	n100 := check(t, shuffled(100, 1))
	t.Logf("100 numbers: %d instructions", n100)
	if n100 > 700 {
		t.Errorf("100 numbers: %d instructions, expected at most 700", n100)
	}

	n500 := check(t, shuffled(500, 1))
	t.Logf("500 numbers: %d instructions", n500)
	if n500 > 5500 {
		t.Errorf("500 numbers: %d instructions, expected at most 5500", n500)
	}
}
