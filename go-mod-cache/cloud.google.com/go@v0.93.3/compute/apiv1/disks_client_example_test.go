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

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package compute_test

import (
	"context"

	compute "cloud.google.com/go/compute/apiv1"
	"google.golang.org/api/iterator"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
)

func ExampleNewDisksRESTClient() {
	ctx := context.Background()
	c, err := compute.NewDisksRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleDisksClient_AddResourcePolicies() {
	ctx := context.Background()
	c, err := compute.NewDisksRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.AddResourcePoliciesDiskRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.AddResourcePolicies(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDisksClient_AggregatedList() {
	ctx := context.Background()
	c, err := compute.NewDisksRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.AggregatedListDisksRequest{
		// TODO: Fill request struct fields.
	}
	it := c.AggregatedList(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleDisksClient_CreateSnapshot() {
	ctx := context.Background()
	c, err := compute.NewDisksRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.CreateSnapshotDiskRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateSnapshot(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDisksClient_Delete() {
	ctx := context.Background()
	c, err := compute.NewDisksRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.DeleteDiskRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.Delete(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDisksClient_Get() {
	ctx := context.Background()
	c, err := compute.NewDisksRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.GetDiskRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.Get(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDisksClient_GetIamPolicy() {
	ctx := context.Background()
	c, err := compute.NewDisksRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.GetIamPolicyDiskRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDisksClient_Insert() {
	ctx := context.Background()
	c, err := compute.NewDisksRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.InsertDiskRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.Insert(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDisksClient_List() {
	ctx := context.Background()
	c, err := compute.NewDisksRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.ListDisksRequest{
		// TODO: Fill request struct fields.
	}
	it := c.List(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleDisksClient_RemoveResourcePolicies() {
	ctx := context.Background()
	c, err := compute.NewDisksRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.RemoveResourcePoliciesDiskRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.RemoveResourcePolicies(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDisksClient_Resize() {
	ctx := context.Background()
	c, err := compute.NewDisksRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.ResizeDiskRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.Resize(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDisksClient_SetIamPolicy() {
	ctx := context.Background()
	c, err := compute.NewDisksRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.SetIamPolicyDiskRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.SetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDisksClient_SetLabels() {
	ctx := context.Background()
	c, err := compute.NewDisksRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.SetLabelsDiskRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.SetLabels(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDisksClient_TestIamPermissions() {
	ctx := context.Background()
	c, err := compute.NewDisksRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.TestIamPermissionsDiskRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.TestIamPermissions(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}