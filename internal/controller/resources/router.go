package resources

import (
	"context"
	"github.com/go-logr/logr"
	v1 "github.com/openshift/api/route/v1"
	"net"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"strings"
)

func GetRouterLBAddress(ctx context.Context, logger logr.Logger, k8sClient client.Client) ([]string, error) {
	route := &v1.Route{}
	if err := k8sClient.Get(ctx, client.ObjectKey{Name: "console", Namespace: "openshift-console"}, route); err != nil {
		logger.Error(err, "Failed to get console route")
		return nil, err
	}
	ips, err := net.LookupHost(route.Spec.Host)
	if err != nil {
		logger.Error(err, "Failed to lookup host")
		return nil, err
	}
	return ips, nil
}

func GetApiServerAddress(ctx context.Context, logger logr.Logger, k8sClient client.Client) ([]string, error) {
	route := &v1.Route{}
	if err := k8sClient.Get(ctx, client.ObjectKey{Name: "console", Namespace: "openshift-console"}, route); err != nil {
		logger.Error(err, "Failed to get console route")
		return nil, err
	}
	apiServerURL := strings.Replace(route.Spec.Host, "console-openshift-console.apps.", "api.", 1)
	ips, err := net.LookupHost(apiServerURL)
	if err != nil {
		logger.Error(err, "Failed to lookup host")
		return nil, err
	}
	return ips, nil
}
