// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	instance "github.com/terasky-oss/provider-cloudstack/internal/controller/machine/instance"
	template "github.com/terasky-oss/provider-cloudstack/internal/controller/machine/template"
	ipaddress "github.com/terasky-oss/provider-cloudstack/internal/controller/network/ipaddress"
	network "github.com/terasky-oss/provider-cloudstack/internal/controller/network/network"
	vpc "github.com/terasky-oss/provider-cloudstack/internal/controller/network/vpc"
	providerconfig "github.com/terasky-oss/provider-cloudstack/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		instance.Setup,
		template.Setup,
		ipaddress.Setup,
		network.Setup,
		vpc.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
