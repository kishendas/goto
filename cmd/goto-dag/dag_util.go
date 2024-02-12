/*
Copyright 2024 The GOTO Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"errors"
)

// TopologicalSort This method takes an adjacency list of tasks with their dependencies as input
// and creates a series of stages, where each stage can contain one or more tasks
// All the tasks in a given stage can be executed in parallel.
func TopologicalSort(adjList map[int][]int) ([]RawStage, error) {

	topOrder := []RawStage{}
	var stage = 0

	// create a map for each node to its corresponding indegree
	inDegree := map[int]int{}

	// set values of all indegrees to 0
	for n := range adjList {
		inDegree[n] = 0
	}

	// for each vertex u
	for _, u := range adjList {
		// for each vertex v, having an edge to u
		for _, v := range u {
			// increment inDegree[v]
			inDegree[v]++
		}
	}

	// start with all vertices, where each vertex u is such that in-degree[u] = 0
	verticesZeroIndegree := []RawStage{}
	for u, v := range inDegree {
		if v != 0 {
			continue
		}
		verticesZeroIndegree = append(verticesZeroIndegree, RawStage{index: u, stage: stage})
	}

	for len(verticesZeroIndegree) > 0 {

		// delete a vertex from VerticesWithZeroIndegree and call it deletedVertex
		deletedVertex := verticesZeroIndegree[0]
		verticesZeroIndegree = verticesZeroIndegree[1:]

		// Add deleted vertex to the end of the topological order
		topOrder = append(topOrder, deletedVertex)

		// Increment the stage value
		stage = stage + 1

		// For each vertex v adjacent to deleted, decrement the inDegree by 1, since the edge will also be deleted.
		for _, v := range adjList[deletedVertex.index] {
			// Decrement inDegree[v]
			inDegree[v]--

			// if inDegree[v] = 0, then insert v into VerticesWithZeroIndegree list
			if inDegree[v] == 0 {
				verticesZeroIndegree = append(verticesZeroIndegree, RawStage{index: v, stage: stage})
			}
		}
	}

	if len(topOrder) < len(adjList) {
		return nil, errors.New("dag-10000: cycle in the DAG")
	}

	// Fix the final stage values
	for i, _ := range topOrder {
		if i == 0 {
			topOrder[i].finalStage = topOrder[i].stage
		} else if topOrder[i].stage != topOrder[i-1].stage {
			topOrder[i].finalStage = topOrder[i-1].finalStage + 1
		} else {
			topOrder[i].finalStage = topOrder[i-1].finalStage
		}
	}

	return topOrder, nil
}

func main() {

}
