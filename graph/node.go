package graph

import "errors"

// node type represents a node (vertex) of a graph and implements the interface INode;
type Node struct {
	label string
}

// MakeNode(l) takes a label l and returns a new node with label l 
func MakeNode(label string) (INode, error) {
	if label == "" {
		return nil, errors.New("a node must have a non-empty label")
	}

	return Node{label: label}, nil
}

// this.Label() returns the label of the node
func (this Node) Label() string {
	return this.label
}