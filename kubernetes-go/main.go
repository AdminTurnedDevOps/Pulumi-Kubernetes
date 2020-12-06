package main

import (
	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	nginxDeploy(2, "nginxdeploy")
}

func nginxDeploy(replicas int, appName string) {
	pulumi.Run(func(ctx *pulumi.Context) error {

		appLabels := pulumi.StringMap{
			"app": pulumi.String(appName),
		}
		deployment, err := appsv1.NewDeployment(ctx, appName+"dep", &appsv1.DeploymentArgs{
			Spec: appsv1.DeploymentSpecArgs{
				Selector: &metav1.LabelSelectorArgs{
					MatchLabels: appLabels,
				},
				Replicas: pulumi.Int(replicas),
				Template: &corev1.PodTemplateSpecArgs{
					Metadata: &metav1.ObjectMetaArgs{
						Labels: appLabels,
					},
					Spec: &corev1.PodSpecArgs{
						Containers: corev1.ContainerArray{
							corev1.ContainerArgs{
								Name:  pulumi.String(appName),
								Image: pulumi.String(appName),
								Ports: corev1.ContainerPortArray{
									&corev1.ContainerPortArgs{
										ContainerPort: pulumi.Int(8080),
									},
								},
							}},
					},
				},
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("name", deployment.Metadata.Elem().Name())

		return nil
	})
}
