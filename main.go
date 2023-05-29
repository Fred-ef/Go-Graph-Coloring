package main

import (
	"fmt"
	. "GraphColoring/graph"
)

func main()  {

	// Complete Triangle graph
	var n1, n2, n3 INode
	var e1, e2, e3 IEdge
	n1,_ = MakeNode("n1")
	n2,_ = MakeNode("n2")
	n3,_ = MakeNode("n3")
	e1,_ = MakeEdge("n1","n2")
	e2,_ = MakeEdge("n1","n3")
	e3,_ = MakeEdge("n2","n3")
	fmt.Println("Nodes: ",n1,n2,n3)
	
	g, err := Make([]INode{n1,n2,n3}, []IEdge{e1,e2,e3})
	fmt.Println("\nGraph: ",g)
	neighbours,_ := g.NeighboursOf("n1")
	fmt.Println("\nn1 neighbours: ", neighbours)
	neighbours,_ = g.NeighboursOf("n2")
	fmt.Println("\nn2 neighbours: ", neighbours)
	neighbours,_ = g.NeighboursOf("n3")
	fmt.Println("\nn3 neighbours: ", neighbours)

	err = g.Paint([]string{"R","G","B"})
	if err != nil { 
		fmt.Println(err)
	} else { 
		fmt.Println("\nColor of the graph: ")
		for _,n := range g.Nodes(){
			fmt.Println("<"+n.Label()+", "+n.Color()+">; ")
		}
	}
	
	// some error cases...

	fmt.Println("\nEmpty-label node construction")
	n,err := MakeNode("")
	if err==nil{ fmt.Println(n) } else{ fmt.Println(err) }

	fmt.Println("\nNil argument in graph creation")
	g,err = Make(nil, []IEdge{})
	if err==nil{ fmt.Println(g) } else{ fmt.Println(err) }
	g,err = Make([]INode{}, nil)
	if err==nil{ fmt.Println(g) } else{ fmt.Println(err) }
	g,err = Make([]INode{nil}, []IEdge{})
	if err==nil{ fmt.Println(g) } else{ fmt.Println(err) }
	g,err = Make([]INode{}, []IEdge{nil})
	if err==nil{ fmt.Println(g) } else{ fmt.Println(err) }

	fmt.Println("\nDuplicate node in graph construction")
	g,err = Make([]INode{n1,n2,n1}, []IEdge{})
	if err==nil{ fmt.Println(g) } else{ fmt.Println(err) }
	
	fmt.Println("\nUnknow label for an edge")
	e,_ := MakeEdge("n1","n3")
	g,err = Make([]INode{n1,n2}, []IEdge{e})
	if err==nil{ fmt.Println(g) } else{ fmt.Println(err) }

	fmt.Println("\nNeighbour of unknow label")
	g,_ = Make([]INode{n1,n2,n3}, []IEdge{e1,e2,e3})
	neighbours,err = g.NeighboursOf("n5")
	if err==nil{ fmt.Println(neighbours) } else{ fmt.Println(err) }

	fmt.Println("\nBad graph coloring")
	err = g.Paint([]string{"R","G"})
	if err != nil { 
		fmt.Println(err)
	} else { 
		fmt.Println("\nColor of the graph: ")
		for _,n := range g.Nodes(){
			fmt.Println("<"+n.Label()+", "+n.Color()+">; ")
		}
	}

}