package k3s

import (
	"fmt"

	"github.com/rancher/rancher/tests/framework/clients/rancher"
	"github.com/rancher/rancher/tests/framework/extensions/cloudcredentials"
	"github.com/rancher/rancher/tests/framework/extensions/cloudcredentials/aws"
	"github.com/rancher/rancher/tests/framework/extensions/cloudcredentials/azure"
	"github.com/rancher/rancher/tests/framework/extensions/cloudcredentials/harvester"
	"github.com/rancher/rancher/tests/framework/extensions/cloudcredentials/linode"
	"github.com/rancher/rancher/tests/framework/extensions/machinepools"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

const (
	awsProviderName       = "aws"
	azureProviderName     = "azure"
	harvesterProviderName = "harvester"
	linodeProviderName    = "linode"
)

type CloudCredFunc func(rancherClient *rancher.Client) (*cloudcredentials.CloudCredential, error)
type MachinePoolFunc func(generatedPoolName, namespace string) *unstructured.Unstructured

type Provider struct {
	Name            string
	MachineConfig   string
	MachinePoolFunc MachinePoolFunc
	CloudCredFunc   CloudCredFunc
}

// CreateProvider returns all machine and cloud credential
// configs in the form of a Provider struct. Accepts a
// string of the name of the provider.
func CreateProvider(name string) Provider {
	switch {
	case name == awsProviderName:
		provider := Provider{
			Name:            name,
			MachineConfig:   machinepools.AWSResourceConfig,
			MachinePoolFunc: machinepools.NewAWSMachineConfig,
			CloudCredFunc:   aws.CreateAWSCloudCredentials,
		}
		return provider
	case name == azureProviderName:
		provider := Provider{
			Name:            name,
			MachineConfig:   machinepools.AzureResourceConfig,
			MachinePoolFunc: machinepools.NewAzureMachineConfig,
			CloudCredFunc:   azure.CreateAzureCloudCredentials,
		}
		return provider
	case name == linodeProviderName:
		provider := Provider{
			Name:            name,
			MachineConfig:   machinepools.LinodeResourceConfig,
			MachinePoolFunc: machinepools.NewLinodeMachineConfig,
			CloudCredFunc:   linode.CreateLinodeCloudCredentials,
		}
		return provider
	case name == harvesterProviderName:
		provider := Provider{
			Name:            name,
			MachineConfig:   machinepools.HarvesterResourceConfig,
			MachinePoolFunc: machinepools.NewHarvesterMachineConfig,
			CloudCredFunc:   harvester.CreateHarvesterCloudCredentials,
		}
		return provider
	default:
		panic(fmt.Sprintf("Provider:%v not found", name))
	}
}
