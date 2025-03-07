// Code generated by protoc-gen-go-aip-test. DO NOT EDIT.

package aiplatformpb

import (
	context "context"
	cmpopts "github.com/google/go-cmp/cmp/cmpopts"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	proto "google.golang.org/protobuf/proto"
	protocmp "google.golang.org/protobuf/testing/protocmp"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	assert "gotest.tools/v3/assert"
	strings "strings"
	testing "testing"
)

// FeatureRegistryServiceTestSuiteConfigProvider is the interface to implement to decide which resources
// that should be tested and how it's configured.
type FeatureRegistryServiceTestSuiteConfigProvider interface {
	// FeatureRegistryServiceFeature should return a config, or nil, which means that the tests will be skipped.
	FeatureRegistryServiceFeature(t *testing.T) *FeatureRegistryServiceFeatureTestSuiteConfig
	// FeatureRegistryServiceFeatureGroup should return a config, or nil, which means that the tests will be skipped.
	FeatureRegistryServiceFeatureGroup(t *testing.T) *FeatureRegistryServiceFeatureGroupTestSuiteConfig
}

// testFeatureRegistryService is the main entrypoint for starting the AIP tests.
func testFeatureRegistryService(t *testing.T, s FeatureRegistryServiceTestSuiteConfigProvider) {
	testFeatureRegistryServiceFeature(t, s)
	testFeatureRegistryServiceFeatureGroup(t, s)
}

func testFeatureRegistryServiceFeature(t *testing.T, s FeatureRegistryServiceTestSuiteConfigProvider) {
	t.Run("Feature", func(t *testing.T) {
		config := s.FeatureRegistryServiceFeature(t)
		if config == nil {
			t.Skip("Method FeatureRegistryServiceFeature not implemented")
		}
		if config.Service == nil {
			t.Skip("Method FeatureRegistryServiceFeature.Service() not implemented")
		}
		if config.Context == nil {
			config.Context = func() context.Context { return context.Background() }
		}
		config.test(t)
	})
}

func testFeatureRegistryServiceFeatureGroup(t *testing.T, s FeatureRegistryServiceTestSuiteConfigProvider) {
	t.Run("FeatureGroup", func(t *testing.T) {
		config := s.FeatureRegistryServiceFeatureGroup(t)
		if config == nil {
			t.Skip("Method FeatureRegistryServiceFeatureGroup not implemented")
		}
		if config.Service == nil {
			t.Skip("Method FeatureRegistryServiceFeatureGroup.Service() not implemented")
		}
		if config.Context == nil {
			config.Context = func() context.Context { return context.Background() }
		}
		config.test(t)
	})
}

type FeatureRegistryServiceTestSuite struct {
	T *testing.T
	// Server to test.
	Server FeatureRegistryServiceServer
}

func (fx FeatureRegistryServiceTestSuite) TestFeature(ctx context.Context, options FeatureRegistryServiceFeatureTestSuiteConfig) {
	fx.T.Run("Feature", func(t *testing.T) {
		options.Context = func() context.Context { return ctx }
		options.Service = func() FeatureRegistryServiceServer { return fx.Server }
		options.test(t)
	})
}

func (fx FeatureRegistryServiceTestSuite) TestFeatureGroup(ctx context.Context, options FeatureRegistryServiceFeatureGroupTestSuiteConfig) {
	fx.T.Run("FeatureGroup", func(t *testing.T) {
		options.Context = func() context.Context { return ctx }
		options.Service = func() FeatureRegistryServiceServer { return fx.Server }
		options.test(t)
	})
}

