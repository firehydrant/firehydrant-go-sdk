# GetV1AlertsQueryParamTagMatchStrategy

The strategy to match tags. `any` will return alerts that have at least one of the supplied tags, `match_all` will return only alerts that have all of the supplied tags, and `exclude` will only return alerts that have none of the supplied tags. This currently only works for Signals alerts.


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `GetV1AlertsQueryParamTagMatchStrategyAny`      | any                                             |
| `GetV1AlertsQueryParamTagMatchStrategyMatchAll` | match_all                                       |
| `GetV1AlertsQueryParamTagMatchStrategyExclude`  | exclude                                         |