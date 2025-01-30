#!/bin/bash

# Create directories and config files for each resource
for resource in account affinity_group autoscale_vm_profile disk disk_offering domain egress_firewall firewall instance ipaddress load_balancer_rule network network_acl network_acl_rule network_offering nic port_forward private_gateway secondary_ipaddress security_group security_group_rule service_offering ssh_keypair static_nat static_route template user volume vpc vpn_connection vpn_customer_gateway vpn_gateway zone; do
    # Create directory
    mkdir -p "config/$resource"
    
    # Create config.go file
    cat > "config/$resource/config.go" << EOF

package $resource

import (
    "github.com/crossplane/upjet/pkg/config"
)

// Configure configures the resource group category resource.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("cloudstack_$resource", func(r *config.Resource) {
        // We need to override the default group that upjet generated for
        // this resource, which would be "github"
        r.ShortGroup = "$resource"
    })
}
EOF

    echo "Created config/$resource/config.go"
done

echo "Done! Don't forget to update config/provider.go with the new imports and Configure() calls."
