# concurrent-KColoring

A concurrent *K-Coloring* solver

To launch an instance of the solver, you can first **build** the program as follows:

`go build main.go`

Then, you can execute the program passing 2 parameters as follows:

`./main FILENAME K`

where `FILENAME` is the name of a JSON file (without the ".json" format) in the "*tests*" folder and `K` is the number of colors for the query.

For example:

`./main graph_8_7 7`

To find out if the graph into the file "*graph_8_7*" is *7-colorable*.

The file must contain the adjacency list of the graph, for example a complete 8-nodes graph is represented by the following JSON object

```json
{
    "list": 
    [
        [1, 2, 3, 4, 5, 6, 7],
        [0, 2, 3, 4, 5, 6, 7],
        [0, 1, 3, 4, 5, 6, 7],
        [0, 1, 2, 4, 5, 6, 7],
        [0, 1, 2, 3, 5, 6, 7],
        [0, 1, 2, 3, 4, 6, 7],
        [0, 1, 2, 3, 4, 5, 7],
        [0, 1, 2, 3, 4, 5, 6]
    ]
 }
```

in this representation, the *i-th* row contains the edge outgoing from the node *i*.

###### NOTE:

The naming of the graphs is also relevant; you will find every graph is named in
the following format: *graph_n_m.json*

This means that the graph has '*n*' nodes and IS *m*-colorable (but **NOT** (*m-1*)-colorable).
So, the graph: *graph_8_7* is a graph made up of 8 nodes which is 7-colorable (but NOT 6-colorable).