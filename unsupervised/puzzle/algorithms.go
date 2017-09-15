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
	return Search(puzzle, NewQueue())
}

func DephFirst(puzzle Puzzle) []Puzzle {
	return Search(puzzle, NewStack(-1))
}

func DephLimited(puzzle Puzzle) []Puzzle {
	return Search(puzzle, NewStack(1000))
}

func IterativeDeepening(puzzle Puzzle) []Puzzle {
	return Search(puzzle, &IterativeStack{data: make(map[int][]*Node, 0), deepeningStep: 10})
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

func BiDirectional(puzzle Puzzle) []Puzzle {
	return SearchBoth(puzzle, NewQueue())
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
	solution = append([]Puzzle{puzzle}, solution...)

	return solution
}

func SearchBoth(puzzle Puzzle, coll Collection) []Puzzle {

	visited := 0
	createdNodes := make(map[int]*Node)
	solutionNodes := make(map[int]*Node)
	solColl := NewQueue()

	// add root node
	node := &Node{Puzzle: puzzle, parent: nil, deph: 0}
	solvedPuzzle := GetSolvedPuzzle()
	solColl.Push(&Node{Puzzle: solvedPuzzle, parent: nil, deph: 0})
	coll.Push(node)
	createdNodes[puzzle.Hash()] = node
	var endNode *Node
	var solutionEndNode *Node
	for {
		cur := coll.Pop()
		solutioncur := solColl.Pop()
		visited += 2

		_, ok := solutionNodes[cur.Puzzle.Hash()]
		// check if current is solved
		if cur.Puzzle.Solved() || ok {

			solutionEndNode = solutionNodes[cur.Puzzle.Hash()]
			endNode = cur
			break
		}

		//expand children
		posMoves := cur.Puzzle.GetPossibleMoves()
		for _, move := range posMoves {
			hash := move.Hash()
			_, ok := createdNodes[hash]
			if !ok {
				node := &Node{Puzzle: move, deph: cur.deph + 1, parent: cur}
				coll.Push(node)
				createdNodes[hash] = node
			}
		}
		posMoves = solutioncur.Puzzle.GetPossibleMoves()
		for _, move := range posMoves {
			hash := move.Hash()
			_, ok := solutionNodes[hash]
			if !ok {
				node := &Node{Puzzle: move, deph: cur.deph + 1, parent: solutioncur}
				solColl.Push(node)
				solutionNodes[hash] = node
			}
		}
	}

	if endNode.Puzzle.Solved() {
		fmt.Println("visited", visited, "nodes", "of", len(createdNodes))
		// build path using links in nodes
		solution := make([]Puzzle, endNode.deph)
		for i := len(solution) - 1; i >= 0; i -= 1 {
			solution[i] = endNode.Puzzle
			endNode = endNode.parent
		}

		return solution
	} else {
		fmt.Println("visited", visited, "nodes", "of", len(createdNodes))

		otherNode := endNode
		// build path using links in nodes
		solution := []Puzzle{}
		otherNode = otherNode.parent
		for i := otherNode.deph; otherNode.parent != nil; i -= 1 {
			//fmt.Println("assigning from othernode", otherNode.Puzzle.ToString())
			solution = append([]Puzzle{otherNode.Puzzle}, solution...)
			//solution[i] = otherNode.Puzzle
			otherNode = otherNode.parent
		}
		solution = append([]Puzzle{otherNode.Puzzle}, solution...)
		//solution[0] = otherNode.Puzzle
		//fmt.Println("assigning from endnode end", otherNode.Puzzle.ToString())
		for i := otherNode.deph + 2; solutionEndNode.parent != nil; i += 1 {
			//fmt.Println("assigning from solutionendnode", solutionEndNode.Puzzle.ToString())
			solution = append(solution, solutionEndNode.Puzzle)
			//solution[i] = solutionEndNode.Puzzle
			solutionEndNode = solutionEndNode.parent
		}
		solution = append(solution, solutionEndNode.Puzzle)
		//solution[len(solution)-1] = solutionEndNode.Puzzle
		//fmt.Println("assigning from solutionendnode end", solutionEndNode.Puzzle.ToString())

		return solution
	}
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
