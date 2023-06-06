package main

import (
	. "GraphColoring/coloring"
	. "GraphColoring/graph"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// ##### GETTING PARAMETERS AND BUILDING THE GRAPH #####

	args := os.Args
	if len(args) < 3 {
		fmt.Println("Error: name of test-graph and parameter k are required")
		return
	}

	testGraph := args[1]
	k, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Error in getting console parameters")
	}

	graph, err := LoadFromFile("./tests/" + testGraph + ".json")
	if err != nil {
		fmt.Println("Shutting down because of graph construction error...")
		return
	}

	// ##### CHECKING THE COLORING #####

	res, err := FindColoring(graph, k)
	if err != nil {
		fmt.Println("Error finding coloring:\n", err)
	}

	if res.Success {
		fmt.Print("Coloring found!\n", res.Coloring, "\n\n")
		for i := 0; i < len(res.Coloring); i++ {
			fmt.Println("Node", i, "-", "color", res.Coloring[i])
		}
	} else {
		fmt.Println("The graph is not", k, "colorable")
	}
}
