// Code generated by protoc-gen-go-aip-test. DO NOT EDIT.

package pubsubpb

import (
	context "context"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	proto "google.golang.org/protobuf/proto"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	assert "gotest.tools/v3/assert"
	strings "strings"
	testing "testing"
)

// PublisherTestSuiteConfigProvider is the interface to implement to decide which resources
// that should be tested and how it's configured.
type PublisherTestSuiteConfigProvider interface {
	// PublisherTopic should return a config, or nil, which means that the tests will be skipped.
	PublisherTopic(t *testing.T) *PublisherTopicTestSuiteConfig
}

// testPublisher is the main entrypoint for starting the AIP tests.
func testPublisher(t *testing.T, s PublisherTestSuiteConfigProvider) {
	testPublisherTopic(t, s)
}

func testPublisherTopic(t *testing.T, s PublisherTestSuiteConfigProvider) {
	t.Run("Topic", func(t *testing.T) {
		config := s.PublisherTopic(t)
		if config == nil {
			t.Skip("Method PublisherTopic not implemented")
		}
		if config.Service == nil {
			t.Skip("Method PublisherTopic.Service() not implemented")
		}
		if config.Context == nil {
			config.Context = func() context.Context { return context.Background() }
		}
		config.test(t)
	})
}

type PublisherTestSuite struct {
	T *testing.T
	// Server to test.
	Server PublisherServer
}

func (fx PublisherTestSuite) TestTopic(ctx context.Context, options PublisherTopicTestSuiteConfig) {
	fx.T.Run("Topic", func(t *testing.T) {
		options.Context = func() context.Context { return ctx }
		options.Service = func() PublisherServer { return fx.Server }
		options.test(t)
	})
}

type PublisherTopicTestSuiteConfig struct {
	currParent int

	// Service should return the service that should be tested.
	// The service will be used for several tests.
	Service func() PublisherServer
	// Context should return a new context.
	// The context will be used for several tests.
	Context func() context.Context
	// The parents to use when creating resources.
	// At least one parent needs to be set. Depending on methods available on the resource,
	// more may be required. If insufficient number of parents are
	// provided the test will fail.
	Parents []string
	// CreateResource should create a Topic and return it.
	// If the field is not set, some tests will be skipped.
	//
	// This method is generated because service does not expose a Create
	// method (or it does not comply with AIP).
	CreateResource func(ctx context.Context, parent string) (*Topic, error)
	// Update should return a resource which is valid to update, i.e.
	// all required fields set.
	Update func(parent string) *Topic
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *PublisherTopicTestSuiteConfig) test(t *testing.T) {
	t.Run("Update", fx.testUpdate)
}

