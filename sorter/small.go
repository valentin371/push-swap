package sorter

import "push-swap/stack"

func sortThree(s *stack.State, ops []string) []string {
	x, y, z := s.A[0], s.A[1], s.A[2]

	if x > y && y < z && x < z {
		ops = do(s, ops, "sa")
	} else if x > y && y > z {
		ops = do(s, ops, "sa")
		ops = do(s, ops, "rra")
	} else if x > y && y < z && x > z {
		ops = do(s, ops, "ra")
	} else if x < y && y > z && x < z {
		ops = do(s, ops, "sa")
		ops = do(s, ops, "ra")
	} else if x < y && y > z && x > z {
		ops = do(s, ops, "rra")
	}
	return ops
}
