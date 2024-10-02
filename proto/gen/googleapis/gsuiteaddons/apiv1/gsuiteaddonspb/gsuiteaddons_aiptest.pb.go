// Code generated by protoc-gen-go-aip-test. DO NOT EDIT.

package gsuiteaddonspb

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

func TestGSuiteAddOns(
	t *testing.T,
	s GSuiteAddOnsTestsConfigSupplier,
) {
	{
		cfg := s.TestAuthorization(t)
		fx := GSuiteAddOnsTestSuite{
			T:      t,
			Server: cfg.Server(),
		}
		fx.TestAuthorization(cfg.Context(), *cfg)
	}
	{
		cfg := s.TestDeployment(t)
		fx := GSuiteAddOnsTestSuite{
			T:      t,
			Server: cfg.Server(),
		}
		fx.TestDeployment(cfg.Context(), *cfg)
	}
	{
		cfg := s.TestInstallStatus(t)
		fx := GSuiteAddOnsTestSuite{
			T:      t,
			Server: cfg.Server(),
		}
		fx.TestInstallStatus(cfg.Context(), *cfg)
	}
}

type GSuiteAddOnsTestsConfigSupplier interface {
	TestAuthorization(t *testing.T) *GSuiteAddOnsAuthorizationTestSuiteConfig
	TestDeployment(t *testing.T) *GSuiteAddOnsDeploymentTestSuiteConfig
	TestInstallStatus(t *testing.T) *GSuiteAddOnsInstallStatusTestSuiteConfig
}
type GSuiteAddOnsTestSuite struct {
	T *testing.T
	// Server to test.
	Server GSuiteAddOnsServer
}

func (fx GSuiteAddOnsTestSuite) TestAuthorization(ctx context.Context, options GSuiteAddOnsAuthorizationTestSuiteConfig) {
	fx.T.Run("Authorization", func(t *testing.T) {
		options.ctx = ctx
		options.service = fx.Server
		options.test(t)
	})
}

func (fx GSuiteAddOnsTestSuite) TestDeployment(ctx context.Context, options GSuiteAddOnsDeploymentTestSuiteConfig) {
	fx.T.Run("Deployment", func(t *testing.T) {
		options.ctx = ctx
		options.service = fx.Server
		options.test(t)
	})
}

func (fx GSuiteAddOnsTestSuite) TestInstallStatus(ctx context.Context, options GSuiteAddOnsInstallStatusTestSuiteConfig) {
	fx.T.Run("InstallStatus", func(t *testing.T) {
		options.ctx = ctx
		options.service = fx.Server
		options.test(t)
	})
}

type GSuiteAddOnsAuthorizationTestSuiteConfig struct {
	ctx        context.Context
	service    GSuiteAddOnsServer
	currParent int

	Server func() GSuiteAddOnsServer
	// Context should return a new context that can be used for each test.
	Context func() context.Context
	// CreateResource should create a Authorization and return it.
	// If the field is not set, some tests will be skipped.
	//
	// This method is generated because service does not expose a Create
	// method (or it does not comply with AIP).
	CreateResource func(ctx context.Context) (*Authorization, error)
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *GSuiteAddOnsAuthorizationTestSuiteConfig) test(t *testing.T) {
	t.Run("Get", fx.testGet)
}

