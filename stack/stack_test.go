package stack

import "testing"

func equal(a, b []int) bool {
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

func TestSwap(t *testing.T) {
	s := []int{1, 2, 3}
	Swap(s)
	if !equal(s, []int{2, 1, 3}) {
		t.Errorf("got %v", s)
	}
	Swap([]int{})
	Swap([]int{1})
}

func TestRotate(t *testing.T) {
	s := []int{1, 2, 3}
	Rotate(s)
	if !equal(s, []int{2, 3, 1}) {
		t.Errorf("rotate: got %v", s)
	}
	ReverseRotate(s)
	if !equal(s, []int{1, 2, 3}) {
		t.Errorf("reverse rotate: got %v", s)
	}
}

func TestPush(t *testing.T) {
	a, b := Push([]int{1, 2}, []int{9})
	if !equal(a, []int{2}) || !equal(b, []int{1, 9}) {
		t.Errorf("got %v %v", a, b)
	}
	a, b = Push([]int{}, []int{9})
	if len(a) != 0 || !equal(b, []int{9}) {
		t.Errorf("got %v %v", a, b)
	}
}

func TestSubjectExample(t *testing.T) {
	s := &State{A: []int{2, 1, 3, 6, 8, 5}}
	for _, instr := range []string{"sa", "pb", "pb", "pb", "rb", "rrr", "pa", "pa", "pa"} {
		if !Apply(s, instr) {
			t.Fatalf("unknown instruction %s", instr)
		}
	}
	if !equal(s.A, []int{1, 2, 3, 5, 6, 8}) || !Solved(s) {
		t.Errorf("got a=%v b=%v", s.A, s.B)
	}
}

func TestUnknownInstruction(t *testing.T) {
	if Apply(&State{}, "xx") {
		t.Error("expected false")
	}
}

func TestParseValid(t *testing.T) {
	got, ok := Parse([]string{"2 1 3"})
	if !ok || !equal(got, []int{2, 1, 3}) {
		t.Errorf("got %v %v", got, ok)
	}
	got, ok = Parse([]string{"2", "-1", "3"})
	if !ok || !equal(got, []int{2, -1, 3}) {
		t.Errorf("got %v %v", got, ok)
	}
}

func TestParseErrors(t *testing.T) {
	bad := [][]string{
		{"0 one 2 3"},
		{"1 2 1"},
		{"1.5"},
	}
	for _, args := range bad {
		if _, ok := Parse(args); ok {
			t.Errorf("%v: expected an error", args)
		}
	}
}
