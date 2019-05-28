package main

// Queue is a queue.
type Queue struct {
	nodes []*Node
}

// Node is a item of a stack.
type Node struct {
	value string
}

// newQueue create a Stack.
func newQueue() *Queue {
	return &Queue{}
}

// enqueue adds an node to the end of the queue.
func (s *Queue) enqueue(n *Node) {
	s.nodes = append(s.nodes, n)
}

// dequeue removes an node from the top of the queue.
func (s *Queue) dequeue() {
	s.nodes = s.nodes[1:len(s.nodes)]
}
