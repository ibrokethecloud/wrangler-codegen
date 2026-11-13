package main

import (
	"os"

	controllergen "github.com/rancher/wrangler/v3/pkg/controller-gen"
	"github.com/rancher/wrangler/v3/pkg/controller-gen/args"
	"k8s.io/gengo/v2/types"
	kubevirtv1 "kubevirt.io/api/core/v1"
)

func main() {
	os.Unsetenv("GOPATH")
	controllergen.Run(args.Options{
		OutputPackage: "github.com/ibrokethecloud/wrangler-codegen/pkg/generated",
		Boilerplate:   "scripts/boilerplate.go.txt",
		ExternalApplyConfigurations: map[types.Name]string{
			{Package: "k8s.io/apimachinery/pkg/apis/meta/v1", Name: "Condition"}:       "k8s.io/client-go/applyconfigurations/meta/v1",
			{Package: "k8s.io/api/core/v1", Name: "PersistentVolumeClaimVolumeSource"}: "k8s.io/client-go/applyconfigurations/core/v1",
			{Package: "k8s.io/api/core/v1", Name: "LocalObjectReference"}:              "k8s.io/client-go/applyconfigurations/core/v1",
		},
		Groups: map[string]args.Group{
			kubevirtv1.SchemeGroupVersion.Group: {
				Types: []interface{}{
					kubevirtv1.UtilityVolume{},
				},
				GenerateClients: true,
			},
		},
	})
}
