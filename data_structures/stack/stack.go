package main

// Stack is a stack.
type Stack struct {
	nodes []*Node
}

// Node is a item of a stack.
type Node struct {
	value string
}

// newStack create a Stack.
func newStack() *Stack {
	return &Stack{}
}

// push adds an node to the top of the stack.
func (s *Stack) push(n *Node) {
	s.nodes = append(s.nodes[:len(s.nodes)], n)
}

// pop removes an node from the top of the stack.
func (s *Stack) pop() {
	s.nodes = s.nodes[:len(s.nodes)-1]
}
