package graph

// INode is the abstract representation of a node of a graph;
// it exposes methods to get the label and the color of the node
// and to update the color of the node
type INode interface {
	Label() string         // returns the label of the node
	Color() string // returns the color of the node
	SetColor(string) error // sets the color of the node
}
