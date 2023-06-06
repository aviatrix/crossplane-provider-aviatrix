/*
Copyright 2021 Upbound Inc.
*/

package null

import (
	"github.com/upbound/upjet/pkg/config"
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("aviatrix_gateway", func(r *config.Resource) {
		// r.LateInitializer = config.LateInitializer{
		// 	IgnoredFields: []string{"software_version"},

		// }
		config.MoveToStatus(r.TerraformResource, "software_version")
		config.MoveToStatus(r.TerraformResource, "image_version")
	})
}
