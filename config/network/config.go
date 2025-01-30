package network

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_network", func(r *config.Resource) {
		r.ShortGroup = "network"
	})
}
