/*
Copyright 2022 Upbound Inc.
*/

package egress_firewall

import (
    "github.com/crossplane/upjet/pkg/config"
)

// Configure configures the resource group category resource.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("cloudstack_egress_firewall", func(r *config.Resource) {
        // We need to override the default group that upjet generated for
        // this resource, which would be "github"
        r.ShortGroup = "network.cloudstack"
    })
}
