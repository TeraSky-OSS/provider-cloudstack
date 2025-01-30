package template

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudstack_template", func(r *config.Resource) {
		r.ShortGroup = "machine"
	})
}
