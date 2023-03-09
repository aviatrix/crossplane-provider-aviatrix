/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"account_user":          config.ParameterAsIdentifier("username"),
	"account":               config.ParameterAsIdentifier("account_name"),
	"aws_guard_duty":        config.TemplatedStringAsIdentifier("account_name", "{{ .externalName }}~~{{ .parameters.region }}"),
	"aws_peer":              config.TemplatedStringAsIdentifier("vpc_id1", "{{ .externalName }}~{{ .parameters.vpc_id2 }}"),
	"aws_tgw_connect_peer":  config.TemplatedStringAsIdentifier("connect_peer_name", "{{.parameters.tgw_name}}~~{{.parameters.connection_name}}~~{{.externalName}}"),
	"aws_tgw_connect":       config.TemplatedStringAsIdentifier("connection_name", "{{.parameters.tgw_name}}~~{{.externalName}}"),
	"aws_tgw_directconnect": config.TemplatedStringAsIdentifier("tgw_name", "{{.externalName}}~{{.parameters.dx_gateway_id}}"),
	"aviatrix_aws_tgw_intra_domain_inspection":   config.TemplatedStringAsIdentifier("route_domain_name", "{{.parameters.tgw_name}}~{{.externalName}}"),
	"aws_tgw_network_domain":                     config.TemplatedStringAsIdentifier("name", "{{.parameters.tgw_name}}~{{.externalName}}"),
	"aws_tgw_peering_domain_conn":                config.TemplatedStringAsIdentifier("tgw_name1", "{{.externalName}}:{{.parameters.domain_name1}}~{{.parameters.tgw_name2}}:{{.parameters.domain_name2}}"),
	"aws_tgw_peering":                            config.TemplatedStringAsIdentifier("tgw_name1", "{{.externalName}}~{{.parameters.tgw_name2}}"),
	"aws_tgw_transit_gateway_attachment":         config.TemplatedStringAsIdentifier("tgw_name", "{{.externalName}}~{{.parameters.vpc_id}}"),
	"aws_tgw_vpc_attachment":                     config.TemplatedStringAsIdentifier("tgw_name", "{{.externalName}}~{{.parameters.network_domain_name}}~{{.parameters.vpc_id}}"),
	"aws_tgw_vpn_conn":                           config.TemplatedStringAsIdentifier("tgw_name", "{{.externalName}}~{{.parameters.vpc_id}}"),
	"aws_tgw":                                    config.ParameterAsIdentifier("tgw_name"),
	"azure_peer":                                 config.TemplatedStringAsIdentifier("vnet_name_resource_group1", "{{.externalName}}~{{.parameters.vnet_name_resource_group2}}"),
	"azure_spoke_native_peering":                 config.TemplatedStringAsIdentifier("transit_gateway_name", "{{.externalName}}~{{.parameters.spoke_account_name}}~{{.parameters.spoke_vpc_id}}"),
	"aviatrix_azure_vng_conn":                    config.ParameterAsIdentifier("connection_name"),
	"aviatrix_centralized_transit_firenet":       config.TemplatedStringAsIdentifier("primary_firenet_gw_name", "{{.externalName}}~{{.parameters.secondary_firenet_gw_name}}"),
	"cloudn_registration":                        config.NameAsIdentifier,
	"aviatrix_cloudn_transit_gateway_attachment": config.ParameterAsIdentifier("connection_name"),
	//	"cloudwatch_agent":		config.ParameterAsIdentifier("cloudwatch_agent"),
	//	"aviatrix_controller_bgp_max_as_limit_config":				config.ParameterAsIdentifier(""), // id is controller IP, which is not a parameter of the obj...
	"datadog_agent":                   config.ParameterAsIdentifier("datadog_agent"),
	"device_interface_config":         config.ParameterAsIdentifier("device_name"),
	"edge_csp":                        config.ParameterAsIdentifier("gw_name"),
	"edge_spoke_external_device_conn": config.TemplatedStringAsIdentifier("connection_name", "{{.externalName}}~{{.parameters.site_id}}"),
	"spoke_transit_attachment":        config.TemplatedStringAsIdentifier("spoke_gw_name", "{{.externalName}}~{{.parameters.transit_gw_name}}"),
	"edge_spoke":                      config.ParameterAsIdentifier("gw_name"),
	//	"filebeat_forwarder":	config.ParameterAsIdentifier("filebeat_forwarder"), // filebeat_forwarder is not a property of the resource
	"firenet":                       config.ParameterAsIdentifier("vpc_id"),
	"firewall_instance_association": config.TemplatedStringAsIdentifier("firenet_gw_name", "{{.parameters.vpc_id}}~~{{.externalName}}~~{{.parameters.instance_id}}"),
	"firewall_instance":             config.ParameterAsIdentifier("instance_id"),
	"firewall_management_access":    config.ParameterAsIdentifier("transit_firenet_gateway_name"),
	"firewall_policy":               config.TemplatedStringAsIdentifier("gw_name", "{{.externalName}}~{{.parameters.src_ip}}~{{.parameters.dst_ip}}~{{.parameters.protocol}}~{{.parameters.port}}~{{.parameters.action}}"),
	"firewall_tag":                  config.ParameterAsIdentifier("firewall_tag"),
	"firewall":                      config.ParameterAsIdentifier("gw_name"),
	"fqdn_pass_through":             config.ParameterAsIdentifier("gw_name"),
	"fqdn_tag_rule":                 config.TemplatedStringAsIdentifier("fqdn_tag_name", "{{.externalName}}~{{.parameters.fqdn}}~{{.parameters.protocol}}~{{.parameters.port}}~{{.parameters.action}}"),
	"fqdn":                          config.ParameterAsIdentifier("fqdn_tag"),
	"gateway_dnat":                  config.ParameterAsIdentifier("gw_name"),
	"gateway_snat":                  config.ParameterAsIdentifier("gw_name"),
	"gateway":                       config.ParameterAsIdentifier("gw_name"),
	"geo_vpn":                       config.TemplatedStringAsIdentifier("service_name", "{{.externalName}}~{{.parameters.domain_name}}"),
	"aviatrix_periodic_ping":        config.ParameterAsIdentifier("gw_name"),
	"aviatrix_private_mode_lb":      config.ParameterAsIdentifier("vpc_id"),
	"aviatrix_private_mode_multicloud_endpoint":              config.ParameterAsIdentifier("vpc_id"),
	"rbac_group_access_account_attachment":                   config.TemplatedStringAsIdentifier("group_name", "{{.externalName}}~{{.parameters.access_account_name}}"),
	"rbac_group_permission_attachment":                       config.TemplatedStringAsIdentifier("group_name", "{{.externalName}}~{{.parameters.permission_name}}"),
	"rbac_group_user_attachment":                             config.TemplatedStringAsIdentifier("group_name", "{{.externalName}}~{{.parameters.user_name}}"),
	"rbac_group":                                             config.ParameterAsIdentifier("group_name"),
	"remote_syslog":                                          config.TemplatedStringAsIdentifier("index", "remote_syslog_{{.externalName}}"),
	"saml_endpoint":                                          config.ParameterAsIdentifier("endpoint_name"),
	"aviatrix_segmentation_network_domain_association":       config.TemplatedStringAsIdentifier("network_domain_name", "{{.externalName}}~{{.parameters.attachment_name}}"),
	"aviatrix_segmentation_network_domain_connection_policy": config.TemplatedStringAsIdentifier("domain_name_1", "{{.externalName}}~{{.parameters.domain_name_2}}"),
	"aviatrix_segmentation_network_domain":                   config.ParameterAsIdentifier("domain_name"),
	"site2cloud_ca_cert_tag":                                 config.ParameterAsIdentifier("tag_name"),
	"site2cloud":                                             config.TemplatedStringAsIdentifier("connection_name", "{{.externalName}}~{{.parameters.vpc_id}}"),
	"aviatrix_smart_group":                                   config.ParameterAsIdentifier("uuid"),
	"splunk_logging":                                         config.ParameterAsIdentifier("splunk_logging"),
	"spoke_external_device_conn":                             config.TemplatedStringAsIdentifier("connection_name", "{{.externalName}}~{{.parameters.vpc_id}}"),
	"spoke_gateway_subnet_group":                             config.TemplatedStringAsIdentifier("name", "{{.parameters.gw_name}}~{{.externalName}}"),
	"spoke_gateway":                                          config.ParameterAsIdentifier("gw_name"),
	"spoke_ha_gateway":                                       config.ParameterAsIdentifier("gw_name"),
	"trans_peer":                                             config.TemplatedStringAsIdentifier("source", "{{.externalName}}~{{.parameters.nexthop}}~{{.parameters.reachable_cidr}}"),
	"transit_external_device_conn":                           config.TemplatedStringAsIdentifier("connection_name", "{{.externalName}}~{{.parameters.vpc_id}}"),
	"transit_firenet_policy":                                 config.TemplatedStringAsIdentifier("transit_firenet_gateway_name", "{{.externalName}}~{{.parameters.inspected_resource_name}}"),
	"transit_gateway_peering":                                config.TemplatedStringAsIdentifier("transit_gateway_name1", "{{.externalName}}~{{.parameters.transit_gateway_name2}}"),
	"transit_gateway":                                        config.ParameterAsIdentifier("gw_name"),
	"tunnel":                                                 config.TemplatedStringAsIdentifier("gw_name1", "{{.externalName}}~{{.parameters.gw_name2}}"),
	"vgw_conn":                                               config.TemplatedStringAsIdentifier("conn_name", "{{.externalName}}~{{.parameters.vpc_id}}"),
	"vpc":                                                    config.NameAsIdentifier,
	"vpn_cert_download":                                      config.ParameterAsIdentifier("vpn_cert_download"),
	"vpn_profile":                                            config.NameAsIdentifier,
	"vpn_user_accelerator":                                   config.ParameterAsIdentifier("elb_name"),
	"vpn_user":                                               config.ParameterAsIdentifier("user_name"),
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
