package puzzle

import (
	"fmt"
)

type Node struct {
	Puzzle Puzzle
	deph   int
	parent *Node
}

func DephFirst(puzzle Puzzle) []Puzzle {
	visited := make([]Node, 0)
	edges := make([]Node, 0)
	createdNodes := make(map[int]bool)

	// add root node
	edges = append(edges, Node{Puzzle: puzzle, parent: nil, deph: 0})
	createdNodes[puzzle.Hash()] = true
	var endNode *Node
	i := 0
	for {
		i += 1
		cur := edges[0]
		edges = edges[1:]
		if i%1000 == 0 {
			fmt.Println("edge removed", cur, len(edges))
		}

		visited = append(visited, cur)
		// check if current is solved
		if cur.Puzzle.Solved() {
			endNode = &cur
			break
		}

		//expand children
		posMoves := cur.Puzzle.GetPossibleMoves()
		for _, move := range posMoves {
			hash := move.Hash()
			_, ok := createdNodes[hash]
			//if !contains(move, visited) && !contains(move, edges) {
			if !ok {
				edges = append(edges, Node{Puzzle: move.Clone(), deph: cur.deph + 1, parent: &cur})
				createdNodes[hash] = true
				//fmt.Println("adding edge")
			}
		}
	}
	fmt.Println("visited", len(visited), "nodes")
	// build path using links in nodes
	solution := make([]Puzzle, 0)
	for endNode != nil {
		//fmt.Println("appending", endNode.Puzzle.ToString())
		solution = append([]Puzzle{endNode.Puzzle}, solution...)
		endNode = endNode.parent
	}
	//solution = append(solution, endNode.Puzzle)

	return solution
}