func (fx *GSuiteAddOnsAuthorizationTestSuiteConfig) testGet(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetAuthorization(fx.ctx, &GetAuthorizationRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetAuthorization(fx.ctx, &GetAuthorizationRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be returned without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		created := fx.create(t)
		msg, err := fx.service.GetAuthorization(fx.ctx, &GetAuthorizationRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, msg, created, protocmp.Transform())
	})

	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		created := fx.create(t)
		_, err := fx.service.GetAuthorization(fx.ctx, &GetAuthorizationRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetAuthorization(fx.ctx, &GetAuthorizationRequest{
			Name: "projects/-/authorization",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *GSuiteAddOnsAuthorizationTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}

func (fx *GSuiteAddOnsAuthorizationTestSuiteConfig) create(t *testing.T) *Authorization {
	t.Helper()
	if fx.CreateResource == nil {
		t.Skip("Test skipped because CreateResource not specified on GSuiteAddOnsAuthorizationTestSuiteConfig")
	}
	created, err := fx.CreateResource(fx.ctx)
	assert.NilError(t, err)
	return created
}

type GSuiteAddOnsDeploymentTestSuiteConfig struct {
	ctx        context.Context
	service    GSuiteAddOnsServer
	currParent int

	Server func() GSuiteAddOnsServer
	// Context should return a new context that can be used for each test.
	Context func() context.Context
	// The parents to use when creating resources.
	// At least one parent needs to be set. Depending on methods available on the resource,
	// more may be required. If insufficient number of parents are
	// provided the test will fail.
	Parents []string
	// Create should return a resource which is valid to create, i.e.
	// all required fields set.
	Create func(parent string) *Deployment
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *GSuiteAddOnsDeploymentTestSuiteConfig) test(t *testing.T) {
	t.Run("Create", fx.testCreate)
	t.Run("Get", fx.testGet)
	t.Run("List", fx.testList)
	t.Run("Delete", fx.testDelete)
}

func (fx *GSuiteAddOnsDeploymentTestSuiteConfig) testCreate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no parent is provided.
	t.Run("missing parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreateDeployment(fx.ctx, &CreateDeploymentRequest{
			Parent:     "",
			Deployment: fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreateDeployment(fx.ctx, &CreateDeploymentRequest{
			Parent:     "invalid resource name",
			Deployment: fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// The created resource should be persisted and reachable with Get.
	t.Run("persisted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg, err := fx.service.CreateDeployment(fx.ctx, &CreateDeploymentRequest{
			Parent:     parent,
			Deployment: fx.Create(parent),
		})
		assert.NilError(t, err)
		persisted, err := fx.service.GetDeployment(fx.ctx, &GetDeploymentRequest{
			Name: msg.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, msg, persisted, protocmp.Transform())
	})

	// The method should fail with InvalidArgument if the resource has any
	// required fields and they are not provided.
	t.Run("required fields", func(t *testing.T) {
		fx.maybeSkip(t)
		t.Run(".add_ons.calendar.event_open_trigger.run_function", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetAddOns().GetCalendar().GetEventOpenTrigger()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("run_function")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateDeployment(fx.ctx, &CreateDeploymentRequest{
				Parent:     parent,
				Deployment: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".add_ons.docs.on_file_scope_granted_trigger.run_function", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetAddOns().GetDocs().GetOnFileScopeGrantedTrigger()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("run_function")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateDeployment(fx.ctx, &CreateDeploymentRequest{
				Parent:     parent,
				Deployment: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".add_ons.sheets.on_file_scope_granted_trigger.run_function", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetAddOns().GetSheets().GetOnFileScopeGrantedTrigger()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("run_function")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateDeployment(fx.ctx, &CreateDeploymentRequest{
				Parent:     parent,
				Deployment: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".add_ons.slides.on_file_scope_granted_trigger.run_function", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetAddOns().GetSlides().GetOnFileScopeGrantedTrigger()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("run_function")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateDeployment(fx.ctx, &CreateDeploymentRequest{
				Parent:     parent,
				Deployment: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
	})

	// Field etag should be populated when the resource is created.
	t.Run("etag populated", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created, _ := fx.service.CreateDeployment(fx.ctx, &CreateDeploymentRequest{
			Parent:     parent,
			Deployment: fx.Create(parent),
		})
		assert.Check(t, created.Etag != "")
	})

}

func (fx *GSuiteAddOnsDeploymentTestSuiteConfig) testGet(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetDeployment(fx.ctx, &GetDeploymentRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetDeployment(fx.ctx, &GetDeploymentRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be returned without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		msg, err := fx.service.GetDeployment(fx.ctx, &GetDeploymentRequest{
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
		_, err := fx.service.GetDeployment(fx.ctx, &GetDeploymentRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetDeployment(fx.ctx, &GetDeploymentRequest{
			Name: "projects/-/deployments/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *GSuiteAddOnsDeploymentTestSuiteConfig) testList(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.ListDeployments(fx.ctx, &ListDeploymentsRequest{
			Parent: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page token is not valid.
	t.Run("invalid page token", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.service.ListDeployments(fx.ctx, &ListDeploymentsRequest{
			Parent:    parent,
			PageToken: "invalid page token",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page size is negative.
	t.Run("negative page size", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.service.ListDeployments(fx.ctx, &ListDeploymentsRequest{
			Parent:   parent,
			PageSize: -10,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	const resourcesCount = 15
	parent := fx.nextParent(t, true)
	parentMsgs := make([]*Deployment, resourcesCount)
	for i := 0; i < resourcesCount; i++ {
		parentMsgs[i] = fx.create(t, parent)
	}

	// If parent is provided the method must only return resources
	// under that parent.
	t.Run("isolation", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.ListDeployments(fx.ctx, &ListDeploymentsRequest{
			Parent:   parent,
			PageSize: 999,
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			parentMsgs,
			response.Deployments,
			cmpopts.SortSlices(func(a, b *Deployment) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

	// If there are no more resources, next_page_token should not be set.
	t.Run("last page", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.ListDeployments(fx.ctx, &ListDeploymentsRequest{
			Parent:   parent,
			PageSize: resourcesCount,
		})
		assert.NilError(t, err)
		assert.Equal(t, "", response.NextPageToken)
	})

	// If there are more resources, next_page_token should be set.
	t.Run("more pages", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.ListDeployments(fx.ctx, &ListDeploymentsRequest{
			Parent:   parent,
			PageSize: resourcesCount - 1,
		})
		assert.NilError(t, err)
		assert.Check(t, response.NextPageToken != "")
	})

	// Listing resource one by one should eventually return all resources.
	t.Run("one by one", func(t *testing.T) {
		fx.maybeSkip(t)
		msgs := make([]*Deployment, 0, resourcesCount)
		var nextPageToken string
		for {
			response, err := fx.service.ListDeployments(fx.ctx, &ListDeploymentsRequest{
				Parent:    parent,
				PageSize:  1,
				PageToken: nextPageToken,
			})
			assert.NilError(t, err)
			assert.Equal(t, 1, len(response.Deployments))
			msgs = append(msgs, response.Deployments...)
			nextPageToken = response.NextPageToken
			if nextPageToken == "" {
				break
			}
		}
		assert.DeepEqual(
			t,
			parentMsgs,
			msgs,
			cmpopts.SortSlices(func(a, b *Deployment) bool {
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
			_, err := fx.service.DeleteDeployment(fx.ctx, &DeleteDeploymentRequest{
				Name: parentMsgs[i].Name,
			})
			assert.NilError(t, err)
		}
		response, err := fx.service.ListDeployments(fx.ctx, &ListDeploymentsRequest{
			Parent:   parent,
			PageSize: 9999,
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			parentMsgs[deleteCount:],
			response.Deployments,
			cmpopts.SortSlices(func(a, b *Deployment) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

}

func (fx *GSuiteAddOnsDeploymentTestSuiteConfig) testDelete(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.DeleteDeployment(fx.ctx, &DeleteDeploymentRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.DeleteDeployment(fx.ctx, &DeleteDeploymentRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be deleted without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		_, err := fx.service.DeleteDeployment(fx.ctx, &DeleteDeploymentRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
	})

	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		_, err := fx.service.DeleteDeployment(fx.ctx, &DeleteDeploymentRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with NotFound if the resource was already deleted. This also applies to soft-deletion.
	t.Run("already deleted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		deleted, err := fx.service.DeleteDeployment(fx.ctx, &DeleteDeploymentRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
		_ = deleted
		_, err = fx.service.DeleteDeployment(fx.ctx, &DeleteDeploymentRequest{
			Name: created.Name,
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.DeleteDeployment(fx.ctx, &DeleteDeploymentRequest{
			Name: "projects/-/deployments/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with Aborted if the supplied etag doesnt match the current etag value.
	t.Run("etag mismatch", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		_, err := fx.service.DeleteDeployment(fx.ctx, &DeleteDeploymentRequest{
			Name: created.Name,
			Etag: `"99999"`,
		})
		assert.Equal(t, codes.Aborted, status.Code(err), err)
	})

}

func (fx *GSuiteAddOnsDeploymentTestSuiteConfig) nextParent(t *testing.T, pristine bool) string {
	if pristine {
		fx.currParent++
	}
	if fx.currParent >= len(fx.Parents) {
		t.Fatal("need at least", fx.currParent+1, "parents")
	}
	return fx.Parents[fx.currParent]
}

func (fx *GSuiteAddOnsDeploymentTestSuiteConfig) peekNextParent(t *testing.T) string {
	next := fx.currParent + 1
	if next >= len(fx.Parents) {
		t.Fatal("need at least", next+1, "parents")
	}
	return fx.Parents[next]
}

func (fx *GSuiteAddOnsDeploymentTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}

func (fx *GSuiteAddOnsDeploymentTestSuiteConfig) create(t *testing.T, parent string) *Deployment {
	t.Helper()
	created, err := fx.service.CreateDeployment(fx.ctx, &CreateDeploymentRequest{
		Parent:     parent,
		Deployment: fx.Create(parent),
	})
	assert.NilError(t, err)
	return created
}

type GSuiteAddOnsInstallStatusTestSuiteConfig struct {
	ctx        context.Context
	service    GSuiteAddOnsServer
	currParent int

	Server func() GSuiteAddOnsServer
	// Context should return a new context that can be used for each test.
	Context func() context.Context
	// The parents to use when creating resources.
	// At least one parent needs to be set. Depending on methods available on the resource,
	// more may be required. If insufficient number of parents are
	// provided the test will fail.
	Parents []string
	// CreateResource should create a InstallStatus and return it.
	// If the field is not set, some tests will be skipped.
	//
	// This method is generated because service does not expose a Create
	// method (or it does not comply with AIP).
	CreateResource func(ctx context.Context, parent string) (*InstallStatus, error)
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *GSuiteAddOnsInstallStatusTestSuiteConfig) test(t *testing.T) {
	t.Run("Get", fx.testGet)
}

func (fx *GSuiteAddOnsInstallStatusTestSuiteConfig) testGet(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetInstallStatus(fx.ctx, &GetInstallStatusRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetInstallStatus(fx.ctx, &GetInstallStatusRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be returned without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		msg, err := fx.service.GetInstallStatus(fx.ctx, &GetInstallStatusRequest{
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
		_, err := fx.service.GetInstallStatus(fx.ctx, &GetInstallStatusRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetInstallStatus(fx.ctx, &GetInstallStatusRequest{
			Name: "projects/-/deployments/-/installStatus",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *GSuiteAddOnsInstallStatusTestSuiteConfig) nextParent(t *testing.T, pristine bool) string {
	if pristine {
		fx.currParent++
	}
	if fx.currParent >= len(fx.Parents) {
		t.Fatal("need at least", fx.currParent+1, "parents")
	}
	return fx.Parents[fx.currParent]
}

func (fx *GSuiteAddOnsInstallStatusTestSuiteConfig) peekNextParent(t *testing.T) string {
	next := fx.currParent + 1
	if next >= len(fx.Parents) {
		t.Fatal("need at least", next+1, "parents")
	}
	return fx.Parents[next]
}

func (fx *GSuiteAddOnsInstallStatusTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}

func (fx *GSuiteAddOnsInstallStatusTestSuiteConfig) create(t *testing.T, parent string) *InstallStatus {
	t.Helper()
	if fx.CreateResource == nil {
		t.Skip("Test skipped because CreateResource not specified on GSuiteAddOnsInstallStatusTestSuiteConfig")
	}
	created, err := fx.CreateResource(fx.ctx, parent)
	assert.NilError(t, err)
	return created
}
