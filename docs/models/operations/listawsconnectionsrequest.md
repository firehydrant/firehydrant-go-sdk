# ListAwsConnectionsRequest


## Fields

| Field                                            | Type                                             | Required                                         | Description                                      |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| `Page`                                           | `*int`                                           | :heavy_minus_sign:                               | N/A                                              |
| `PerPage`                                        | `*int`                                           | :heavy_minus_sign:                               | N/A                                              |
| `AwsAccountID`                                   | `*string`                                        | :heavy_minus_sign:                               | AWS account ID containing the role to be assumed |
| `TargetArn`                                      | `*string`                                        | :heavy_minus_sign:                               | ARN of the role to be assumed                    |
| `ExternalID`                                     | `*string`                                        | :heavy_minus_sign:                               | The external ID supplied when assuming the role  |