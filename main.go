package main

import "fmt"

func main()  {
	var n node
	n1 := makeNode("1")
	n2 := makeNode("2")
	fmt.Println("nodes:\n",n,"\n",n1,"\n",n2)
	
	n1.setColor("white")
	fmt.Println("n1 color: ", n1.color)

	fmt.Println("n1 neighbours:")
	n1.addNeighbour(n2)
	for _,v := range n1.adjs {
		fmt.Println(v)
	}

	// zero-init graph
	fmt.Println("zero-init graph (new):")
	g := new(Graph)
	fmt.Println(g)

	// constructed graph
	fmt.Println("constructed graph (make):")
	g1 := Make([]node{n,n1,n2})
	fmt.Println(g1)
}