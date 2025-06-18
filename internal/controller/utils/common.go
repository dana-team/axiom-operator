package utils

import (
	"context"
	"k8s.io/client-go/util/retry"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func RetryOnConflictUpdate[T client.Object](ctx context.Context, obj T, k8sClient client.Client, name, namespace string, updateOperation func(T) error) error {
	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		if err := k8sClient.Get(ctx, client.ObjectKey{Name: name, Namespace: namespace}, obj); err != nil {
			return err
		}
		return updateOperation(obj)
	})
}
