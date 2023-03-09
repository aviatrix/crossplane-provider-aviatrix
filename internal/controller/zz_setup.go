/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	user "github.com/aviatrix/provider-aviatrix/internal/controller/account/user"
	account "github.com/aviatrix/provider-aviatrix/internal/controller/aviatrix/account"
	firenet "github.com/aviatrix/provider-aviatrix/internal/controller/aviatrix/firenet"
	firewall "github.com/aviatrix/provider-aviatrix/internal/controller/aviatrix/firewall"
	fqdn "github.com/aviatrix/provider-aviatrix/internal/controller/aviatrix/fqdn"
	gateway "github.com/aviatrix/provider-aviatrix/internal/controller/aviatrix/gateway"
	site2cloud "github.com/aviatrix/provider-aviatrix/internal/controller/aviatrix/site2cloud"
	tunnel "github.com/aviatrix/provider-aviatrix/internal/controller/aviatrix/tunnel"
	vpc "github.com/aviatrix/provider-aviatrix/internal/controller/aviatrix/vpc"
	guardduty "github.com/aviatrix/provider-aviatrix/internal/controller/aws/guardduty"
	peer "github.com/aviatrix/provider-aviatrix/internal/controller/aws/peer"
	tgw "github.com/aviatrix/provider-aviatrix/internal/controller/aws/tgw"
	tgwconnect "github.com/aviatrix/provider-aviatrix/internal/controller/aws/tgwconnect"
	tgwconnectpeer "github.com/aviatrix/provider-aviatrix/internal/controller/aws/tgwconnectpeer"
	tgwdirectconnect "github.com/aviatrix/provider-aviatrix/internal/controller/aws/tgwdirectconnect"
	tgwintradomaininspection "github.com/aviatrix/provider-aviatrix/internal/controller/aws/tgwintradomaininspection"
	tgwnetworkdomain "github.com/aviatrix/provider-aviatrix/internal/controller/aws/tgwnetworkdomain"
	tgwpeering "github.com/aviatrix/provider-aviatrix/internal/controller/aws/tgwpeering"
	tgwpeeringdomainconn "github.com/aviatrix/provider-aviatrix/internal/controller/aws/tgwpeeringdomainconn"
	tgwtransitgatewayattachment "github.com/aviatrix/provider-aviatrix/internal/controller/aws/tgwtransitgatewayattachment"
	tgwvpcattachment "github.com/aviatrix/provider-aviatrix/internal/controller/aws/tgwvpcattachment"
	tgwvpnconn "github.com/aviatrix/provider-aviatrix/internal/controller/aws/tgwvpnconn"
	peerazure "github.com/aviatrix/provider-aviatrix/internal/controller/azure/peer"
	spokenativepeering "github.com/aviatrix/provider-aviatrix/internal/controller/azure/spokenativepeering"
	vngconn "github.com/aviatrix/provider-aviatrix/internal/controller/azure/vngconn"
	transitfirenet "github.com/aviatrix/provider-aviatrix/internal/controller/centralized/transitfirenet"
	registration "github.com/aviatrix/provider-aviatrix/internal/controller/cloudn/registration"
	transitgatewayattachment "github.com/aviatrix/provider-aviatrix/internal/controller/cloudn/transitgatewayattachment"
	agent "github.com/aviatrix/provider-aviatrix/internal/controller/datadog/agent"
	interfaceconfig "github.com/aviatrix/provider-aviatrix/internal/controller/device/interfaceconfig"
	firewallingintravpc "github.com/aviatrix/provider-aviatrix/internal/controller/distributed/firewallingintravpc"
	csp "github.com/aviatrix/provider-aviatrix/internal/controller/edge/csp"
	spoke "github.com/aviatrix/provider-aviatrix/internal/controller/edge/spoke"
	spokeexternaldeviceconn "github.com/aviatrix/provider-aviatrix/internal/controller/edge/spokeexternaldeviceconn"
	spoketransitattachment "github.com/aviatrix/provider-aviatrix/internal/controller/edge/spoketransitattachment"
	instance "github.com/aviatrix/provider-aviatrix/internal/controller/firewall/instance"
	instanceassociation "github.com/aviatrix/provider-aviatrix/internal/controller/firewall/instanceassociation"
	managementaccess "github.com/aviatrix/provider-aviatrix/internal/controller/firewall/managementaccess"
	policy "github.com/aviatrix/provider-aviatrix/internal/controller/firewall/policy"
	tag "github.com/aviatrix/provider-aviatrix/internal/controller/firewall/tag"
	passthrough "github.com/aviatrix/provider-aviatrix/internal/controller/fqdn/passthrough"
	tagrule "github.com/aviatrix/provider-aviatrix/internal/controller/fqdn/tagrule"
	dnat "github.com/aviatrix/provider-aviatrix/internal/controller/gateway/dnat"
	snat "github.com/aviatrix/provider-aviatrix/internal/controller/gateway/snat"
	vpn "github.com/aviatrix/provider-aviatrix/internal/controller/geo/vpn"
	ping "github.com/aviatrix/provider-aviatrix/internal/controller/periodic/ping"
	modelb "github.com/aviatrix/provider-aviatrix/internal/controller/private/modelb"
	modemulticloudendpoint "github.com/aviatrix/provider-aviatrix/internal/controller/private/modemulticloudendpoint"
	providerconfig "github.com/aviatrix/provider-aviatrix/internal/controller/providerconfig"
	group "github.com/aviatrix/provider-aviatrix/internal/controller/rbac/group"
	groupaccessaccountattachment "github.com/aviatrix/provider-aviatrix/internal/controller/rbac/groupaccessaccountattachment"
	grouppermissionattachment "github.com/aviatrix/provider-aviatrix/internal/controller/rbac/grouppermissionattachment"
	groupuserattachment "github.com/aviatrix/provider-aviatrix/internal/controller/rbac/groupuserattachment"
	syslog "github.com/aviatrix/provider-aviatrix/internal/controller/remote/syslog"
	endpoint "github.com/aviatrix/provider-aviatrix/internal/controller/saml/endpoint"
	networkdomain "github.com/aviatrix/provider-aviatrix/internal/controller/segmentation/networkdomain"
	networkdomainassociation "github.com/aviatrix/provider-aviatrix/internal/controller/segmentation/networkdomainassociation"
	networkdomainconnectionpolicy "github.com/aviatrix/provider-aviatrix/internal/controller/segmentation/networkdomainconnectionpolicy"
	cacerttag "github.com/aviatrix/provider-aviatrix/internal/controller/site2cloud/cacerttag"
	groupsmart "github.com/aviatrix/provider-aviatrix/internal/controller/smart/group"
	logging "github.com/aviatrix/provider-aviatrix/internal/controller/splunk/logging"
	externaldeviceconn "github.com/aviatrix/provider-aviatrix/internal/controller/spoke/externaldeviceconn"
	gatewayspoke "github.com/aviatrix/provider-aviatrix/internal/controller/spoke/gateway"
	gatewaysubnetgroup "github.com/aviatrix/provider-aviatrix/internal/controller/spoke/gatewaysubnetgroup"
	hagateway "github.com/aviatrix/provider-aviatrix/internal/controller/spoke/hagateway"
	transitattachment "github.com/aviatrix/provider-aviatrix/internal/controller/spoke/transitattachment"
	peertrans "github.com/aviatrix/provider-aviatrix/internal/controller/trans/peer"
	externaldeviceconntransit "github.com/aviatrix/provider-aviatrix/internal/controller/transit/externaldeviceconn"
	firenetpolicy "github.com/aviatrix/provider-aviatrix/internal/controller/transit/firenetpolicy"
	gatewaytransit "github.com/aviatrix/provider-aviatrix/internal/controller/transit/gateway"
	gatewaypeering "github.com/aviatrix/provider-aviatrix/internal/controller/transit/gatewaypeering"
	conn "github.com/aviatrix/provider-aviatrix/internal/controller/vgw/conn"
	certdownload "github.com/aviatrix/provider-aviatrix/internal/controller/vpn/certdownload"
	profile "github.com/aviatrix/provider-aviatrix/internal/controller/vpn/profile"
	uservpn "github.com/aviatrix/provider-aviatrix/internal/controller/vpn/user"
	useraccelerator "github.com/aviatrix/provider-aviatrix/internal/controller/vpn/useraccelerator"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		user.Setup,
		account.Setup,
		firenet.Setup,
		firewall.Setup,
		fqdn.Setup,
		gateway.Setup,
		site2cloud.Setup,
		tunnel.Setup,
		vpc.Setup,
		guardduty.Setup,
		peer.Setup,
		tgw.Setup,
		tgwconnect.Setup,
		tgwconnectpeer.Setup,
		tgwdirectconnect.Setup,
		tgwintradomaininspection.Setup,
		tgwnetworkdomain.Setup,
		tgwpeering.Setup,
		tgwpeeringdomainconn.Setup,
		tgwtransitgatewayattachment.Setup,
		tgwvpcattachment.Setup,
		tgwvpnconn.Setup,
		peerazure.Setup,
		spokenativepeering.Setup,
		vngconn.Setup,
		transitfirenet.Setup,
		registration.Setup,
		transitgatewayattachment.Setup,
		agent.Setup,
		interfaceconfig.Setup,
		firewallingintravpc.Setup,
		csp.Setup,
		spoke.Setup,
		spokeexternaldeviceconn.Setup,
		spoketransitattachment.Setup,
		instance.Setup,
		instanceassociation.Setup,
		managementaccess.Setup,
		policy.Setup,
		tag.Setup,
		passthrough.Setup,
		tagrule.Setup,
		dnat.Setup,
		snat.Setup,
		vpn.Setup,
		ping.Setup,
		modelb.Setup,
		modemulticloudendpoint.Setup,
		providerconfig.Setup,
		group.Setup,
		groupaccessaccountattachment.Setup,
		grouppermissionattachment.Setup,
		groupuserattachment.Setup,
		syslog.Setup,
		endpoint.Setup,
		networkdomain.Setup,
		networkdomainassociation.Setup,
		networkdomainconnectionpolicy.Setup,
		cacerttag.Setup,
		groupsmart.Setup,
		logging.Setup,
		externaldeviceconn.Setup,
		gatewayspoke.Setup,
		gatewaysubnetgroup.Setup,
		hagateway.Setup,
		transitattachment.Setup,
		peertrans.Setup,
		externaldeviceconntransit.Setup,
		firenetpolicy.Setup,
		gatewaytransit.Setup,
		gatewaypeering.Setup,
		conn.Setup,
		certdownload.Setup,
		profile.Setup,
		uservpn.Setup,
		useraccelerator.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
