package stack

func Apply(s *State, instr string) bool {
	switch instr {
	case "sa":
		Swap(s.A)
	case "sb":
		Swap(s.B)
	case "ss":
		Swap(s.A)
		Swap(s.B)
	case "pa":
		s.B, s.A = Push(s.B, s.A)
	case "pb":
		s.A, s.B = Push(s.A, s.B)
	case "ra":
		Rotate(s.A)
	case "rb":
		Rotate(s.B)
	case "rr":
		Rotate(s.A)
		Rotate(s.B)
	case "rra":
		ReverseRotate(s.A)
	case "rrb":
		ReverseRotate(s.B)
	case "rrr":
		ReverseRotate(s.A)
		ReverseRotate(s.B)
	default:
		return false
	}
	return true
}

func Solved(s *State) bool {
	return len(s.B) == 0 && IsSorted(s.A)
}
