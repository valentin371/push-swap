package main

import (
	"fmt"
	"os"

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
	s := &stack.State{A: nums}

	for {
		var instr string
		_, err := fmt.Scan(&instr)
		if err != nil {
			break
		}
		if !stack.Apply(s, instr) {
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
	}

	if stack.Solved(s) {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
