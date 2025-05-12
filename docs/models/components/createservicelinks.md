# CreateServiceLinks

Creates a service with the appropriate integration for each external service ID passed


## Fields

| Field                                                                                   | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `ExternalServiceIds`                                                                    | *string*                                                                                | :heavy_check_mark:                                                                      | ID of the service                                                                       |
| `ConnectionID`                                                                          | *string*                                                                                | :heavy_check_mark:                                                                      | ID for the integration. This can be found by going to the edit page for the integration |
| `Integration`                                                                           | [components.Integration](../../models/components/integration.md)                        | :heavy_check_mark:                                                                      | The name of the service                                                                 |