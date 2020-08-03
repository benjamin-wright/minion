# Version Sidecar

Pairs with a resource image to create a new ResourceVersion when the monitored resource is updated

## Usage

Expects a file containing the version information at `/input/version.txt`

## Environment

| Variable  | Description                     |
| ---       | ---                             |
| RESOURCE  | The name of the origin resource |
| LOG_LEVEL | The logrus logging level to use |