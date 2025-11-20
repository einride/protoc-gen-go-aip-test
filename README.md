# protoc-gen-go-aip-test

Generate test suites for protobuf services implementing
[standard AIP methods](https://google.aip.dev/121#methods).

The generated test suites are based on guidance for standard methods, and
experience from implementing these methods in practice. See [Suites](#suites)
for a list of the generated tests.

**Experimental**: This plugin is experimental, and breaking changes with regard
to the generated tests suites should be expected.

## Usage

### Step 1: Declare a service with AIP standard methods

```protobuf
service FreightService {
  // Get a shipper.
  // See: https://google.aip.dev/131 (Standard methods: Get).
  rpc GetShipper(GetShipperRequest) returns (Shipper) {
    option (google.api.http) = {
      get: "/v1/{name=shippers/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // ...
}
```

### Step 2: Install the generator

Either install using `go install`:

```bash
go install github.com/einride/protoc-gen-go-aip-test@latest
```

Or download a prebuilt binary from
[releases](https://github.com/einride/protoc-gen-go-aip-test/releases) and put
it in your PATH.

The generator can also be built from source using Go.

### Step 3: Generate test suites

Include the plugin in `protoc` invocation

```bash
protoc
  --go-aip-test_out=[OUTPUT DIR] \
  --go-aip-test_opt=module=[OUTPUT MODULE] \
  [.proto files ...]
```

This can also be done via a
[buf generate](https://docs.buf.build/generate/usage) template. See
[buf.gen.yaml](./proto/buf.gen.yaml) for an example.

### Step 4: Run tests

There are two alternative ways of bootstrapping the tests.

#### Alternative 1:

Instantiate the generated test suites and call the methods you want to test.

```go
package example

func Test_FreightService(t *testing.T) {
	t.Skip("this is just an example, the service is not implemented.")
	// setup server before test
	server := examplefreightv1.UnimplementedFreightServiceServer{}
	// setup test suite
	suite := examplefreightv1.FreightServiceTestSuite{
		T:      t,
		Server: server,
	}

	// run tests for each resource in the service
	ctx := context.Background()
	suite.TestShipper(ctx, examplefreightv1.ShipperTestSuiteConfig{
		// Create should return a resource which is valid to create, i.e.
		// all required fields set.
		Create: func() *examplefreightv1.Shipper {
			return &examplefreightv1.Shipper{
				DisplayName:    "Example shipper",
				BillingAccount: "billingAccounts/12345",
			}
		},
		// Update should return a resource which is valid to update, i.e.
		// all required fields set.
		Update: func() *examplefreightv1.Shipper {
			return &examplefreightv1.Shipper{
				DisplayName:    "Updated example shipper",
				BillingAccount: "billingAccounts/54321",
			}
		},
	})
}
```

#### Alternative 2:

Implement the generated configure provider interface
(`FreightServiceTestSuiteConfigProvider`) and pass the implementation to
`TestServices` to start the tests.

A benefit of using `TestServices` (over alternative 1) is that as new services
or resources are added to the API the test code won't compile until the required
inputs are also added (or explicitly ignored). This makes it harder to forget to
add the test implementations for new services/resources.

```go
package example

import "testing"

func Test_FreightService(t *testing.T) {
	// Even though no implementation exists, the tests will pass but be skipped.
	examplefreightv1.TestServices(t, &aipTests{})
}

type aipTests struct{}

var _ examplefreightv1.FreightServiceTestSuiteConfigProvider = &aipTests{}

func (a aipTests) FreightServiceShipper(_ *testing.T) *examplefreightv1.FreightServiceShipperTestSuiteConfig {
	// Returns nil to indicate that it's not ready to be tested.
	return nil
}

func (a aipTests) FreightServiceSite(_ *testing.T) *examplefreightv1.FreightServiceSiteTestSuiteConfig {
	// Returns nil to indicate that it's not ready to be tested.
	return nil
}
```

### Skipping tests

There may be multiple reasons for an API to deviate from the guidance for
standard methods (for examples see [AIP-200](https://google.aip.dev/200)). This
plugin supports skipping individual or groups of tests using the `Skip` field
generated for each test suite config.

Each test are compared, using `strings.Contains`, against a list of skipped test
patterns. The full name of each test will follow the format
`[resource]/[method type]/[test_name]`.

Sample skips:

- `"Get/invalid_name"` skips the "invalid name" test for Get standard method.
- `"Get"` skips all tests for a Get standard method.

## Suites

<!-- BEGIN suites -->

### Create

| Name                     | Description                                                                                                                                              | Generated only if all are true:                                                                                                  |
| ------------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| missing parent           | Method should fail with InvalidArgument if no parent is provided.                                                                                        | <ul><li>has Create method</li><li>resource has a parent</li></ul>                                                                |
| invalid parent           | Method should fail with InvalidArgument if provided parent is invalid.                                                                                   | <ul><li>has Create method</li><li>resource has a parent</li></ul>                                                                |
| create time              | Field create_time should be populated when the resource is created.                                                                                      | <ul><li>Create method does not return long-running operation</li><li>has Create method</li><li>has field 'create_time'</li></ul> |
| persisted                | The created resource should be persisted and reachable with Get.                                                                                         | <ul><li>Create method does not return long-running operation</li><li>has Create method</li><li>has Get method</li></ul>          |
| user settable id         | If method support user settable IDs, when set the resource should be returned with the provided ID.                                                      | <ul><li>Create method does not return long-running operation</li><li>has Create method</li><li>has user settable ID</li></ul>    |
| invalid user settable id | Method should fail with InvalidArgument if the user settable id doesn't conform to RFC-1034, see [doc](https://google.aip.dev/122#resource-id-segments). | <ul><li>Create method does not return long-running operation</li><li>has Create method</li><li>has user settable ID</li></ul>    |
| already exists           | If method support user settable IDs and the same ID is reused the method should return AlreadyExists.                                                    | <ul><li>Create method does not return long-running operation</li><li>has Create method</li><li>has user settable ID</li></ul>    |
| required fields          | The method should fail with InvalidArgument if the resource has any required fields and they are not provided.                                           | <ul><li>has Create method</li><li>resource has any required fields</li></ul>                                                     |
| resource references      | The method should fail with InvalidArgument if the resource has any resource references and they are invalid.                                            | <ul><li>has Create method</li><li>resource has any mutable resource references</li></ul>                                         |
| etag populated           | Field etag should be populated when the resource is created.                                                                                             | <ul><li>Create method does not return long-running operation</li><li>has Create method</li><li>has field 'etag'</li></ul>        |

### Get

| Name           | Description                                                                                | Generated only if all are true:                                                            |
| -------------- | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| missing name   | Method should fail with InvalidArgument if no name is provided.                            | <ul><li>has Get method</li></ul>                                                           |
| invalid name   | Method should fail with InvalidArgument if the provided name is not valid.                 | <ul><li>has Get method</li></ul>                                                           |
| exists         | Resource should be returned without errors if it exists.                                   | <ul><li>has Get method</li></ul>                                                           |
| not found      | Method should fail with NotFound if the resource does not exist.                           | <ul><li>has Get method</li><li>resource is not a singleton</li></ul>                       |
| only wildcards | Method should fail with InvalidArgument if the provided name only contains wildcards ('-') | <ul><li>has Get method</li><li>resource name pattern contains variables</li></ul>          |
| soft-deleted   | A soft-deleted resource should be returned without errors.                                 | <ul><li>has Delete method</li><li>has Get method</li><li>has field 'delete_time'</li></ul> |

### BatchGet

| Name            | Description                                                                                                                                 | Generated only if all are true:                                                                                      |
| --------------- | ------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| invalid parent  | Method should fail with InvalidArgument if provided parent is invalid.                                                                      | <ul><li>has BatchGet method</li><li>is not alternative batch request message</li><li>resource has a parent</li></ul> |
| names missing   | Method should fail with InvalidArgument if no names are provided.                                                                           | <ul><li>has BatchGet method</li><li>is not alternative batch request message</li></ul>                               |
| invalid names   | Method should fail with InvalidArgument if a provided name is not valid.                                                                    | <ul><li>has BatchGet method</li><li>is not alternative batch request message</li></ul>                               |
| wildcard name   | Method should fail with InvalidArgument if a provided name only contains wildcards (-)                                                      | <ul><li>has BatchGet method</li><li>is not alternative batch request message</li></ul>                               |
| all exists      | Resources should be returned without errors if they exist.                                                                                  | <ul><li>has BatchGet method</li><li>is not alternative batch request message</li></ul>                               |
| atomic          | The method must be atomic; it must fail for all resources or succeed for all resources (no partial success).                                | <ul><li>has BatchGet method</li><li>is not alternative batch request message</li></ul>                               |
| parent mismatch | If a caller sets the "parent", and the parent collection in the name of any resource being retrieved does not match, the request must fail. | <ul><li>has BatchGet method</li><li>is not alternative batch request message</li><li>resource has a parent</li></ul> |
| ordered         | The order of resources in the response must be the same as the names in the request.                                                        | <ul><li>has BatchGet method</li><li>is not alternative batch request message</li></ul>                               |
| duplicate names | If a caller provides duplicate names, the service should return duplicate resources.                                                        | <ul><li>has BatchGet method</li><li>is not alternative batch request message</li></ul>                               |

### Update

| Name                 | Description                                                                                                 | Generated only if all are true:                                                                                                                                                                                         |
| -------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| missing name         | Method should fail with InvalidArgument if no name is provided.                                             | <ul><li>has Update method</li></ul>                                                                                                                                                                                     |
| invalid name         | Method should fail with InvalidArgument if provided name is not valid.                                      | <ul><li>has Update method</li></ul>                                                                                                                                                                                     |
| update time          | Field update_time should be updated when the resource is updated.                                           | <ul><li>Create method does not return long-running operation</li><li>Update method does not return long-running operation</li><li>has Create method</li><li>has Update method</li><li>has field 'update_time'</li></ul> |
| persisted            | The updated resource should be persisted and reachable with Get.                                            | <ul><li>Update method does not return long-running operation</li><li>has Get method</li><li>has Update method</li></ul>                                                                                                 |
| preserve create_time | The field create_time should be preserved when a '\*'-update mask is used.                                  | <ul><li>Update method does not return long-running operation</li><li>has Update method</li><li>has field 'create_time'</li><li>resource has any required fields</li></ul>                                               |
| etag mismatch        | Method should fail with Aborted if the supplied etag doesnt match the current etag value.                   | <ul><li>has Update method</li><li>has field 'etag'</li></ul>                                                                                                                                                            |
| etag updated         | Field etag should have a new value when the resource is successfully updated.                               | <ul><li>has Update method</li><li>has field 'etag'</li></ul>                                                                                                                                                            |
| not found            | Method should fail with NotFound if the resource does not exist.                                            | <ul><li>has Update method</li><li>resource is not a singleton</li></ul>                                                                                                                                                 |
| invalid update mask  | The method should fail with InvalidArgument if the update_mask is invalid.                                  | <ul><li>Update method has update_mask</li><li>has Update method</li></ul>                                                                                                                                               |
| required fields      | Method should fail with InvalidArgument if any required field is missing when called with '\*' update_mask. | <ul><li>has Update method</li><li>resource has any required fields</li></ul>                                                                                                                                            |

### List

| Name               | Description                                                                    | Generated only if all are true:                                                           |
| ------------------ | ------------------------------------------------------------------------------ | ----------------------------------------------------------------------------------------- |
| invalid parent     | Method should fail with InvalidArgument if provided parent is invalid.         | <ul><li>has List method</li><li>resource has a parent</li></ul>                           |
| invalid page token | Method should fail with InvalidArgument is provided page token is not valid.   | <ul><li>has List method</li></ul>                                                         |
| negative page size | Method should fail with InvalidArgument is provided page size is negative.     | <ul><li>has List method</li></ul>                                                         |
| isolation          | If parent is provided the method must only return resources under that parent. | <ul><li>has List method</li><li>resource has a parent</li></ul>                           |
| last page          | If there are no more resources, next_page_token should not be set.             | <ul><li>has List method</li><li>resource has a parent</li></ul>                           |
| more pages         | If there are more resources, next_page_token should be set.                    | <ul><li>has List method</li><li>resource has a parent</li></ul>                           |
| one by one         | Listing resource one by one should eventually return all resources.            | <ul><li>has List method</li><li>resource has a parent</li></ul>                           |
| deleted            | Method should not return deleted resources.                                    | <ul><li>has Delete method</li><li>has List method</li><li>resource has a parent</li></ul> |
| page size zero     | Listing resource with page size zero should eventually return all resources.   | <ul><li>has List method</li><li>resource has a parent</li></ul>                           |

### Search

| Name               | Description                                                                    | Generated only if all are true:                                                             |
| ------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------- |
| invalid parent     | Method should fail with InvalidArgument if provided parent is invalid.         | <ul><li>has Search method</li><li>resource has a parent</li></ul>                           |
| invalid page token | Method should fail with InvalidArgument is provided page token is not valid.   | <ul><li>has Search method</li></ul>                                                         |
| negative page size | Method should fail with InvalidArgument is provided page size is negative.     | <ul><li>has Search method</li></ul>                                                         |
| isolation          | If parent is provided the method must only return resources under that parent. | <ul><li>has Search method</li><li>resource has a parent</li></ul>                           |
| last page          | If there are no more resources, next_page_token should not be set.             | <ul><li>has Search method</li><li>resource has a parent</li></ul>                           |
| more pages         | If there are more resources, next_page_token should be set.                    | <ul><li>has Search method</li><li>resource has a parent</li></ul>                           |
| one by one         | Searching resource one by one should eventually return all resources.          | <ul><li>has Search method</li><li>resource has a parent</li></ul>                           |
| deleted            | Method should not return deleted resources.                                    | <ul><li>has Delete method</li><li>has Search method</li><li>resource has a parent</li></ul> |

### Delete

| Name                     | Description                                                                                               | Generated only if all are true:                                                                                                                                                              |
| ------------------------ | --------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| missing name             | Method should fail with InvalidArgument if no name is provided.                                           | <ul><li>has Delete method</li></ul>                                                                                                                                                          |
| invalid name             | Method should fail with InvalidArgument if the provided name is not valid.                                | <ul><li>has Delete method</li></ul>                                                                                                                                                          |
| exists                   | Resource should be deleted without errors if it exists.                                                   | <ul><li>has Delete method</li></ul>                                                                                                                                                          |
| not found                | Method should fail with NotFound if the resource does not exist.                                          | <ul><li>has Delete method</li></ul>                                                                                                                                                          |
| already deleted          | Method should fail with NotFound if the resource was already deleted. This also applies to soft-deletion. | <ul><li>has Delete method</li></ul>                                                                                                                                                          |
| only wildcards           | Method should fail with InvalidArgument if the provided name only contains wildcards ('-')                | <ul><li>has Delete method</li></ul>                                                                                                                                                          |
| etag mismatch            | Method should fail with Aborted if the supplied etag doesnt match the current etag value.                 | <ul><li>has Delete method</li><li>has field 'etag'</li><li>request has etag field</li></ul>                                                                                                  |
| soft-deleted delete_time | A soft-deleted resource should have delete_time assigned.                                                 | <ul><li>Delete method does not return google.protobuf.Empty</li><li>Delete method does not return long-running operation</li><li>has Delete method</li><li>has field 'delete_time'</li></ul> |

<!-- END suites -->
