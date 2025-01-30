/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"cloudstack_account":          config.IdentifierFromProvider,
	"cloudstack_affinity_group":   config.IdentifierFromProvider,
	"cloudstack_autoscale_vm_profile":   config.IdentifierFromProvider,
	"cloudstack_disk":             config.IdentifierFromProvider,
	"cloudstack_disk_offering": config.IdentifierFromProvider,
	"cloudstack_domain":           config.IdentifierFromProvider,
	"cloudstack_egress_firewall":  config.IdentifierFromProvider,
	"cloudstack_firewall":         config.IdentifierFromProvider,
	"cloudstack_instance":         config.IdentifierFromProvider,
	"cloudstack_ipaddress":        config.IdentifierFromProvider,
	"cloudstack_load_balancer_rule":    config.IdentifierFromProvider,
	"cloudstack_network":          config.IdentifierFromProvider,
	"cloudstack_network_acl":      config.IdentifierFromProvider,
	"cloudstack_network_acl_rule":     config.IdentifierFromProvider,
	"cloudstack_network_offering": config.IdentifierFromProvider,
	"cloudstack_nic":                 config.IdentifierFromProvider,
	"cloudstack_port_forward": config.IdentifierFromProvider,
	"cloudstack_private_gateway": config.IdentifierFromProvider,
	"cloudstack_secondary_ipaddress": config.IdentifierFromProvider,
	"cloudstack_security_group": config.IdentifierFromProvider,
	"cloudstack_security_group_rule": config.IdentifierFromProvider,
	"cloudstack_service_offering": config.IdentifierFromProvider,
	"cloudstack_ssh_keypair": config.IdentifierFromProvider,
	"cloudstack_static_nat": config.IdentifierFromProvider,
	"cloudstack_static_route": config.IdentifierFromProvider,
	"cloudstack_template":         config.IdentifierFromProvider,
	"cloudstack_user": config.IdentifierFromProvider,
	"cloudstack_volume": config.IdentifierFromProvider,
	"cloudstack_vpc":              config.IdentifierFromProvider,
	"cloudstack_vpn_connection": config.IdentifierFromProvider,
	"cloudstack_vpn_customer_gateway": config.IdentifierFromProvider,
	"cloudstack_vpn_gateway": config.IdentifierFromProvider,
	"cloudstack_zone": config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
