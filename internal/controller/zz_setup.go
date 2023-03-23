/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	account "github.com/aviatrix/provider-aviatrix/internal/controller/aviatrix/account"
	gateway "github.com/aviatrix/provider-aviatrix/internal/controller/aviatrix/gateway"
	vpc "github.com/aviatrix/provider-aviatrix/internal/controller/aviatrix/vpc"
	dnat "github.com/aviatrix/provider-aviatrix/internal/controller/gateway/dnat"
	snat "github.com/aviatrix/provider-aviatrix/internal/controller/gateway/snat"
	providerconfig "github.com/aviatrix/provider-aviatrix/internal/controller/providerconfig"
	networkdomain "github.com/aviatrix/provider-aviatrix/internal/controller/segmentation/networkdomain"
	networkdomainassociation "github.com/aviatrix/provider-aviatrix/internal/controller/segmentation/networkdomainassociation"
	networkdomainconnectionpolicy "github.com/aviatrix/provider-aviatrix/internal/controller/segmentation/networkdomainconnectionpolicy"
	gatewayspoke "github.com/aviatrix/provider-aviatrix/internal/controller/spoke/gateway"
	transitattachment "github.com/aviatrix/provider-aviatrix/internal/controller/spoke/transitattachment"
	gatewaytransit "github.com/aviatrix/provider-aviatrix/internal/controller/transit/gateway"
	gatewaypeering "github.com/aviatrix/provider-aviatrix/internal/controller/transit/gatewaypeering"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.Setup,
		gateway.Setup,
		vpc.Setup,
		dnat.Setup,
		snat.Setup,
		providerconfig.Setup,
		networkdomain.Setup,
		networkdomainassociation.Setup,
		networkdomainconnectionpolicy.Setup,
		gatewayspoke.Setup,
		transitattachment.Setup,
		gatewaytransit.Setup,
		gatewaypeering.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
