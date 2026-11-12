package main

import (
	"os"

	controllergen "github.com/rancher/wrangler/v3/pkg/controller-gen"
	"github.com/rancher/wrangler/v3/pkg/controller-gen/args"
	corev1 "k8s.io/api/core/v1"
	kubevirtv1 "kubevirt.io/api/core/v1"
	cdiv1 "kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
	cdiuploadv1 "kubevirt.io/containerized-data-importer-api/pkg/apis/upload/v1beta1"

	kubeovnv1 "github.com/kubeovn/kube-ovn/pkg/apis/kubeovn/v1"
)

func main() {
	os.Unsetenv("GOPATH")
	controllergen.Run(args.Options{
		OutputPackage: "github.com/harvester/harvester/pkg/generated",
		Boilerplate:   "scripts/boilerplate.go.txt",
		Groups: map[string]args.Group{
			kubevirtv1.SchemeGroupVersion.Group: {
				Types: []interface{}{
					kubevirtv1.VirtualMachine{},
					kubevirtv1.VirtualMachineInstance{},
					kubevirtv1.VirtualMachineInstanceMigration{},
					kubevirtv1.KubeVirt{},
				},
				GenerateTypes:   false,
				GenerateClients: true,
			},
			corev1.GroupName: {
				Types: []interface{}{
					corev1.Node{},
					corev1.PersistentVolume{},
					corev1.PersistentVolumeClaim{},
					corev1.ResourceQuota{},
				},
				GenerateTypes:   false,
				GenerateClients: true,
			},
			cdiv1.CDIGroupVersionKind.Group: {
				Types: []interface{}{
					cdiv1.DataVolume{},
					cdiv1.StorageProfile{},
					cdiv1.VolumeImportSource{},
					cdiv1.CDI{},
				},
				GenerateTypes:   false,
				GenerateClients: true,
			},
			cdiuploadv1.SchemeGroupVersion.Group: {
				Types: []interface{}{
					cdiuploadv1.UploadTokenRequest{},
				},
				GenerateTypes:   false,
				GenerateClients: true,
			},
			kubeovnv1.SchemeGroupVersion.Group: {
				Types: []interface{}{
					kubeovnv1.Subnet{},
				},
				GenerateTypes:   false,
				GenerateClients: true,
			},
		},
	})
}
