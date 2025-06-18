package resources

import (
	"context"
	"github.com/go-logr/logr"
	configv1 "github.com/openshift/api/config/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func GetIdentityProviders(ctx context.Context, logger logr.Logger, k8sClient client.Client) ([]string, error) {
	oauth := &configv1.OAuth{}
	if err := k8sClient.Get(ctx, client.ObjectKey{Name: "cluster"}, oauth); err != nil {
		logger.Error(err, "Failed to get cluster OAuth")
	}
	idps := []string{}
	for _, provider := range oauth.Spec.IdentityProviders {
		idps = append(idps, provider.Name)
	}
	return idps, nil
}
