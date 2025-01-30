package instance

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_instance", func(r *config.Resource) {
		r.ShortGroup = "machine"
	})
}
