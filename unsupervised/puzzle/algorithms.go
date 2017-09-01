package puzzle

import (
	"fmt"
)

type Node struct {
	Puzzle Puzzle
	deph   int
	parent *Node
}

func BreadthFirst(puzzle Puzzle) []Puzzle {
	return Search(puzzle, NewStack())
}

func DephFirst(puzzle Puzzle) []Puzzle {
	return Search(puzzle, NewQueue())
}

func Search(puzzle Puzzle, coll Collection) []Puzzle {
	visited := 0
	createdNodes := make(map[int]bool)

	// add root node
	coll.Push(&Node{Puzzle: puzzle, parent: nil, deph: 0})
	createdNodes[puzzle.Hash()] = true
	var endNode *Node
	for {
		cur := coll.Pop()
		visited += 1

		// check if current is solved
		if cur.Puzzle.Solved() {
			endNode = cur
			break
		}

		//expand children
		posMoves := cur.Puzzle.GetPossibleMoves()
		for _, move := range posMoves {
			hash := move.Hash()
			_, ok := createdNodes[hash]
			if !ok {
				coll.Push(&Node{Puzzle: move, deph: cur.deph + 1, parent: cur})
				createdNodes[hash] = true
			}
		}
	}
	fmt.Println("visited", visited, "nodes")
	// build path using links in nodes
	solution := make([]Puzzle, endNode.deph)
	for i := len(solution) - 1; i >= 0; i -= 1 {
		solution[i] = endNode.Puzzle
		endNode = endNode.parent
	}

	return solution
}

type Collection interface {
	Push(*Node)
	Pop() *Node
	Size() int
}

func NewStack() *Stack {
	return &Stack{data: make([]*Node, 0)}
}

func NewQueue() *Queue {
	return &Queue{data: make([]*Node, 0)}
}

type Stack struct {
	data []*Node
}

func (s *Stack) Push(n *Node) {
	s.data = append(s.data, n)
}

func (s *Stack) Pop() *Node {
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}

func (s *Stack) Size() int {
	return len(s.data)
}

type Queue struct {
	data []*Node
}

func (q *Queue) Push(n *Node) {
	q.data = append(q.data, n)
}

func (q *Queue) Pop() *Node {
	val := q.data[0]
	q.data = q.data[1:]
	return val
}

func (q *Queue) Size() int {
	return len(q.data)
}
