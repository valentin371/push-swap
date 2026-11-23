package stack

import (
	"strconv"
	"strings"
)

func Parse(args []string) ([]int, bool) {
	var nums []int
	for _, arg := range args {
		for _, word := range strings.Fields(arg) {
			n, err := strconv.Atoi(word)
			if err != nil {
				return nil, false
			}
			for _, x := range nums {
				if x == n {
					return nil, false
				}
			}
			nums = append(nums, n)
		}
	}
	return nums, true
}
