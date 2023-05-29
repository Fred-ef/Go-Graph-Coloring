package graph

// IEdge is the abstract representation of an indirect edge between two nodes of a graph;
// it exposes methods to get the label of the two nodes
type IEdge interface{
	SourceLabel() string	// returns the label of the source node
	TargetLabel() string	// returns the lebal of the target node
}