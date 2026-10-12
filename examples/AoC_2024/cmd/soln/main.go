package main

import (
	"fmt"
	"os"

	"github.com/SegniAT/monkey-language-interpreter/examples/AoC_2024/go_soln"
)

func main() {
	path := "input.txt"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	col1, col2 := soln.Parse(path)
	fmt.Println(soln.Part1(col1, col2))
	fmt.Println(soln.Part2(col1, col2))
}