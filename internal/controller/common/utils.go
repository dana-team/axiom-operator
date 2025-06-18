package common

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/api/resource"

	"k8s.io/client-go/util/retry"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// RetryOnConflictUpdate attempts to update a Kubernetes object with retry mechanism in case of conflicts.
// It retrieves the object first and then applies the update operation with built-in retry logic
func RetryOnConflictUpdate[T client.Object](ctx context.Context, obj T, k8sClient client.Client, name, namespace string, updateOperation func(T) error) error {
	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		if err := k8sClient.Get(ctx, client.ObjectKey{Name: name, Namespace: namespace}, obj); err != nil {
			return err
		}
		return updateOperation(obj)
	})
}

// FormatMiB converts a Kubernetes resource.Quantity value to a human-readable string in MiB units
func FormatMiB(q *resource.Quantity) string {
	bytes := q.Value()
	mib := float64(bytes) / (1024 * 1024)
	return fmt.Sprintf("%.0fMi", mib)
}
