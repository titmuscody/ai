package puzzle

import (
	"io/ioutil"
	"regexp"
	"strconv"
)

type Puzzle struct {
	state []int
	empty int
}

func ReadPuzzle(filename string) Puzzle {
	// initialize puzzle
	puzzle := Puzzle{}
	puzzle.state = make([]int, 9)
	// read in file and remove white space
	fileText, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	re := regexp.MustCompile("\n|\r")
	text := re.ReplaceAllString(string(fileText), "")
	// read the integers into the puzzle
	for i := 0; i < len(puzzle.state); i += 1 {
		if string(text[i]) == "_" {
			puzzle.state[i] = 9
			puzzle.empty = i
		} else {
			val, err := strconv.Atoi(string(text[i]))
			if err != nil {
				panic(err)
			}
			puzzle.state[i] = val
		}
	}
	return puzzle
}

func (p *Puzzle) Clone() Puzzle {
	puz := Puzzle{empty: p.empty}
	puz.state = make([]int, 9)
	copy(puz.state, p.state)
	return puz
}

func (p *Puzzle) Hash() int {
	hash := 0
	place := 1
	for _, val := range p.state {
		hash += val * place
		place *= 10
	}
	return hash
}

func (p *Puzzle) Equals(p2 Puzzle) bool {
	for i := range p.state {
		if p.state[i] != p2.state[i] {
			return false
		}
	}
	return true
}

func (p *Puzzle) Solved() bool {
	last := 0
	for i := 0; i < len(p.state); i += 1 {
		if p.state[i] <= last {
			return false
		}
		last = p.state[i]
	}
	return true
}

func (p *Puzzle) GetPossibleMoves() []Puzzle {
	moves := []Puzzle{}
	// check if you can go down
	if p.empty < 6 {
		newMove := p.Clone()
		newMove.state[p.empty] = newMove.state[p.empty+3]
		newMove.state[p.empty+3] = 9
		newMove.empty += 3
		moves = append(moves, newMove)
	}
	// check if you can go up
	if p.empty > 2 {
		newMove := p.Clone()
		newMove.state[p.empty] = newMove.state[p.empty-3]
		newMove.state[p.empty-3] = 9
		newMove.empty -= 3
		moves = append(moves, newMove)
	}
	// check if you can go left
	if p.empty%3 != 0 {
		newMove := p.Clone()
		newMove.state[p.empty] = newMove.state[p.empty-1]
		newMove.state[p.empty-1] = 9
		newMove.empty -= 1
		moves = append(moves, newMove)
	}
	// check if you can go right
	if p.empty%3 != 2 {
		newMove := p.Clone()
		newMove.state[p.empty] = newMove.state[p.empty+1]
		newMove.state[p.empty+1] = 9
		newMove.empty += 1
		moves = append(moves, newMove)
	}
	return moves
}

func (p *Puzzle) ToString() string {
	str := ""
	for i := 0; i < len(p.state); i += 1 {
		if i%3 == 0 {
			str += "\n"
		}
		if p.state[i] == 9 {
			str += "_"
		} else {
			str += strconv.Itoa(p.state[i])
		}
	}
	return str
}
