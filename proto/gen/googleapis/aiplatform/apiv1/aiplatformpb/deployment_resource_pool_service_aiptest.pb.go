// Code generated by protoc-gen-go-aip-test. DO NOT EDIT.

package aiplatformpb

import (
	context "context"
	cmpopts "github.com/google/go-cmp/cmp/cmpopts"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protocmp "google.golang.org/protobuf/testing/protocmp"
	assert "gotest.tools/v3/assert"
	strings "strings"
	testing "testing"
)

// DeploymentResourcePoolServiceTestSuiteConfigProvider is the interface to implement to decide which resources
// that should be tested and configured.
type DeploymentResourcePoolServiceTestSuiteConfigProvider interface {
	DeploymentResourcePoolTestSuiteConfig(t *testing.T) *DeploymentResourcePoolServiceDeploymentResourcePoolTestSuiteConfig
}

// TestDeploymentResourcePoolService is the main entrypoint for starting the AIP tests.
func TestDeploymentResourcePoolService(t *testing.T, s DeploymentResourcePoolServiceTestSuiteConfigProvider) {
	testDeploymentResourcePoolServiceDeploymentResourcePoolTestSuiteConfig(t, s)
}

func testDeploymentResourcePoolServiceDeploymentResourcePoolTestSuiteConfig(t *testing.T, s DeploymentResourcePoolServiceTestSuiteConfigProvider) {
	t.Run("DeploymentResourcePool", func(t *testing.T) {
		config := s.DeploymentResourcePoolTestSuiteConfig(t)
		if config == nil {
			t.Skip("Method DeploymentResourcePoolTestSuiteConfig not implemented")
		}
		if config.Service == nil {
			t.Skip("Method DeploymentResourcePoolServiceDeploymentResourcePoolTestSuiteConfig.Service() not implemented")
		}
		if config.Context == nil {
			config.Context = func() context.Context { return context.Background() }
		}
		config.test(t)
	})
}

type DeploymentResourcePoolServiceTestSuite struct {
	T *testing.T
	// Server to test.
	Server DeploymentResourcePoolServiceServer
}

func (fx DeploymentResourcePoolServiceTestSuite) TestDeploymentResourcePool(ctx context.Context, options DeploymentResourcePoolServiceDeploymentResourcePoolTestSuiteConfig) {
	fx.T.Run("DeploymentResourcePool", func(t *testing.T) {
		options.Context = func() context.Context { return ctx }
		options.Service = func() DeploymentResourcePoolServiceServer { return fx.Server }
		options.test(t)
	})
}

type DeploymentResourcePoolServiceDeploymentResourcePoolTestSuiteConfig struct {
	currParent int

	// Service should return the service that should be tested.
	// The service will be used for several tests.
	Service func() DeploymentResourcePoolServiceServer
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
	Create func(parent string) *DeploymentResourcePool
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *DeploymentResourcePoolServiceDeploymentResourcePoolTestSuiteConfig) test(t *testing.T) {
	t.Run("Create", fx.testCreate)
	t.Run("Get", fx.testGet)
	t.Run("List", fx.testList)
	t.Run("Delete", fx.testDelete)
}

