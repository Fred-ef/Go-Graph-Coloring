package graph

// IGraph is the abstract representation of an indirected graph;
// it exposes methods to get the nodes of the graph and the adjacents of a given node
// and for coloring the graph with a given color assignment
type IGraph interface {
	Nodes() []INode                       // returns the list of nodes of the graph
	NeighboursOf(string) ([]INode, error) // returns the list of adjacents of the node with a given label
	Paint([]string) error                 // assigns a color to each node, orderly;
}
