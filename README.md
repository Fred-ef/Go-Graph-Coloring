# concurrent-KColoring
A concurrent K-Coloring solver

To launch an instance of the solver, you can first build the program as follows:
- go build main.go

Then, you can execute the program passing 2 parameters as follows:
- ./main TEST_GRAPH_NAME K

where "TEST_GRAPH_NAME" is the name of a JSON graph in the "tests" folder (without the ".json")
and K is the number of colors for the query.
For example:
- ./main graph_8_7 7

To find out if the graph "graph_8_7" is 7-colorable.