func (fx *DeploymentResourcePoolServiceDeploymentResourcePoolTestSuiteConfig) testCreate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no parent is provided.
	t.Run("missing parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().CreateDeploymentResourcePool(fx.Context(), &CreateDeploymentResourcePoolRequest{
			Parent:                 "",
			DeploymentResourcePool: fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().CreateDeploymentResourcePool(fx.Context(), &CreateDeploymentResourcePoolRequest{
			Parent:                 "invalid resource name",
			DeploymentResourcePool: fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// The method should fail with InvalidArgument if the resource has any
	// required fields and they are not provided.
	t.Run("required fields", func(t *testing.T) {
		fx.maybeSkip(t)
		t.Run(".dedicated_resources", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("dedicated_resources")
			container.ProtoReflect().Clear(fd)
			_, err := fx.Service().CreateDeploymentResourcePool(fx.Context(), &CreateDeploymentResourcePoolRequest{
				Parent:                 parent,
				DeploymentResourcePool: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".dedicated_resources.machine_spec", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetDedicatedResources()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("machine_spec")
			container.ProtoReflect().Clear(fd)
			_, err := fx.Service().CreateDeploymentResourcePool(fx.Context(), &CreateDeploymentResourcePoolRequest{
				Parent:                 parent,
				DeploymentResourcePool: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".dedicated_resources.min_replica_count", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetDedicatedResources()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("min_replica_count")
			container.ProtoReflect().Clear(fd)
			_, err := fx.Service().CreateDeploymentResourcePool(fx.Context(), &CreateDeploymentResourcePoolRequest{
				Parent:                 parent,
				DeploymentResourcePool: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
	})

}

func (fx *DeploymentResourcePoolServiceDeploymentResourcePoolTestSuiteConfig) testGet(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetDeploymentResourcePool(fx.Context(), &GetDeploymentResourcePoolRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetDeploymentResourcePool(fx.Context(), &GetDeploymentResourcePoolRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be returned without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		msg, err := fx.Service().GetDeploymentResourcePool(fx.Context(), &GetDeploymentResourcePoolRequest{
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
		_, err := fx.Service().GetDeploymentResourcePool(fx.Context(), &GetDeploymentResourcePoolRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetDeploymentResourcePool(fx.Context(), &GetDeploymentResourcePoolRequest{
			Name: "projects/-/locations/-/deploymentResourcePools/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *DeploymentResourcePoolServiceDeploymentResourcePoolTestSuiteConfig) testList(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().ListDeploymentResourcePools(fx.Context(), &ListDeploymentResourcePoolsRequest{
			Parent: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page token is not valid.
	t.Run("invalid page token", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.Service().ListDeploymentResourcePools(fx.Context(), &ListDeploymentResourcePoolsRequest{
			Parent:    parent,
			PageToken: "invalid page token",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page size is negative.
	t.Run("negative page size", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.Service().ListDeploymentResourcePools(fx.Context(), &ListDeploymentResourcePoolsRequest{
			Parent:   parent,
			PageSize: -10,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	const resourcesCount = 15
	parent := fx.nextParent(t, true)
	parentMsgs := make([]*DeploymentResourcePool, resourcesCount)
	for i := 0; i < resourcesCount; i++ {
		parentMsgs[i] = fx.create(t, parent)
	}

	// If parent is provided the method must only return resources
	// under that parent.
	t.Run("isolation", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.Service().ListDeploymentResourcePools(fx.Context(), &ListDeploymentResourcePoolsRequest{
			Parent:   parent,
			PageSize: 999,
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			parentMsgs,
			response.DeploymentResourcePools,
			cmpopts.SortSlices(func(a, b *DeploymentResourcePool) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

	// If there are no more resources, next_page_token should not be set.
	t.Run("last page", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.Service().ListDeploymentResourcePools(fx.Context(), &ListDeploymentResourcePoolsRequest{
			Parent:   parent,
			PageSize: resourcesCount,
		})
		assert.NilError(t, err)
		assert.Equal(t, "", response.NextPageToken)
	})

	// If there are more resources, next_page_token should be set.
	t.Run("more pages", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.Service().ListDeploymentResourcePools(fx.Context(), &ListDeploymentResourcePoolsRequest{
			Parent:   parent,
			PageSize: resourcesCount - 1,
		})
		assert.NilError(t, err)
		assert.Check(t, response.NextPageToken != "")
	})

	// Listing resource one by one should eventually return all resources.
	t.Run("one by one", func(t *testing.T) {
		fx.maybeSkip(t)
		msgs := make([]*DeploymentResourcePool, 0, resourcesCount)
		var nextPageToken string
		for {
			response, err := fx.Service().ListDeploymentResourcePools(fx.Context(), &ListDeploymentResourcePoolsRequest{
				Parent:    parent,
				PageSize:  1,
				PageToken: nextPageToken,
			})
			assert.NilError(t, err)
			assert.Equal(t, 1, len(response.DeploymentResourcePools))
			msgs = append(msgs, response.DeploymentResourcePools...)
			nextPageToken = response.NextPageToken
			if nextPageToken == "" {
				break
			}
		}
		assert.DeepEqual(
			t,
			parentMsgs,
			msgs,
			cmpopts.SortSlices(func(a, b *DeploymentResourcePool) bool {
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
			_, err := fx.Service().DeleteDeploymentResourcePool(fx.Context(), &DeleteDeploymentResourcePoolRequest{
				Name: parentMsgs[i].Name,
			})
			assert.NilError(t, err)
		}
		response, err := fx.Service().ListDeploymentResourcePools(fx.Context(), &ListDeploymentResourcePoolsRequest{
			Parent:   parent,
			PageSize: 9999,
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			parentMsgs[deleteCount:],
			response.DeploymentResourcePools,
			cmpopts.SortSlices(func(a, b *DeploymentResourcePool) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

}

func (fx *DeploymentResourcePoolServiceDeploymentResourcePoolTestSuiteConfig) testDelete(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().DeleteDeploymentResourcePool(fx.Context(), &DeleteDeploymentResourcePoolRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().DeleteDeploymentResourcePool(fx.Context(), &DeleteDeploymentResourcePoolRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be deleted without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		_, err := fx.Service().DeleteDeploymentResourcePool(fx.Context(), &DeleteDeploymentResourcePoolRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
	})

	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		_, err := fx.Service().DeleteDeploymentResourcePool(fx.Context(), &DeleteDeploymentResourcePoolRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with NotFound if the resource was already deleted. This also applies to soft-deletion.
	t.Run("already deleted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		deleted, err := fx.Service().DeleteDeploymentResourcePool(fx.Context(), &DeleteDeploymentResourcePoolRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
		_ = deleted
		_, err = fx.Service().DeleteDeploymentResourcePool(fx.Context(), &DeleteDeploymentResourcePoolRequest{
			Name: created.Name,
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().DeleteDeploymentResourcePool(fx.Context(), &DeleteDeploymentResourcePoolRequest{
			Name: "projects/-/locations/-/deploymentResourcePools/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *DeploymentResourcePoolServiceDeploymentResourcePoolTestSuiteConfig) nextParent(t *testing.T, pristine bool) string {
	if pristine {
		fx.currParent++
	}
	if fx.currParent >= len(fx.Parents) {
		t.Fatal("need at least", fx.currParent+1, "parents")
	}
	return fx.Parents[fx.currParent]
}

func (fx *DeploymentResourcePoolServiceDeploymentResourcePoolTestSuiteConfig) peekNextParent(t *testing.T) string {
	next := fx.currParent + 1
	if next >= len(fx.Parents) {
		t.Fatal("need at least", next+1, "parents")
	}
	return fx.Parents[next]
}

func (fx *DeploymentResourcePoolServiceDeploymentResourcePoolTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}

func (fx *DeploymentResourcePoolServiceDeploymentResourcePoolTestSuiteConfig) create(t *testing.T, parent string) *DeploymentResourcePool {
	t.Helper()
	t.Skip("Long running create method not supported")
	return nil
}
