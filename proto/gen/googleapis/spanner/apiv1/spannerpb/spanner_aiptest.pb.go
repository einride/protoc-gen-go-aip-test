// Code generated by protoc-gen-go-aip-test. DO NOT EDIT.

package spannerpb

import (
	context "context"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protocmp "google.golang.org/protobuf/testing/protocmp"
	assert "gotest.tools/v3/assert"
	strings "strings"
	testing "testing"
)

func TestSpanner(
	t *testing.T,
	s SpannerTestsConfigSupplier,
) {
	{
		cfg := s.TestSession(t)
		fx := SpannerTestSuite{
			T:      t,
			Server: cfg.Server(),
		}
		fx.TestSession(cfg.Context(), *cfg)
	}
}

type SpannerTestsConfigSupplier interface {
	TestSession(t *testing.T) *SpannerSessionTestSuiteConfig
}
type SpannerTestSuite struct {
	T *testing.T
	// Server to test.
	Server SpannerServer
}

func (fx SpannerTestSuite) TestSession(ctx context.Context, options SpannerSessionTestSuiteConfig) {
	fx.T.Run("Session", func(t *testing.T) {
		options.ctx = ctx
		options.service = fx.Server
		options.test(t)
	})
}

type SpannerSessionTestSuiteConfig struct {
	ctx        context.Context
	service    SpannerServer
	currParent int

	Server func() SpannerServer
	// Context should return a new context that can be used for each test.
	Context func() context.Context
	// The parents to use when creating resources.
	// At least one parent needs to be set. Depending on methods available on the resource,
	// more may be required. If insufficient number of parents are
	// provided the test will fail.
	Parents []string
	// CreateResource should create a Session and return it.
	// If the field is not set, some tests will be skipped.
	//
	// This method is generated because service does not expose a Create
	// method (or it does not comply with AIP).
	CreateResource func(ctx context.Context, parent string) (*Session, error)
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *SpannerSessionTestSuiteConfig) test(t *testing.T) {
	t.Run("Get", fx.testGet)
	t.Run("Delete", fx.testDelete)
}

func (fx *SpannerSessionTestSuiteConfig) testGet(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetSession(fx.ctx, &GetSessionRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetSession(fx.ctx, &GetSessionRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be returned without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		msg, err := fx.service.GetSession(fx.ctx, &GetSessionRequest{
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
		_, err := fx.service.GetSession(fx.ctx, &GetSessionRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.GetSession(fx.ctx, &GetSessionRequest{
			Name: "projects/-/instances/-/databases/-/sessions/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *SpannerSessionTestSuiteConfig) testDelete(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.DeleteSession(fx.ctx, &DeleteSessionRequest{
			Name: "",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.DeleteSession(fx.ctx, &DeleteSessionRequest{
			Name: "invalid resource name",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Resource should be deleted without errors if it exists.
	t.Run("exists", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		_, err := fx.service.DeleteSession(fx.ctx, &DeleteSessionRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
	})

	// Method should fail with NotFound if the resource does not exist.
	t.Run("not found", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		_, err := fx.service.DeleteSession(fx.ctx, &DeleteSessionRequest{
			Name: created.Name + "notfound",
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with NotFound if the resource was already deleted. This also applies to soft-deletion.
	t.Run("already deleted", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		deleted, err := fx.service.DeleteSession(fx.ctx, &DeleteSessionRequest{
			Name: created.Name,
		})
		assert.NilError(t, err)
		_ = deleted
		_, err = fx.service.DeleteSession(fx.ctx, &DeleteSessionRequest{
			Name: created.Name,
		})
		assert.Equal(t, codes.NotFound, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if the provided name only contains wildcards ('-')
	t.Run("only wildcards", func(t *testing.T) {
		fx.maybeSkip(t)
		_, err := fx.service.DeleteSession(fx.ctx, &DeleteSessionRequest{
			Name: "projects/-/instances/-/databases/-/sessions/-",
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

}

func (fx *SpannerSessionTestSuiteConfig) nextParent(t *testing.T, pristine bool) string {
	if pristine {
		fx.currParent++
	}
	if fx.currParent >= len(fx.Parents) {
		t.Fatal("need at least", fx.currParent+1, "parents")
	}
	return fx.Parents[fx.currParent]
}

func (fx *SpannerSessionTestSuiteConfig) peekNextParent(t *testing.T) string {
	next := fx.currParent + 1
	if next >= len(fx.Parents) {
		t.Fatal("need at least", next+1, "parents")
	}
	return fx.Parents[next]
}

func (fx *SpannerSessionTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}

func (fx *SpannerSessionTestSuiteConfig) create(t *testing.T, parent string) *Session {
	t.Helper()
	if fx.CreateResource == nil {
		t.Skip("Test skipped because CreateResource not specified on SpannerSessionTestSuiteConfig")
	}
	created, err := fx.CreateResource(fx.ctx, parent)
	assert.NilError(t, err)
	return created
}
