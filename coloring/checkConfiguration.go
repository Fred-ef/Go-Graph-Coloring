package coloring

import (
	"errors"	
	"reflect"
	"strconv"
	. "GraphColoring/graph"
)

// checkConfiguration(g, a, b, k) takes a graph g, two array of integers a and b and an int k.
// a and b represent the digits of two integers in the base k;
// this function scans the range [a,b) until it finds a valid graph coloring, that is
// an assignment <node, digit> such that two adjacent nodes never have the same digit;
// return true if it finds that assignment, false otherwise
func CheckConfiguration(g IGraph, a, b []int, k int) (found bool, err error){
	
	// check parameter correctness
	if g==nil || a==nil || b==nil { return false, errors.New("Nil argument.") }
	valid, err := validRange(a,b,k)
	if !valid { return false,err }
	
	current := make([]int,len(a))
	for copy(current, a); !reflect.DeepEqual(current,b); incr(current, k) {

		// Assign colors to nodes
		err = g.Paint(ToColors(current))
		if err != nil { return false, err }

		// if valid assignment -> return true
		found, err = validAssignment(g)
		if err != nil { return false, err }
		if found{ return true,nil }
	}

	return
}

// validRange(a, b, k) return true if
// len(a) == len(b) and a[i] < b[i] < k for each i in [0,len(a));
// return false, err otherwise, where err is an error that explain the error-case
func validRange(a,b []int, k int) (valid bool, err error) {
	if len(a) != len(b) { // return with priority
		return false, errors.New("Intervals of different length.") 
	} 
	for i := 0; i<len(a); i++{
		if a[i] >= k || b[i] >= k { // return with priority
			return false, errors.New("Wrong representation for base "+(strconv.Itoa(k))) 
		} 
		if b[i]>a[i]{ valid = true } // return when the valid checks end
	}

	return
}

// validAssignment(g) returns true if each node of g has a different color from 
// its adjacents' color; false otherwise
func validAssignment(g IGraph) (bool, error) {
	if g == nil { return false, errors.New("Nil argument.") }
	
	// for each node, check if its color is the same as any of its neighbours
	for _,n := range g.Nodes(){
		neighbours, e := g.NeighboursOf(n.Label())
		if e != nil { return false, e }
		
		for _, nn := range neighbours{
			if n.Color() == nn.Color() { return false, nil }
		}
	}

	return true, nil
}

// incr takes an array of integer n and an int k, 
// that represent the digit of a number in base k, 
// and returns the successor of n in base k
func incr(n []int, k int){
	for i := len(n)-1; i>=0; i--{
		if n[i] < k-1 { n[i]++; return } 
		n[i] = 0 // n[i] = k
	}
}