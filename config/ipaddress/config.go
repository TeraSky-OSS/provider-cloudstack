package ipaddress

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_ipaddress", func(r *config.Resource) {
		r.ShortGroup = "network"
		r.References["network_id"] = config.Reference{
			Type: "github.com/terasky-oss/provider-cloudstack/apis/network/v1alpha1.Network",
		}
		r.References["vpc_id"] = config.Reference{
			Type: "github.com/terasky-oss/provider-cloudstack/apis/network/v1alpha1.VPC",
		}
	})
}
