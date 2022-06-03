// Code generated by protoc-gen-go-aip-test. DO NOT EDIT.

package scheduler

import (
	context "context"
	cmpopts "github.com/google/go-cmp/cmp/cmpopts"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protocmp "google.golang.org/protobuf/testing/protocmp"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	assert "gotest.tools/v3/assert"
	strings "strings"
	testing "testing"
)

type CloudSchedulerTestSuite struct {
	T *testing.T
	// Server to test.
	Server CloudSchedulerServer
}

func (fx CloudSchedulerTestSuite) TestJob(ctx context.Context, options JobTestSuiteConfig) {
	fx.T.Run("Job", func(t *testing.T) {
		options.ctx = ctx
		options.service = fx.Server
		options.test(t)
	})
}

type JobTestSuiteConfig struct {
	ctx        context.Context
	service    CloudSchedulerServer
	currParent int

	// The parents to use when creating resources.
	// At least one parent needs to be set. Depending on methods available on the resource,
	// more may be required. If insufficient number of parents are
	// provided the test will fail.
	Parents []string
	// Create should return a resource which is valid to create, i.e.
	// all required fields set.
	Create func(parent string) *Job
	// Update should return a resource which is valid to update, i.e.
	// all required fields set.
	Update func(parent string) *Job
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *JobTestSuiteConfig) test(t *testing.T) {
	t.Run("Create", fx.testCreate)
	t.Run("Get", fx.testGet)
	t.Run("Update", fx.testUpdate)
	t.Run("List", fx.testList)
}

