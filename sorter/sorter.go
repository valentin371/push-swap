package sorter

import "push-swap/stack"

const smallLimit = 13

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

	var ops []string
	if len(a) == 2 {
		ops = do(s, ops, "sa")
	} else if len(a) == 3 {
		ops = sortThree(s, ops)
	} else if len(a) <= smallLimit {
		ops = selectionSort(s)
	} else {
		s.A = toRanks(a)
		ops = chunkSort(s)
	}
	return simplify(ops)
}

func simplify(ops []string) []string {
	var result []string
	for _, op := range ops {
		n := len(result)
		if n > 0 {
			last := result[n-1]
			if cancel(last, op) {
				result = result[:n-1]
				continue
			}
			if merged := merge(last, op); merged != "" {
				result[n-1] = merged
				continue
			}
		}
		result = append(result, op)
	}
	return result
}

func cancel(a, b string) bool {
	switch a + " " + b {
	case "ra rra", "rra ra", "rb rrb", "rrb rb", "pa pb", "pb pa":
		return true
	}
	return false
}

func merge(a, b string) string {
	switch a + " " + b {
	case "ra rb", "rb ra":
		return "rr"
	case "rra rrb", "rrb rra":
		return "rrr"
	}
	return ""
}
