// Code generated by protoc-gen-go-aip-test. DO NOT EDIT.

package aiplatform

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

type FeaturestoreServiceTestSuite struct {
	T *testing.T
	// Server to test.
	Server FeaturestoreServiceServer
}

func (fx FeaturestoreServiceTestSuite) TestEntityType(ctx context.Context, options EntityTypeTestSuiteConfig) {
	fx.T.Run("EntityType", func(t *testing.T) {
		options.ctx = ctx
		options.service = fx.Server
		options.test(t)
	})
}

func (fx FeaturestoreServiceTestSuite) TestFeature(ctx context.Context, options FeatureTestSuiteConfig) {
	fx.T.Run("Feature", func(t *testing.T) {
		options.ctx = ctx
		options.service = fx.Server
		options.test(t)
	})
}

func (fx FeaturestoreServiceTestSuite) TestFeaturestore(ctx context.Context, options FeaturestoreTestSuiteConfig) {
	fx.T.Run("Featurestore", func(t *testing.T) {
		options.ctx = ctx
		options.service = fx.Server
		options.test(t)
	})
}

type EntityTypeTestSuiteConfig struct {
	ctx        context.Context
	service    FeaturestoreServiceServer
	currParent int

	// The parents to use when creating resources.
	// At least one parent needs to be set. Depending on methods available on the resource,
	// more may be required. If insufficient number of parents are
	// provided the test will fail.
	Parents []string
	// Create should return a resource which is valid to create, i.e.
	// all required fields set.
	Create func(parent string) *EntityType
	// Update should return a resource which is valid to update, i.e.
	// all required fields set.
	Update func(parent string) *EntityType
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *EntityTypeTestSuiteConfig) test(t *testing.T) {
	t.Run("Create", fx.testCreate)
	t.Run("Get", fx.testGet)
	t.Run("Update", fx.testUpdate)
	t.Run("List", fx.testList)
}

