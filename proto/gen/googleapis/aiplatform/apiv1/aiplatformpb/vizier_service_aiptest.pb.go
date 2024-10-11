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
	time "time"
)

// VizierServiceTestSuiteConfigProvider is the interface to implement to decide which resources
// that should be tested and how it's configured.
type VizierServiceTestSuiteConfigProvider interface {
	// VizierServiceStudy should return a config, or nil, which means that the tests will be skipped.
	VizierServiceStudy(t *testing.T) *VizierServiceStudyTestSuiteConfig
	// VizierServiceTrial should return a config, or nil, which means that the tests will be skipped.
	VizierServiceTrial(t *testing.T) *VizierServiceTrialTestSuiteConfig
}

// testVizierService is the main entrypoint for starting the AIP tests.
func testVizierService(t *testing.T, s VizierServiceTestSuiteConfigProvider) {
	testVizierServiceStudy(t, s)
	testVizierServiceTrial(t, s)
}

func testVizierServiceStudy(t *testing.T, s VizierServiceTestSuiteConfigProvider) {
	t.Run("Study", func(t *testing.T) {
		config := s.VizierServiceStudy(t)
		if config == nil {
			t.Skip("Method VizierServiceStudy not implemented")
		}
		if config.Service == nil {
			t.Skip("Method VizierServiceStudy.Service() not implemented")
		}
		if config.Context == nil {
			config.Context = func() context.Context { return context.Background() }
		}
		config.test(t)
	})
}

func testVizierServiceTrial(t *testing.T, s VizierServiceTestSuiteConfigProvider) {
	t.Run("Trial", func(t *testing.T) {
		config := s.VizierServiceTrial(t)
		if config == nil {
			t.Skip("Method VizierServiceTrial not implemented")
		}
		if config.Service == nil {
			t.Skip("Method VizierServiceTrial.Service() not implemented")
		}
		if config.Context == nil {
			config.Context = func() context.Context { return context.Background() }
		}
		config.test(t)
	})
}

type VizierServiceTestSuite struct {
	T *testing.T
	// Server to test.
	Server VizierServiceServer
}

func (fx VizierServiceTestSuite) TestStudy(ctx context.Context, options VizierServiceStudyTestSuiteConfig) {
	fx.T.Run("Study", func(t *testing.T) {
		options.Context = func() context.Context { return ctx }
		options.Service = func() VizierServiceServer { return fx.Server }
		options.test(t)
	})
}

func (fx VizierServiceTestSuite) TestTrial(ctx context.Context, options VizierServiceTrialTestSuiteConfig) {
	fx.T.Run("Trial", func(t *testing.T) {
		options.Context = func() context.Context { return ctx }
		options.Service = func() VizierServiceServer { return fx.Server }
		options.test(t)
	})
}

type VizierServiceStudyTestSuiteConfig struct {
	currParent int

	// Service should return the service that should be tested.
	// The service will be used for several tests.
	Service func() VizierServiceServer
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
	Create func(parent string) *Study
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *VizierServiceStudyTestSuiteConfig) test(t *testing.T) {
	t.Run("Create", fx.testCreate)
	t.Run("Get", fx.testGet)
	t.Run("List", fx.testList)
	t.Run("Delete", fx.testDelete)
}

