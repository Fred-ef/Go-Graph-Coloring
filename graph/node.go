package graph

import "errors"

// node type represents a node (vertex) of a graph and implements the interface INode;
// a node has a label, a color and a list of neighbours (adjacent nodes)
type Node struct {
	label string
	color *string // mutable
}

// MakeNode(l, ns) takes a label l and a list of neighbours ns
// and returns a new node with label l and with neighbours ns
func MakeNode(label string) (INode, error) {
	if label=="" { return nil, errors.New("A node must have a non-empty label.") }

	return Node{label: label, color:new(string)}, nil
}

// this.Label() returns the label of the node
func (this Node) Label() string {
	return this.label
}

// this.Color() returns the color of the node
func (this Node) Color() string {
	return *(this.color)
}

// this.SetColor(c) sets this.color to c if c is a valid color and returns an error otherwise
func (this Node) SetColor(color string) error {
	if IsColor(color) {
		*(this.color) = color
		return nil
	}
	return errors.New(color + " is not a valid color.")
}