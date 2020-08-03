# Minion CI - Resource-Monitor

[Back to top](../../README.md)

Monitors `resource` CRDs and creates resource monitoring cronjobs with `resource-sidecar` as required.

## Environment

| Variable      | Description                              |
| ---           | ---                                      |
| LOG_LEVEL     | The logrus logging level                 |
| SIDECAR_IMAGE | The image to use for the version sidecar |