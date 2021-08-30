## Suites

<!-- SUITES_SNIPPET -->

### Create

| Name                | Description                                                                                                    |
| ------------------- | -------------------------------------------------------------------------------------------------------------- |
| missing parent      | Method should fail with InvalidArgument if no parent is provided.                                              |
| invalid parent      | Method should fail with InvalidArgument if provided parent is invalid.                                         |
| create time         | Field create_time should be populated when the resource is created.                                            |
| persisted           | The created resource should be persisted and reachable with Get.                                               |
| user settable id    | If method support user settable IDs, when set the resource should be returned with the provided ID.            |
| already exists      | If method support user settable IDs and the same ID is reused the method should return AlreadyExists.          |
| required fields     | The method should fail with InvalidArgument if the resource has any required fields and they are not provided. |
| resource references | The method should fail with InvalidArgument if the resource has any resource references and they are invalid.  |

### Get

| Name         | Description                                                            |
| ------------ | ---------------------------------------------------------------------- |
| missing name | Method should fail with InvalidArgument if no name is provided.        |
| invalid name | Method should fail with InvalidArgument is provided name is not valid. |
| exists       | Resource should be returned without errors if it exists.               |
| not found    | Method should fail with NotFound if the resource does not exist.       |

### BatchGet

| Name            | Description                                                                                                                                 |
| --------------- | ------------------------------------------------------------------------------------------------------------------------------------------- |
| invalid parent  | Method should fail with InvalidArgument if provided parent is invalid.                                                                      |
| names missing   | Method should fail with InvalidArgument if no names are provided.                                                                           |
| names missing   | Method should fail with InvalidArgument if a provided name is not valid.                                                                    |
| all exists      | Resources should be returned without errors if they exist.                                                                                  |
| atomic          | The method must be atomic; it must fail for all resources or succeed for all resources (no partial success).                                |
| parent mismatch | If a caller sets the "parent", and the parent collection in the name of any resource being retrieved does not match, the request must fail. |
| ordered         | The order of resources in the response must be the same as the names in the request.                                                        |
| duplicate names | If a caller provides duplicate names, the service should return duplicate resources.                                                        |

### Update

| Name                | Description                                                                                                 |
| ------------------- | ----------------------------------------------------------------------------------------------------------- |
| missing name        | Method should fail with InvalidArgument if no name is provided.                                             |
| invalid name        | Method should fail with InvalidArgument is provided name is not valid.                                      |
| update time         | Field update_time should be updated when the resource is updated.                                           |
| persisted           | The updated resource should be persisted and reachable with Get.                                            |
| not found           | Method should fail with NotFound if the resource does not exist.                                            |
| invalid update mask | The method should fail with InvalidArgument if the update_mask is invalid.                                  |
| required fields     | Method should fail with InvalidArgument if any required field is missing when called with '\*' update_mask. |

### List

| Name               | Description                                                                    |
| ------------------ | ------------------------------------------------------------------------------ |
| invalid parent     | Method should fail with InvalidArgument if provided parent is invalid.         |
| invalid page token | Method should fail with InvalidArgument is provided page token is not valid.   |
| negative page size | Method should fail with InvalidArgument is provided page size is negative.     |
| isolation          | If parent is provided the method must only return resources under that parent. |
| last page          | If there are no more resources, next_page_token should not be set.             |
| more pages         | If there are more resources, next_page_token should be set.                    |
| one by one         | Listing resource one by one should eventually return all resources.            |
| deleted            | Method should not return deleted resources.                                    |

<!-- SUITES_SNIPPET -->
