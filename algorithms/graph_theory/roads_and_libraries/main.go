package main

import (
	"bufio"
	"fmt"
	"os"
)

type QueryDetail struct {
	N     int // the number of cities
	M     int // the number of roads
	Clib  int // the cost to build a library
	Croad int // the cost to repair a road
}

type VisitedVertex []bool

type Vertex []int

type Graph []Vertex

func (graph Graph) push(node int, adjNode int) {
	graph[node] = append(graph[node], adjNode)
	graph[adjNode] = append(graph[adjNode], node)
}

func dfs(node int, graph Graph, visited VisitedVertex, count *int) {
	// total roads to connect to each vertex in the compomnent
	visited[node] = true
	for _, adjNode := range graph[node] {
		if visited[adjNode] == false {
			// fmt.Println(adjNode)
			*count += 1
			dfs(adjNode, graph, visited, count)
		}
	}
}

// https://stackoverflow.com/questions/28081486/golang-multiple-files-in-main-package
func main() {
	bufin := bufio.NewReader(os.Stdin)
	var numOfQueries int
	fmt.Fscanf(bufin, "%d\n", &numOfQueries)

	for i := 1; i <= numOfQueries; i++ {

		var qd QueryDetail

		// Initialize vertex
		fmt.Fscanf(bufin, "%d %d %d %d\n", &qd.N, &qd.M, &qd.Clib, &qd.Croad)
		graph := make(Graph, qd.N)
		visited := make(VisitedVertex, qd.N)

		for i := 1; i <= qd.M; i++ {
			// fmt.Println(i)
			// The integers u, v describe a bidirectional road connecting cities.
			var u, v int
			fmt.Fscanf(bufin, "%d %d\n", &u, &v)

			graph.push(u-1, v-1)
		}

		totalCost := 0

		if qd.Clib <= qd.Croad {

			totalCost = qd.N * qd.Clib

		} else {
			for i := 0; i < len(graph); i++ {
				if visited[i] == false {
					count := 0
					dfs(i, graph, visited, &count)
					totalCost += qd.Croad*(count) + qd.Clib
				}
			}
		}

		fmt.Println(totalCost)
	}
}
