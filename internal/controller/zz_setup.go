// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	instance "github.com/terasky-oss/provider-cloudstack/internal/controller/machines/instance"
	template "github.com/terasky-oss/provider-cloudstack/internal/controller/machines/template"
	acl "github.com/terasky-oss/provider-cloudstack/internal/controller/network/acl"
	aclrule "github.com/terasky-oss/provider-cloudstack/internal/controller/network/aclrule"
	connection "github.com/terasky-oss/provider-cloudstack/internal/controller/network/connection"
	customergateway "github.com/terasky-oss/provider-cloudstack/internal/controller/network/customergateway"
	firewall "github.com/terasky-oss/provider-cloudstack/internal/controller/network/firewall"
	forward "github.com/terasky-oss/provider-cloudstack/internal/controller/network/forward"
	gateway "github.com/terasky-oss/provider-cloudstack/internal/controller/network/gateway"
	ipaddress "github.com/terasky-oss/provider-cloudstack/internal/controller/network/ipaddress"
	nat "github.com/terasky-oss/provider-cloudstack/internal/controller/network/nat"
	network "github.com/terasky-oss/provider-cloudstack/internal/controller/network/network"
	nic "github.com/terasky-oss/provider-cloudstack/internal/controller/network/nic"
	offering "github.com/terasky-oss/provider-cloudstack/internal/controller/network/offering"
	route "github.com/terasky-oss/provider-cloudstack/internal/controller/network/route"
	vpc "github.com/terasky-oss/provider-cloudstack/internal/controller/network/vpc"
	providerconfig "github.com/terasky-oss/provider-cloudstack/internal/controller/providerconfig"
	group "github.com/terasky-oss/provider-cloudstack/internal/controller/security/group"
	grouprule "github.com/terasky-oss/provider-cloudstack/internal/controller/security/grouprule"
	keypair "github.com/terasky-oss/provider-cloudstack/internal/controller/security/keypair"
	disk "github.com/terasky-oss/provider-cloudstack/internal/controller/storage/disk"
	offeringstorage "github.com/terasky-oss/provider-cloudstack/internal/controller/storage/offering"
	volume "github.com/terasky-oss/provider-cloudstack/internal/controller/storage/volume"
	account "github.com/terasky-oss/provider-cloudstack/internal/controller/system/account"
	domain "github.com/terasky-oss/provider-cloudstack/internal/controller/system/domain"
	groupsystem "github.com/terasky-oss/provider-cloudstack/internal/controller/system/group"
	offeringsystem "github.com/terasky-oss/provider-cloudstack/internal/controller/system/offering"
	user "github.com/terasky-oss/provider-cloudstack/internal/controller/system/user"
	vmprofile "github.com/terasky-oss/provider-cloudstack/internal/controller/system/vmprofile"
	zone "github.com/terasky-oss/provider-cloudstack/internal/controller/system/zone"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		instance.Setup,
		template.Setup,
		acl.Setup,
		aclrule.Setup,
		connection.Setup,
		customergateway.Setup,
		firewall.Setup,
		firewall.Setup,
		forward.Setup,
		gateway.Setup,
		gateway.Setup,
		ipaddress.Setup,
		ipaddress.Setup,
		nat.Setup,
		network.Setup,
		nic.Setup,
		offering.Setup,
		route.Setup,
		vpc.Setup,
		providerconfig.Setup,
		group.Setup,
		grouprule.Setup,
		keypair.Setup,
		disk.Setup,
		offeringstorage.Setup,
		volume.Setup,
		account.Setup,
		domain.Setup,
		groupsystem.Setup,
		offeringsystem.Setup,
		user.Setup,
		vmprofile.Setup,
		zone.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
