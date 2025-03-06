// Code generated by protoc-gen-go-aip-test. DO NOT EDIT.

package schedulerpb

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

// CloudSchedulerTestSuiteConfigProvider is the interface to implement to decide which resources
// that should be tested and how it's configured.
type CloudSchedulerTestSuiteConfigProvider interface {
	// CloudSchedulerJob should return a config, or nil, which means that the tests will be skipped.
	CloudSchedulerJob(t *testing.T) *CloudSchedulerJobTestSuiteConfig
}

// testCloudScheduler is the main entrypoint for starting the AIP tests.
func testCloudScheduler(t *testing.T, s CloudSchedulerTestSuiteConfigProvider) {
	testCloudSchedulerJob(t, s)
}

func testCloudSchedulerJob(t *testing.T, s CloudSchedulerTestSuiteConfigProvider) {
	t.Run("Job", func(t *testing.T) {
		config := s.CloudSchedulerJob(t)
		if config == nil {
			t.Skip("Method CloudSchedulerJob not implemented")
		}
		if config.Service == nil {
			t.Skip("Method CloudSchedulerJob.Service() not implemented")
		}
		if config.Context == nil {
			config.Context = func() context.Context { return context.Background() }
		}
		config.test(t)
	})
}

type CloudSchedulerTestSuite struct {
	T *testing.T
	// Server to test.
	Server CloudSchedulerServer
}

func (fx CloudSchedulerTestSuite) TestJob(ctx context.Context, options CloudSchedulerJobTestSuiteConfig) {
	fx.T.Run("Job", func(t *testing.T) {
		options.Context = func() context.Context { return ctx }
		options.Service = func() CloudSchedulerServer { return fx.Server }
		options.test(t)
	})
}