func (fx *JobTestSuiteConfig) testCreate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no parent is provided.
	t.Run("missing parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreateJob(fx.ctx, &CreateJobRequest{
			Parent: "",
			Job:    fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreateJob(fx.ctx, &CreateJobRequest{
			Parent: "invalid resource name",
			Job:    fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// The created resource should be persisted and reachable with Get.
	t.Run("persisted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg, err := fx.service.CreateJob(fx.ctx, &CreateJobRequest{
			Parent: parent,
			Job:    fx.Create(parent),
		})
		assert.NilError(t, err)
		persisted, err := fx.service.GetJob(fx.ctx, &GetJobRequest{
			Name: msg.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, msg, persisted, protocmp.Transform())
	})

	// The method should fail with InvalidArgument if the resource has any
	// resource references and they are invalid.
	t.Run("resource references", func(t *testing.T) {
		fx.maybeSkip(t)
		t.Run(".pubsub_target.topic_name", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetPubsubTarget()
			if container == nil {
				t.Skip("not reachable")
			}
			container.TopicName = "invalid resource name"
			_, err := fx.service.CreateJob(fx.ctx, &CreateJobRequest{
				Parent: parent,
				Job:    msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
	})

}

func (fx *JobTestSuiteConfig) testGet(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetJob(fx.ctx, &GetJobRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetJob(fx.ctx, &GetJobRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be returned without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created, err := fx.service.CreateJob(fx.ctx, &CreateJobRequest{
			Parent: parent,
			Job:    fx.Create(parent),
		})
		assert.NilError(t, err)
		msg, err := fx.service.GetJob(fx.ctx, &GetJobRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, msg, created, protocmp.Transform())
	})

	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created, err := fx.service.CreateJob(fx.ctx, &CreateJobRequest{
			Parent: parent,
			Job:    fx.Create(parent),
		})
		assert.NilError(t, err)
		_, err = fx.service.GetJob(fx.ctx, &GetJobRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetJob(fx.ctx, &GetJobRequest{
			Name: "projects/-/locations/-/jobs/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *JobTestSuiteConfig) testUpdate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = ""
		_, err := fx.service.UpdateJob(fx.ctx, &UpdateJobRequest{
			Job: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = "invalid resource name"
		_, err := fx.service.UpdateJob(fx.ctx, &UpdateJobRequest{
			Job: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// The updated resource should be persisted and reachable with Get.
	t.Run("persisted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created, err := fx.service.CreateJob(fx.ctx, &CreateJobRequest{
			Parent: parent,
			Job:    fx.Create(parent),
		})
		assert.NilError(t, err)
		updated, err := fx.service.UpdateJob(fx.ctx, &UpdateJobRequest{
			Job: created,
		})
		assert.NilError(t, err)
		persisted, err := fx.service.GetJob(fx.ctx, &GetJobRequest{
			Name: updated.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, updated, persisted, protocmp.Transform())
	})

	parent := fx.nextParent(t, false)
	created, err := fx.service.CreateJob(fx.ctx, &CreateJobRequest{
		Parent: parent,
		Job:    fx.Create(parent),
	})
	assert.NilError(t, err)
	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		msg := fx.Update(parent)
		msg.Name = created.Name + "notfound"
		_, err := fx.service.UpdateJob(fx.ctx, &UpdateJobRequest{
			Job: msg,
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// The method should fail with InvalidArgument if the update_mask is invalid.
	t.Run("invalid update mask", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.UpdateJob(fx.ctx, &UpdateJobRequest{
			Job: created,
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"invalid_field_xyz",
				},
			},
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *JobTestSuiteConfig) testList(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.ListJobs(fx.ctx, &ListJobsRequest{
			Parent: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page token is not valid.
	t.Run("invalid page token", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.service.ListJobs(fx.ctx, &ListJobsRequest{
			Parent:    parent,
			PageToken: "invalid page token",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page size is negative.
	t.Run("negative page size", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.service.ListJobs(fx.ctx, &ListJobsRequest{
			Parent:   parent,
			PageSize: -10,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	const resourcesCount = 15
	parent := fx.nextParent(t, true)
	parentMsgs := make([]*Job, resourcesCount)
	for i := 0; i < resourcesCount; i++ {
		msg, err := fx.service.CreateJob(fx.ctx, &CreateJobRequest{
			Parent: parent,
			Job:    fx.Create(parent),
		})
		assert.NilError(t, err)
		parentMsgs[i] = msg
	}

	// If parent is provided the method must only return resources
	// under that parent.
	t.Run("isolation", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.ListJobs(fx.ctx, &ListJobsRequest{
			Parent:   parent,
			PageSize: 999,
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			parentMsgs,
			response.Jobs,
			cmpopts.SortSlices(func(a, b *Job) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

	// If there are no more resources, next_page_token should not be set.
	t.Run("last page", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.ListJobs(fx.ctx, &ListJobsRequest{
			Parent:   parent,
			PageSize: resourcesCount,
		})
		assert.NilError(t, err)
		assert.Equal(t, "", response.NextPageToken)
	})

	// If there are more resources, next_page_token should be set.
	t.Run("more pages", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.ListJobs(fx.ctx, &ListJobsRequest{
			Parent:   parent,
			PageSize: resourcesCount - 1,
		})
		assert.NilError(t, err)
		assert.Check(t, response.NextPageToken != "")
	})

	// Listing resource one by one should eventually return all resources.
	t.Run("one by one", func(t *testing.T) {
		fx.maybeSkip(t)
		msgs := make([]*Job, 0, resourcesCount)
		var nextPageToken string
		for {
			response, err := fx.service.ListJobs(fx.ctx, &ListJobsRequest{
				Parent:    parent,
				PageSize:  1,
				PageToken: nextPageToken,
			})
			assert.NilError(t, err)
			assert.Equal(t, 1, len(response.Jobs))
			msgs = append(msgs, response.Jobs...)
			nextPageToken = response.NextPageToken
			if nextPageToken == "" {
				break
			}
		}
		assert.DeepEqual(
			t,
			parentMsgs,
			msgs,
			cmpopts.SortSlices(func(a, b *Job) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

	// Method should not return deleted resources.
	t.Run("deleted", func(t *testing.T) {
		fx.maybeSkip(t)
		const deleteCount = 5
		for i := 0; i < deleteCount; i++ {
			_, err := fx.service.DeleteJob(fx.ctx, &DeleteJobRequest{
				Name: parentMsgs[i].Name,
			})
			assert.NilError(t, err)
		}
		response, err := fx.service.ListJobs(fx.ctx, &ListJobsRequest{
			Parent:   parent,
			PageSize: 9999,
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			parentMsgs[deleteCount:],
			response.Jobs,
			cmpopts.SortSlices(func(a, b *Job) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

}

func (fx *JobTestSuiteConfig) nextParent(t *testing.T, pristine bool) string {
	if pristine {
		fx.currParent++
	}
	if fx.currParent >= len(fx.Parents) {
		t.Fatal("need at least", fx.currParent+1, "parents")
	}
	return fx.Parents[fx.currParent]
}

func (fx *JobTestSuiteConfig) peekNextParent(t *testing.T) string {
	next := fx.currParent + 1
	if next >= len(fx.Parents) {
		t.Fatal("need at least", next+1, "parents")
	}
	return fx.Parents[next]
}

func (fx *JobTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}
