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

NOTE: The naming of the graphs is also relevant; you will find every graph is named in
the following format:
- graph_n_m.json

This means that the graph has 'n' nodes and IS m-colorable (but NOT (m-1)-colorable).
So, the graph:
- graph_8_7

is a graph made up of 8 nodes which is 7-colorable (but NOT 6-colorable).