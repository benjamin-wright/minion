# Minion CI

A monorepo for the Minion CI project

| Components                                                  | Description                                                                         |
| ---                                                         | ---                                                                                 |
| [CRDS](./crds/README.md)                                    | Custom Resource Definitions                                                         |
| [CRD-lib](./components/crd-lib/README.md)                   | Go library for interacting with custom resource definitions                         |
| [Resource-Monitor](./components/resource-monitor/README.md) | Monitors resources and creates resource and `version-sidecar` instances as required |
| [Version-Sidecar](./components/version-sidecar/README.md)   | Sits alongside resource images and creates `version`s as required                   |
