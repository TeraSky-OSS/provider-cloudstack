/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/terasky-oss/provider-cloudstack/config/instance"
	"github.com/terasky-oss/provider-cloudstack/config/ipaddress"
	"github.com/terasky-oss/provider-cloudstack/config/network"
	"github.com/terasky-oss/provider-cloudstack/config/template"
	"github.com/terasky-oss/provider-cloudstack/config/vpc"
)

const (
	resourcePrefix = "cloudstack"
	modulePath     = "github.com/terasky-oss/provider-cloudstack"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("terasky.com"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		instance.Configure,
		ipaddress.Configure,
		network.Configure,
		template.Configure,
		vpc.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
