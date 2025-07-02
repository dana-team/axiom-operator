package resources

import (
	"context"

	nmstatev1 "github.com/dana-team/axiom-operator/api/nmstate/v1"
	"github.com/dana-team/axiom-operator/api/v1alpha1"
	"github.com/go-logr/logr"
	yamlv3 "gopkg.in/yaml.v3"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// DnsResolverConfig represents the DNS resolver configuration structure
// that matches the YAML format used in NodeNetworkConfigurationPolicy
type DnsResolverConfig struct {
	Config struct {
		Server []string `yaml:"server"`
		Search []string `yaml:"search"`
	} `yaml:"config"`
}

// GetClusterDnsConfiguration retrieves DNS configuration from NodeNetworkConfigurationPolicy
// and converts it to ClusterDnsConfig format
func GetClusterDnsConfiguration(ctx context.Context, logger logr.Logger, k8sClient client.Client) (v1alpha1.ClusterDnsConfig, error) {
	var state struct {
		DNSResolver *DnsResolverConfig `yaml:"dns-resolver"`
	}

	nncp := &nmstatev1.NodeNetworkConfigurationPolicy{}
	if err := k8sClient.Get(ctx, client.ObjectKey{Name: "node-resolver"}, nncp); err != nil {
		logger.Error(err, "failed to get NodeNetworkConfigurationPolicy")
		return v1alpha1.ClusterDnsConfig{}, err
	}

	if err := yamlv3.Unmarshal(nncp.Spec.DesiredState.Raw, &state); err != nil {
		logger.Error(err, "failed to unmarshal DesiredState")
	}

	return v1alpha1.ClusterDnsConfig{
			SearchDomains: state.DNSResolver.Config.Search,
			Servers:       state.DNSResolver.Config.Server,
		},
		nil
}