func (fx *EntityTypeTestSuiteConfig) testCreate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no parent is provided.
	t.Run("missing parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreateEntityType(fx.ctx, &CreateEntityTypeRequest{
			Parent:     "",
			EntityType: fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreateEntityType(fx.ctx, &CreateEntityTypeRequest{
			Parent:     "invalid resource name",
			EntityType: fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *EntityTypeTestSuiteConfig) testGet(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetEntityType(fx.ctx, &GetEntityTypeRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetEntityType(fx.ctx, &GetEntityTypeRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be returned without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		msg, err := fx.service.GetEntityType(fx.ctx, &GetEntityTypeRequest{
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
		_, err := fx.service.GetEntityType(fx.ctx, &GetEntityTypeRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetEntityType(fx.ctx, &GetEntityTypeRequest{
			Name: "projects/-/locations/-/featurestores/-/entityTypes/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *EntityTypeTestSuiteConfig) testUpdate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = ""
		_, err := fx.service.UpdateEntityType(fx.ctx, &UpdateEntityTypeRequest{
			EntityType: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = "invalid resource name"
		_, err := fx.service.UpdateEntityType(fx.ctx, &UpdateEntityTypeRequest{
			EntityType: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// The updated resource should be persisted and reachable with Get.
	t.Run("persisted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		updated, err := fx.service.UpdateEntityType(fx.ctx, &UpdateEntityTypeRequest{
			EntityType: created,
		})
		assert.NilError(t, err)
		persisted, err := fx.service.GetEntityType(fx.ctx, &GetEntityTypeRequest{
			Name: updated.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, updated, persisted, protocmp.Transform())
	})

	parent := fx.nextParent(t, false)
	created := fx.create(t, parent)
	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		msg := fx.Update(parent)
		msg.Name = created.Name + "notfound"
		_, err := fx.service.UpdateEntityType(fx.ctx, &UpdateEntityTypeRequest{
			EntityType: msg,
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// The method should fail with InvalidArgument if the update_mask is invalid.
	t.Run("invalid update mask", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.UpdateEntityType(fx.ctx, &UpdateEntityTypeRequest{
			EntityType: created,
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"invalid_field_xyz",
				},
			},
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *EntityTypeTestSuiteConfig) testList(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.ListEntityTypes(fx.ctx, &ListEntityTypesRequest{
			Parent: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page token is not valid.
	t.Run("invalid page token", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.service.ListEntityTypes(fx.ctx, &ListEntityTypesRequest{
			Parent:    parent,
			PageToken: "invalid page token",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page size is negative.
	t.Run("negative page size", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.service.ListEntityTypes(fx.ctx, &ListEntityTypesRequest{
			Parent:   parent,
			PageSize: -10,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	const resourcesCount = 15
	parent := fx.nextParent(t, true)
	parentMsgs := make([]*EntityType, resourcesCount)
	for i := 0; i < resourcesCount; i++ {
		parentMsgs[i] = fx.create(t, parent)
	}

	// If parent is provided the method must only return resources
	// under that parent.
	t.Run("isolation", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.ListEntityTypes(fx.ctx, &ListEntityTypesRequest{
			Parent:   parent,
			PageSize: 999,
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			parentMsgs,
			response.EntityTypes,
			cmpopts.SortSlices(func(a, b *EntityType) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

	// If there are no more resources, next_page_token should not be set.
	t.Run("last page", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.ListEntityTypes(fx.ctx, &ListEntityTypesRequest{
			Parent:   parent,
			PageSize: resourcesCount,
		})
		assert.NilError(t, err)
		assert.Equal(t, "", response.NextPageToken)
	})

	// If there are more resources, next_page_token should be set.
	t.Run("more pages", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.ListEntityTypes(fx.ctx, &ListEntityTypesRequest{
			Parent:   parent,
			PageSize: resourcesCount - 1,
		})
		assert.NilError(t, err)
		assert.Check(t, response.NextPageToken != "")
	})

	// Listing resource one by one should eventually return all resources.
	t.Run("one by one", func(t *testing.T) {
		fx.maybeSkip(t)
		msgs := make([]*EntityType, 0, resourcesCount)
		var nextPageToken string
		for {
			response, err := fx.service.ListEntityTypes(fx.ctx, &ListEntityTypesRequest{
				Parent:    parent,
				PageSize:  1,
				PageToken: nextPageToken,
			})
			assert.NilError(t, err)
			assert.Equal(t, 1, len(response.EntityTypes))
			msgs = append(msgs, response.EntityTypes...)
			nextPageToken = response.NextPageToken
			if nextPageToken == "" {
				break
			}
		}
		assert.DeepEqual(
			t,
			parentMsgs,
			msgs,
			cmpopts.SortSlices(func(a, b *EntityType) bool {
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
			_, err := fx.service.DeleteEntityType(fx.ctx, &DeleteEntityTypeRequest{
				Name: parentMsgs[i].Name,
			})
			assert.NilError(t, err)
		}
		response, err := fx.service.ListEntityTypes(fx.ctx, &ListEntityTypesRequest{
			Parent:   parent,
			PageSize: 9999,
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			parentMsgs[deleteCount:],
			response.EntityTypes,
			cmpopts.SortSlices(func(a, b *EntityType) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

}

func (fx *EntityTypeTestSuiteConfig) nextParent(t *testing.T, pristine bool) string {
	if pristine {
		fx.currParent++
	}
	if fx.currParent >= len(fx.Parents) {
		t.Fatal("need at least", fx.currParent+1, "parents")
	}
	return fx.Parents[fx.currParent]
}

func (fx *EntityTypeTestSuiteConfig) peekNextParent(t *testing.T) string {
	next := fx.currParent + 1
	if next >= len(fx.Parents) {
		t.Fatal("need at least", next+1, "parents")
	}
	return fx.Parents[next]
}

func (fx *EntityTypeTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}

func (fx *EntityTypeTestSuiteConfig) create(t *testing.T, parent string) *EntityType {
	t.Helper()
	t.Skip("Long running create method not supported")
	return nil
}

type FeatureTestSuiteConfig struct {
	ctx        context.Context
	service    FeaturestoreServiceServer
	currParent int

	// The parents to use when creating resources.
	// At least one parent needs to be set. Depending on methods available on the resource,
	// more may be required. If insufficient number of parents are
	// provided the test will fail.
	Parents []string
	// Create should return a resource which is valid to create, i.e.
	// all required fields set.
	Create func(parent string) *Feature
	// Update should return a resource which is valid to update, i.e.
	// all required fields set.
	Update func(parent string) *Feature
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *FeatureTestSuiteConfig) test(t *testing.T) {
	t.Run("Create", fx.testCreate)
	t.Run("Get", fx.testGet)
	t.Run("Update", fx.testUpdate)
	t.Run("List", fx.testList)
}

func (fx *FeatureTestSuiteConfig) testCreate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no parent is provided.
	t.Run("missing parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreateFeature(fx.ctx, &CreateFeatureRequest{
			Parent:  "",
			Feature: fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreateFeature(fx.ctx, &CreateFeatureRequest{
			Parent:  "invalid resource name",
			Feature: fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// The method should fail with InvalidArgument if the resource has any
	// required fields and they are not provided.
	t.Run("required fields", func(t *testing.T) {
		fx.maybeSkip(t)
		t.Run(".value_type", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("value_type")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateFeature(fx.ctx, &CreateFeatureRequest{
				Parent:  parent,
				Feature: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
	})

}

func (fx *FeatureTestSuiteConfig) testGet(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetFeature(fx.ctx, &GetFeatureRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetFeature(fx.ctx, &GetFeatureRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be returned without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		msg, err := fx.service.GetFeature(fx.ctx, &GetFeatureRequest{
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
		_, err := fx.service.GetFeature(fx.ctx, &GetFeatureRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetFeature(fx.ctx, &GetFeatureRequest{
			Name: "projects/-/locations/-/featurestores/-/entityTypes/-/features/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *FeatureTestSuiteConfig) testUpdate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = ""
		_, err := fx.service.UpdateFeature(fx.ctx, &UpdateFeatureRequest{
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
		_, err := fx.service.UpdateFeature(fx.ctx, &UpdateFeatureRequest{
			Feature: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// The updated resource should be persisted and reachable with Get.
	t.Run("persisted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		updated, err := fx.service.UpdateFeature(fx.ctx, &UpdateFeatureRequest{
			Feature: created,
		})
		assert.NilError(t, err)
		persisted, err := fx.service.GetFeature(fx.ctx, &GetFeatureRequest{
			Name: updated.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, updated, persisted, protocmp.Transform())
	})

	// The field create_time should be preserved when a '*'-update mask is used.
	t.Run("preserve create_time", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		originalCreateTime := created.CreateTime
		updated, err := fx.service.UpdateFeature(fx.ctx, &UpdateFeatureRequest{
			Feature: created,
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"*",
				},
			},
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, originalCreateTime, updated.CreateTime, protocmp.Transform())
	})

	parent := fx.nextParent(t, false)
	created := fx.create(t, parent)
	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		msg := fx.Update(parent)
		msg.Name = created.Name + "notfound"
		_, err := fx.service.UpdateFeature(fx.ctx, &UpdateFeatureRequest{
			Feature: msg,
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// The method should fail with InvalidArgument if the update_mask is invalid.
	t.Run("invalid update mask", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.UpdateFeature(fx.ctx, &UpdateFeatureRequest{
			Feature: created,
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
		t.Run(".value_type", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created).(*Feature)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("value_type")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.UpdateFeature(fx.ctx, &UpdateFeatureRequest{
				Feature: msg,
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

func (fx *FeatureTestSuiteConfig) testList(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.ListFeatures(fx.ctx, &ListFeaturesRequest{
			Parent: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page token is not valid.
	t.Run("invalid page token", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.service.ListFeatures(fx.ctx, &ListFeaturesRequest{
			Parent:    parent,
			PageToken: "invalid page token",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page size is negative.
	t.Run("negative page size", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.service.ListFeatures(fx.ctx, &ListFeaturesRequest{
			Parent:   parent,
			PageSize: -10,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

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
		response, err := fx.service.ListFeatures(fx.ctx, &ListFeaturesRequest{
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
		response, err := fx.service.ListFeatures(fx.ctx, &ListFeaturesRequest{
			Parent:   parent,
			PageSize: resourcesCount,
		})
		assert.NilError(t, err)
		assert.Equal(t, "", response.NextPageToken)
	})

	// If there are more resources, next_page_token should be set.
	t.Run("more pages", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.ListFeatures(fx.ctx, &ListFeaturesRequest{
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
			response, err := fx.service.ListFeatures(fx.ctx, &ListFeaturesRequest{
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
			_, err := fx.service.DeleteFeature(fx.ctx, &DeleteFeatureRequest{
				Name: parentMsgs[i].Name,
			})
			assert.NilError(t, err)
		}
		response, err := fx.service.ListFeatures(fx.ctx, &ListFeaturesRequest{
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

func (fx *FeatureTestSuiteConfig) nextParent(t *testing.T, pristine bool) string {
	if pristine {
		fx.currParent++
	}
	if fx.currParent >= len(fx.Parents) {
		t.Fatal("need at least", fx.currParent+1, "parents")
	}
	return fx.Parents[fx.currParent]
}

func (fx *FeatureTestSuiteConfig) peekNextParent(t *testing.T) string {
	next := fx.currParent + 1
	if next >= len(fx.Parents) {
		t.Fatal("need at least", next+1, "parents")
	}
	return fx.Parents[next]
}

func (fx *FeatureTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}

func (fx *FeatureTestSuiteConfig) create(t *testing.T, parent string) *Feature {
	t.Helper()
	t.Skip("Long running create method not supported")
	return nil
}

type FeaturestoreTestSuiteConfig struct {
	ctx        context.Context
	service    FeaturestoreServiceServer
	currParent int

	// The parents to use when creating resources.
	// At least one parent needs to be set. Depending on methods available on the resource,
	// more may be required. If insufficient number of parents are
	// provided the test will fail.
	Parents []string
	// Create should return a resource which is valid to create, i.e.
	// all required fields set.
	Create func(parent string) *Featurestore
	// Update should return a resource which is valid to update, i.e.
	// all required fields set.
	Update func(parent string) *Featurestore
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *FeaturestoreTestSuiteConfig) test(t *testing.T) {
	t.Run("Create", fx.testCreate)
	t.Run("Get", fx.testGet)
	t.Run("Update", fx.testUpdate)
	t.Run("List", fx.testList)
}

func (fx *FeaturestoreTestSuiteConfig) testCreate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no parent is provided.
	t.Run("missing parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreateFeaturestore(fx.ctx, &CreateFeaturestoreRequest{
			Parent:       "",
			Featurestore: fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreateFeaturestore(fx.ctx, &CreateFeaturestoreRequest{
			Parent:       "invalid resource name",
			Featurestore: fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// The method should fail with InvalidArgument if the resource has any
	// required fields and they are not provided.
	t.Run("required fields", func(t *testing.T) {
		fx.maybeSkip(t)
		t.Run(".online_serving_config.scaling.min_node_count", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetOnlineServingConfig().GetScaling()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("min_node_count")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateFeaturestore(fx.ctx, &CreateFeaturestoreRequest{
				Parent:       parent,
				Featurestore: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".encryption_spec.kms_key_name", func(t *testing.T) {
			fx.maybeSkip(t)
			parent := fx.nextParent(t, false)
			msg := fx.Create(parent)
			container := msg.GetEncryptionSpec()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("kms_key_name")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateFeaturestore(fx.ctx, &CreateFeaturestoreRequest{
				Parent:       parent,
				Featurestore: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
	})

}

func (fx *FeaturestoreTestSuiteConfig) testGet(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetFeaturestore(fx.ctx, &GetFeaturestoreRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetFeaturestore(fx.ctx, &GetFeaturestoreRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be returned without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		msg, err := fx.service.GetFeaturestore(fx.ctx, &GetFeaturestoreRequest{
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
		_, err := fx.service.GetFeaturestore(fx.ctx, &GetFeaturestoreRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetFeaturestore(fx.ctx, &GetFeaturestoreRequest{
			Name: "projects/-/locations/-/featurestores/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *FeaturestoreTestSuiteConfig) testUpdate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = ""
		_, err := fx.service.UpdateFeaturestore(fx.ctx, &UpdateFeaturestoreRequest{
			Featurestore: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = "invalid resource name"
		_, err := fx.service.UpdateFeaturestore(fx.ctx, &UpdateFeaturestoreRequest{
			Featurestore: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	parent := fx.nextParent(t, false)
	created := fx.create(t, parent)
	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		msg := fx.Update(parent)
		msg.Name = created.Name + "notfound"
		_, err := fx.service.UpdateFeaturestore(fx.ctx, &UpdateFeaturestoreRequest{
			Featurestore: msg,
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// The method should fail with InvalidArgument if the update_mask is invalid.
	t.Run("invalid update mask", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.UpdateFeaturestore(fx.ctx, &UpdateFeaturestoreRequest{
			Featurestore: created,
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
		t.Run(".online_serving_config.scaling.min_node_count", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created).(*Featurestore)
			container := msg.GetOnlineServingConfig().GetScaling()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("min_node_count")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.UpdateFeaturestore(fx.ctx, &UpdateFeaturestoreRequest{
				Featurestore: msg,
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{
						"*",
					},
				},
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".encryption_spec.kms_key_name", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created).(*Featurestore)
			container := msg.GetEncryptionSpec()
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("kms_key_name")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.UpdateFeaturestore(fx.ctx, &UpdateFeaturestoreRequest{
				Featurestore: msg,
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

func (fx *FeaturestoreTestSuiteConfig) testList(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.ListFeaturestores(fx.ctx, &ListFeaturestoresRequest{
			Parent: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page token is not valid.
	t.Run("invalid page token", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.service.ListFeaturestores(fx.ctx, &ListFeaturestoresRequest{
			Parent:    parent,
			PageToken: "invalid page token",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page size is negative.
	t.Run("negative page size", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.service.ListFeaturestores(fx.ctx, &ListFeaturestoresRequest{
			Parent:   parent,
			PageSize: -10,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	const resourcesCount = 15
	parent := fx.nextParent(t, true)
	parentMsgs := make([]*Featurestore, resourcesCount)
	for i := 0; i < resourcesCount; i++ {
		parentMsgs[i] = fx.create(t, parent)
	}

	// If parent is provided the method must only return resources
	// under that parent.
	t.Run("isolation", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.ListFeaturestores(fx.ctx, &ListFeaturestoresRequest{
			Parent:   parent,
			PageSize: 999,
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			parentMsgs,
			response.Featurestores,
			cmpopts.SortSlices(func(a, b *Featurestore) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

	// If there are no more resources, next_page_token should not be set.
	t.Run("last page", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.ListFeaturestores(fx.ctx, &ListFeaturestoresRequest{
			Parent:   parent,
			PageSize: resourcesCount,
		})
		assert.NilError(t, err)
		assert.Equal(t, "", response.NextPageToken)
	})

	// If there are more resources, next_page_token should be set.
	t.Run("more pages", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.ListFeaturestores(fx.ctx, &ListFeaturestoresRequest{
			Parent:   parent,
			PageSize: resourcesCount - 1,
		})
		assert.NilError(t, err)
		assert.Check(t, response.NextPageToken != "")
	})

	// Listing resource one by one should eventually return all resources.
	t.Run("one by one", func(t *testing.T) {
		fx.maybeSkip(t)
		msgs := make([]*Featurestore, 0, resourcesCount)
		var nextPageToken string
		for {
			response, err := fx.service.ListFeaturestores(fx.ctx, &ListFeaturestoresRequest{
				Parent:    parent,
				PageSize:  1,
				PageToken: nextPageToken,
			})
			assert.NilError(t, err)
			assert.Equal(t, 1, len(response.Featurestores))
			msgs = append(msgs, response.Featurestores...)
			nextPageToken = response.NextPageToken
			if nextPageToken == "" {
				break
			}
		}
		assert.DeepEqual(
			t,
			parentMsgs,
			msgs,
			cmpopts.SortSlices(func(a, b *Featurestore) bool {
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
			_, err := fx.service.DeleteFeaturestore(fx.ctx, &DeleteFeaturestoreRequest{
				Name: parentMsgs[i].Name,
			})
			assert.NilError(t, err)
		}
		response, err := fx.service.ListFeaturestores(fx.ctx, &ListFeaturestoresRequest{
			Parent:   parent,
			PageSize: 9999,
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			parentMsgs[deleteCount:],
			response.Featurestores,
			cmpopts.SortSlices(func(a, b *Featurestore) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

}

func (fx *FeaturestoreTestSuiteConfig) nextParent(t *testing.T, pristine bool) string {
	if pristine {
		fx.currParent++
	}
	if fx.currParent >= len(fx.Parents) {
		t.Fatal("need at least", fx.currParent+1, "parents")
	}
	return fx.Parents[fx.currParent]
}

func (fx *FeaturestoreTestSuiteConfig) peekNextParent(t *testing.T) string {
	next := fx.currParent + 1
	if next >= len(fx.Parents) {
		t.Fatal("need at least", next+1, "parents")
	}
	return fx.Parents[next]
}

func (fx *FeaturestoreTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}

func (fx *FeaturestoreTestSuiteConfig) create(t *testing.T, parent string) *Featurestore {
	t.Helper()
	t.Skip("Long running create method not supported")
	return nil
}
