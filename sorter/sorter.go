package sorter

import "push-swap/stack"

func do(s *stack.State, ops []string, instr string) []string {
	stack.Apply(s, instr)
	return append(ops, instr)
}

func Sort(nums []int) []string {
	a := make([]int, len(nums))
	copy(a, nums)
	s := &stack.State{A: a}

	if stack.IsSorted(a) {
		return nil
	}
	if len(a) == 2 {
		return do(s, nil, "sa")
	}
	if len(a) == 3 {
		return sortThree(s, nil)
	}
	return selectionSort(s)
}
