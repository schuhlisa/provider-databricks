package mws_ncc_binding

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_mws_ncc_binding", func(r *config.Resource) {
		r.ShortGroup = "databricks"
	})
}
