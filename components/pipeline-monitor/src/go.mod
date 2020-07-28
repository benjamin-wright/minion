module ponglehub.co.uk/pipeline-monitor

go 1.14

require github.com/sirupsen/logrus v1.6.0

require (
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	k8s.io/apimachinery v0.16.9
	k8s.io/utils v0.0.0-20200724153422-f32512634ab7 // indirect
	ponglehub.co.uk/crd-lib v1.0.0
)

replace ponglehub.co.uk/crd-lib => ../../crd-lib
