/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"aviatrix_account_user":          config.ParameterAsIdentifier("username"),
	"aviatrix_account":               config.ParameterAsIdentifier("account_name"),
	"aviatrix_aws_guard_duty":        config.TemplatedStringAsIdentifier("account_name", "{{ .externalName }}~~{{ .parameters.region }}"),
	"aviatrix_aws_peer":              config.TemplatedStringAsIdentifier("vpc_id1", "{{ .externalName }}~{{ .parameters.vpc_id2 }}"),
	"aviatrix_aws_tgw_connect_peer":  config.TemplatedStringAsIdentifier("connect_peer_name", "{{.parameters.tgw_name}}~~{{.parameters.connection_name}}~~{{.externalName}}"),
	"aviatrix_aws_tgw_connect":       config.TemplatedStringAsIdentifier("connection_name", "{{.parameters.tgw_name}}~~{{.externalName}}"),
	"aviatrix_aws_tgw_directconnect": config.TemplatedStringAsIdentifier("tgw_name", "{{.externalName}}~{{.parameters.dx_gateway_id}}"),
	"aviatrix_aws_tgw_intra_domain_inspection":   config.TemplatedStringAsIdentifier("route_domain_name", "{{.parameters.tgw_name}}~{{.externalName}}"),
	"aviatrix_aws_tgw_network_domain":                     config.TemplatedStringAsIdentifier("name", "{{.parameters.tgw_name}}~{{.externalName}}"),
	"aviatrix_aws_tgw_peering_domain_conn":                config.TemplatedStringAsIdentifier("tgw_name1", "{{.externalName}}:{{.parameters.domain_name1}}~{{.parameters.tgw_name2}}:{{.parameters.domain_name2}}"),
	"aviatrix_aws_tgw_peering":                            config.TemplatedStringAsIdentifier("tgw_name1", "{{.externalName}}~{{.parameters.tgw_name2}}"),
	"aviatrix_aws_tgw_transit_gateway_attachment":         config.TemplatedStringAsIdentifier("tgw_name", "{{.externalName}}~{{.parameters.vpc_id}}"),
	"aviatrix_aws_tgw_vpc_attachment":                     config.TemplatedStringAsIdentifier("tgw_name", "{{.externalName}}~{{.parameters.network_domain_name}}~{{.parameters.vpc_id}}"),
	"aviatrix_aws_tgw_vpn_conn":                           config.TemplatedStringAsIdentifier("tgw_name", "{{.externalName}}~{{.parameters.vpc_id}}"),
	"aviatrix_aws_tgw":                                    config.ParameterAsIdentifier("tgw_name"),
	"aviatrix_azure_peer":                                 config.TemplatedStringAsIdentifier("vnet_name_resource_group1", "{{.externalName}}~{{.parameters.vnet_name_resource_group2}}"),
	"aviatrix_azure_spoke_native_peering":                 config.TemplatedStringAsIdentifier("transit_gateway_name", "{{.externalName}}~{{.parameters.spoke_account_name}}~{{.parameters.spoke_vpc_id}}"),
	"aviatrix_azure_vng_conn":                    config.ParameterAsIdentifier("connection_name"),
	"aviatrix_centralized_transit_firenet":       config.TemplatedStringAsIdentifier("primary_firenet_gw_name", "{{.externalName}}~{{.parameters.secondary_firenet_gw_name}}"),
	"aviatrix_cloudn_registration":                        config.NameAsIdentifier,
	"aviatrix_cloudn_transit_gateway_attachment": config.ParameterAsIdentifier("connection_name"),
	//	"cloudwatch_agent":		config.ParameterAsIdentifier("cloudwatch_agent"),
	//	"aviatrix_controller_bgp_max_as_limit_config":				config.ParameterAsIdentifier(""), // id is controller IP, which is not a parameter of the obj...
	"aviatrix_datadog_agent":                   config.ParameterAsIdentifier("datadog_agent"),
	"aviatrix_device_interface_config":         config.ParameterAsIdentifier("device_name"),
	"aviatrix_edge_csp":                        config.ParameterAsIdentifier("gw_name"),
	"aviatrix_edge_spoke_external_device_conn": config.TemplatedStringAsIdentifier("connection_name", "{{.externalName}}~{{.parameters.site_id}}"),
	"aviatrix_spoke_transit_attachment":        config.TemplatedStringAsIdentifier("spoke_gw_name", "{{.externalName}}~{{.parameters.transit_gw_name}}"),
	"aviatrix_edge_spoke":                      config.ParameterAsIdentifier("gw_name"),
	//	"filebeat_forwarder":	config.ParameterAsIdentifier("filebeat_forwarder"), // filebeat_forwarder is not a property of the resource
	"aviatrix_firenet":                       config.ParameterAsIdentifier("vpc_id"),
	"aviatrix_firewall_instance_association": config.TemplatedStringAsIdentifier("firenet_gw_name", "{{.parameters.vpc_id}}~~{{.externalName}}~~{{.parameters.instance_id}}"),
	"aviatrix_firewall_instance":             config.ParameterAsIdentifier("instance_id"),
	"aviatrix_firewall_management_access":    config.ParameterAsIdentifier("transit_firenet_gateway_name"),
	"aviatrix_firewall_policy":               config.TemplatedStringAsIdentifier("gw_name", "{{.externalName}}~{{.parameters.src_ip}}~{{.parameters.dst_ip}}~{{.parameters.protocol}}~{{.parameters.port}}~{{.parameters.action}}"),
	"aviatrix_firewall_tag":                  config.ParameterAsIdentifier("firewall_tag"),
	"aviatrix_firewall":                      config.ParameterAsIdentifier("gw_name"),
	"aviatrix_fqdn_pass_through":             config.ParameterAsIdentifier("gw_name"),
	"aviatrix_fqdn_tag_rule":                 config.TemplatedStringAsIdentifier("fqdn_tag_name", "{{.externalName}}~{{.parameters.fqdn}}~{{.parameters.protocol}}~{{.parameters.port}}~{{.parameters.action}}"),
	"aviatrix_fqdn":                          config.ParameterAsIdentifier("fqdn_tag"),
	"aviatrix_gateway_dnat":                  config.ParameterAsIdentifier("gw_name"),
	"aviatrix_gateway_snat":                  config.ParameterAsIdentifier("gw_name"),
	"aviatrix_gateway":                       config.ParameterAsIdentifier("gw_name"),
	"aviatrix_geo_vpn":                       config.TemplatedStringAsIdentifier("service_name", "{{.externalName}}~{{.parameters.domain_name}}"),
	"aviatrix_periodic_ping":        config.ParameterAsIdentifier("gw_name"),
	"aviatrix_private_mode_lb":      config.ParameterAsIdentifier("vpc_id"),
	"aviatrix_private_mode_multicloud_endpoint":              config.ParameterAsIdentifier("vpc_id"),
	"aviatrix_rbac_group_access_account_attachment":                   config.TemplatedStringAsIdentifier("group_name", "{{.externalName}}~{{.parameters.access_account_name}}"),
	"aviatrix_rbac_group_permission_attachment":                       config.TemplatedStringAsIdentifier("group_name", "{{.externalName}}~{{.parameters.permission_name}}"),
	"aviatrix_rbac_group_user_attachment":                             config.TemplatedStringAsIdentifier("group_name", "{{.externalName}}~{{.parameters.user_name}}"),
	"aviatrix_rbac_group":                                             config.ParameterAsIdentifier("group_name"),
	"aviatrix_remote_syslog":                                          config.TemplatedStringAsIdentifier("index", "remote_syslog_{{.externalName}}"),
	"aviatrix_saml_endpoint":                                          config.ParameterAsIdentifier("endpoint_name"),
	"aviatrix_segmentation_network_domain_association":       config.TemplatedStringAsIdentifier("network_domain_name", "{{.externalName}}~{{.parameters.attachment_name}}"),
	"aviatrix_segmentation_network_domain_connection_policy": config.TemplatedStringAsIdentifier("domain_name_1", "{{.externalName}}~{{.parameters.domain_name_2}}"),
	"aviatrix_segmentation_network_domain":                   config.ParameterAsIdentifier("domain_name"),
	"aviatrix_site2cloud_ca_cert_tag":                                 config.ParameterAsIdentifier("tag_name"),
	"aviatrix_site2cloud":                                             config.TemplatedStringAsIdentifier("connection_name", "{{.externalName}}~{{.parameters.vpc_id}}"),
	"aviatrix_smart_group":                                   config.ParameterAsIdentifier("uuid"),
	"aviatrix_splunk_logging":                                         config.ParameterAsIdentifier("splunk_logging"),
	"aviatrix_spoke_external_device_conn":                             config.TemplatedStringAsIdentifier("connection_name", "{{.externalName}}~{{.parameters.vpc_id}}"),
	"aviatrix_spoke_gateway_subnet_group":                             config.TemplatedStringAsIdentifier("name", "{{.parameters.gw_name}}~{{.externalName}}"),
	"aviatrix_spoke_gateway":                                          config.ParameterAsIdentifier("gw_name"),
	"aviatrix_spoke_ha_gateway":                                       config.ParameterAsIdentifier("gw_name"),
	"aviatrix_trans_peer":                                             config.TemplatedStringAsIdentifier("source", "{{.externalName}}~{{.parameters.nexthop}}~{{.parameters.reachable_cidr}}"),
	"aviatrix_transit_external_device_conn":                           config.TemplatedStringAsIdentifier("connection_name", "{{.externalName}}~{{.parameters.vpc_id}}"),
	"aviatrix_transit_firenet_policy":                                 config.TemplatedStringAsIdentifier("transit_firenet_gateway_name", "{{.externalName}}~{{.parameters.inspected_resource_name}}"),
	"aviatrix_transit_gateway_peering":                                config.TemplatedStringAsIdentifier("transit_gateway_name1", "{{.externalName}}~{{.parameters.transit_gateway_name2}}"),
	"aviatrix_transit_gateway":                                        config.ParameterAsIdentifier("gw_name"),
	"aviatrix_tunnel":                                                 config.TemplatedStringAsIdentifier("gw_name1", "{{.externalName}}~{{.parameters.gw_name2}}"),
	"aviatrix_vgw_conn":                                               config.TemplatedStringAsIdentifier("conn_name", "{{.externalName}}~{{.parameters.vpc_id}}"),
	"aviatrix_vpc":                                                    config.NameAsIdentifier,
	"aviatrix_vpn_cert_download":                                      config.ParameterAsIdentifier("vpn_cert_download"),
	"aviatrix_vpn_profile":                                            config.NameAsIdentifier,
	"aviatrix_vpn_user_accelerator":                                   config.ParameterAsIdentifier("elb_name"),
	"aviatrix_vpn_user":                                               config.ParameterAsIdentifier("user_name"),
	// "":				config.TemplatedStringAsIdentifier("", "{{.externalName}}~{{.parameters.}}"),
	// "":				config.ParameterAsIdentifier(""),
	// "":	config.NameAsIdentifier,
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
