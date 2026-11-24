package sorter

import "push-swap/stack"

func toRanks(nums []int) []int {
	ranks := make([]int, len(nums))
	for i, x := range nums {
		for _, y := range nums {
			if y < x {
				ranks[i]++
			}
		}
	}
	return ranks
}

func intSqrt(n int) int {
	r := 0
	for (r+1)*(r+1) <= n {
		r++
	}
	return r
}

func chunkSort(s *stack.State) []string {
	var ops []string
	chunkSize := intSqrt(len(s.A) * 9 / 4)
	pushed := 0

	for len(s.A) > 0 {
		rank := s.A[0]
		if rank <= pushed {
			ops = do(s, ops, "pb")
			ops = do(s, ops, "rb")
			pushed++
		} else if rank <= pushed+chunkSize {
			ops = do(s, ops, "pb")
			pushed++
		} else {
			ops = do(s, ops, "ra")
		}
	}

	for len(s.B) > 0 {
		ops = moveMaxToTopOfB(s, ops)
		ops = do(s, ops, "pa")
	}
	return ops
}

func moveMaxToTopOfB(s *stack.State, ops []string) []string {
	size := len(s.B)

	maxIndex := 0
	for i := 1; i < size; i++ {
		if s.B[i] > s.B[maxIndex] {
			maxIndex = i
		}
	}

	if maxIndex <= size/2 {
		for i := 0; i < maxIndex; i++ {
			ops = do(s, ops, "rb")
		}
	} else {
		for i := 0; i < size-maxIndex; i++ {
			ops = do(s, ops, "rrb")
		}
	}
	return ops
}
