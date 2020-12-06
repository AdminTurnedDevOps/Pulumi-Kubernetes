package main

import (
	"fmt"
	"log"

	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/containerservice"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	k8sCluster("pulumiaks92", "dev2", "3d99da44-0493-4420-ad16-49388eaa87fb")
}

func k8sCluster(clusterName, resourceGroupName, clientID string) {
	pulumi.Run(func(ctx *pulumi.Context) error {
		k8s, err := containerservice.NewKubernetesCluster(ctx, clusterName, &containerservice.KubernetesClusterArgs{
			Location:          pulumi.String("eastus"),
			ResourceGroupName: pulumi.String(resourceGroupName),
			DnsPrefix:         pulumi.String(clusterName + "dns"),
			ServicePrincipal: &containerservice.KubernetesClusterServicePrincipalArgs{
				ClientId:     pulumi.String(clientID),
				ClientSecret: pulumi.String(""),
			},
			DefaultNodePool: &containerservice.KubernetesClusterDefaultNodePoolArgs{
				Name:                pulumi.String("default"),
				VmSize:              pulumi.String("Standard_D2_v2"),
				MinCount:            pulumi.Int(1),
				MaxCount:            pulumi.Int(3),
				EnableAutoScaling:   pulumi.Bool(true),
				OrchestratorVersion: pulumi.String("1.19.3"),
			},
		})
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println(k8s)
		}
		return nil

	},
	)
}
