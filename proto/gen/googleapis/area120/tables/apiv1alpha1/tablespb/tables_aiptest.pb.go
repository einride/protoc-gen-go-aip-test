// Code generated by protoc-gen-go-aip-test. DO NOT EDIT.

package tablespb

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

// TablesServiceTestSuiteConfigProvider is the interface to implement to decide which resources
// that should be tested and how it's configured.
type TablesServiceTestSuiteConfigProvider interface {
	// TablesServiceRow should return a config, or nil, which means that the tests will be skipped.
	TablesServiceRow(t *testing.T) *TablesServiceRowTestSuiteConfig
	// TablesServiceTable should return a config, or nil, which means that the tests will be skipped.
	TablesServiceTable(t *testing.T) *TablesServiceTableTestSuiteConfig
	// TablesServiceWorkspace should return a config, or nil, which means that the tests will be skipped.
	TablesServiceWorkspace(t *testing.T) *TablesServiceWorkspaceTestSuiteConfig
}

// testTablesService is the main entrypoint for starting the AIP tests.
func testTablesService(t *testing.T, s TablesServiceTestSuiteConfigProvider) {
	testTablesServiceRow(t, s)
	testTablesServiceTable(t, s)
	testTablesServiceWorkspace(t, s)
}

func testTablesServiceRow(t *testing.T, s TablesServiceTestSuiteConfigProvider) {
	t.Run("Row", func(t *testing.T) {
		config := s.TablesServiceRow(t)
		if config == nil {
			t.Skip("Method TablesServiceRow not implemented")
		}
		if config.Service == nil {
			t.Skip("Method TablesServiceRow.Service() not implemented")
		}
		if config.Context == nil {
			config.Context = func() context.Context { return context.Background() }
		}
		config.test(t)
	})
}

func testTablesServiceTable(t *testing.T, s TablesServiceTestSuiteConfigProvider) {
	t.Run("Table", func(t *testing.T) {
		config := s.TablesServiceTable(t)
		if config == nil {
			t.Skip("Method TablesServiceTable not implemented")
		}
		if config.Service == nil {
			t.Skip("Method TablesServiceTable.Service() not implemented")
		}
		if config.Context == nil {
			config.Context = func() context.Context { return context.Background() }
		}
		config.test(t)
	})
}

func testTablesServiceWorkspace(t *testing.T, s TablesServiceTestSuiteConfigProvider) {
	t.Run("Workspace", func(t *testing.T) {
		config := s.TablesServiceWorkspace(t)
		if config == nil {
			t.Skip("Method TablesServiceWorkspace not implemented")
		}
		if config.Service == nil {
			t.Skip("Method TablesServiceWorkspace.Service() not implemented")
		}
		if config.Context == nil {
			config.Context = func() context.Context { return context.Background() }
		}
		config.test(t)
	})
}

type TablesServiceTestSuite struct {
	T *testing.T
	// Server to test.
	Server TablesServiceServer
}

func (fx TablesServiceTestSuite) TestRow(ctx context.Context, options TablesServiceRowTestSuiteConfig) {
	fx.T.Run("Row", func(t *testing.T) {
		options.Context = func() context.Context { return ctx }
		options.Service = func() TablesServiceServer { return fx.Server }
		options.test(t)
	})
}

func (fx TablesServiceTestSuite) TestTable(ctx context.Context, options TablesServiceTableTestSuiteConfig) {
	fx.T.Run("Table", func(t *testing.T) {
		options.Context = func() context.Context { return ctx }
		options.Service = func() TablesServiceServer { return fx.Server }
		options.test(t)
	})
}

func (fx TablesServiceTestSuite) TestWorkspace(ctx context.Context, options TablesServiceWorkspaceTestSuiteConfig) {
	fx.T.Run("Workspace", func(t *testing.T) {
		options.Context = func() context.Context { return ctx }
		options.Service = func() TablesServiceServer { return fx.Server }
		options.test(t)
	})
}

type TablesServiceRowTestSuiteConfig struct {
	currParent int

	// Service should return the service that should be tested.
	// The service will be used for several tests.
	Service func() TablesServiceServer
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
	Create func(parent string) *Row
	// Update should return a resource which is valid to update, i.e.
	// all required fields set.
	Update func(parent string) *Row
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *TablesServiceRowTestSuiteConfig) test(t *testing.T) {
	t.Run("Create", fx.testCreate)
	t.Run("Get", fx.testGet)
	t.Run("Update", fx.testUpdate)
	t.Run("List", fx.testList)
	t.Run("Delete", fx.testDelete)
}

