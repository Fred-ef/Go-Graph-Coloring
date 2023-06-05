package graph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// isColor(c) returns true if c is a valid color, false otherwise
// STUB: dipendente dalla rappresentazione scelta per i colori;
//
//	enumerazione, lista di color-names, codifica RGB/CMYK, hex code, interi,...
func IsColor(color string) bool {
	return true
}

// toColor(c) returns the color associated to the integer c
// STUB: dipendente dalla rappresentazione scelta per i colori;
//
//	enumerazione, lista di color-names, codifica RGB/CMYK, hex code, interi,...
func ToColor(c int) string {
	return strconv.Itoa(c)
}

// toColors(cs) takes an array of integers cs and returns an array of strings s such that
// s[i] = toColor(cs[i])
func ToColors(colors []int) (s []string) {
	s = make([]string, len(colors))
	for i, color := range colors {
		s[i] = ToColor(color)
	}

	return
}

// contains(e, v) returns true if v is an element of e, false otherwise
func contains[T comparable](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

// starts the construction of a graph from a JSON file in the "examples" folder
func LoadFromFile(path string) (IGraph, error) {
	// opening the JSON file
	graphFile, err := os.Open(path)
	if err != nil {
		fmt.Println("Error in opening the JSON file: ", err)
		return nil, err
	}
	defer graphFile.Close()

	// reading the JSON file as a byte array
	bytes, _ := ioutil.ReadAll(graphFile)

	// sending the json to the building function and getting back a graph
	graph, err := buildFromJson(bytes)
	if err != nil {
		return nil, err
	}

	return graph, nil
}

// TEMP VERSION: starts the construction of a graph from a frontend request
func LoadFromUI(bytes []byte) (IGraph, error) {
	graph, err := buildFromJson(bytes)
	if err != nil {
		return nil, err
	}

	return graph, nil
}

// constructs a graph from a JSON representation of an adjacency list
func buildFromJson(bytes []byte) (IGraph, error) {

	var adList AdList

	// unmarshaling the JSON file
	err := json.Unmarshal(bytes, &adList)
	if err != nil {
		fmt.Println("Error in JSON parsing: wrong format", err)
		return nil, err
	}

	var nodes []INode
	var edges []IEdge

	for i := 0; i < len(adList.List); i++ {
		currNode, _ := MakeNode(strconv.Itoa(i))
		nodes = append(nodes, currNode)
		currNodeEdges := adList.List[i]

		for j := 0; j < len(currNodeEdges); j++ {
			currEdge, _ := MakeEdge(strconv.Itoa(i), strconv.Itoa(currNodeEdges[j]))
			edges = append(edges, currEdge)
		}
	}

	graph, err := Make(nodes, edges)
	if err != nil {
		fmt.Println("Error in graph construction: ", err)
		return nil, err
	}

	return graph, nil
}
