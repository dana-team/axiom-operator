package resources

import (
	"context"
	nmstatev1 "github.com/dana-team/axiom-operator/api/nmstate/v1"
	"github.com/go-logr/logr"
	yamlv3 "gopkg.in/yaml.v3"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type DnsResolverConfig struct {
	Config struct {
		Server []string `yaml:"server"`
		Search []string `yaml:"search"`
	} `yaml:"config"`
}

func GetClusterDnsConfiguration(ctx context.Context, logger logr.Logger, k8sClient client.Client) (*DnsResolverConfig, error) {
	var state struct {
		DNSResolver *DnsResolverConfig `yaml:"dns-resolver"`
	}

	nncp := &nmstatev1.NodeNetworkConfigurationPolicy{}
	if err := k8sClient.Get(ctx, client.ObjectKey{Name: "node-resolver"}, nncp); err != nil {
		logger.Error(err, "failed to get NodeNetworkConfigurationPolicy")
		return nil, err
	}

	if err := yamlv3.Unmarshal(nncp.Spec.DesiredState.Raw, &state); err != nil {
		logger.Error(err, "failed to unmarshal DesiredState")
	}

	return state.DNSResolver, nil
}