func (fx *PublisherTopicTestSuiteConfig) testUpdate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = ""
		_, err := fx.Service().UpdateTopic(fx.Context(), &UpdateTopicRequest{
			Topic: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = "invalid resource name"
		_, err := fx.Service().UpdateTopic(fx.Context(), &UpdateTopicRequest{
			Topic: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	{
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		// Method should fail with NotFound if the resource does not exist.
		t.Run("not found", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := fx.Update(parent)
			msg.Name = created.Name + "notfound"
			_, err := fx.Service().UpdateTopic(fx.Context(), &UpdateTopicRequest{
				Topic: msg,
			})
			assert.Equal(t, codes.NotFound, status.Code(err), err)
		})

		// The method should fail with InvalidArgument if the update_mask is invalid.
		t.Run("invalid update mask", func(t *testing.T) {
			fx.maybeSkip(t)
			_, err := fx.Service().UpdateTopic(fx.Context(), &UpdateTopicRequest{
				Topic: created,
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
			t.Run(".name", func(t *testing.T) {
				fx.maybeSkip(t)
				msg := proto.Clone(created).(*Topic)
				container := msg
				if container == nil {
					t.Skip("not reachable")
				}
				fd := container.ProtoReflect().Descriptor().Fields().ByName("name")
				container.ProtoReflect().Clear(fd)
				_, err := fx.Service().UpdateTopic(fx.Context(), &UpdateTopicRequest{
					Topic: msg,
					UpdateMask: &fieldmaskpb.FieldMask{
						Paths: []string{
							"*",
						},
					},
				})
				assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
			})
			t.Run(".schema_settings.schema", func(t *testing.T) {
				fx.maybeSkip(t)
				msg := proto.Clone(created).(*Topic)
				container := msg.GetSchemaSettings()
				if container == nil {
					t.Skip("not reachable")
				}
				fd := container.ProtoReflect().Descriptor().Fields().ByName("schema")
				container.ProtoReflect().Clear(fd)
				_, err := fx.Service().UpdateTopic(fx.Context(), &UpdateTopicRequest{
					Topic: msg,
					UpdateMask: &fieldmaskpb.FieldMask{
						Paths: []string{
							"*",
						},
					},
				})
				assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
			})
			t.Run(".ingestion_data_source_settings.aws_kinesis.stream_arn", func(t *testing.T) {
				fx.maybeSkip(t)
				msg := proto.Clone(created).(*Topic)
				container := msg.GetIngestionDataSourceSettings().GetAwsKinesis()
				if container == nil {
					t.Skip("not reachable")
				}
				fd := container.ProtoReflect().Descriptor().Fields().ByName("stream_arn")
				container.ProtoReflect().Clear(fd)
				_, err := fx.Service().UpdateTopic(fx.Context(), &UpdateTopicRequest{
					Topic: msg,
					UpdateMask: &fieldmaskpb.FieldMask{
						Paths: []string{
							"*",
						},
					},
				})
				assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
			})
			t.Run(".ingestion_data_source_settings.aws_kinesis.consumer_arn", func(t *testing.T) {
				fx.maybeSkip(t)
				msg := proto.Clone(created).(*Topic)
				container := msg.GetIngestionDataSourceSettings().GetAwsKinesis()
				if container == nil {
					t.Skip("not reachable")
				}
				fd := container.ProtoReflect().Descriptor().Fields().ByName("consumer_arn")
				container.ProtoReflect().Clear(fd)
				_, err := fx.Service().UpdateTopic(fx.Context(), &UpdateTopicRequest{
					Topic: msg,
					UpdateMask: &fieldmaskpb.FieldMask{
						Paths: []string{
							"*",
						},
					},
				})
				assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
			})
			t.Run(".ingestion_data_source_settings.aws_kinesis.aws_role_arn", func(t *testing.T) {
				fx.maybeSkip(t)
				msg := proto.Clone(created).(*Topic)
				container := msg.GetIngestionDataSourceSettings().GetAwsKinesis()
				if container == nil {
					t.Skip("not reachable")
				}
				fd := container.ProtoReflect().Descriptor().Fields().ByName("aws_role_arn")
				container.ProtoReflect().Clear(fd)
				_, err := fx.Service().UpdateTopic(fx.Context(), &UpdateTopicRequest{
					Topic: msg,
					UpdateMask: &fieldmaskpb.FieldMask{
						Paths: []string{
							"*",
						},
					},
				})
				assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
			})
			t.Run(".ingestion_data_source_settings.aws_kinesis.gcp_service_account", func(t *testing.T) {
				fx.maybeSkip(t)
				msg := proto.Clone(created).(*Topic)
				container := msg.GetIngestionDataSourceSettings().GetAwsKinesis()
				if container == nil {
					t.Skip("not reachable")
				}
				fd := container.ProtoReflect().Descriptor().Fields().ByName("gcp_service_account")
				container.ProtoReflect().Clear(fd)
				_, err := fx.Service().UpdateTopic(fx.Context(), &UpdateTopicRequest{
					Topic: msg,
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

func (fx *PublisherTopicTestSuiteConfig) nextParent(t *testing.T, pristine bool) string {
	if pristine {
		fx.currParent++
	}
	if fx.currParent >= len(fx.Parents) {
		t.Fatal("need at least", fx.currParent+1, "parents")
	}
	return fx.Parents[fx.currParent]
}

func (fx *PublisherTopicTestSuiteConfig) peekNextParent(t *testing.T) string {
	next := fx.currParent + 1
	if next >= len(fx.Parents) {
		t.Fatal("need at least", next+1, "parents")
	}
	return fx.Parents[next]
}

func (fx *PublisherTopicTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}

func (fx *PublisherTopicTestSuiteConfig) create(t *testing.T, parent string) *Topic {
	t.Helper()
	if fx.CreateResource == nil {
		t.Skip("Test skipped because CreateResource not specified on PublisherTopicTestSuiteConfig")
	}
	created, err := fx.CreateResource(fx.Context(), parent)
	assert.NilError(t, err)
	return created
}

// SubscriberTestSuiteConfigProvider is the interface to implement to decide which resources
// that should be tested and how it's configured.
type SubscriberTestSuiteConfigProvider interface {
	// SubscriberSnapshot should return a config, or nil, which means that the tests will be skipped.
	SubscriberSnapshot(t *testing.T) *SubscriberSnapshotTestSuiteConfig
	// SubscriberSubscription should return a config, or nil, which means that the tests will be skipped.
	SubscriberSubscription(t *testing.T) *SubscriberSubscriptionTestSuiteConfig
}

// testSubscriber is the main entrypoint for starting the AIP tests.
func testSubscriber(t *testing.T, s SubscriberTestSuiteConfigProvider) {
	testSubscriberSnapshot(t, s)
	testSubscriberSubscription(t, s)
}

func testSubscriberSnapshot(t *testing.T, s SubscriberTestSuiteConfigProvider) {
	t.Run("Snapshot", func(t *testing.T) {
		config := s.SubscriberSnapshot(t)
		if config == nil {
			t.Skip("Method SubscriberSnapshot not implemented")
		}
		if config.Service == nil {
			t.Skip("Method SubscriberSnapshot.Service() not implemented")
		}
		if config.Context == nil {
			config.Context = func() context.Context { return context.Background() }
		}
		config.test(t)
	})
}

func testSubscriberSubscription(t *testing.T, s SubscriberTestSuiteConfigProvider) {
	t.Run("Subscription", func(t *testing.T) {
		config := s.SubscriberSubscription(t)
		if config == nil {
			t.Skip("Method SubscriberSubscription not implemented")
		}
		if config.Service == nil {
			t.Skip("Method SubscriberSubscription.Service() not implemented")
		}
		if config.Context == nil {
			config.Context = func() context.Context { return context.Background() }
		}
		config.test(t)
	})
}

type SubscriberTestSuite struct {
	T *testing.T
	// Server to test.
	Server SubscriberServer
}

func (fx SubscriberTestSuite) TestSnapshot(ctx context.Context, options SubscriberSnapshotTestSuiteConfig) {
	fx.T.Run("Snapshot", func(t *testing.T) {
		options.Context = func() context.Context { return ctx }
		options.Service = func() SubscriberServer { return fx.Server }
		options.test(t)
	})
}

func (fx SubscriberTestSuite) TestSubscription(ctx context.Context, options SubscriberSubscriptionTestSuiteConfig) {
	fx.T.Run("Subscription", func(t *testing.T) {
		options.Context = func() context.Context { return ctx }
		options.Service = func() SubscriberServer { return fx.Server }
		options.test(t)
	})
}

type SubscriberSnapshotTestSuiteConfig struct {
	currParent int

	// Service should return the service that should be tested.
	// The service will be used for several tests.
	Service func() SubscriberServer
	// Context should return a new context.
	// The context will be used for several tests.
	Context func() context.Context
	// The parents to use when creating resources.
	// At least one parent needs to be set. Depending on methods available on the resource,
	// more may be required. If insufficient number of parents are
	// provided the test will fail.
	Parents []string
	// CreateResource should create a Snapshot and return it.
	// If the field is not set, some tests will be skipped.
	//
	// This method is generated because service does not expose a Create
	// method (or it does not comply with AIP).
	CreateResource func(ctx context.Context, parent string) (*Snapshot, error)
	// Update should return a resource which is valid to update, i.e.
	// all required fields set.
	Update func(parent string) *Snapshot
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *SubscriberSnapshotTestSuiteConfig) test(t *testing.T) {
	t.Run("Update", fx.testUpdate)
}

func (fx *SubscriberSnapshotTestSuiteConfig) testUpdate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = ""
		_, err := fx.Service().UpdateSnapshot(fx.Context(), &UpdateSnapshotRequest{
			Snapshot: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = "invalid resource name"
		_, err := fx.Service().UpdateSnapshot(fx.Context(), &UpdateSnapshotRequest{
			Snapshot: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	{
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		// Method should fail with NotFound if the resource does not exist.
		t.Run("not found", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := fx.Update(parent)
			msg.Name = created.Name + "notfound"
			_, err := fx.Service().UpdateSnapshot(fx.Context(), &UpdateSnapshotRequest{
				Snapshot: msg,
			})
			assert.Equal(t, codes.NotFound, status.Code(err), err)
		})

		// The method should fail with InvalidArgument if the update_mask is invalid.
		t.Run("invalid update mask", func(t *testing.T) {
			fx.maybeSkip(t)
			_, err := fx.Service().UpdateSnapshot(fx.Context(), &UpdateSnapshotRequest{
				Snapshot: created,
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

func (fx *SubscriberSnapshotTestSuiteConfig) nextParent(t *testing.T, pristine bool) string {
	if pristine {
		fx.currParent++
	}
	if fx.currParent >= len(fx.Parents) {
		t.Fatal("need at least", fx.currParent+1, "parents")
	}
	return fx.Parents[fx.currParent]
}

func (fx *SubscriberSnapshotTestSuiteConfig) peekNextParent(t *testing.T) string {
	next := fx.currParent + 1
	if next >= len(fx.Parents) {
		t.Fatal("need at least", next+1, "parents")
	}
	return fx.Parents[next]
}

func (fx *SubscriberSnapshotTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}

func (fx *SubscriberSnapshotTestSuiteConfig) create(t *testing.T, parent string) *Snapshot {
	t.Helper()
	if fx.CreateResource == nil {
		t.Skip("Test skipped because CreateResource not specified on SubscriberSnapshotTestSuiteConfig")
	}
	created, err := fx.CreateResource(fx.Context(), parent)
	assert.NilError(t, err)
	return created
}

type SubscriberSubscriptionTestSuiteConfig struct {
	currParent int

	// Service should return the service that should be tested.
	// The service will be used for several tests.
	Service func() SubscriberServer
	// Context should return a new context.
	// The context will be used for several tests.
	Context func() context.Context
	// The parents to use when creating resources.
	// At least one parent needs to be set. Depending on methods available on the resource,
	// more may be required. If insufficient number of parents are
	// provided the test will fail.
	Parents []string
	// CreateResource should create a Subscription and return it.
	// If the field is not set, some tests will be skipped.
	//
	// This method is generated because service does not expose a Create
	// method (or it does not comply with AIP).
	CreateResource func(ctx context.Context, parent string) (*Subscription, error)
	// Update should return a resource which is valid to update, i.e.
	// all required fields set.
	Update func(parent string) *Subscription
	// Patterns of tests to skip.
	// For example if a service has a Get method:
	// Skip: ["Get"] will skip all tests for Get.
	// Skip: ["Get/persisted"] will only skip the subtest called "persisted" of Get.
	Skip []string
}

func (fx *SubscriberSubscriptionTestSuiteConfig) test(t *testing.T) {
	t.Run("Update", fx.testUpdate)
}

func (fx *SubscriberSubscriptionTestSuiteConfig) testUpdate(t *testing.T) {
	fx.maybeSkip(t)
	// Method should fail with InvalidArgument if no name is provided.
	t.Run("missing name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = ""
		_, err := fx.Service().UpdateSubscription(fx.Context(), &UpdateSubscriptionRequest{
			Subscription: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	// Method should fail with InvalidArgument if provided name is not valid.
	t.Run("invalid name", func(t *testing.T) {
		fx.maybeSkip(t)
		parent := fx.nextParent(t, false)
		msg := fx.Update(parent)
		msg.Name = "invalid resource name"
		_, err := fx.Service().UpdateSubscription(fx.Context(), &UpdateSubscriptionRequest{
			Subscription: msg,
		})
		assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
	})

	{
		parent := fx.nextParent(t, false)
		created := fx.create(t, parent)
		// Method should fail with NotFound if the resource does not exist.
		t.Run("not found", func(t *testing.T) {
			fx.maybeSkip(t)
			msg := fx.Update(parent)
			msg.Name = created.Name + "notfound"
			_, err := fx.Service().UpdateSubscription(fx.Context(), &UpdateSubscriptionRequest{
				Subscription: msg,
			})
			assert.Equal(t, codes.NotFound, status.Code(err), err)
		})

		// The method should fail with InvalidArgument if the update_mask is invalid.
		t.Run("invalid update mask", func(t *testing.T) {
			fx.maybeSkip(t)
			_, err := fx.Service().UpdateSubscription(fx.Context(), &UpdateSubscriptionRequest{
				Subscription: created,
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
			t.Run(".name", func(t *testing.T) {
				fx.maybeSkip(t)
				msg := proto.Clone(created).(*Subscription)
				container := msg
				if container == nil {
					t.Skip("not reachable")
				}
				fd := container.ProtoReflect().Descriptor().Fields().ByName("name")
				container.ProtoReflect().Clear(fd)
				_, err := fx.Service().UpdateSubscription(fx.Context(), &UpdateSubscriptionRequest{
					Subscription: msg,
					UpdateMask: &fieldmaskpb.FieldMask{
						Paths: []string{
							"*",
						},
					},
				})
				assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
			})
			t.Run(".topic", func(t *testing.T) {
				fx.maybeSkip(t)
				msg := proto.Clone(created).(*Subscription)
				container := msg
				if container == nil {
					t.Skip("not reachable")
				}
				fd := container.ProtoReflect().Descriptor().Fields().ByName("topic")
				container.ProtoReflect().Clear(fd)
				_, err := fx.Service().UpdateSubscription(fx.Context(), &UpdateSubscriptionRequest{
					Subscription: msg,
					UpdateMask: &fieldmaskpb.FieldMask{
						Paths: []string{
							"*",
						},
					},
				})
				assert.Equal(t, codes.InvalidArgument, status.Code(err), err)
			})
			t.Run(".cloud_storage_config.bucket", func(t *testing.T) {
				fx.maybeSkip(t)
				msg := proto.Clone(created).(*Subscription)
				container := msg.GetCloudStorageConfig()
				if container == nil {
					t.Skip("not reachable")
				}
				fd := container.ProtoReflect().Descriptor().Fields().ByName("bucket")
				container.ProtoReflect().Clear(fd)
				_, err := fx.Service().UpdateSubscription(fx.Context(), &UpdateSubscriptionRequest{
					Subscription: msg,
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

func (fx *SubscriberSubscriptionTestSuiteConfig) nextParent(t *testing.T, pristine bool) string {
	if pristine {
		fx.currParent++
	}
	if fx.currParent >= len(fx.Parents) {
		t.Fatal("need at least", fx.currParent+1, "parents")
	}
	return fx.Parents[fx.currParent]
}

func (fx *SubscriberSubscriptionTestSuiteConfig) peekNextParent(t *testing.T) string {
	next := fx.currParent + 1
	if next >= len(fx.Parents) {
		t.Fatal("need at least", next+1, "parents")
	}
	return fx.Parents[next]
}

func (fx *SubscriberSubscriptionTestSuiteConfig) maybeSkip(t *testing.T) {
	for _, skip := range fx.Skip {
		if strings.Contains(t.Name(), skip) {
			t.Skip("skipped because of .Skip")
		}
	}
}

func (fx *SubscriberSubscriptionTestSuiteConfig) create(t *testing.T, parent string) *Subscription {
	t.Helper()
	if fx.CreateResource == nil {
		t.Skip("Test skipped because CreateResource not specified on SubscriberSubscriptionTestSuiteConfig")
	}
	created, err := fx.CreateResource(fx.Context(), parent)
	assert.NilError(t, err)
	return created
}
