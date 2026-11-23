package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if !stack.Apply(s, line) {
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
