package main

import (
	"ai/unsupervised/puzzle"
	"fmt"
	"os"
	"runtime/pprof"
)

func main() {
	f, err := os.Create("profile.prof")
	if err != nil {
		panic(err)
	}
	pprof.StartCPUProfile(f)

	problem := puzzle.ReadPuzzle("puzzle.txt")
	var solution []puzzle.Puzzle
	solution = puzzle.BreadthFirst(problem)
	solution = puzzle.DephFirst(problem)
	pprof.StopCPUProfile()
	maxLength := 20
	if len(solution) > maxLength {
		fmt.Printf("Solution was %v long reducing to last %v moves", len(solution), maxLength)
		solution = solution[len(solution)-maxLength:]

	}
	for _, val := range solution {
		fmt.Println(val.ToString())
	}
}
