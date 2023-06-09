package coloring

import (
	. "GraphColoring/graph"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"sync"
)

// function used to find the K coloring of the graph, taking in input a graph,
// the start and the end of the interval to check and the channels on which to write
// the result/from which to get the stop signal from the manager thread
func FindColoringWorker(g IGraph, intervalStart, intervalEnd []int, k, digitsNumber int, done chan<- Result, stop chan bool, wg *sync.WaitGroup, mu *sync.Mutex) {

	defer wg.Done()

	found, config, err := checkInterval(g, intervalStart, intervalEnd, k, stop)
	if err != nil {
		fmt.Println("error checking configuration: ", err)
		return
	}

	if found {
		done <- MakeResult(found, config)
		mu.Lock()

		select {
		case _, ok := <-stop:
			if ok {
				fmt.Println("error - stop channel is not writable")
			}
		default:
			close(stop)
		}

		mu.Unlock()
	}
}

// checkInterval(g, a, b, k) takes a graph g, two array of integers a and b and an int k.
// a and b represent the digits of two integers in the base k;
// this function scans the range [a,b) until it finds a valid graph coloring, that is
// an assignment <node, digit> such that two adjacent nodes never have the same digit;
// return true if it finds that assignment, false otherwise
func checkInterval(g IGraph, a, b []int, k int, stop <-chan bool) (found bool, config []int, err error) {

	// check parameter correctness
	if g == nil || a == nil || b == nil {
		return false, nil, errors.New("nil argument")
	}
	valid, err := validRange(a, b, k)
	if !valid {
		return false, nil, err
	}

	current := make([]int, len(a))
	for copy(current, a); !reflect.DeepEqual(current, b); incr(current, k) {
		select {
		case _, ok := <-stop:
			if !ok {
				return false, nil, nil
			}
		default:

			// if valid assignment -> return true
			found, err = validAssignment(g, current)
			if err != nil {
				return false, nil, err
			}
			if found {
				return true, current, nil
			}

		}
	}

	return false, nil, nil
}

// colorOf(coloring, nodes, node) returns the i-th color in coloring
// if node is the i-th element of nodes.
// if node is not in nodes, returns an error
func colorOf(coloring []int, nodes []INode, n INode) (int, error) {
	for i, nn := range nodes {
		if nn.Label() == n.Label() {
			return coloring[i], nil
		}
	}

	return -1, errors.New("Error: node " + n.Label() + " not found")
}

// validRange(a, b, k) return true if
// len(a) == len(b) and a[i] < b[i] < k for each i in [0,len(a));
// return false, err otherwise, where err is an error that explain the error-case
func validRange(a, b []int, k int) (valid bool, err error) {
	if len(a) != len(b) { // return with priority
		return false, errors.New("intervals of different length")
	}
	for i := 0; i < len(a); i++ {
		if a[i] >= k || b[i] >= k { // return with priority
			return false, errors.New("Wrong representation for base " + (strconv.Itoa(k)))
		}
		if b[i] > a[i] {
			valid = true
		} // return when the valid checks end
	}

	return
}

// validAssignment(g) returns true if each node of g has a different color from
// its adjacents' color; false otherwise
func validAssignment(g IGraph, coloring []int) (bool, error) {
	if g == nil {
		return false, errors.New("nil argument")
	}

	// coloring[i] is the color of g.nodes[i] (pointwise coloring)
	for i, n := range g.Nodes() {
		neighbours, err := g.NeighboursOf(n.Label())
		if err != nil {
			fmt.Println(err.Error())
		}

		// peek the color of a neighbour of n
		for _, nn := range neighbours {
			nn_color, err := colorOf(coloring, g.Nodes(), nn)
			if err != nil {
				fmt.Println(err.Error())
			}

			// n.color == nn.color -> false
			if coloring[i] == nn_color {
				return false, nil
			}
		}
	}

	return true, nil
}

// incr takes an array of integer n and an int k,
// that represent the digit of a number in base k,
// and returns the successor of n in base k
func incr(n []int, k int) {
	for i := len(n) - 1; i >= 0; i-- {
		if n[i] < k-1 {
			n[i]++
			return
		}
		n[i] = 0 // n[i] = k
	}
}
