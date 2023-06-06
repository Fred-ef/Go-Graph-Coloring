package main

import (
	. "GraphColoring/coloring"
	. "GraphColoring/graph"
	"fmt"
)

func main() {

	g, err := LoadFromFile("./tests/graph_10_10.json")
	if err != nil {
		fmt.Println("Shutting down because of graph construction error...")
		return
	}

	// ##### SERVER CREATION #####

	k := 8

	res, err := FindColoring(g, k)
	if err != nil {
		fmt.Println("Error finding coloring:\n", err)
	}

	if res.Success {
		fmt.Println("Found coloring:\n", res.Coloring)
	} else {
		fmt.Println("The graph is not", k, "colorable")
	}
}
