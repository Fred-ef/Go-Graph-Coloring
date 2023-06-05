// Intermediate representation used in the graph construction
// before building the final graph, the file containing its adjacency list is parsed and this representation is built
// this representation is then given to the graph-building function to build the final version of the graph

package graph

// struct representing an adjacency list
type AdList struct {
	List [][]int `json:"list"`
}
