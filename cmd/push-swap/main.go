package main

import (
	"fmt"
	"os"

	"push-swap/sorter"
	"push-swap/stack"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	nums, ok := stack.Parse(os.Args[1:])
	if !ok {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	for _, op := range sorter.Sort(nums) {
		fmt.Println(op)
	}
}
