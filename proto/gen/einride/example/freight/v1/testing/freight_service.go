// Code generated by protoc-gen-go-aiptest. DO NOT EDIT.

package examplefreightv1test

import (
	context "context"
	v1 "github.com/einride/protoc-gen-go-aiptest/proto/gen/einride/example/freight/v1"
	cmpopts "github.com/google/go-cmp/cmp/cmpopts"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	proto "google.golang.org/protobuf/proto"
	protocmp "google.golang.org/protobuf/testing/protocmp"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	assert "gotest.tools/v3/assert"
	strings "strings"
	testing "testing"
	time "time"
)

type FreightService struct {
	T *testing.T
	// Server to test.
	Server v1.FreightServiceServer
}

func (fx *FreightService) TestShipper(ctx context.Context, options Shipper) {
	fx.T.Run("Shipper", func(t *testing.T) {
		options.ctx = ctx
		options.service = fx.Server
		options.test(t)
	})
}

func (fx *FreightService) TestSite(ctx context.Context, options Site) {
	fx.T.Run("Site", func(t *testing.T) {
		options.ctx = ctx
		options.service = fx.Server
		options.test(t)
	})
}

type Shipper struct {
	ctx        context.Context
	service    v1.FreightServiceServer
	currParent int

	// Create should return a resource which is valid to create, i.e.
	// all required fields set.
	Create func() *v1.Shipper
	// Update should return a resource which is valid to update, i.e.
	// all required fields set.
	Update func() *v1.Shipper
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *Shipper) test(t *testing.T) {
	t.Run("Create", fx.testCreate)
	t.Run("Get", fx.testGet)
	t.Run("Update", fx.testUpdate)
	t.Run("List", fx.testList)
}

