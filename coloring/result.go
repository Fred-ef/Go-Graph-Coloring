package coloring

// struct used to represent the result of a checkConfiguration computation:
// a struct containing 2 values, one indicating whether the check has succeeded or failed
// (success) and one indicating, where relevant (success case), the coloring of the graph
// (coloring)
type Result struct {
	Success  bool
	Coloring string
}

// creates a new Result object
func MakeResult(success bool, coloring string) Result {
	return Result{Success: success, Coloring: coloring}
}
