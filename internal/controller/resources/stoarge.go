package resources

import (
	"context"
	"github.com/go-logr/logr"
	storagev1 "k8s.io/api/storage/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func GetStorageClasses(ctx context.Context, logger logr.Logger, k8sClient client.Client) ([]storagev1.StorageClass, error) {
	scList := &storagev1.StorageClassList{}
	if err := k8sClient.List(ctx, scList); err != nil {
		logger.Error(err, "Failed to list StorageClasses")
		return nil, err
	}
	return scList.Items, nil
}
