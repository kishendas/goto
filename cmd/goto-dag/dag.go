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
	"time"
)

type DAG struct {
	DAGId             string
	ClientId          string
	SandboxId         string
	DAGAuthorId       string
	ShardId           string
	Phase             []Stage
	CreatedTime       time.Time
	LastHeartbeatTime time.Time
	LastRunTime       time.Time
	TimeSLA           int64
}

type Stage struct {
	DAGId       string
	PhaseId     string
	Tasks       []Task
	CreatedTime time.Time
	LastRunTime time.Time
	TimeSLA     int64
}

type Task struct {
	DAGId             string
	StageId           string
	TaskId            string
	Query             string
	CreatedTime       time.Time
	LastHeartbeatTime time.Time
	LastRunTime       time.Time
	TimeSLA           int64
}

type RawStage struct {
	index      int
	stage      int
	finalStage int
}
