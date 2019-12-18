// Copyright 2019 Google LLC
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

package querying

// [START bigquery_query_clustered_table]
import (
	"cloud.google.com/go/bigquery"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"io"
)

// queryClusteredTable demonstrates querying a table that has a clustering specification.
func queryClusteredTable(w io.Writer, projectID, datasetID, tableID string) error {
	// projectID := "my-project-id"
	// datasetID := "mydataset"
	// tableID := "mytable"
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("bigquery.NewClient: %v", err)
	}

	q := client.Query(fmt.Sprintf(`
	SELECT
	  COUNT(1) as transactions,
	  SUM(amount) as total_paid,
	  COUNT(DISTINCT destination) as distinct_recipients
    FROM
	  `+"`%s.%s`"+`
	 WHERE
	    timestamp > TIMESTAMP('2015-01-01')
		AND origin = @wallet`, datasetID, tableID))
	q.Parameters = []bigquery.QueryParameter{
		{
			Name:  "wallet",
			Value: "wallet00001866cb7e0f09a890",
		},
	}
	// Run the query and print results when the query job is completed.
	job, err := q.Run(ctx)
	if err != nil {
		return err
	}
	status, err := job.Wait(ctx)
	if err != nil {
		return err
	}
	if err := status.Err(); err != nil {
		return err
	}
	it, err := job.Read(ctx)
	for {
		var row []bigquery.Value
		err := it.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		fmt.Fprintln(w, row)
	}
	return nil
}

// [END bigquery_query_clustered_table]
