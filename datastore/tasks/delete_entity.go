// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

// [START datastore_delete_entity]
import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
)

// DeleteTask deletes the task with the given ID.
func DeleteTask(projectID string, databaseID string, taskID int64) error {
	ctx := context.Background()
	client, err := datastore.NewClientWithDatabase(ctx, projectID, databaseID)
	if err != nil {
		log.Fatalf("Could not create datastore client: %v", err)
	}
	return client.Delete(ctx, datastore.IDKey("Task", taskID, nil))
}

// [END datastore_delete_entity]
