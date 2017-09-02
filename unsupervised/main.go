package main

import (
	"ai/unsupervised/puzzle"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	f, err := os.Create("profile.prof")
	if err != nil {
		panic(err)
	}
	pprof.StartCPUProfile(f)

	problem := puzzle.ReadPuzzle("extra_puzzle.txt")
	var solution []puzzle.Puzzle
	start := time.Now()
	solution = puzzle.BreadthFirst(problem)
	fmt.Println("breadth time", time.Since(start))
	fmt.Printf("Solution was %v long\n\n", len(solution))

	start = time.Now()
	solution = puzzle.DephFirst(problem)
	fmt.Println("depth time", time.Since(start))
	fmt.Printf("Solution was %v long\n\n", len(solution))

	start = time.Now()
	solution = puzzle.BookGreedy(problem)
	fmt.Println("book greedy time", time.Since(start))
	fmt.Printf("Solution was %v long\n\n", len(solution))

	start = time.Now()
	solution = puzzle.BookAStar(problem)
	fmt.Println("book a* time", time.Since(start))
	fmt.Printf("Solution was %v long\n\n", len(solution))

	start = time.Now()
	solution = puzzle.MyGreedy(problem)
	fmt.Println("my greedy time", time.Since(start))
	fmt.Printf("Solution was %v long\n\n", len(solution))

	start = time.Now()
	solution = puzzle.MyAStar(problem)
	fmt.Println("my a* time", time.Since(start))
	fmt.Printf("Solution was %v long\n\n", len(solution))

	pprof.StopCPUProfile()
	maxLength := 20
	if len(solution) > maxLength {
		fmt.Printf("Solution was %v long reducing to last %v moves\n", len(solution), maxLength)
		solution = solution[len(solution)-maxLength:]

	}
	for _, val := range solution {
		fmt.Println(val.ToString())
	}
}