func (fx *Shipper) testCreate(t *testing.T) {
	// Standard methods: Create
	// https://google.aip.dev/133

	// Field create_time should be populated when the resource is created.
	t.Run("create time", func(t *testing.T) {
		fx.maybeSkip(t)
		msg, err := fx.service.CreateShipper(fx.ctx, &v1.CreateShipperRequest{
			Shipper: fx.Create(),
		})
		assert.NilError(t, err)
		assert.Check(t, time.Since(msg.CreateTime.AsTime()) < time.Second)
	})

	// The created resource should be persisted and reachable with Get.
	t.Run("persisted", func(t *testing.T) {
		fx.maybeSkip(t)
		msg, err := fx.service.CreateShipper(fx.ctx, &v1.CreateShipperRequest{
			Shipper: fx.Create(),
		})
		assert.NilError(t, err)
		persisted, err := fx.service.GetShipper(fx.ctx, &v1.GetShipperRequest{
			Name: msg.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, msg, persisted, protocmp.Transform())
	})

	// If method support user settable IDs, when set the resource should
	// returned with the provided ID.
	t.Run("user settable id", func(t *testing.T) {
		fx.maybeSkip(t)
		msg, err := fx.service.CreateShipper(fx.ctx, &v1.CreateShipperRequest{
			Shipper:   fx.Create(),
			ShipperId: "usersetid",
		})
		assert.NilError(t, err)
		assert.Check(t, strings.HasSuffix(msg.GetName(), "usersetid"))
	})

	// If method support user settable IDs and the same ID is reused
	// the method should return AlreadyExists.
	t.Run("already exists", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreateShipper(fx.ctx, &v1.CreateShipperRequest{
			Shipper:   fx.Create(),
			ShipperId: "alreadyexists",
		})
		assert.NilError(t, err)
		_, err = fx.service.CreateShipper(fx.ctx, &v1.CreateShipperRequest{
			Shipper:   fx.Create(),
			ShipperId: "alreadyexists",
		})
		assert.Equal(t, codes.AlreadyExists, status.Code(err), err)
	})

	t.Run("required fields", func(t *testing.T) {
		fx.maybeSkip(t)
		t.Run(".display_name", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := fx.Create()
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("display_name")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateShipper(fx.ctx, &v1.CreateShipperRequest{
				Shipper: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".billing_account", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := fx.Create()
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("billing_account")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateShipper(fx.ctx, &v1.CreateShipperRequest{
				Shipper: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
	})

	// If resource references are accepted on the resource, they must be validated.
	t.Run("resource references", func(t *testing.T) {
		fx.maybeSkip(t)
		t.Run(".billing_account", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := fx.Create()
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			container.BillingAccount = "invalid resource name"
			_, err := fx.service.CreateShipper(fx.ctx, &v1.CreateShipperRequest{
				Shipper: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
	})

	_ = time.Second
	_ = strings.HasSuffix
	_ = codes.InvalidArgument
	_ = protocmp.Transform
}

func (fx *Shipper) testGet(t *testing.T) {
	// Standard methods: Get
	// https://google.aip.dev/131
	created00, err := fx.service.CreateShipper(fx.ctx, &v1.CreateShipperRequest{
		Shipper: fx.Create(),
	})
	assert.NilError(t, err)

	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetShipper(fx.ctx, &v1.GetShipperRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetShipper(fx.ctx, &v1.GetShipperRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be returned without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		msg, err := fx.service.GetShipper(fx.ctx, &v1.GetShipperRequest{
			Name: created00.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, msg, created00, protocmp.Transform())
	})

	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetShipper(fx.ctx, &v1.GetShipperRequest{
			Name: created00.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})
	_ = codes.NotFound
	_ = protocmp.Transform
}

func (fx *Shipper) testUpdate(t *testing.T) {
	// Standard methods: Update
	// https://google.aip.dev/134
	created00, err := fx.service.CreateShipper(fx.ctx, &v1.CreateShipperRequest{
		Shipper: fx.Create(),
	})
	assert.NilError(t, err)

	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		msg := fx.Update()
		msg.Name = ""
		_, err := fx.service.UpdateShipper(fx.ctx, &v1.UpdateShipperRequest{
			Shipper: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		msg := fx.Update()
		msg.Name = "invalid resource name"
		_, err := fx.service.UpdateShipper(fx.ctx, &v1.UpdateShipperRequest{
			Shipper: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Field update_time should be updated when the resource is updated.
	t.Run("update time", func(t *testing.T) {
		fx.maybeSkip(t)
		initial, err := fx.service.CreateShipper(fx.ctx, &v1.CreateShipperRequest{
			Shipper: fx.Create(),
		})
		assert.NilError(t, err)
		updated, err := fx.service.UpdateShipper(fx.ctx, &v1.UpdateShipperRequest{
			Shipper: initial,
		})
		assert.NilError(t, err)
		assert.Check(t, updated.UpdateTime.AsTime().After(initial.UpdateTime.AsTime()))
	})

	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		msg := fx.Update()
		msg.Name = created00.Name + "notfound"
		_, err := fx.service.UpdateShipper(fx.ctx, &v1.UpdateShipperRequest{
			Shipper: msg,
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// The updated resource should be persisted and reachable with Get.
	t.Run("persisted", func(t *testing.T) {
		fx.maybeSkip(t)
		initial, err := fx.service.CreateShipper(fx.ctx, &v1.CreateShipperRequest{
			Shipper: fx.Create(),
		})
		assert.NilError(t, err)
		updated, err := fx.service.UpdateShipper(fx.ctx, &v1.UpdateShipperRequest{
			Shipper: initial,
		})
		persisted, err := fx.service.GetShipper(fx.ctx, &v1.GetShipperRequest{
			Name: updated.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, updated, persisted, protocmp.Transform())
	})

	// The method should fail with InvalidArgument if the update_mask is invalid.
	t.Run("invalid update mask", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.UpdateShipper(fx.ctx, &v1.UpdateShipperRequest{
			Shipper: created00,
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
		t.Run(".display_name", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created00).(*v1.Shipper)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("display_name")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateShipper(fx.ctx, &v1.CreateShipperRequest{
				Shipper: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
		t.Run(".billing_account", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created00).(*v1.Shipper)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("billing_account")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateShipper(fx.ctx, &v1.CreateShipperRequest{
				Shipper: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
	})
	_ = codes.NotFound
	_ = protocmp.Transform
	_ = proto.Clone
}

func (fx *Shipper) testList(t *testing.T) {
	// Standard methods: List
	// https://google.aip.dev/132
	const n = 15

	parent02msgs := make([]*v1.Shipper, n)
	for i := 0; i < n; i++ {
		msg, err := fx.service.CreateShipper(fx.ctx, &v1.CreateShipperRequest{
			Shipper: fx.Create(),
		})
		assert.NilError(t, err)
		parent02msgs[i] = msg
	}

	// Method should fail with InvalidArgument is provided page token is not valid.
	t.Run("invalid page token", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.ListShippers(fx.ctx, &v1.ListShippersRequest{
			PageToken: "invalid page token",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page size is negative.
	t.Run("negative page size", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.ListShippers(fx.ctx, &v1.ListShippersRequest{
			PageSize: -10,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})
	_ = codes.NotFound
	_ = protocmp.Transform
	_ = cmpopts.SortSlices
}

func (fx *Shipper) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}

type Site struct {
	ctx        context.Context
	service    v1.FreightServiceServer
	currParent int

	// The parents to use when creating resources.
	// At least one parent needs to be set. Depending on methods available on the resource,
	// more may be required. If insufficient number of parents are
	// provided the test will fail.
	Parents []string
	// Create should return a resource which is valid to create, i.e.
	// all required fields set.
	Create func(parent string) *v1.Site
	// Update should return a resource which is valid to update, i.e.
	// all required fields set.
	Update func(parent string) *v1.Site
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *Site) test(t *testing.T) {
	t.Run("Create", fx.testCreate)
	t.Run("Get", fx.testGet)
	t.Run("BatchGet", fx.testBatchGet)
	t.Run("Update", fx.testUpdate)
	t.Run("List", fx.testList)
}

func (fx *Site) testCreate(t *testing.T) {
	// Standard methods: Create
	// https://google.aip.dev/133

	parent := fx.nextParent(t, false)

	// Method should fail with InvalidArgument if no parent is provided.
	t.Run("missing parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreateSite(fx.ctx, &v1.CreateSiteRequest{
			Parent: "",
			Site:   fx.Create(""),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided parent is not valid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.CreateSite(fx.ctx, &v1.CreateSiteRequest{
			Parent: "invalid resource name",
			Site:   fx.Create("invalid resource name"),
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Field create_time should be populated when the resource is created.
	t.Run("create time", func(t *testing.T) {
		fx.maybeSkip(t)
		msg, err := fx.service.CreateSite(fx.ctx, &v1.CreateSiteRequest{
			Parent: parent,
			Site:   fx.Create(parent),
		})
		assert.NilError(t, err)
		assert.Check(t, time.Since(msg.CreateTime.AsTime()) < time.Second)
	})

	// The created resource should be persisted and reachable with Get.
	t.Run("persisted", func(t *testing.T) {
		fx.maybeSkip(t)
		msg, err := fx.service.CreateSite(fx.ctx, &v1.CreateSiteRequest{
			Parent: parent,
			Site:   fx.Create(parent),
		})
		assert.NilError(t, err)
		persisted, err := fx.service.GetSite(fx.ctx, &v1.GetSiteRequest{
			Name: msg.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, msg, persisted, protocmp.Transform())
	})

	t.Run("required fields", func(t *testing.T) {
		fx.maybeSkip(t)
		t.Run(".display_name", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := fx.Create(parent)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("display_name")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateSite(fx.ctx, &v1.CreateSiteRequest{
				Parent: parent,
				Site:   msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
	})

	// If resource references are accepted on the resource, they must be validated.
	t.Run("resource references", func(t *testing.T) {
		fx.maybeSkip(t)
		t.Run(".billing.billing_account", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := fx.Create(parent)
			container := msg.GetBilling()
			if container == nil {
				t.Skip("not reachable")
			}
			container.BillingAccount = "invalid resource name"
			_, err := fx.service.CreateSite(fx.ctx, &v1.CreateSiteRequest{
				Parent: parent,
				Site:   msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
	})

	_ = time.Second
	_ = strings.HasSuffix
	_ = codes.InvalidArgument
	_ = protocmp.Transform
}

func (fx *Site) testGet(t *testing.T) {
	// Standard methods: Get
	// https://google.aip.dev/131

	parent := fx.nextParent(t, false)
	created00, err := fx.service.CreateSite(fx.ctx, &v1.CreateSiteRequest{
		Parent: parent,
		Site:   fx.Create(parent),
	})
	assert.NilError(t, err)

	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetSite(fx.ctx, &v1.GetSiteRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetSite(fx.ctx, &v1.GetSiteRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be returned without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		msg, err := fx.service.GetSite(fx.ctx, &v1.GetSiteRequest{
			Name: created00.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, msg, created00, protocmp.Transform())
	})

	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetSite(fx.ctx, &v1.GetSiteRequest{
			Name: created00.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})
	_ = codes.NotFound
	_ = protocmp.Transform
}

func (fx *Site) testBatchGet(t *testing.T) {
	// Batch methods: Get
	// https://google.aip.dev/231

	parent := fx.nextParent(t, false)
	created00, err := fx.service.CreateSite(fx.ctx, &v1.CreateSiteRequest{
		Parent: parent,
		Site:   fx.Create(parent),
	})
	assert.NilError(t, err)
	created01, err := fx.service.CreateSite(fx.ctx, &v1.CreateSiteRequest{
		Parent: parent,
		Site:   fx.Create(parent),
	})
	assert.NilError(t, err)
	created02, err := fx.service.CreateSite(fx.ctx, &v1.CreateSiteRequest{
		Parent: parent,
		Site:   fx.Create(parent),
	})
	assert.NilError(t, err)

	// Method should fail with InvalidArgument if provided parent is not valid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.BatchGetSites(fx.ctx, &v1.BatchGetSitesRequest{
			Parent: "invalid resource name",
			Names: []string{
				created00.Name,
			},
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if no names are provided.
	t.Run("no names", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.BatchGetSites(fx.ctx, &v1.BatchGetSitesRequest{
			Parent: parent,
			Names:  []string{},
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if a provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.BatchGetSites(fx.ctx, &v1.BatchGetSitesRequest{
			Parent: parent,
			Names: []string{
				"invalid resource name",
			},
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resources should be returned without errors if they exist.
	t.Run("all exists", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.BatchGetSites(fx.ctx, &v1.BatchGetSitesRequest{
			Parent: parent,
			Names: []string{
				created00.Name,
				created01.Name,
				created02.Name,
			},
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			[]*v1.Site{
				created00,
				created01,
				created02,
			},
			response.Sites,
			protocmp.Transform(),
		)
	})

	// The method must be atomic; it must fail for all resources
	// or succeed for all resources (no partial success).
	t.Run("atomic", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.BatchGetSites(fx.ctx, &v1.BatchGetSitesRequest{
			Parent: parent,
			Names: []string{
				created00.Name,
				created01.Name + "notfound",
				created02.Name,
			},
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// If a caller sets the "parent", and the parent collection in the name of any resource
	// being retrieved does not match, the request must fail.
	t.Run("parent mismatch", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.BatchGetSites(fx.ctx, &v1.BatchGetSitesRequest{
			Parent: fx.peekNextParent(t),
			Names: []string{
				created00.Name,
			},
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// The order of resources in the response must be the same as the names in the request.
	t.Run("ordered", func(t *testing.T) {
		fx.maybeSkip(t)
		for _, order := range [][]*v1.Site{
			{created00, created01, created02},
			{created01, created00, created02},
			{created02, created01, created00},
		} {
			response, err := fx.service.BatchGetSites(fx.ctx, &v1.BatchGetSitesRequest{
				Parent: parent,
				Names: []string{
					order[0].GetName(),
					order[1].GetName(),
					order[2].GetName(),
				},
			})
			assert.NilError(t, err)
			assert.DeepEqual(t, order, response.Sites, protocmp.Transform())
		}
	})

	// If a caller provides duplicate names, the service should return
	// duplicate resources.
	t.Run("duplicate names", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.BatchGetSites(fx.ctx, &v1.BatchGetSitesRequest{
			Parent: parent,
			Names: []string{
				created00.Name,
				created00.Name,
			},
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			[]*v1.Site{
				created00,
				created00,
			},
			response.Sites,
			protocmp.Transform(),
		)
	})

	_ = codes.NotFound
	_ = protocmp.Transform
}

func (fx *Site) testUpdate(t *testing.T) {
	// Standard methods: Update
	// https://google.aip.dev/134

	parent := fx.nextParent(t, false)
	created00, err := fx.service.CreateSite(fx.ctx, &v1.CreateSiteRequest{
		Parent: parent,
		Site:   fx.Create(parent),
	})
	assert.NilError(t, err)

	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		msg := fx.Update(parent)
		msg.Name = ""
		_, err := fx.service.UpdateSite(fx.ctx, &v1.UpdateSiteRequest{
			Site: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		msg := fx.Update(parent)
		msg.Name = "invalid resource name"
		_, err := fx.service.UpdateSite(fx.ctx, &v1.UpdateSiteRequest{
			Site: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Field update_time should be updated when the resource is updated.
	t.Run("update time", func(t *testing.T) {
		fx.maybeSkip(t)
		initial, err := fx.service.CreateSite(fx.ctx, &v1.CreateSiteRequest{
			Parent: parent,
			Site:   fx.Create(parent),
		})
		assert.NilError(t, err)
		updated, err := fx.service.UpdateSite(fx.ctx, &v1.UpdateSiteRequest{
			Site: initial,
		})
		assert.NilError(t, err)
		assert.Check(t, updated.UpdateTime.AsTime().After(initial.UpdateTime.AsTime()))
	})

	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		msg := fx.Update(parent)
		msg.Name = created00.Name + "notfound"
		_, err := fx.service.UpdateSite(fx.ctx, &v1.UpdateSiteRequest{
			Site: msg,
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// The updated resource should be persisted and reachable with Get.
	t.Run("persisted", func(t *testing.T) {
		fx.maybeSkip(t)
		initial, err := fx.service.CreateSite(fx.ctx, &v1.CreateSiteRequest{
			Parent: parent,
			Site:   fx.Create(parent),
		})
		assert.NilError(t, err)
		updated, err := fx.service.UpdateSite(fx.ctx, &v1.UpdateSiteRequest{
			Site: initial,
		})
		persisted, err := fx.service.GetSite(fx.ctx, &v1.GetSiteRequest{
			Name: updated.Name,
		})
		assert.NilError(t, err)
		assert.DeepEqual(t, updated, persisted, protocmp.Transform())
	})

	// The method should fail with InvalidArgument if the update_mask is invalid.
	t.Run("invalid update mask", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.UpdateSite(fx.ctx, &v1.UpdateSiteRequest{
			Site: created00,
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
		t.Run(".display_name", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := proto.Clone(created00).(*v1.Site)
			container := msg
			if container == nil {
				t.Skip("not reachable")
			}
			fd := container.ProtoReflect().Descriptor().Fields().ByName("display_name")
			container.ProtoReflect().Clear(fd)
			_, err := fx.service.CreateSite(fx.ctx, &v1.CreateSiteRequest{
				Site: msg,
			})
			assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
		})
	})
	_ = codes.NotFound
	_ = protocmp.Transform
	_ = proto.Clone
}

func (fx *Site) testList(t *testing.T) {
	// Standard methods: List
	// https://google.aip.dev/132
	parent01 := fx.nextParent(t, false)
	parent02 := fx.nextParent(t, true)

	const n = 15

	parent01msgs := make([]*v1.Site, n)
	for i := 0; i < n; i++ {
		msg, err := fx.service.CreateSite(fx.ctx, &v1.CreateSiteRequest{
			Parent: parent01,
			Site:   fx.Create(parent01),
		})
		assert.NilError(t, err)
		parent01msgs[i] = msg
	}

	parent02msgs := make([]*v1.Site, n)
	for i := 0; i < n; i++ {
		msg, err := fx.service.CreateSite(fx.ctx, &v1.CreateSiteRequest{
			Parent: parent02,
			Site:   fx.Create(parent02),
		})
		assert.NilError(t, err)
		parent02msgs[i] = msg
	}

	// Method should fail with InvalidArgument is provided parent is not valid.
	t.Run("invalid parent", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.ListSites(fx.ctx, &v1.ListSitesRequest{
			Parent: "invalid parent",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page token is not valid.
	t.Run("invalid page token", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.ListSites(fx.ctx, &v1.ListSitesRequest{
			Parent:    parent01,
			PageToken: "invalid page token",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument is provided page size is negative.
	t.Run("negative page size", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.ListSites(fx.ctx, &v1.ListSitesRequest{
			Parent:   parent01,
			PageSize: -10,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// If parent is provided the method must only return resources
	// under that parent.
	t.Run("isolation", func(t *testing.T) {
		fx.maybeSkip(t)
		response, err := fx.service.ListSites(fx.ctx, &v1.ListSitesRequest{
			Parent:   parent02,
			PageSize: 999,
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			parent02msgs,
			response.Sites,
			cmpopts.SortSlices(func(a, b *v1.Site) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})

	t.Run("pagination", func(t *testing.T) {
		fx.maybeSkip(t)

		// If there are no more resources, next_page_token should be unset.
		t.Run("next page token", func(t *testing.T) {
			fx.maybeSkip(t)
			response, err := fx.service.ListSites(fx.ctx, &v1.ListSitesRequest{
				Parent:   parent02,
				PageSize: 999,
			})
			assert.NilError(t, err)
			assert.Equal(t, "", response.NextPageToken)
		})

		// Listing resource one by one should eventually return all resources created.
		t.Run("one by one", func(t *testing.T) {
			fx.maybeSkip(t)
			msgs := make([]*v1.Site, 0, n)
			var nextPageToken string
			for {
				response, err := fx.service.ListSites(fx.ctx, &v1.ListSitesRequest{
					Parent:    parent02,
					PageSize:  1,
					PageToken: nextPageToken,
				})
				assert.NilError(t, err)
				assert.Equal(t, 1, len(response.Sites))
				msgs = append(msgs, response.Sites...)
				nextPageToken = response.NextPageToken
				if nextPageToken == "" {
					break
				}
			}
			assert.DeepEqual(
				t,
				parent02msgs,
				msgs,
				cmpopts.SortSlices(func(a, b *v1.Site) bool {
					return a.Name < b.Name
				}),
				protocmp.Transform(),
			)
		})
	})

	// Method should not return deleted resources.
	t.Run("deleted", func(t *testing.T) {
		fx.maybeSkip(t)
		const nDelete = 5
		for i := 0; i < nDelete; i++ {
			_, err := fx.service.DeleteSite(fx.ctx, &v1.DeleteSiteRequest{
				Name: parent02msgs[i].Name,
			})
			assert.NilError(t, err)
		}
		response, err := fx.service.ListSites(fx.ctx, &v1.ListSitesRequest{
			Parent:   parent02,
			PageSize: 9999,
		})
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			parent02msgs[nDelete:],
			response.Sites,
			cmpopts.SortSlices(func(a, b *v1.Site) bool {
				return a.Name < b.Name
			}),
			protocmp.Transform(),
		)
	})
	_ = codes.NotFound
	_ = protocmp.Transform
	_ = cmpopts.SortSlices
}

func (fx *Site) nextParent(t *testing.T, pristine bool) string {
	if pristine {
		fx.currParent++
	}
	if fx.currParent >= len(fx.Parents) {
		t.Fatal("need at least", fx.currParent+1, "parents")
	}
	return fx.Parents[fx.currParent]
}

func (fx *Site) peekNextParent(t *testing.T) string {
	next := fx.currParent + 1
	if next >= len(fx.Parents) {
		t.Fatal("need at least", next+1, "parents")
	}
	return fx.Parents[next]
}

func (fx *Site) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}
