package main

import (
	"ai/unsupervised/puzzle"
	"fmt"
)

func main() {
	problem := puzzle.ReadPuzzle("puzzle.txt")
	solution := puzzle.DephFirst(problem)
	maxLength := 20
	if len(solution) > maxLength {
		fmt.Printf("Solution was %v long reducing to last %v moves", len(solution), maxLength)
		solution = solution[len(solution)-maxLength:]

	}
	for _, val := range solution {
		fmt.Println(val.ToString())
	}
}
