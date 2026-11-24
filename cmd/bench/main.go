package main

import (
	"fmt"
	"math/rand"
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

func main() {
	size := readArg(1, 100)
	runs := readArg(2, 10)

	success := 0
	total := 0
	lowest := 0
	highest := 0

	for i := 0; i < runs; i++ {
		nums := rand.Perm(size)
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
