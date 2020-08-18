# Minion CI

A monorepo for the Minion CI project

| Components                                                       | Description                                                                         |
| ---                                                              | ---                                                                                 |
| [CRDS](./crds/README.md)                                         | Custom Resource Definitions                                                         |
| [CRD-lib](./modules/crd-lib/README.md)                           | NPM library for interacting with custom resource definitions                        |
| [async](./modules/async/README.md)                               | NPM library for async helper tools                                                  |
| [eslint-config-minion](./modules/eslint-config-minion/README.md) | NPM shared eslint config                                                            |
| [CRD-lib](./libraries/crd-lib/README.md)                         | Go library for interacting with custom resource definitions                         |
| [Resource-Monitor](./components/resource-monitor/README.md)      | Monitors resources and creates resource and `version-sidecar` instances as required |
| [Resource-Webhooks](./components/resource-webhooks/README.md)    | Exposes webhook endpoints allowing repos to create version resources directly       |
| [Version-Sidecar](./components/version-sidecar/README.md)        | Sits alongside resource images and creates `version`s as required                   |


## Layout

---

<img src="docs/Minion Overview.svg"/>
