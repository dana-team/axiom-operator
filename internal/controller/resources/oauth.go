package resources

import (
	"context"
	"github.com/go-logr/logr"
	configv1 "github.com/openshift/api/config/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func GetIdentityProviders(ctx context.Context, logger logr.Logger, k8sClient client.Client) ([]configv1.IdentityProvider, error) {
	oauth := &configv1.OAuth{}
	if err := k8sClient.Get(ctx, client.ObjectKey{Name: "cluster"}, oauth); err != nil {
		logger.Error(err, "Failed to get cluster OAuth")
	}
	return oauth.Spec.IdentityProviders, nil
}
