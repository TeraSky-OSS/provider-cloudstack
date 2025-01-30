/*
Copyright 2022 Upbound Inc.
*/

package affinity_group

import (
    "github.com/crossplane/upjet/pkg/config"
)

// Configure configures the resource group category resource.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("cloudstack_affinity_group", func(r *config.Resource) {
        // We need to override the default group that upjet generated for
        // this resource, which would be "github"
        r.ShortGroup = "system"
    })
}
