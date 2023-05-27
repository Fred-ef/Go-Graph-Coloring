package main

//	Graph is a mutable representation of an undirected graph;
//	it has a root node,
//	a set of nodes and
//	a set of edges
type Graph struct{
	nodes map[string]node // nodes indexed by label
	edges map[string]([]string) // set of edges: map from source node to target nodes
}

// Make(nodes) returns a graph with the given nodes and edges
func Make(nodes []node) (g Graph) {
	g = *new(Graph)
	g.nodes = make(map[string]node)
	g.edges = make(map[string]([]string))

	for _,n := range nodes{
		g.nodes[n.label] = n
		
		// set edges
		for _,adj := range n.adjs{
			g.edges[n.label] = append(g.edges[n.label], adj.label)
		}

	}

	return
}