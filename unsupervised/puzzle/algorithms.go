package puzzle

import (
	"fmt"
)

type Node struct {
	Puzzle Puzzle
	deph   int
	parent *Node
	rating int
}

func BreadthFirst(puzzle Puzzle) []Puzzle {
	return Search(puzzle, NewStack())
}

func DephFirst(puzzle Puzzle) []Puzzle {
	return Search(puzzle, NewQueue())
}

func BookGreedy(puzzle Puzzle) []Puzzle {
	queue := &PriorityQueue{}
	queue.data = make(map[int][]*Node, 0)
	queue.evalFunc = bookGreedy
	return Search(puzzle, queue)
}

func BookAStar(puzzle Puzzle) []Puzzle {
	queue := &PriorityQueue{}
	queue.data = make(map[int][]*Node, 0)
	queue.evalFunc = bookAStar
	return Search(puzzle, queue)
}

func MyAStar(puzzle Puzzle) []Puzzle {
	queue := &PriorityQueue{}
	queue.data = make(map[int][]*Node, 0)
	queue.evalFunc = myAStar
	return Search(puzzle, queue)
}

func MyGreedy(puzzle Puzzle) []Puzzle {
	queue := &PriorityQueue{}
	queue.data = make(map[int][]*Node, 0)
	queue.evalFunc = myGreedy
	return Search(puzzle, queue)
}

func Search(puzzle Puzzle, coll Collection) []Puzzle {
	visited := 0
	createdNodes := make(map[int]*Node)

	// add root node
	node := &Node{Puzzle: puzzle, parent: nil, deph: 0}
	coll.Push(node)
	createdNodes[puzzle.Hash()] = node
	var endNode *Node
	for {
		cur := coll.Pop()
		visited += 1

		// check if current is solved
		if cur.Puzzle.Solved() {
			//fmt.Println("done", cur.Puzzle.ToString(), cur.Puzzle.Solved())
			endNode = cur
			break
		}

		//expand children
		posMoves := cur.Puzzle.GetPossibleMoves()
		for _, move := range posMoves {
			hash := move.Hash()
			oldNode, ok := createdNodes[hash]
			if !ok {
				node := &Node{Puzzle: move, deph: cur.deph + 1, parent: cur}
				coll.Push(node)
				createdNodes[hash] = node
			} else {
				// check if you are better than current state found
				if oldNode.deph > cur.deph+1 {
					//fmt.Println("duplicate found", oldNode.deph, cur.deph+1)
					oldNode.parent = cur
					oldNode.deph = cur.deph + 1
				}
			}
		}
	}
	fmt.Println("visited", visited, "nodes", "of", len(createdNodes))
	// build path using links in nodes
	solution := make([]Puzzle, endNode.deph)
	for i := len(solution) - 1; i >= 0; i -= 1 {
		solution[i] = endNode.Puzzle
		endNode = endNode.parent
	}

	return solution
}

func bookGreedy(n *Node) int {
	rating := 0
	for i, val := range n.Puzzle.state {
		width := abs(i%3 - (val-1)%3)
		height := abs(i/3 - (val-1)/3)
		rating += width + height
	}
	return rating
}

func bookAStar(n *Node) int {
	return bookGreedy(n) + n.deph
}

func myGreedy(n *Node) int {
	rating := 0
	state := n.Puzzle.state
	for i := 1; i < len(state); i += 1 {
		if state[i-1] > state[i] {
			rating += 1
		}
	}
	return rating
}

func myAStar(n *Node) int {
	return myGreedy(n) + n.deph
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	} else {
		return i
	}
}
