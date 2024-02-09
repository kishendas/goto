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
	"context"
	"database/sql"
	"fmt"
)

func ExecuteTx(
	ctx context.Context, opts *sql.TxOptions,
) error {
	return nil
}

func Exec(_ context.Context, q string, args ...interface{}) error {
	return nil
}

func Commit(_ context.Context) error {
	return nil
}

func Rollback(_ context.Context) error {
	return nil
}

func main() {
	fmt.Println("db_client")
}
