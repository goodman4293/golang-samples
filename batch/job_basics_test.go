// Copyright 2022 Google LLC
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

package snippets

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/golang-samples/internal/testutil"
)

func TestBatchJobCRUD(t *testing.T) {
	var r *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	tc := testutil.SystemTest(t)
	region := "us-central1"
	jobName := fmt.Sprintf("test-job-go-script-%v-%v", time.Now().Format("2006-12-25"), r.Int())

	buf := &bytes.Buffer{}

	if err := createScriptJob(buf, tc.ProjectID, region, jobName); err != nil {
		t.Errorf("createScriptJob got err: %v", err)
	}

	buf.Reset()

	if err := getJob(buf, tc.ProjectID, region, jobName); err != nil {
		t.Errorf("getJob got err: %v", err)
	}

	buf.Reset()

	if err := listJobs(buf, tc.ProjectID, region); err != nil {
		t.Errorf("listJobs got err: %v", err)
	}
	if got := buf.String(); !strings.Contains(got, jobName) {
		t.Errorf("listJobs got %q, expected %q", got, jobName)
	}

	buf.Reset()

	// Tasks take a couple of seconds to be created on the server side.
	// We're going to poll until they're created, or give up if the errors are persistent.
	var attempts uint = 0
	var loop_err = getTask(buf, tc.ProjectID, region, jobName, "group0", 0)
	for loop_err != nil {
		attempts += 1
		// tasks usually appear in a couple of seconds, 20 seconds is way more than enough
		if attempts > 20 {
			t.Errorf("getTask got err: %v", loop_err)
			break
		}
		time.Sleep(1 * time.Second)
		loop_err = getTask(buf, tc.ProjectID, region, jobName, "group0", 0)
	}
	if got := buf.String(); !strings.Contains(got, "status:") {
		t.Errorf("getTask got %q, expected %q", got, "status:")
	}

	buf.Reset()

	if err := listTasks(buf, tc.ProjectID, region, jobName, "group0"); err != nil {
		t.Errorf("listTasks got err: %v", err)
	}
	if got := buf.String(); !strings.Contains(got, "status:") {
		t.Errorf("listTasks got %q, expected %q", got, "status:")
	}

	buf.Reset()

	if err := deleteJob(buf, tc.ProjectID, region, jobName); err != nil {
		t.Errorf("deleteJob got err: %v", err)
	}
}

func TestBatchContainerJob(t *testing.T) {
	var r *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	tc := testutil.SystemTest(t)
	region := "us-central1"
	jobName := fmt.Sprintf("test-job-go-docker-%v-%v", time.Now().Format("2006-12-25"), r.Int())

	buf := &bytes.Buffer{}

	if err := createContainerJob(buf, tc.ProjectID, region, jobName); err != nil {
		t.Errorf("createContainerJob got err: %v", err)
	}
}