func (fx *VizierServiceStudyTestSuiteConfig) testCreate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no parent is provided.
	t.Run("missing parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().CreateStudy(fx.Context(), &CreateStudyRequest{
			Parent: "",
			Study:  fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().CreateStudy(fx.Context(), &CreateStudyRequest{
			Parent: "invalid resource name",
			Study:  fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Field create_time should be populated when the resource is created.
	t.Run("create time", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		beforeCreate := time.Now()
		msg, err := fx.Service().CreateStudy(fx.Context(), &CreateStudyRequest{
			Parent: parent,
			Study:  fx.Create(parent),
		})
		assert.NilError(t, err)
		assert.Check(t, msg.CreateTime != nil)
		assert.Check(t, !msg.CreateTime.AsTime().IsZero())
		assert.Check(t, msg.CreateTime.AsTime().After(beforeCreate), "msg.CreateTime (%v) is not after beforeCreate (%v)", msg.CreateTime.AsTime(), beforeCreate)
	})

	// The created resource should be persisted and reachable with Get.
	t.Run("persisted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg, err := fx.Service().CreateStudy(fx.Context(), &CreateStudyRequest{
			Parent: parent,
			Study:  fx.Create(parent),
		})
		assert.NilError(t, err)
		persisted, err := fx.Service().GetStudy(fx.Context(), &GetStudyRequest{
			Name: msg.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, msg, persisted, protocmp.Transform())
	})

	// The method should fail with InvalidArgument if the resource has any
	// required fields and they are not provided.
	t.Run("required fields", func(t *testing.T) {
		fx.maybeSkip(t)
		t.Run(".display_name", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("display_name")
			container.ProtoReflect().Clear(fd)
			_, err := fx.Service().CreateStudy(fx.Context(), &CreateStudyRequest{
				Parent: parent,
				Study:  msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".study_spec", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("study_spec")
			container.ProtoReflect().Clear(fd)
			_, err := fx.Service().CreateStudy(fx.Context(), &CreateStudyRequest{
				Parent: parent,
				Study:  msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".study_spec.metrics", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetStudySpec()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("metrics")
			container.ProtoReflect().Clear(fd)
			_, err := fx.Service().CreateStudy(fx.Context(), &CreateStudyRequest{
				Parent: parent,
				Study:  msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".study_spec.parameters", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetStudySpec()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("parameters")
			container.ProtoReflect().Clear(fd)
			_, err := fx.Service().CreateStudy(fx.Context(), &CreateStudyRequest{
				Parent: parent,
				Study:  msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
	})

}

func (fx *VizierServiceStudyTestSuiteConfig) testGet(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetStudy(fx.Context(), &GetStudyRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetStudy(fx.Context(), &GetStudyRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be returned without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		msg, err := fx.Service().GetStudy(fx.Context(), &GetStudyRequest{
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
		_, err := fx.Service().GetStudy(fx.Context(), &GetStudyRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetStudy(fx.Context(), &GetStudyRequest{
			Name: "projects/-/locations/-/studies/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *VizierServiceStudyTestSuiteConfig) testList(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().ListStudies(fx.Context(), &ListStudiesRequest{
			Parent: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page token is not valid.
	t.Run("invalid page token", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.Service().ListStudies(fx.Context(), &ListStudiesRequest{
			Parent:    parent,
			PageToken: "invalid page token",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page size is negative.
	t.Run("negative page size", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.Service().ListStudies(fx.Context(), &ListStudiesRequest{
			Parent:   parent,
			PageSize: -10,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	const resourcesCount = 15
	parent := fx.nextParent(t, true)
	parentMsgs := make([]*Study, resourcesCount)
	for i := 0; i < resourcesCount; i++ {
		parentMsgs[i] = fx.create(t, parent)
	}

	// If parent is provided the method must only return resources
	// under that parent.
	t.Run("isolation", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.Service().ListStudies(fx.Context(), &ListStudiesRequest{
			Parent:   parent,
			PageSize: 999,
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			parentMsgs,
			response.Studies,
			cmpopts.SortSlices(func(a, b *Study) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

	// If there are no more resources, next_page_token should not be set.
	t.Run("last page", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.Service().ListStudies(fx.Context(), &ListStudiesRequest{
			Parent:   parent,
			PageSize: resourcesCount,
		})
		assert.NilError(t, err)
		assert.Equal(t, "", response.NextPageToken)
	})

	// If there are more resources, next_page_token should be set.
	t.Run("more pages", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.Service().ListStudies(fx.Context(), &ListStudiesRequest{
			Parent:   parent,
			PageSize: resourcesCount - 1,
		})
		assert.NilError(t, err)
		assert.Check(t, response.NextPageToken != "")
	})

	// Listing resource one by one should eventually return all resources.
	t.Run("one by one", func(t *testing.T) {
		fx.maybeSkip(t)
		msgs := make([]*Study, 0, resourcesCount)
		var nextPageToken string
		for {
			response, err := fx.Service().ListStudies(fx.Context(), &ListStudiesRequest{
				Parent:    parent,
				PageSize:  1,
				PageToken: nextPageToken,
			})
			assert.NilError(t, err)
			assert.Equal(t, 1, len(response.Studies))
			msgs = append(msgs, response.Studies...)
			nextPageToken = response.NextPageToken
			if nextPageToken == "" {
				break
			}
		}
		assert.DeepEqual(
			t,
			parentMsgs,
			msgs,
			cmpopts.SortSlices(func(a, b *Study) bool {
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
			_, err := fx.Service().DeleteStudy(fx.Context(), &DeleteStudyRequest{
				Name: parentMsgs[i].Name,
			})
			assert.NilError(t, err)
		}
		response, err := fx.Service().ListStudies(fx.Context(), &ListStudiesRequest{
			Parent:   parent,
			PageSize: 9999,
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			parentMsgs[deleteCount:],
			response.Studies,
			cmpopts.SortSlices(func(a, b *Study) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

}

func (fx *VizierServiceStudyTestSuiteConfig) testDelete(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().DeleteStudy(fx.Context(), &DeleteStudyRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().DeleteStudy(fx.Context(), &DeleteStudyRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be deleted without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		_, err := fx.Service().DeleteStudy(fx.Context(), &DeleteStudyRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
	})

	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		_, err := fx.Service().DeleteStudy(fx.Context(), &DeleteStudyRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with NotFound if the resource was already deleted. This also applies to soft-deletion.
	t.Run("already deleted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		deleted, err := fx.Service().DeleteStudy(fx.Context(), &DeleteStudyRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
		_ = deleted
		_, err = fx.Service().DeleteStudy(fx.Context(), &DeleteStudyRequest{
			Name: created.Name,
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().DeleteStudy(fx.Context(), &DeleteStudyRequest{
			Name: "projects/-/locations/-/studies/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *VizierServiceStudyTestSuiteConfig) nextParent(t *testing.T, pristine bool) string {
	if pristine {
		fx.currParent++
	}
	if fx.currParent >= len(fx.Parents) {
		t.Fatal("need at least", fx.currParent+1, "parents")
	}
	return fx.Parents[fx.currParent]
}

func (fx *VizierServiceStudyTestSuiteConfig) peekNextParent(t *testing.T) string {
	next := fx.currParent + 1
	if next >= len(fx.Parents) {
		t.Fatal("need at least", next+1, "parents")
	}
	return fx.Parents[next]
}

func (fx *VizierServiceStudyTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}

func (fx *VizierServiceStudyTestSuiteConfig) create(t *testing.T, parent string) *Study {
	t.Helper()
	created, err := fx.Service().CreateStudy(fx.Context(), &CreateStudyRequest{
		Parent: parent,
		Study:  fx.Create(parent),
	})
	assert.NilError(t, err)
	return created
}

type VizierServiceTrialTestSuiteConfig struct {
	currParent int

	// Service should return the service that should be tested.
	// The service will be used for several tests.
	Service func() VizierServiceServer
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
	Create func(parent string) *Trial
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *VizierServiceTrialTestSuiteConfig) test(t *testing.T) {
	t.Run("Create", fx.testCreate)
	t.Run("Get", fx.testGet)
	t.Run("List", fx.testList)
	t.Run("Delete", fx.testDelete)
}

func (fx *VizierServiceTrialTestSuiteConfig) testCreate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no parent is provided.
	t.Run("missing parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().CreateTrial(fx.Context(), &CreateTrialRequest{
			Parent: "",
			Trial:  fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().CreateTrial(fx.Context(), &CreateTrialRequest{
			Parent: "invalid resource name",
			Trial:  fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// The created resource should be persisted and reachable with Get.
	t.Run("persisted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg, err := fx.Service().CreateTrial(fx.Context(), &CreateTrialRequest{
			Parent: parent,
			Trial:  fx.Create(parent),
		})
		assert.NilError(t, err)
		persisted, err := fx.Service().GetTrial(fx.Context(), &GetTrialRequest{
			Name: msg.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, msg, persisted, protocmp.Transform())
	})

}

func (fx *VizierServiceTrialTestSuiteConfig) testGet(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetTrial(fx.Context(), &GetTrialRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetTrial(fx.Context(), &GetTrialRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be returned without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		msg, err := fx.Service().GetTrial(fx.Context(), &GetTrialRequest{
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
		_, err := fx.Service().GetTrial(fx.Context(), &GetTrialRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetTrial(fx.Context(), &GetTrialRequest{
			Name: "projects/-/locations/-/studies/-/trials/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *VizierServiceTrialTestSuiteConfig) testList(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().ListTrials(fx.Context(), &ListTrialsRequest{
			Parent: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page token is not valid.
	t.Run("invalid page token", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.Service().ListTrials(fx.Context(), &ListTrialsRequest{
			Parent:    parent,
			PageToken: "invalid page token",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page size is negative.
	t.Run("negative page size", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.Service().ListTrials(fx.Context(), &ListTrialsRequest{
			Parent:   parent,
			PageSize: -10,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	const resourcesCount = 15
	parent := fx.nextParent(t, true)
	parentMsgs := make([]*Trial, resourcesCount)
	for i := 0; i < resourcesCount; i++ {
		parentMsgs[i] = fx.create(t, parent)
	}

	// If parent is provided the method must only return resources
	// under that parent.
	t.Run("isolation", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.Service().ListTrials(fx.Context(), &ListTrialsRequest{
			Parent:   parent,
			PageSize: 999,
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			parentMsgs,
			response.Trials,
			cmpopts.SortSlices(func(a, b *Trial) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

	// If there are no more resources, next_page_token should not be set.
	t.Run("last page", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.Service().ListTrials(fx.Context(), &ListTrialsRequest{
			Parent:   parent,
			PageSize: resourcesCount,
		})
		assert.NilError(t, err)
		assert.Equal(t, "", response.NextPageToken)
	})

	// If there are more resources, next_page_token should be set.
	t.Run("more pages", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.Service().ListTrials(fx.Context(), &ListTrialsRequest{
			Parent:   parent,
			PageSize: resourcesCount - 1,
		})
		assert.NilError(t, err)
		assert.Check(t, response.NextPageToken != "")
	})

	// Listing resource one by one should eventually return all resources.
	t.Run("one by one", func(t *testing.T) {
		fx.maybeSkip(t)
		msgs := make([]*Trial, 0, resourcesCount)
		var nextPageToken string
		for {
			response, err := fx.Service().ListTrials(fx.Context(), &ListTrialsRequest{
				Parent:    parent,
				PageSize:  1,
				PageToken: nextPageToken,
			})
			assert.NilError(t, err)
			assert.Equal(t, 1, len(response.Trials))
			msgs = append(msgs, response.Trials...)
			nextPageToken = response.NextPageToken
			if nextPageToken == "" {
				break
			}
		}
		assert.DeepEqual(
			t,
			parentMsgs,
			msgs,
			cmpopts.SortSlices(func(a, b *Trial) bool {
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
			_, err := fx.Service().DeleteTrial(fx.Context(), &DeleteTrialRequest{
				Name: parentMsgs[i].Name,
			})
			assert.NilError(t, err)
		}
		response, err := fx.Service().ListTrials(fx.Context(), &ListTrialsRequest{
			Parent:   parent,
			PageSize: 9999,
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			parentMsgs[deleteCount:],
			response.Trials,
			cmpopts.SortSlices(func(a, b *Trial) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

}

func (fx *VizierServiceTrialTestSuiteConfig) testDelete(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().DeleteTrial(fx.Context(), &DeleteTrialRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().DeleteTrial(fx.Context(), &DeleteTrialRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be deleted without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		_, err := fx.Service().DeleteTrial(fx.Context(), &DeleteTrialRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
	})

	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		_, err := fx.Service().DeleteTrial(fx.Context(), &DeleteTrialRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with NotFound if the resource was already deleted. This also applies to soft-deletion.
	t.Run("already deleted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		deleted, err := fx.Service().DeleteTrial(fx.Context(), &DeleteTrialRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
		_ = deleted
		_, err = fx.Service().DeleteTrial(fx.Context(), &DeleteTrialRequest{
			Name: created.Name,
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().DeleteTrial(fx.Context(), &DeleteTrialRequest{
			Name: "projects/-/locations/-/studies/-/trials/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *VizierServiceTrialTestSuiteConfig) nextParent(t *testing.T, pristine bool) string {
	if pristine {
		fx.currParent++
	}
	if fx.currParent >= len(fx.Parents) {
		t.Fatal("need at least", fx.currParent+1, "parents")
	}
	return fx.Parents[fx.currParent]
}

func (fx *VizierServiceTrialTestSuiteConfig) peekNextParent(t *testing.T) string {
	next := fx.currParent + 1
	if next >= len(fx.Parents) {
		t.Fatal("need at least", next+1, "parents")
	}
	return fx.Parents[next]
}

func (fx *VizierServiceTrialTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}

func (fx *VizierServiceTrialTestSuiteConfig) create(t *testing.T, parent string) *Trial {
	t.Helper()
	created, err := fx.Service().CreateTrial(fx.Context(), &CreateTrialRequest{
		Parent: parent,
		Trial:  fx.Create(parent),
	})
	assert.NilError(t, err)
	return created
}