type FeatureRegistryServiceFeatureTestSuiteConfig struct {
	currParent int

	// Service should return the service that should be tested.
	// The service will be used for several tests.
	Service func() FeatureRegistryServiceServer
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
	Create func(parent string) *Feature
	// IDGenerator should return a valid and unique ID to use in the Create call.
	// If non-nil, this function will be called to set the ID on all Create calls.
	// If the ID field is required, tests will fail if this is nil.
	IDGenerator func() string
	// Update should return a resource which is valid to update, i.e.
	// all required fields set.
	Update func(parent string) *Feature
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *FeatureRegistryServiceFeatureTestSuiteConfig) test(t *testing.T) {
	t.Run("Create", fx.testCreate)
	t.Run("Get", fx.testGet)
	t.Run("Update", fx.testUpdate)
	t.Run("List", fx.testList)
	t.Run("Delete", fx.testDelete)
}

func (fx *FeatureRegistryServiceFeatureTestSuiteConfig) testCreate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no parent is provided.
	t.Run("missing parent", func(t *testing.T) {
		fx.maybeSkip(t)
		userSetID := ""
		if fx.IDGenerator != nil {
			userSetID = fx.IDGenerator()
		}
		_, err := fx.Service().CreateFeature(fx.Context(), &CreateFeatureRequest{
			Parent:    "",
			Feature:   fx.Create(fx.nextParent(t, false)),
			FeatureId: userSetID,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		userSetID := ""
		if fx.IDGenerator != nil {
			userSetID = fx.IDGenerator()
		}
		_, err := fx.Service().CreateFeature(fx.Context(), &CreateFeatureRequest{
			Parent:    "invalid resource name",
			Feature:   fx.Create(fx.nextParent(t, false)),
			FeatureId: userSetID,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *FeatureRegistryServiceFeatureTestSuiteConfig) testGet(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetFeature(fx.Context(), &GetFeatureRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetFeature(fx.Context(), &GetFeatureRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be returned without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		msg, err := fx.Service().GetFeature(fx.Context(), &GetFeatureRequest{
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
		_, err := fx.Service().GetFeature(fx.Context(), &GetFeatureRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetFeature(fx.Context(), &GetFeatureRequest{
			Name: "projects/-/locations/-/featurestores/-/entityTypes/-/features/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *FeatureRegistryServiceFeatureTestSuiteConfig) testUpdate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = ""
		_, err := fx.Service().UpdateFeature(fx.Context(), &UpdateFeatureRequest{
			Feature: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = "invalid resource name"
		_, err := fx.Service().UpdateFeature(fx.Context(), &UpdateFeatureRequest{
			Feature: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with Aborted if the supplied etag doesnt match the current etag value.
	t.Run("etag mismatch", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		msg := fx.Update(parent)
		msg.Name = created.Name
		_, err := fx.Service().UpdateFeature(fx.Context(), &UpdateFeatureRequest{
			Feature: msg,
		})
		assert.Equal(t, codes.Aborted, status.Code(err), err)
	})

	// Field etag should have a new value when the resource is successfully updated.
	t.Run("etag updated", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		msg := fx.Update(parent)
		msg.Name = created.Name
		updated, err := fx.Service().UpdateFeature(fx.Context(), &UpdateFeatureRequest{
			Feature: msg,
		})
		assert.NilError(t, err)
		_ = updated
	})

	{
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		// Method should fail with NotFound if the resource does not exist.
		t.Run("not found", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := fx.Update(parent)
			msg.Name = created.Name + "notfound"
			_, err := fx.Service().UpdateFeature(fx.Context(), &UpdateFeatureRequest{
				Feature: msg,
			})
			assert.Equal(t, codes.NotFound, status.Code(err), err)
		})

		// The method should fail with InvalidArgument if the update_mask is invalid.
		t.Run("invalid update mask", func(t *testing.T) {
			fx.maybeSkip(t)
			_, err := fx.Service().UpdateFeature(fx.Context(), &UpdateFeatureRequest{
				Feature: created,
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

func (fx *FeatureRegistryServiceFeatureTestSuiteConfig) testList(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().ListFeatures(fx.Context(), &ListFeaturesRequest{
			Parent: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page token is not valid.
	t.Run("invalid page token", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.Service().ListFeatures(fx.Context(), &ListFeaturesRequest{
			Parent:    parent,
			PageToken: "invalid page token",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page size is negative.
	t.Run("negative page size", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.Service().ListFeatures(fx.Context(), &ListFeaturesRequest{
			Parent:   parent,
			PageSize: -10,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	{
		const resourcesCount = 15
		parent := fx.nextParent(t, true)
		parentMsgs := make([]*Feature, resourcesCount)
		for i := 0; i < resourcesCount; i++ {
			parentMsgs[i] = fx.create(t, parent)
		}

		// If parent is provided the method must only return resources
		// under that parent.
		t.Run("isolation", func(t *testing.T) {
			fx.maybeSkip(t)
			response, err := fx.Service().ListFeatures(fx.Context(), &ListFeaturesRequest{
				Parent:   parent,
				PageSize: 999,
			})
			assert.NilError(t, err)
			assert.DeepEqual(
				t,
				parentMsgs,
				response.Features,
				cmpopts.SortSlices(func(a, b *Feature) bool {
					return a.Name < b.Name
				}),
				protocmp.Transform(),
			)
		})

		// If there are no more resources, next_page_token should not be set.
		t.Run("last page", func(t *testing.T) {
			fx.maybeSkip(t)
			response, err := fx.Service().ListFeatures(fx.Context(), &ListFeaturesRequest{
				Parent:   parent,
				PageSize: resourcesCount,
			})
			assert.NilError(t, err)
			assert.Equal(t, "", response.NextPageToken)
		})

		// If there are more resources, next_page_token should be set.
		t.Run("more pages", func(t *testing.T) {
			fx.maybeSkip(t)
			response, err := fx.Service().ListFeatures(fx.Context(), &ListFeaturesRequest{
				Parent:   parent,
				PageSize: resourcesCount - 1,
			})
			assert.NilError(t, err)
			assert.Check(t, response.NextPageToken != "")
		})

		// Listing resource one by one should eventually return all resources.
		t.Run("one by one", func(t *testing.T) {
			fx.maybeSkip(t)
			msgs := make([]*Feature, 0, resourcesCount)
			var nextPageToken string
			for {
				response, err := fx.Service().ListFeatures(fx.Context(), &ListFeaturesRequest{
					Parent:    parent,
					PageSize:  1,
					PageToken: nextPageToken,
				})
				assert.NilError(t, err)
				assert.Equal(t, 1, len(response.Features))
				msgs = append(msgs, response.Features...)
				nextPageToken = response.NextPageToken
				if nextPageToken == "" {
					break
				}
			}
			assert.DeepEqual(
				t,
				parentMsgs,
				msgs,
				cmpopts.SortSlices(func(a, b *Feature) bool {
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
				_, err := fx.Service().DeleteFeature(fx.Context(), &DeleteFeatureRequest{
					Name: parentMsgs[i].Name,
				})
				assert.NilError(t, err)
			}
			response, err := fx.Service().ListFeatures(fx.Context(), &ListFeaturesRequest{
				Parent:   parent,
				PageSize: 9999,
			})
			assert.NilError(t, err)
			assert.DeepEqual(
				t,
				parentMsgs[deleteCount:],
				response.Features,
				cmpopts.SortSlices(func(a, b *Feature) bool {
					return a.Name < b.Name
				}),
				protocmp.Transform(),
			)
		})

	}
	{
		const resourcesCount = 101
		parent := fx.nextParent(t, true)
		parentMsgs := make([]*Feature, resourcesCount)
		for i := 0; i < resourcesCount; i++ {
			parentMsgs[i] = fx.create(t, parent)
		}

		// Listing resource with page size zero should eventually return all resources.
		t.Run("page size zero", func(t *testing.T) {
			fx.maybeSkip(t)
			msgs := make([]*Feature, 0, resourcesCount)
			var nextPageToken string
			for {
				page, err := fx.Service().ListFeatures(fx.Context(), &ListFeaturesRequest{
					Parent:    parent,
					PageSize:  0,
					PageToken: nextPageToken,
				})
				assert.NilError(t, err)
				msgs = append(msgs, page.Features...)
				nextPageToken = page.NextPageToken
				if nextPageToken == "" {
					break
				}
			}
			assert.DeepEqual(
				t,
				parentMsgs,
				msgs,
				cmpopts.SortSlices(func(a, b *Feature) bool {
					return a.Name < b.Name
				}),
				protocmp.Transform(),
			)
		})

	}
}

func (fx *FeatureRegistryServiceFeatureTestSuiteConfig) testDelete(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().DeleteFeature(fx.Context(), &DeleteFeatureRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().DeleteFeature(fx.Context(), &DeleteFeatureRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be deleted without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		_, err := fx.Service().DeleteFeature(fx.Context(), &DeleteFeatureRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
	})

	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		_, err := fx.Service().DeleteFeature(fx.Context(), &DeleteFeatureRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with NotFound if the resource was already deleted. This also applies to soft-deletion.
	t.Run("already deleted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		deleted, err := fx.Service().DeleteFeature(fx.Context(), &DeleteFeatureRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
		_ = deleted
		_, err = fx.Service().DeleteFeature(fx.Context(), &DeleteFeatureRequest{
			Name: created.Name,
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().DeleteFeature(fx.Context(), &DeleteFeatureRequest{
			Name: "projects/-/locations/-/featurestores/-/entityTypes/-/features/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *FeatureRegistryServiceFeatureTestSuiteConfig) nextParent(t *testing.T, pristine bool) string {
	if pristine {
		fx.currParent++
	}
	if fx.currParent >= len(fx.Parents) {
		t.Fatal("need at least", fx.currParent+1, "parents")
	}
	return fx.Parents[fx.currParent]
}

func (fx *FeatureRegistryServiceFeatureTestSuiteConfig) peekNextParent(t *testing.T) string {
	next := fx.currParent + 1
	if next >= len(fx.Parents) {
		t.Fatal("need at least", next+1, "parents")
	}
	return fx.Parents[next]
}

func (fx *FeatureRegistryServiceFeatureTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}

func (fx *FeatureRegistryServiceFeatureTestSuiteConfig) create(t *testing.T, parent string) *Feature {
	t.Helper()
	t.Skip("Long running create method not supported")
	return nil
}

type FeatureRegistryServiceFeatureGroupTestSuiteConfig struct {
	currParent int

	// Service should return the service that should be tested.
	// The service will be used for several tests.
	Service func() FeatureRegistryServiceServer
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
	Create func(parent string) *FeatureGroup
	// IDGenerator should return a valid and unique ID to use in the Create call.
	// If non-nil, this function will be called to set the ID on all Create calls.
	// If the ID field is required, tests will fail if this is nil.
	IDGenerator func() string
	// Update should return a resource which is valid to update, i.e.
	// all required fields set.
	Update func(parent string) *FeatureGroup
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *FeatureRegistryServiceFeatureGroupTestSuiteConfig) test(t *testing.T) {
	t.Run("Create", fx.testCreate)
	t.Run("Get", fx.testGet)
	t.Run("Update", fx.testUpdate)
	t.Run("List", fx.testList)
	t.Run("Delete", fx.testDelete)
}

func (fx *FeatureRegistryServiceFeatureGroupTestSuiteConfig) testCreate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no parent is provided.
	t.Run("missing parent", func(t *testing.T) {
		fx.maybeSkip(t)
		userSetID := ""
		if fx.IDGenerator != nil {
			userSetID = fx.IDGenerator()
		}
		_, err := fx.Service().CreateFeatureGroup(fx.Context(), &CreateFeatureGroupRequest{
			Parent:         "",
			FeatureGroup:   fx.Create(fx.nextParent(t, false)),
			FeatureGroupId: userSetID,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		userSetID := ""
		if fx.IDGenerator != nil {
			userSetID = fx.IDGenerator()
		}
		_, err := fx.Service().CreateFeatureGroup(fx.Context(), &CreateFeatureGroupRequest{
			Parent:         "invalid resource name",
			FeatureGroup:   fx.Create(fx.nextParent(t, false)),
			FeatureGroupId: userSetID,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// The method should fail with InvalidArgument if the resource has any
	// required fields and they are not provided.
	t.Run("required fields", func(t *testing.T) {
		fx.maybeSkip(t)
		t.Run(".big_query.big_query_source", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetBigQuery()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("big_query_source")
			container.ProtoReflect().Clear(fd)
			userSetID := ""
			if fx.IDGenerator != nil {
				userSetID = fx.IDGenerator()
			}
			_, err := fx.Service().CreateFeatureGroup(fx.Context(), &CreateFeatureGroupRequest{
				Parent:         parent,
				FeatureGroup:   msg,
				FeatureGroupId: userSetID,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".big_query.big_query_source.input_uri", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetBigQuery().GetBigQuerySource()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("input_uri")
			container.ProtoReflect().Clear(fd)
			userSetID := ""
			if fx.IDGenerator != nil {
				userSetID = fx.IDGenerator()
			}
			_, err := fx.Service().CreateFeatureGroup(fx.Context(), &CreateFeatureGroupRequest{
				Parent:         parent,
				FeatureGroup:   msg,
				FeatureGroupId: userSetID,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
	})

}

func (fx *FeatureRegistryServiceFeatureGroupTestSuiteConfig) testGet(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetFeatureGroup(fx.Context(), &GetFeatureGroupRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetFeatureGroup(fx.Context(), &GetFeatureGroupRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be returned without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		msg, err := fx.Service().GetFeatureGroup(fx.Context(), &GetFeatureGroupRequest{
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
		_, err := fx.Service().GetFeatureGroup(fx.Context(), &GetFeatureGroupRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetFeatureGroup(fx.Context(), &GetFeatureGroupRequest{
			Name: "projects/-/locations/-/featureGroups/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *FeatureRegistryServiceFeatureGroupTestSuiteConfig) testUpdate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = ""
		_, err := fx.Service().UpdateFeatureGroup(fx.Context(), &UpdateFeatureGroupRequest{
			FeatureGroup: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = "invalid resource name"
		_, err := fx.Service().UpdateFeatureGroup(fx.Context(), &UpdateFeatureGroupRequest{
			FeatureGroup: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with Aborted if the supplied etag doesnt match the current etag value.
	t.Run("etag mismatch", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		msg := fx.Update(parent)
		msg.Name = created.Name
		_, err := fx.Service().UpdateFeatureGroup(fx.Context(), &UpdateFeatureGroupRequest{
			FeatureGroup: msg,
		})
		assert.Equal(t, codes.Aborted, status.Code(err), err)
	})

	// Field etag should have a new value when the resource is successfully updated.
	t.Run("etag updated", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		msg := fx.Update(parent)
		msg.Name = created.Name
		updated, err := fx.Service().UpdateFeatureGroup(fx.Context(), &UpdateFeatureGroupRequest{
			FeatureGroup: msg,
		})
		assert.NilError(t, err)
		_ = updated
	})

	{
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		// Method should fail with NotFound if the resource does not exist.
		t.Run("not found", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := fx.Update(parent)
			msg.Name = created.Name + "notfound"
			_, err := fx.Service().UpdateFeatureGroup(fx.Context(), &UpdateFeatureGroupRequest{
				FeatureGroup: msg,
			})
			assert.Equal(t, codes.NotFound, status.Code(err), err)
		})

		// The method should fail with InvalidArgument if the update_mask is invalid.
		t.Run("invalid update mask", func(t *testing.T) {
			fx.maybeSkip(t)
			_, err := fx.Service().UpdateFeatureGroup(fx.Context(), &UpdateFeatureGroupRequest{
				FeatureGroup: created,
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"invalid_field_xyz",
					},
				},
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})

		// Method should fail with InvalidArgument if any required field is missing
		// when called with '*' update_mask.
		t.Run("required fields", func(t *testing.T) {
			fx.maybeSkip(t)
			t.Run(".big_query.big_query_source.input_uri", func(t *testing.T) {
				fx.maybeSkip(t)
				msg := proto.Clone(created).(*FeatureGroup)
				container := msg.GetBigQuery().GetBigQuerySource()
				if container == nil {
					t.Skip("not reachable")
				}
				fd := container.ProtoReflect().Descriptor().Fields().ByName("input_uri")
				container.ProtoReflect().Clear(fd)
				_, err := fx.Service().UpdateFeatureGroup(fx.Context(), &UpdateFeatureGroupRequest{
					FeatureGroup: msg,
					UpdateMask: &fieldmaskpb.FieldMask{
						Paths: []string{
							"*",
						},
					},
				})
				assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
			})
		})

	}
}

func (fx *FeatureRegistryServiceFeatureGroupTestSuiteConfig) testList(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().ListFeatureGroups(fx.Context(), &ListFeatureGroupsRequest{
			Parent: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page token is not valid.
	t.Run("invalid page token", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.Service().ListFeatureGroups(fx.Context(), &ListFeatureGroupsRequest{
			Parent:    parent,
			PageToken: "invalid page token",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page size is negative.
	t.Run("negative page size", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.Service().ListFeatureGroups(fx.Context(), &ListFeatureGroupsRequest{
			Parent:   parent,
			PageSize: -10,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	{
		const resourcesCount = 15
		parent := fx.nextParent(t, true)
		parentMsgs := make([]*FeatureGroup, resourcesCount)
		for i := 0; i < resourcesCount; i++ {
			parentMsgs[i] = fx.create(t, parent)
		}

		// If parent is provided the method must only return resources
		// under that parent.
		t.Run("isolation", func(t *testing.T) {
			fx.maybeSkip(t)
			response, err := fx.Service().ListFeatureGroups(fx.Context(), &ListFeatureGroupsRequest{
				Parent:   parent,
				PageSize: 999,
			})
			assert.NilError(t, err)
			assert.DeepEqual(
				t,
				parentMsgs,
				response.FeatureGroups,
				cmpopts.SortSlices(func(a, b *FeatureGroup) bool {
					return a.Name < b.Name
				}),
				protocmp.Transform(),
			)
		})

		// If there are no more resources, next_page_token should not be set.
		t.Run("last page", func(t *testing.T) {
			fx.maybeSkip(t)
			response, err := fx.Service().ListFeatureGroups(fx.Context(), &ListFeatureGroupsRequest{
				Parent:   parent,
				PageSize: resourcesCount,
			})
			assert.NilError(t, err)
			assert.Equal(t, "", response.NextPageToken)
		})

		// If there are more resources, next_page_token should be set.
		t.Run("more pages", func(t *testing.T) {
			fx.maybeSkip(t)
			response, err := fx.Service().ListFeatureGroups(fx.Context(), &ListFeatureGroupsRequest{
				Parent:   parent,
				PageSize: resourcesCount - 1,
			})
			assert.NilError(t, err)
			assert.Check(t, response.NextPageToken != "")
		})

		// Listing resource one by one should eventually return all resources.
		t.Run("one by one", func(t *testing.T) {
			fx.maybeSkip(t)
			msgs := make([]*FeatureGroup, 0, resourcesCount)
			var nextPageToken string
			for {
				response, err := fx.Service().ListFeatureGroups(fx.Context(), &ListFeatureGroupsRequest{
					Parent:    parent,
					PageSize:  1,
					PageToken: nextPageToken,
				})
				assert.NilError(t, err)
				assert.Equal(t, 1, len(response.FeatureGroups))
				msgs = append(msgs, response.FeatureGroups...)
				nextPageToken = response.NextPageToken
				if nextPageToken == "" {
					break
				}
			}
			assert.DeepEqual(
				t,
				parentMsgs,
				msgs,
				cmpopts.SortSlices(func(a, b *FeatureGroup) bool {
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
				_, err := fx.Service().DeleteFeatureGroup(fx.Context(), &DeleteFeatureGroupRequest{
					Name: parentMsgs[i].Name,
				})
				assert.NilError(t, err)
			}
			response, err := fx.Service().ListFeatureGroups(fx.Context(), &ListFeatureGroupsRequest{
				Parent:   parent,
				PageSize: 9999,
			})
			assert.NilError(t, err)
			assert.DeepEqual(
				t,
				parentMsgs[deleteCount:],
				response.FeatureGroups,
				cmpopts.SortSlices(func(a, b *FeatureGroup) bool {
					return a.Name < b.Name
				}),
				protocmp.Transform(),
			)
		})

	}
	{
		const resourcesCount = 101
		parent := fx.nextParent(t, true)
		parentMsgs := make([]*FeatureGroup, resourcesCount)
		for i := 0; i < resourcesCount; i++ {
			parentMsgs[i] = fx.create(t, parent)
		}

		// Listing resource with page size zero should eventually return all resources.
		t.Run("page size zero", func(t *testing.T) {
			fx.maybeSkip(t)
			msgs := make([]*FeatureGroup, 0, resourcesCount)
			var nextPageToken string
			for {
				page, err := fx.Service().ListFeatureGroups(fx.Context(), &ListFeatureGroupsRequest{
					Parent:    parent,
					PageSize:  0,
					PageToken: nextPageToken,
				})
				assert.NilError(t, err)
				msgs = append(msgs, page.FeatureGroups...)
				nextPageToken = page.NextPageToken
				if nextPageToken == "" {
					break
				}
			}
			assert.DeepEqual(
				t,
				parentMsgs,
				msgs,
				cmpopts.SortSlices(func(a, b *FeatureGroup) bool {
					return a.Name < b.Name
				}),
				protocmp.Transform(),
			)
		})

	}
}

func (fx *FeatureRegistryServiceFeatureGroupTestSuiteConfig) testDelete(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().DeleteFeatureGroup(fx.Context(), &DeleteFeatureGroupRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().DeleteFeatureGroup(fx.Context(), &DeleteFeatureGroupRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be deleted without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		_, err := fx.Service().DeleteFeatureGroup(fx.Context(), &DeleteFeatureGroupRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
	})

	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		_, err := fx.Service().DeleteFeatureGroup(fx.Context(), &DeleteFeatureGroupRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with NotFound if the resource was already deleted. This also applies to soft-deletion.
	t.Run("already deleted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		deleted, err := fx.Service().DeleteFeatureGroup(fx.Context(), &DeleteFeatureGroupRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
		_ = deleted
		_, err = fx.Service().DeleteFeatureGroup(fx.Context(), &DeleteFeatureGroupRequest{
			Name: created.Name,
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().DeleteFeatureGroup(fx.Context(), &DeleteFeatureGroupRequest{
			Name: "projects/-/locations/-/featureGroups/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *FeatureRegistryServiceFeatureGroupTestSuiteConfig) nextParent(t *testing.T, pristine bool) string {
	if pristine {
		fx.currParent++
	}
	if fx.currParent >= len(fx.Parents) {
		t.Fatal("need at least", fx.currParent+1, "parents")
	}
	return fx.Parents[fx.currParent]
}

func (fx *FeatureRegistryServiceFeatureGroupTestSuiteConfig) peekNextParent(t *testing.T) string {
	next := fx.currParent + 1
	if next >= len(fx.Parents) {
		t.Fatal("need at least", next+1, "parents")
	}
	return fx.Parents[next]
}

func (fx *FeatureRegistryServiceFeatureGroupTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}

func (fx *FeatureRegistryServiceFeatureGroupTestSuiteConfig) create(t *testing.T, parent string) *FeatureGroup {
	t.Helper()
	t.Skip("Long running create method not supported")
	return nil
}
