## Suites

<!-- SUITES_SNIPPET -->

### Create

| Name                | Description                                                                                                    | Only if                                                                                                |
| ------------------- | -------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------ |
| missing parent      | Method should fail with InvalidArgument if no parent is provided.                                              | has Create method and resource has a parent                                                            |
| invalid parent      | Method should fail with InvalidArgument if provided parent is invalid.                                         | has Create method and resource has a parent                                                            |
| create time         | Field create_time should be populated when the resource is created.                                            | has Create method and Create method does not return long-running operation and has field 'create_time' |
| persisted           | The created resource should be persisted and reachable with Get.                                               | has Create method and Create method does not return long-running operation and has Get method          |
| user settable id    | If method support user settable IDs, when set the resource should be returned with the provided ID.            | has Create method and Create method does not return long-running operation and has user settable ID    |
| already exists      | If method support user settable IDs and the same ID is reused the method should return AlreadyExists.          | has Create method and Create method does not return long-running operation and has user settable ID    |
| required fields     | The method should fail with InvalidArgument if the resource has any required fields and they are not provided. | has Create method and resource has any required fields                                                 |
| resource references | The method should fail with InvalidArgument if the resource has any resource references and they are invalid.  | has Create method and resource has any mutable resource references                                     |

### Get

| Name         | Description                                                            | Only if                                                                                       |
| ------------ | ---------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| missing name | Method should fail with InvalidArgument if no name is provided.        | has Get method                                                                                |
| invalid name | Method should fail with InvalidArgument is provided name is not valid. | has Get method                                                                                |
| exists       | Resource should be returned without errors if it exists.               | has Create method and Create method does not return long-running operation and has Get method |
| not found    | Method should fail with NotFound if the resource does not exist.       | has Create method and Create method does not return long-running operation and has Get method |

### BatchGet

| Name            | Description                                                                                                                                 | Only if                                                                                                                                                                   |
| --------------- | ------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| invalid parent  | Method should fail with InvalidArgument if provided parent is invalid.                                                                      | resource has a parent and has BatchGet method and is not alternative batch request message                                                                                |
| names missing   | Method should fail with InvalidArgument if no names are provided.                                                                           | has BatchGet method and is not alternative batch request message                                                                                                          |
| names missing   | Method should fail with InvalidArgument if a provided name is not valid.                                                                    | has BatchGet method and is not alternative batch request message                                                                                                          |
| all exists      | Resources should be returned without errors if they exist.                                                                                  | has Create method and Create method does not return long-running operation and has BatchGet method and is not alternative batch request message                           |
| atomic          | The method must be atomic; it must fail for all resources or succeed for all resources (no partial success).                                | has Create method and Create method does not return long-running operation and has BatchGet method and is not alternative batch request message                           |
| parent mismatch | If a caller sets the "parent", and the parent collection in the name of any resource being retrieved does not match, the request must fail. | has Create method and Create method does not return long-running operation and resource has a parent and has BatchGet method and is not alternative batch request message |
| ordered         | The order of resources in the response must be the same as the names in the request.                                                        | has Create method and Create method does not return long-running operation and has BatchGet method and is not alternative batch request message                           |
| duplicate names | If a caller provides duplicate names, the service should return duplicate resources.                                                        | has Create method and Create method does not return long-running operation and has BatchGet method and is not alternative batch request message                           |

### Update

| Name                | Description                                                                                                 | Only if                                                                                                                                                                               |
| ------------------- | ----------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| missing name        | Method should fail with InvalidArgument if no name is provided.                                             | has Update method                                                                                                                                                                     |
| invalid name        | Method should fail with InvalidArgument if provided name is not valid.                                      | has Update method                                                                                                                                                                     |
| update time         | Field update_time should be updated when the resource is updated.                                           | has Create method and Create method does not return long-running operation and has Update method and Update method does not return long-running operation and has field 'update_time' |
| persisted           | The updated resource should be persisted and reachable with Get.                                            | has Create method and Create method does not return long-running operation and has Update method and Update method does not return long-running operation and has Get method          |
| not found           | Method should fail with NotFound if the resource does not exist.                                            | has Create method and Create method does not return long-running operation and has Update method                                                                                      |
| invalid update mask | The method should fail with InvalidArgument if the update_mask is invalid.                                  | has Create method and Create method does not return long-running operation and has Update method and Update method has update_mask                                                    |
| required fields     | Method should fail with InvalidArgument if any required field is missing when called with '\*' update_mask. | has Create method and Create method does not return long-running operation and has Update method and resource has any required fields                                                 |

### List

| Name               | Description                                                                    | Only if                                                                                                                                                                  |
| ------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| invalid parent     | Method should fail with InvalidArgument if provided parent is invalid.         | has List method and resource has a parent                                                                                                                                |
| invalid page token | Method should fail with InvalidArgument is provided page token is not valid.   | has List method                                                                                                                                                          |
| negative page size | Method should fail with InvalidArgument is provided page size is negative.     | has List method                                                                                                                                                          |
| isolation          | If parent is provided the method must only return resources under that parent. | resource has a parent and has Create method and Create method does not return long-running operation and has List method and resource has a parent                       |
| last page          | If there are no more resources, next_page_token should not be set.             | resource has a parent and has Create method and Create method does not return long-running operation and has List method and resource has a parent                       |
| more pages         | If there are more resources, next_page_token should be set.                    | resource has a parent and has Create method and Create method does not return long-running operation and has List method and resource has a parent                       |
| one by one         | Listing resource one by one should eventually return all resources.            | resource has a parent and has Create method and Create method does not return long-running operation and has List method and resource has a parent                       |
| deleted            | Method should not return deleted resources.                                    | resource has a parent and has Create method and Create method does not return long-running operation and has List method and has Delete method and resource has a parent |

### Search

| Name               | Description                                                                    | Only if                                                                                                                                                                    |
| ------------------ | ------------------------------------------------------------------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| invalid parent     | Method should fail with InvalidArgument if provided parent is invalid.         | has Search method and resource has a parent                                                                                                                                |
| invalid page token | Method should fail with InvalidArgument is provided page token is not valid.   | has Search method                                                                                                                                                          |
| negative page size | Method should fail with InvalidArgument is provided page size is negative.     | has Search method                                                                                                                                                          |
| isolation          | If parent is provided the method must only return resources under that parent. | resource has a parent and has Create method and Create method does not return long-running operation and has Search method and resource has a parent                       |
| last page          | If there are no more resources, next_page_token should not be set.             | resource has a parent and has Create method and Create method does not return long-running operation and has Search method and resource has a parent                       |
| more pages         | If there are more resources, next_page_token should be set.                    | resource has a parent and has Create method and Create method does not return long-running operation and has Search method and resource has a parent                       |
| one by one         | Searching resource one by one should eventually return all resources.          | resource has a parent and has Create method and Create method does not return long-running operation and has Search method and resource has a parent                       |
| deleted            | Method should not return deleted resources.                                    | resource has a parent and has Create method and Create method does not return long-running operation and has Search method and has Delete method and resource has a parent |

<!-- SUITES_SNIPPET -->