type CloudSchedulerJobTestSuiteConfig struct {
	currParent int

	// Service should return the service that should be tested.
	// The service will be used for several tests.
	Service func() CloudSchedulerServer
	// Context should return a new context.
	// The context will be used for several tests.
	Context func() context.Context
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

func (fx *CloudSchedulerJobTestSuiteConfig) test(t *testing.T) {
	t.Run("Create", fx.testCreate)
	t.Run("Get", fx.testGet)
	t.Run("Update", fx.testUpdate)
	t.Run("List", fx.testList)
	t.Run("Delete", fx.testDelete)
}

func (fx *CloudSchedulerJobTestSuiteConfig) testCreate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no parent is provided.
	t.Run("missing parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().CreateJob(fx.Context(), &CreateJobRequest{
			Parent: "",
			Job:    fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().CreateJob(fx.Context(), &CreateJobRequest{
			Parent: "invalid resource name",
			Job:    fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// The created resource should be persisted and reachable with Get.
	t.Run("persisted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg, err := fx.Service().CreateJob(fx.Context(), &CreateJobRequest{
			Parent: parent,
			Job:    fx.Create(parent),
		})
		assert.NilError(t, err)
		persisted, err := fx.Service().GetJob(fx.Context(), &GetJobRequest{
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
			_, err := fx.Service().CreateJob(fx.Context(), &CreateJobRequest{
				Parent: parent,
				Job:    msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
	})

}

func (fx *CloudSchedulerJobTestSuiteConfig) testGet(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetJob(fx.Context(), &GetJobRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetJob(fx.Context(), &GetJobRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be returned without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		msg, err := fx.Service().GetJob(fx.Context(), &GetJobRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, msg, created, protocmp.Transform())
	})

	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		_, err := fx.Service().GetJob(fx.Context(), &GetJobRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetJob(fx.Context(), &GetJobRequest{
			Name: "projects/-/locations/-/jobs/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *CloudSchedulerJobTestSuiteConfig) testUpdate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = ""
		_, err := fx.Service().UpdateJob(fx.Context(), &UpdateJobRequest{
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
		_, err := fx.Service().UpdateJob(fx.Context(), &UpdateJobRequest{
			Job: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// The updated resource should be persisted and reachable with Get.
	t.Run("persisted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		updated, err := fx.Service().UpdateJob(fx.Context(), &UpdateJobRequest{
			Job: created,
		})
		assert.NilError(t, err)
		persisted, err := fx.Service().GetJob(fx.Context(), &GetJobRequest{
			Name: updated.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, updated, persisted, protocmp.Transform())
	})

	{
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		// Method should fail with NotFound if the resource does not exist.
		t.Run("not found", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := fx.Update(parent)
			msg.Name = created.Name + "notfound"
			_, err := fx.Service().UpdateJob(fx.Context(), &UpdateJobRequest{
				Job: msg,
			})
			assert.Equal(t, codes.NotFound, status.Code(err), err)
		})

		// The method should fail with InvalidArgument if the update_mask is invalid.
		t.Run("invalid update mask", func(t *testing.T) {
			fx.maybeSkip(t)
			_, err := fx.Service().UpdateJob(fx.Context(), &UpdateJobRequest{
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
}

func (fx *CloudSchedulerJobTestSuiteConfig) testList(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().ListJobs(fx.Context(), &ListJobsRequest{
			Parent: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page token is not valid.
	t.Run("invalid page token", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.Service().ListJobs(fx.Context(), &ListJobsRequest{
			Parent:    parent,
			PageToken: "invalid page token",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page size is negative.
	t.Run("negative page size", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.Service().ListJobs(fx.Context(), &ListJobsRequest{
			Parent:   parent,
			PageSize: -10,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	{
		const resourcesCount = 15
		parent := fx.nextParent(t, true)
		parentMsgs := make([]*Job, resourcesCount)
		for i := 0; i < resourcesCount; i++ {
			parentMsgs[i] = fx.create(t, parent)
		}

		// If parent is provided the method must only return resources
		// under that parent.
		t.Run("isolation", func(t *testing.T) {
			fx.maybeSkip(t)
			response, err := fx.Service().ListJobs(fx.Context(), &ListJobsRequest{
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
			response, err := fx.Service().ListJobs(fx.Context(), &ListJobsRequest{
				Parent:   parent,
				PageSize: resourcesCount,
			})
			assert.NilError(t, err)
			assert.Equal(t, "", response.NextPageToken)
		})

		// If there are more resources, next_page_token should be set.
		t.Run("more pages", func(t *testing.T) {
			fx.maybeSkip(t)
			response, err := fx.Service().ListJobs(fx.Context(), &ListJobsRequest{
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
				response, err := fx.Service().ListJobs(fx.Context(), &ListJobsRequest{
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
				_, err := fx.Service().DeleteJob(fx.Context(), &DeleteJobRequest{
					Name: parentMsgs[i].Name,
				})
				assert.NilError(t, err)
			}
			response, err := fx.Service().ListJobs(fx.Context(), &ListJobsRequest{
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
}

func (fx *CloudSchedulerJobTestSuiteConfig) testDelete(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().DeleteJob(fx.Context(), &DeleteJobRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().DeleteJob(fx.Context(), &DeleteJobRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be deleted without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		_, err := fx.Service().DeleteJob(fx.Context(), &DeleteJobRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
	})

	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		_, err := fx.Service().DeleteJob(fx.Context(), &DeleteJobRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with NotFound if the resource was already deleted. This also applies to soft-deletion.
	t.Run("already deleted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		deleted, err := fx.Service().DeleteJob(fx.Context(), &DeleteJobRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
		_ = deleted
		_, err = fx.Service().DeleteJob(fx.Context(), &DeleteJobRequest{
			Name: created.Name,
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().DeleteJob(fx.Context(), &DeleteJobRequest{
			Name: "projects/-/locations/-/jobs/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *CloudSchedulerJobTestSuiteConfig) nextParent(t *testing.T, pristine bool) string {
	if pristine {
		fx.currParent++
	}
	if fx.currParent >= len(fx.Parents) {
		t.Fatal("need at least", fx.currParent+1, "parents")
	}
	return fx.Parents[fx.currParent]
}

func (fx *CloudSchedulerJobTestSuiteConfig) peekNextParent(t *testing.T) string {
	next := fx.currParent + 1
	if next >= len(fx.Parents) {
		t.Fatal("need at least", next+1, "parents")
	}
	return fx.Parents[next]
}

func (fx *CloudSchedulerJobTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}

func (fx *CloudSchedulerJobTestSuiteConfig) create(t *testing.T, parent string) *Job {
	t.Helper()
	created, err := fx.Service().CreateJob(fx.Context(), &CreateJobRequest{
		Parent: parent,
		Job:    fx.Create(parent),
	})
	assert.NilError(t, err)
	return created
}