func (fx *TablesServiceRowTestSuiteConfig) testCreate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no parent is provided.
	t.Run("missing parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().CreateRow(fx.Context(), &CreateRowRequest{
			Parent: "",
			Row:    fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().CreateRow(fx.Context(), &CreateRowRequest{
			Parent: "invalid resource name",
			Row:    fx.Create(fx.nextParent(t, false)),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// The created resource should be persisted and reachable with Get.
	t.Run("persisted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg, err := fx.Service().CreateRow(fx.Context(), &CreateRowRequest{
			Parent: parent,
			Row:    fx.Create(parent),
		})
		assert.NilError(t, err)
		persisted, err := fx.Service().GetRow(fx.Context(), &GetRowRequest{
			Name: msg.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, msg, persisted, protocmp.Transform())
	})

}

func (fx *TablesServiceRowTestSuiteConfig) testGet(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetRow(fx.Context(), &GetRowRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetRow(fx.Context(), &GetRowRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be returned without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		msg, err := fx.Service().GetRow(fx.Context(), &GetRowRequest{
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
		_, err := fx.Service().GetRow(fx.Context(), &GetRowRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetRow(fx.Context(), &GetRowRequest{
			Name: "tables/-/rows/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *TablesServiceRowTestSuiteConfig) testUpdate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = ""
		_, err := fx.Service().UpdateRow(fx.Context(), &UpdateRowRequest{
			Row: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = "invalid resource name"
		_, err := fx.Service().UpdateRow(fx.Context(), &UpdateRowRequest{
			Row: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// The updated resource should be persisted and reachable with Get.
	t.Run("persisted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		updated, err := fx.Service().UpdateRow(fx.Context(), &UpdateRowRequest{
			Row: created,
		})
		assert.NilError(t, err)
		persisted, err := fx.Service().GetRow(fx.Context(), &GetRowRequest{
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
			_, err := fx.Service().UpdateRow(fx.Context(), &UpdateRowRequest{
				Row: msg,
			})
			assert.Equal(t, codes.NotFound, status.Code(err), err)
		})

		// The method should fail with InvalidArgument if the update_mask is invalid.
		t.Run("invalid update mask", func(t *testing.T) {
			fx.maybeSkip(t)
			_, err := fx.Service().UpdateRow(fx.Context(), &UpdateRowRequest{
				Row: created,
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

func (fx *TablesServiceRowTestSuiteConfig) testList(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if provided parent is invalid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().ListRows(fx.Context(), &ListRowsRequest{
			Parent: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page token is not valid.
	t.Run("invalid page token", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.Service().ListRows(fx.Context(), &ListRowsRequest{
			Parent:    parent,
			PageToken: "invalid page token",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page size is negative.
	t.Run("negative page size", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		_, err := fx.Service().ListRows(fx.Context(), &ListRowsRequest{
			Parent:   parent,
			PageSize: -10,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	{
		const resourcesCount = 15
		parent := fx.nextParent(t, true)
		parentMsgs := make([]*Row, resourcesCount)
		for i := 0; i < resourcesCount; i++ {
			parentMsgs[i] = fx.create(t, parent)
		}

		// If parent is provided the method must only return resources
		// under that parent.
		t.Run("isolation", func(t *testing.T) {
			fx.maybeSkip(t)
			response, err := fx.Service().ListRows(fx.Context(), &ListRowsRequest{
				Parent:   parent,
				PageSize: 999,
			})
			assert.NilError(t, err)
			assert.DeepEqual(
				t,
				parentMsgs,
				response.Rows,
				cmpopts.SortSlices(func(a, b *Row) bool {
					return a.Name < b.Name
				}),
				protocmp.Transform(),
			)
		})

		// If there are no more resources, next_page_token should not be set.
		t.Run("last page", func(t *testing.T) {
			fx.maybeSkip(t)
			response, err := fx.Service().ListRows(fx.Context(), &ListRowsRequest{
				Parent:   parent,
				PageSize: resourcesCount,
			})
			assert.NilError(t, err)
			assert.Equal(t, "", response.NextPageToken)
		})

		// If there are more resources, next_page_token should be set.
		t.Run("more pages", func(t *testing.T) {
			fx.maybeSkip(t)
			response, err := fx.Service().ListRows(fx.Context(), &ListRowsRequest{
				Parent:   parent,
				PageSize: resourcesCount - 1,
			})
			assert.NilError(t, err)
			assert.Check(t, response.NextPageToken != "")
		})

		// Listing resource one by one should eventually return all resources.
		t.Run("one by one", func(t *testing.T) {
			fx.maybeSkip(t)
			msgs := make([]*Row, 0, resourcesCount)
			var nextPageToken string
			for {
				response, err := fx.Service().ListRows(fx.Context(), &ListRowsRequest{
					Parent:    parent,
					PageSize:  1,
					PageToken: nextPageToken,
				})
				assert.NilError(t, err)
				assert.Equal(t, 1, len(response.Rows))
				msgs = append(msgs, response.Rows...)
				nextPageToken = response.NextPageToken
				if nextPageToken == "" {
					break
				}
			}
			assert.DeepEqual(
				t,
				parentMsgs,
				msgs,
				cmpopts.SortSlices(func(a, b *Row) bool {
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
				_, err := fx.Service().DeleteRow(fx.Context(), &DeleteRowRequest{
					Name: parentMsgs[i].Name,
				})
				assert.NilError(t, err)
			}
			response, err := fx.Service().ListRows(fx.Context(), &ListRowsRequest{
				Parent:   parent,
				PageSize: 9999,
			})
			assert.NilError(t, err)
			assert.DeepEqual(
				t,
				parentMsgs[deleteCount:],
				response.Rows,
				cmpopts.SortSlices(func(a, b *Row) bool {
					return a.Name < b.Name
				}),
				protocmp.Transform(),
			)
		})

	}
	{
		const resourcesCount = 101
		parent := fx.nextParent(t, true)
		parentMsgs := make([]*Row, resourcesCount)
		for i := 0; i < resourcesCount; i++ {
			parentMsgs[i] = fx.create(t, parent)
		}

		// Listing resource with page size zero should eventually return all resources.
		t.Run("page size zero", func(t *testing.T) {
			fx.maybeSkip(t)
			msgs := make([]*Row, 0, resourcesCount)
			var nextPageToken string
			for {
				page, err := fx.Service().ListRows(fx.Context(), &ListRowsRequest{
					Parent:    parent,
					PageSize:  0,
					PageToken: nextPageToken,
				})
				assert.NilError(t, err)
				msgs = append(msgs, page.Rows...)
				nextPageToken = page.NextPageToken
				if nextPageToken == "" {
					break
				}
			}
			assert.DeepEqual(
				t,
				parentMsgs,
				msgs,
				cmpopts.SortSlices(func(a, b *Row) bool {
					return a.Name < b.Name
				}),
				protocmp.Transform(),
			)
		})

	}
}

func (fx *TablesServiceRowTestSuiteConfig) testDelete(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().DeleteRow(fx.Context(), &DeleteRowRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().DeleteRow(fx.Context(), &DeleteRowRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be deleted without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		_, err := fx.Service().DeleteRow(fx.Context(), &DeleteRowRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
	})

	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		_, err := fx.Service().DeleteRow(fx.Context(), &DeleteRowRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with NotFound if the resource was already deleted. This also applies to soft-deletion.
	t.Run("already deleted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		deleted, err := fx.Service().DeleteRow(fx.Context(), &DeleteRowRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
		_ = deleted
		_, err = fx.Service().DeleteRow(fx.Context(), &DeleteRowRequest{
			Name: created.Name,
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().DeleteRow(fx.Context(), &DeleteRowRequest{
			Name: "tables/-/rows/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *TablesServiceRowTestSuiteConfig) nextParent(t *testing.T, pristine bool) string {
	if pristine {
		fx.currParent++
	}
	if fx.currParent >= len(fx.Parents) {
		t.Fatal("need at least", fx.currParent+1, "parents")
	}
	return fx.Parents[fx.currParent]
}

func (fx *TablesServiceRowTestSuiteConfig) peekNextParent(t *testing.T) string {
	next := fx.currParent + 1
	if next >= len(fx.Parents) {
		t.Fatal("need at least", next+1, "parents")
	}
	return fx.Parents[next]
}

func (fx *TablesServiceRowTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}

func (fx *TablesServiceRowTestSuiteConfig) create(t *testing.T, parent string) *Row {
	t.Helper()
	created, err := fx.Service().CreateRow(fx.Context(), &CreateRowRequest{
		Parent: parent,
		Row:    fx.Create(parent),
	})
	assert.NilError(t, err)
	return created
}

type TablesServiceTableTestSuiteConfig struct {
	currParent int

	// Service should return the service that should be tested.
	// The service will be used for several tests.
	Service func() TablesServiceServer
	// Context should return a new context.
	// The context will be used for several tests.
	Context func() context.Context
	// CreateResource should create a Table and return it.
	// If the field is not set, some tests will be skipped.
	//
	// This method is generated because service does not expose a Create
	// method (or it does not comply with AIP).
	CreateResource func(ctx context.Context) (*Table, error)
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *TablesServiceTableTestSuiteConfig) test(t *testing.T) {
	t.Run("Get", fx.testGet)
	t.Run("List", fx.testList)
}

func (fx *TablesServiceTableTestSuiteConfig) testGet(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetTable(fx.Context(), &GetTableRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetTable(fx.Context(), &GetTableRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be returned without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		created := fx.create(t)
		msg, err := fx.Service().GetTable(fx.Context(), &GetTableRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, msg, created, protocmp.Transform())
	})

	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		created := fx.create(t)
		_, err := fx.Service().GetTable(fx.Context(), &GetTableRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetTable(fx.Context(), &GetTableRequest{
			Name: "tables/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *TablesServiceTableTestSuiteConfig) testList(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument is provided page token is not valid.
	t.Run("invalid page token", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().ListTables(fx.Context(), &ListTablesRequest{
			PageToken: "invalid page token",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page size is negative.
	t.Run("negative page size", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().ListTables(fx.Context(), &ListTablesRequest{
			PageSize: -10,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *TablesServiceTableTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}

func (fx *TablesServiceTableTestSuiteConfig) create(t *testing.T) *Table {
	t.Helper()
	if fx.CreateResource == nil {
		t.Skip("Test skipped because CreateResource not specified on TablesServiceTableTestSuiteConfig")
	}
	created, err := fx.CreateResource(fx.Context())
	assert.NilError(t, err)
	return created
}

type TablesServiceWorkspaceTestSuiteConfig struct {
	currParent int

	// Service should return the service that should be tested.
	// The service will be used for several tests.
	Service func() TablesServiceServer
	// Context should return a new context.
	// The context will be used for several tests.
	Context func() context.Context
	// CreateResource should create a Workspace and return it.
	// If the field is not set, some tests will be skipped.
	//
	// This method is generated because service does not expose a Create
	// method (or it does not comply with AIP).
	CreateResource func(ctx context.Context) (*Workspace, error)
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *TablesServiceWorkspaceTestSuiteConfig) test(t *testing.T) {
	t.Run("Get", fx.testGet)
	t.Run("List", fx.testList)
}

func (fx *TablesServiceWorkspaceTestSuiteConfig) testGet(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetWorkspace(fx.Context(), &GetWorkspaceRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetWorkspace(fx.Context(), &GetWorkspaceRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be returned without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		created := fx.create(t)
		msg, err := fx.Service().GetWorkspace(fx.Context(), &GetWorkspaceRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, msg, created, protocmp.Transform())
	})

	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		created := fx.create(t)
		_, err := fx.Service().GetWorkspace(fx.Context(), &GetWorkspaceRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().GetWorkspace(fx.Context(), &GetWorkspaceRequest{
			Name: "workspaces/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *TablesServiceWorkspaceTestSuiteConfig) testList(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument is provided page token is not valid.
	t.Run("invalid page token", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().ListWorkspaces(fx.Context(), &ListWorkspacesRequest{
			PageToken: "invalid page token",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page size is negative.
	t.Run("negative page size", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.Service().ListWorkspaces(fx.Context(), &ListWorkspacesRequest{
			PageSize: -10,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *TablesServiceWorkspaceTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}

func (fx *TablesServiceWorkspaceTestSuiteConfig) create(t *testing.T) *Workspace {
	t.Helper()
	if fx.CreateResource == nil {
		t.Skip("Test skipped because CreateResource not specified on TablesServiceWorkspaceTestSuiteConfig")
	}
	created, err := fx.CreateResource(fx.Context())
	assert.NilError(t, err)
	return created
}
