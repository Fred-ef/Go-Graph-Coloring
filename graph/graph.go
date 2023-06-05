package graph

import (
	"errors"
	"strconv"
)

// Graph is an immutable representation of an undirected graph and implements IGraph;
// this implementation just memorizes a set of nodes and
// a mapping from the label of source nodes the target nodes for the edges;
type Graph struct {
	nodes map[string]INode      // nodes indexed by label
	edges map[string]([]string) // set of edges: map from source node to target nodes
}

// Make(nodes, edges) returns a graph with the given nodes and edges
// it returns an error if some parameter is null, or
// if a duplicate node is passed, or
// if an edge refers label that are not in the given nodes list
func Make(nodes []INode, edges []IEdge) (IGraph, error) {

	// check parameter != nil
	if nodes == nil || edges == nil {
		return nil, errors.New("Error: nil parameter.")
	}
	for i, n := range nodes {
		if n == nil {
			return nil, errors.New("The " + (strconv.Itoa(i)) + "-th node is nil.")
		}
	}
	for i, e := range edges {
		if e == nil {
			return nil, errors.New("The " + (strconv.Itoa(i)) + "-th edge is nil.")
		}
	}

	g := *new(Graph)
	g.nodes = make(map[string]INode)
	g.edges = make(map[string]([]string))

	// set nodes
	for _, n := range nodes {
		err := g.addNode(n)
		if err != nil {
			return nil, err
		}
	}

	// set edges
	for _, e := range edges {
		err := g.addEdge(e.SourceLabel(), e.TargetLabel())
		if err != nil {
			return nil, err
		}
	}

	return g, nil
}

// this.Nodes() returns the list of nodes of this
func (this Graph) Nodes() (nodes []INode) {
	for _, v := range this.nodes {
		nodes = append(nodes, v)
	}

	return
}

// this.NeighboursOf(n) returns the list of adjacent nodes of n
func (this Graph) NeighboursOf(n string) (nodes []INode, e error) {
	if !this.member(n) {
		return nil, errors.New("There is no node labelled with " + n)
	}

	for _, k := range this.edges[n] {
		nodes = append(nodes, this.nodes[k])
	}

	return
}

// this.Paint(colors) compute a pointwise assignment of the colors to the nodes of the graph;
// if the length of the nodes and the length of the colors is not the same, returns an error
func (this Graph) Paint(colors []string) (e error) {
	if colors == nil {
		return errors.New("Error: nil parameter.")
	}

	if len(this.nodes) != len(colors) {
		return errors.New("This assignment of colors is no possible." +
			"There are too many or too few colors")
	}

	i := 0
	for _, n := range this.nodes {
		n.SetColor(colors[i])
		i++
	}

	return
}

/******************** private methods ********************/

// this.member(n) checks if there is a node with label n
func (this Graph) member(n string) bool {
	_, ok := this.nodes[n]

	return ok
}

// this.AddNode(n) adds n to this.nodes
// if this.nodes already contains n, it returns an error
func (this Graph) addNode(n INode) error {
	if this.member(n.Label()) {
		return errors.New("Node " + n.Label() + " is already present.")
	}

	this.nodes[n.Label()] = n
	return nil
}

// this.AddEdge(n1, n2) adds an edge between the nodes with labels n1 and n2
// if n1 or n2 are not nodes of this, returns an error
func (this Graph) addEdge(n1, n2 string) (e error) {
	_, ok1 := this.nodes[n1]
	_, ok2 := this.nodes[n2]
	if ok1 && ok2 {
		if !contains(this.edges[n1], n2) {
			this.edges[n1] = append(this.edges[n1], n2)
		}
		if !contains(this.edges[n2], n1) {
			this.edges[n2] = append(this.edges[n2], n1)
		}
	} else {
		e = errors.New("At least one of the given nodes (" + n1 + ", " + n2 + ") is not in the graph.")
	}

	return
}
