# ServiceWithAllDependenciesEntity

ServiceWithAllDependenciesEntity model


## Fields

| Field                                                                                                  | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ChildServiceDependencies`                                                                             | [][components.ServiceChildDependencyEntity](../../models/components/servicechilddependencyentity.md)   | :heavy_minus_sign:                                                                                     | Services that depend on this service                                                                   |
| `ParentServiceDependencies`                                                                            | [][components.ServiceParentDependencyEntity](../../models/components/serviceparentdependencyentity.md) | :heavy_minus_sign:                                                                                     | Services that this service is dependent on                                                             |
| `ServiceDependencies`                                                                                  | [][components.ServiceDependency](../../models/components/servicedependency.md)                         | :heavy_minus_sign:                                                                                     | All dependencies. Can be one of: ServiceChildDependencyEntity, ServiceParentDependencyEntity           |