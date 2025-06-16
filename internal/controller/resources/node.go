package resources

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func GetClusterNodes(ctx context.Context, logger logr.Logger, k8sClient client.Client) ([]corev1.Node, error) {
	nodeList := &corev1.NodeList{}
	if err := k8sClient.List(ctx, nodeList); err != nil {
		logger.Error(err, "failed to list nodes")
		return nil, err
	}
	return nodeList.Items, nil
}

func formatMiB(q *resource.Quantity) string {
	bytes := q.Value()
	mib := float64(bytes) / (1024 * 1024)
	return fmt.Sprintf("%.0fMi", mib)
}

func CalculateClusterCompute(nodes []corev1.Node) map[string]string {
	totalCPU := resource.NewQuantity(0, resource.BinarySI)
	totalMemory := resource.NewQuantity(0, resource.BinarySI)
	totalPods := resource.NewQuantity(0, resource.DecimalSI)
	totalStorage := resource.NewQuantity(0, resource.BinarySI)
	var totalGPU int64

	for _, node := range nodes {
		if cpu := node.Status.Capacity.Cpu(); cpu != nil {
			totalCPU.Add(*cpu)
		}
		if memory := node.Status.Capacity.Memory(); memory != nil {
			totalMemory.Add(*memory)
		}
		if storage := node.Status.Capacity.StorageEphemeral(); storage != nil {
			totalStorage.Add(*storage)
		}
		if pods := node.Status.Allocatable.Pods(); pods != nil {
			totalPods.Add(*pods)
		}
		if gpuQuantity, ok := node.Status.Capacity["nvidia.com/gpu"]; ok {
			totalGPU += gpuQuantity.Value()
		}
	}

	return map[string]string{
		"cpu":     totalCPU.String(),
		"memory":  formatMiB(totalMemory),
		"pods":    totalPods.String(),
		"storage": formatMiB(totalStorage),
		"gpu":     fmt.Sprintf("%d", totalGPU),
	}
}
