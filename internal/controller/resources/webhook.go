package resources

import (
	"context"
	"github.com/go-logr/logr"
	addmissionv1 "k8s.io/api/admissionregistration/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func GetMutatingWebhooks(ctx context.Context, logger logr.Logger, k8sClient client.Client) ([]addmissionv1.MutatingWebhookConfiguration, error) {
	mutatingWebhooks := &addmissionv1.MutatingWebhookConfigurationList{}
	if err := k8sClient.List(ctx, mutatingWebhooks); err != nil {
		logger.Error(err, "Failed to list MutatingWebhookConfiguration")
		return nil, err
	}
	return mutatingWebhooks.Items, nil
}

func GetValidatingWebhooks(ctx context.Context, logger logr.Logger, k8sClient client.Client) ([]addmissionv1.ValidatingWebhookConfiguration, error) {
	validatingWebhooks := addmissionv1.ValidatingWebhookConfigurationList{}
	if err := k8sClient.List(ctx, &validatingWebhooks); err != nil {
		logger.Error(err, "Failed to list ValidatingWebhookConfiguration")
		return nil, err
	}
	return validatingWebhooks.Items, nil
}
