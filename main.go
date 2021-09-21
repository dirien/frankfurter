package main

import (
	"github.com/pulumi/pulumi-civo/sdk/go/civo"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cluster, err := civo.NewKubernetesCluster(ctx, "frankfurter", &civo.KubernetesClusterArgs{
			Name:            pulumi.StringPtr("frankfurter"),
			NumTargetNodes:  pulumi.IntPtr(1),
			TargetNodesSize: pulumi.StringPtr("g3.k3s.medium"),
			Region:          pulumi.StringPtr("FRA1"),
		})
		if err != nil {
			return err
		}

		ctx.Export("name", cluster.Name)
		return nil
	})
}
