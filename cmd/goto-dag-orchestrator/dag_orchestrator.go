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
	"fmt"
)

type DAGOrchestrator struct {
	DagId string
}

func (d DAGOrchestrator) ExecuteDAG(DAGId string) error {
	return nil
}

func (d DAGOrchestrator) ExecutePhase(DAGId string, PhaseId string) error {
	return nil
}

func (d DAGOrchestrator) InvokeTaskManager(DAGId string, PhaseId string) error {
	return nil
}

func main() {
	fmt.Println("dag_orchestrator")
}
