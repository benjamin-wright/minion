module ponglehub.co.uk/version-sidecar

go 1.14

require (
	github.com/sirupsen/logrus v1.6.0
	k8s.io/apimachinery v0.16.9
	ponglehub.co.uk/crd-lib v1.0.0
)

replace ponglehub.co.uk/crd-lib => ../../../libraries/crd-lib
