package main

import "errors"

//	Node type represents a node (vertex) of a graph;
//	a node has a label,
//	a color and
//	a list of (references to) adjacent nodes.
//
//	Node is an internal representation used by graph's methods; it must not be exported;
type node struct{
	label string
	color string
	adjs []*node
}

//	makeNode(l) returns a new node with label l and no color nor adjacent nodes
func makeNode(label string) node{
	return node{label:label}
}

//	setColor(c) sets this' color if c is a valid color and returns an error otherwise
func (this *node) setColor(color string) error { 
	if isColor(color){
		this.color=color 
		return nil
	} 
	return errors.New(color+" is not a valid color.")
}

//	addNeighbour(n) appends n to this' adjacents list
func (this *node) addNeighbour(n node) {
	this.adjs = append(this.adjs, &n)
}