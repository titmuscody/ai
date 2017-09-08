package puzzle

import "fmt"

const (
	allocSize = 1000
)

type Collection interface {
	Push(*Node)
	Pop() *Node
	Size() int
}

func NewStack(dephLimit int) *Stack {
	return &Stack{data: make([]*Node, 0), dephLimit: dephLimit}
}

func NewQueue() *Queue {
	return &Queue{data: make([]*Node, 0)}
}

type Stack struct {
	data      []*Node
	dephLimit int
}

func (s *Stack) Push(n *Node) {
	if n.deph < s.dephLimit || s.dephLimit == -1 {
		s.data = append(s.data, n)
	}
}

func (s *Stack) Pop() *Node {
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}

func (s *Stack) Size() int {
	return len(s.data)
}

type IterativeStack struct {
	data          map[int][]*Node
	deepeningStep int
	count         int
}

func (s *IterativeStack) Push(n *Node) {
	s.count += 1
	step := n.deph / s.deepeningStep
	if nodes, ok := s.data[step]; ok {
		nodes = append(nodes, n)
		s.data[step] = nodes
	} else {
		s.data[step] = []*Node{n}
	}
}

func (s *IterativeStack) Pop() *Node {
	s.count -= 1
	min := 1000000000
	for k := range s.data {
		if k < min {
			min = k
		}
	}
	nodes := s.data[min]
	val := nodes[len(nodes)-1]
	if len(nodes) == 1 {
		delete(s.data, min)
		return val
	} else {
		nodes = nodes[:len(nodes)-1]
		s.data[min] = nodes
		return val
	}
}

func (s *IterativeStack) Size() int {
	return s.count
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
	fmt.Println("created queue")
	return len(q.data)
}

// this is used to keep the lowest ranked numbers at the front
type PriorityQueue struct {
	data     map[int][]*Node
	count    int
	evalFunc func(*Node) int
}

func (q *PriorityQueue) Push(n *Node) {
	q.count += 1
	rank := q.evalFunc(n)
	if nodes, ok := q.data[rank]; ok {
		nodes = append(nodes, n)
		q.data[rank] = nodes
	} else {
		q.data[rank] = []*Node{n}
	}
}

func (q *PriorityQueue) Pop() *Node {
	q.count -= 1
	min := 1000000000
	for k := range q.data {
		if k < min {
			min = k
		}
		//fmt.Println("checking val", k)
	}
	nodes := q.data[min]
	val := nodes[0]
	//fmt.Println("popping", val, min)
	if len(nodes) == 1 {
		delete(q.data, min)
		return val
	} else {
		nodes = nodes[1:]
		q.data[min] = nodes
		return val
	}

}

func (q *PriorityQueue) Size() int {
	return q.count
}
