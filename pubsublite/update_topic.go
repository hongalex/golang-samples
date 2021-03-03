// Copyright 2021 Google LLC
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

package pslite

import (
	"context"
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/pubsublite"
)

// [START pubsublite_update_topic]

func updateTopic(w io.Writer, projectID, region, zone, topicID string) error {
	// subID := "my-subscription"
	// projectID := "my-project-id"
	// region := "us-central1"
	// zone := "us-central1-a"
	// topicID := "my-topic"
	ctx := context.Background()
	client, err := pubsublite.NewAdminClient(ctx, region)
	if err != nil {
		return fmt.Errorf("pubsublite.NewAdminClient: %v", err)
	}
	defer client.Close()

	topicPath := fmt.Sprintf("projects/%s/locations/%s/topics/%s", projectID, zone, topicID)
	config := pubsublite.TopicConfigToUpdate{
		Name:                       topicPath,
		PartitionCount:             3,                       // Partition count cannot decrease.
		PublishCapacityMiBPerSec:   8,                       // Must be >= 4 and <= 16.
		SubscribeCapacityMiBPerSec: 16,                      // Must be >= 4 and <= 32.
		PerPartitionBytes:          60 * 1024 * 1024 * 1024, // Must be between 30 GiB-10 TiB.
		RetentionDuration:          24 * time.Hour,
	}
	updatedCfg, err := client.UpdateTopic(ctx, config)
	if err != nil {
		return fmt.Errorf("client.UpdateTopic got err: %v", err)
	}
	fmt.Fprintf(w, "Updated topic: %#v\n", *updatedCfg)
	return nil
}

// [END pubsublite_update_topic]
