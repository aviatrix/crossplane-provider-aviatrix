/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"aviatrix_account":                                       config.ParameterAsIdentifier("account_name"),
	"aviatrix_spoke_transit_attachment":                      config.TemplatedStringAsIdentifier("", "{{.parameters.spoke_gw_name}}~{{.parameters.transit_gw_name}}"),
	"aviatrix_gateway_dnat":                                  config.ParameterAsIdentifier("gw_name"),
	"aviatrix_gateway_snat":                                  config.ParameterAsIdentifier("gw_name"),
	"aviatrix_gateway":                                       config.ParameterAsIdentifier("gw_name"),
	"aviatrix_segmentation_network_domain_association":       config.TemplatedStringAsIdentifier("", "{{.parameters.network_domain_name}}~{{.parameters.attachment_name}}"),
	"aviatrix_segmentation_network_domain_connection_policy": config.TemplatedStringAsIdentifier("", "{{.parameters.domain_name_1}}~{{.parameters.domain_name_2}}"),
	"aviatrix_segmentation_network_domain":                   config.ParameterAsIdentifier("domain_name"),
	"aviatrix_spoke_gateway":                                 config.ParameterAsIdentifier("gw_name"),
	"aviatrix_transit_gateway_peering":                       config.TemplatedStringAsIdentifier("", "{{.parameters.transit_gateway_name1}}~{{.parameters.transit_gateway_name2}}"),
	"aviatrix_transit_gateway":                               config.ParameterAsIdentifier("gw_name"),
	"aviatrix_vpc":                                           config.NameAsIdentifier,
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
