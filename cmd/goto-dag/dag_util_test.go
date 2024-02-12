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
	"reflect"
	"testing"
)

func TestTopologicalSort(t *testing.T) {

	//      ------>  t2
	//  t1 			    ------> t4
	//      ------>  t3
	// stage0 = t1
	// stage1 = t2, t3
	// stage3 = t4
	// stage0 -> stage1 -> stage2

	adjList := map[int][]int{
		1: []int{2, 3},
		2: []int{4},
		3: []int{4},
		4: []int{},
	}

	var stages, _ = TopologicalSort(adjList)

	expected := []RawStage{
		RawStage{index: 1, stage: 0, finalStage: 0},
		RawStage{index: 2, stage: 1, finalStage: 1},
		RawStage{index: 3, stage: 1, finalStage: 1},
		RawStage{index: 4, stage: 3, finalStage: 2},
	}

	if !reflect.DeepEqual(expected, stages) {
		t.Fatalf(`expected = %v, result %v ""`, expected, stages)
	}

}

func TestTopologicalSortDetectCycle(t *testing.T) {

	adjList := map[int][]int{
		1: []int{4},
		2: []int{4},
		3: []int{2, 4, 5},
		4: []int{6},
		5: []int{},
		6: []int{7},
		7: []int{1},
	}

	var _, err = TopologicalSort(adjList)

	var ErrMsg = "dag-10000: cycle in the DAG"

	if err.Error() != ErrMsg {
		t.Fatalf(`should fail with message = %q, %v, want "", error`, ErrMsg, err)
	}
}
