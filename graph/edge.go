package graph

import "errors"

// edge type represents an edge between two nodes in a graph and implements IEdge;
// that's just a pair of labels, but it's a convenient way to represent an edge.
//
// Note: since the interface just provide methods for getting the label of the nodes,
// we can just store the label of the nodes
type edge struct{
	source string
	target string
}

// MakeEdge(s, t) returns a new edge initialized with s as source's label
//  and t as target's label
func MakeEdge(source, target string) (IEdge, error){
	if source=="" || target==""{ return nil, errors.New("A node must have a non-empty label.") }

	return edge{source:source, target:target}, nil
}

// SourceLabel() returns the source's label
func (this edge) SourceLabel() string{
	return this.source
}

// TargetLabel() returns the target's label
func (this edge) TargetLabel() string{
	return this.target
}