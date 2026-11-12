package main

import (
	"os"

	controllergen "github.com/rancher/wrangler/v3/pkg/controller-gen"
	"github.com/rancher/wrangler/v3/pkg/controller-gen/args"
	corev1 "k8s.io/api/core/v1"
	kubevirtv1 "kubevirt.io/api/core/v1"
)

func main() {
	os.Unsetenv("GOPATH")
	controllergen.Run(args.Options{
		OutputPackage: "github.com/ibrokethecloud/wrangler-codegen/pkg/generated",
		Boilerplate:   "scripts/boilerplate.go.txt",
		Groups: map[string]args.Group{
			kubevirtv1.SchemeGroupVersion.Group: {
				Types: []interface{}{
					kubevirtv1.VirtualMachine{},
					kubevirtv1.VirtualMachineInstance{},
					kubevirtv1.VirtualMachineInstanceMigration{},
					kubevirtv1.KubeVirt{},
				},
				GenerateTypes:     false,
				GenerateClients:   true,
				GenerateInformers: true,
			},
			corev1.GroupName: {
				Types: []interface{}{
					corev1.Node{},
					corev1.PersistentVolume{},
					corev1.PersistentVolumeClaim{},
					corev1.ResourceQuota{},
				},
				GenerateTypes:     false,
				GenerateClients:   true,
				GenerateInformers: true,
			},
		},
	})
}
