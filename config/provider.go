/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/terasky-oss/provider-cloudstack/config/account"
	"github.com/terasky-oss/provider-cloudstack/config/affinity_group"
	"github.com/terasky-oss/provider-cloudstack/config/autoscale_vm_profile"
	"github.com/terasky-oss/provider-cloudstack/config/disk"
	"github.com/terasky-oss/provider-cloudstack/config/disk_offering"
	"github.com/terasky-oss/provider-cloudstack/config/domain"
	"github.com/terasky-oss/provider-cloudstack/config/egress_firewall"
	"github.com/terasky-oss/provider-cloudstack/config/firewall"
	"github.com/terasky-oss/provider-cloudstack/config/instance"
	"github.com/terasky-oss/provider-cloudstack/config/ipaddress"
	"github.com/terasky-oss/provider-cloudstack/config/load_balancer_rule"
	"github.com/terasky-oss/provider-cloudstack/config/network"
	"github.com/terasky-oss/provider-cloudstack/config/network_acl"
	"github.com/terasky-oss/provider-cloudstack/config/network_acl_rule"
	"github.com/terasky-oss/provider-cloudstack/config/network_offering"
	"github.com/terasky-oss/provider-cloudstack/config/nic"
	"github.com/terasky-oss/provider-cloudstack/config/port_forward"
	"github.com/terasky-oss/provider-cloudstack/config/private_gateway"
	"github.com/terasky-oss/provider-cloudstack/config/secondary_ipaddress"
	"github.com/terasky-oss/provider-cloudstack/config/security_group"
	"github.com/terasky-oss/provider-cloudstack/config/security_group_rule"
	"github.com/terasky-oss/provider-cloudstack/config/service_offering"
	"github.com/terasky-oss/provider-cloudstack/config/ssh_keypair"
	"github.com/terasky-oss/provider-cloudstack/config/static_nat"
	"github.com/terasky-oss/provider-cloudstack/config/static_route"
	"github.com/terasky-oss/provider-cloudstack/config/template"
	"github.com/terasky-oss/provider-cloudstack/config/user"
	"github.com/terasky-oss/provider-cloudstack/config/volume"
	"github.com/terasky-oss/provider-cloudstack/config/vpc"
	"github.com/terasky-oss/provider-cloudstack/config/vpn_connection"
	"github.com/terasky-oss/provider-cloudstack/config/vpn_customer_gateway"
	"github.com/terasky-oss/provider-cloudstack/config/vpn_gateway"
	"github.com/terasky-oss/provider-cloudstack/config/zone"
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
		account.Configure,
		affinity_group.Configure,
		autoscale_vm_profile.Configure,
		disk.Configure,
		disk_offering.Configure,
		domain.Configure,
		egress_firewall.Configure,
		firewall.Configure,
		instance.Configure,
		ipaddress.Configure,
		load_balancer_rule.Configure,
		network.Configure,
		network_acl.Configure,
		network_acl_rule.Configure,
		network_offering.Configure,
		nic.Configure,
		port_forward.Configure,
		private_gateway.Configure,
		secondary_ipaddress.Configure,
		security_group.Configure,
		security_group_rule.Configure,
		service_offering.Configure,
		ssh_keypair.Configure,
		static_nat.Configure,
		static_route.Configure,
		template.Configure,
		user.Configure,
		volume.Configure,
		vpc.Configure,
		vpn_connection.Configure,
		vpn_customer_gateway.Configure,
		vpn_gateway.Configure,
		zone.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
