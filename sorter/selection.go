package sorter

import "push-swap/stack"

func selectionSort(s *stack.State) []string {
	var ops []string
	for len(s.A) > 3 {
		ops = moveMinToTop(s, ops)
		ops = do(s, ops, "pb")
	}
	ops = sortThree(s, ops)
	for len(s.B) > 0 {
		ops = do(s, ops, "pa")
	}
	return ops
}

func moveMinToTop(s *stack.State, ops []string) []string {
	size := len(s.A)

	minIndex := 0
	for i := 1; i < size; i++ {
		if s.A[i] < s.A[minIndex] {
			minIndex = i
		}
	}

	if minIndex <= size/2 {
		for i := 0; i < minIndex; i++ {
			ops = do(s, ops, "ra")
		}
	} else {
		for i := 0; i < size-minIndex; i++ {
			ops = do(s, ops, "rra")
		}
	}
	return ops
}
