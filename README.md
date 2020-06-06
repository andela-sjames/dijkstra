# dijkstra

Let's implement the Dijkstra Algorithm.

<https://www.geeksforgeeks.org/dijkstras-algorithm-for-adjacency-list-representation-greedy-algo-8/>

`Distance table: Example usecase`
|||||||
| :--- | :--- | :--- | :--- | :--- | :--- |
| Nodes| 0 | 1 | 2 | 3 | 4 |
| Distance from node | 0| 3| 1| 4| 7|

## Example Graph

![Directed Acyclic Graph](images/Dag_dsp.png)

## Adjacency list

![Adjacency List](images/adjList_dsp.png)

## Result of Dijkstra shortest path

![djikstra](images/djikstra.png)

## Result

```shell
dijkstra git:(master) go run main.go
Distance graph:  [0 3 1 4 7]
Shortest Path:  [2 1 3 4]
```
