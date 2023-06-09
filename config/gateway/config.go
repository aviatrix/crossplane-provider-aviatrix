/*
Copyright 2021 Upbound Inc.
*/

package gateway

import (
	"github.com/upbound/upjet/pkg/config"
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("aviatrix_gateway", func(r *config.Resource) {
		config.MoveToStatus(r.TerraformResource, "software_version")
		config.MoveToStatus(r.TerraformResource, "image_version")
	})
}
