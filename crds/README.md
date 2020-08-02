# Minion CI - CRDS

[Back to top](../README.md)

## Custom Resources For The User

### Resource
---
A resource is a way of checking an external data source for the latest version, loading data and publishing data. It is basically just a docker image, and must implement the following commands as executables within its working directory:

| Executable | Description                                                               |
| ---        | ---                                                                       |
| version    | Writes a string value into the /input/version.txt file                    |
| load       | Writes the loaded data to the /input directory (optional)                 |
| push       | Updates the resource with the content of the /output directory (optional) |

| Variable     | Description                                                   |
| ---          | ---                                                           |
| LOAD_VERSION | The version to checkout, as provided by the `version` command |

```yaml
apiVersion: "minion.ponglehub.co.uk/v1alpha1"
kind: Resource
metadata:
  name: my-resource
spec:
  image: docker.io/resource-checker
  env:
    - name: REPO
      value: git@github.com:username/repo.git
    - name: BRANCH
      value: master
  secrets:
  - name: my-config
    keys:
    - key: id-rsa.pub
      path: /root/.ssh
```

> NB: Should try to implement `git` and `docker` resource types.

### Pipeline
---
```yaml
apiVersion: "minion.ponglehub.co.uk/v1alpha1"
kind: Pipeline
metadata:
  name: my-pipeline
spec:
  resources:
  - name: my-resource
    trigger: true
  steps:
  - name: Load source          # Resource step schema
    resource: my-resource
    action: GET
    path: some/sub/path
  - name: Install deps         # Task schema
    image: docker.io/node
    command:
    - npm run install
  - name: Run tests
    image: docker.io/node
    command:
    - npm run test
```

## Custom Resources For Internal Use

### PipelineRun
---
```yaml
apiVersion: "minion.ponglehub.co.uk/v1alpha1"
kind: PipelineRun
metadata:
  name: my-pipeline-1
spec:
  status: "Pending" / "Running" / "Error" / "Complete"
  currentTask: task-name / None
```

### Task
---
```yaml
apiVersion: "minion.ponglehub.co.uk/v1alpha1"
kind: Task
metadata:
  name: my-task
spec:
  pipeline: parent-pipeline
  run: 1
  image: docker.io/task-image
  status: "Pending" / "Running" / "Error" / "Complete"
```

### Version
---
```yaml
apiVersion: "minion.ponglehub.co.uk/v1alpha1"
kind: Version
metadata:
  name: my-version
spec:
  resource: my-resource
  version: v1.1.0
```
