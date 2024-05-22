package scope

import (
	"context"
	"fmt"
	"strconv"

	"sigs.k8s.io/controller-runtime/pkg/client"

	infrav1beta2 "sigs.k8s.io/cluster-api-provider-ibmcloud/api/v1beta2"
)

// GetClusterByName finds and return a Cluster object using the specified params.
func GetClusterByName(ctx context.Context, c client.Client, namespace, name string) (*infrav1beta2.IBMPowerVSCluster, error) {
	cluster := &infrav1beta2.IBMPowerVSCluster{}
	key := client.ObjectKey{
		Namespace: namespace,
		Name:      name,
	}

	if err := c.Get(ctx, key, cluster); err != nil {
		return nil, fmt.Errorf("failed to get Cluster/%s: %w", name, err)
	}

	return cluster, nil
}

// CheckCreateInfraAnnotation checks for annotations set on IBMPowerVSCluster object to determine cluster creation workflow.
func CheckCreateInfraAnnotation(cluster infrav1beta2.IBMPowerVSCluster) bool {
	annotations := cluster.GetAnnotations()
	if len(annotations) == 0 {
		return false
	}
	value, found := annotations[infrav1beta2.CreateInfrastructureAnnotation]
	if !found {
		return false
	}
	createInfra, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}
	return createInfra
}
