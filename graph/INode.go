package graph

// INode is the abstract representation of a node of a graph;
// it exposes a method to get the label
type INode interface {
	Label() string         // returns the label of the node
}
