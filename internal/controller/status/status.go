package status

import (
	"context"
	"reflect"

	"github.com/dana-team/axiom-operator/api/v1alpha1"
	"github.com/dana-team/axiom-operator/internal/controller/common"
	"github.com/dana-team/axiom-operator/internal/controller/resources"
	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// UpdateClusterInfoStatus updates the status of a ClusterInfo resource by collecting and comparing
// the latest cluster information with the existing status. If there are differences, it updates
// the status field of the ClusterInfo resource.
func UpdateClusterInfoStatus(ctx context.Context, logger logr.Logger, clusterInfo v1alpha1.ClusterInfo, k8sClient client.Client) error {
	updatedStatus, err := collectClusterInfo(ctx, logger, k8sClient)
	if err != nil {
		return err
	}
	return common.RetryOnConflictUpdate(ctx, &clusterInfo, k8sClient, clusterInfo.Name, clusterInfo.Namespace, func(obj *v1alpha1.ClusterInfo) error {
		desiredCopy := updatedStatus.DeepCopy()
		existingCopy := obj.DeepCopy()
		desiredCopy.Normalize()
		existingCopy.Status.Normalize()

		if !reflect.DeepEqual(*desiredCopy, existingCopy.Status) {
			obj.Status = updatedStatus
			return k8sClient.Status().Update(ctx, obj)
		}
		return nil
	})
}

// collectClusterInfo gathers various information about the cluster.
func collectClusterInfo(ctx context.Context, logger logr.Logger, k8sClient client.Client) (v1alpha1.ClusterInfoStatus, error) {
	clusterInfo := v1alpha1.ClusterInfoStatus{}
	nodes, err := resources.GetClusterNodes(ctx, logger, k8sClient)
	if err != nil {
		return clusterInfo, err
	}
	nodeInfo := resources.FormatNodesInfo(nodes)
	clusterResources := resources.CalculateClusterCompute(nodes)
	k8sVersion, clusterID, err := resources.GetClusterVersionAndID(ctx, logger, k8sClient)
	if err != nil {
		return clusterInfo, err
	}

	clusterDnsConfig, err := resources.GetClusterDnsConfiguration(ctx, logger, k8sClient)
	if err != nil {
		return clusterInfo, err
	}

	routerLBAddresses, err := resources.GetRouterLBAddress(ctx, logger, k8sClient)
	if err != nil {
		return clusterInfo, err
	}

	apiServerAddresses, err := resources.GetApiServerAddress(ctx, logger, k8sClient)
	if err != nil {
		return clusterInfo, err
	}

	identityProviders, err := resources.GetIdentityProviders(ctx, logger, k8sClient)
	if err != nil {
		return clusterInfo, err
	}

	storageProvisioners, err := resources.GetStorageProvisioners(ctx, logger, k8sClient)
	if err != nil {
		return clusterInfo, err
	}

	validatingWebhooks, err := resources.GetValidatingWebhooks(ctx, logger, k8sClient)
	if err != nil {
		return clusterInfo, err
	}

	mutatingWebhooks, err := resources.GetMutatingWebhooks(ctx, logger, k8sClient)
	if err != nil {
		return clusterInfo, err
	}

	clusterInfo.ClusterID = clusterID
	clusterInfo.KubernetesVersion = k8sVersion
	clusterInfo.NodeInfo = nodeInfo
	clusterInfo.ClusterDnsConfig = clusterDnsConfig
	clusterInfo.ClusterResources = clusterResources
	clusterInfo.RouterLBAddresses = routerLBAddresses
	clusterInfo.ApiServerAddresses = apiServerAddresses
	clusterInfo.IdentityProviders = identityProviders
	clusterInfo.StorageProvisioners = storageProvisioners
	clusterInfo.MutatingWebhooks = mutatingWebhooks
	clusterInfo.ValidatingWebhooks = validatingWebhooks
	return clusterInfo, nil
}
