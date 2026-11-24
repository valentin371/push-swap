package main

import (
	"fmt"
	"os"
	"strconv"

	"push-swap/sorter"
	"push-swap/stack"
)

func readArg(index int, def int) int {
	if len(os.Args) > index {
		n, err := strconv.Atoi(os.Args[index])
		if err == nil && n > 0 {
			return n
		}
	}
	return def
}

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

func main() {
	size := readArg(1, 100)
	runs := readArg(2, 10)
	seed := readArg(3, 1)

	success := 0
	total := 0
	lowest := 0
	highest := 0

	for i := 0; i < runs; i++ {
		nums := shuffled(size, seed*1000+i+1)
		ops := sorter.Sort(nums)

		s := &stack.State{A: append([]int{}, nums...)}
		for _, op := range ops {
			stack.Apply(s, op)
		}
		if stack.Solved(s) {
			success++
		}

		count := len(ops)
		total += count
		if i == 0 || count < lowest {
			lowest = count
		}
		if i == 0 || count > highest {
			highest = count
		}
		if size <= 20 {
			fmt.Println("liste :", nums, "->", count, "instructions")
		}
	}

	fmt.Println("nombres par liste :", size)
	fmt.Println("listes testées    :", runs)
	fmt.Printf("triées (OK)       : %d/%d (%d%%)\n", success, runs, success*100/runs)
	fmt.Println("instructions min  :", lowest)
	fmt.Println("instructions moy  :", total/runs)
	fmt.Println("instructions max  :", highest)
}
