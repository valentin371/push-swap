package stack

type State struct {
	A []int
	B []int
}

func IsSorted(s []int) bool {
	for i := 1; i < len(s); i++ {
		if s[i-1] > s[i] {
			return false
		}
	}
	return true
}

func Swap(s []int) {
	if len(s) < 2 {
		return
	}
	s[0], s[1] = s[1], s[0]
}

func Rotate(s []int) {
	if len(s) < 2 {
		return
	}
	first := s[0]
	for i := 0; i < len(s)-1; i++ {
		s[i] = s[i+1]
	}
	s[len(s)-1] = first
}

func ReverseRotate(s []int) {
	if len(s) < 2 {
		return
	}
	last := s[len(s)-1]
	for i := len(s) - 1; i > 0; i-- {
		s[i] = s[i-1]
	}
	s[0] = last
}

func Push(from, to []int) ([]int, []int) {
	if len(from) == 0 {
		return from, to
	}
	newTo := []int{from[0]}
	newTo = append(newTo, to...)
	return from[1:], newTo
}
