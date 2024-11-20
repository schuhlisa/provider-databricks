package model_serving

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_model_serving", func(r *config.Resource) {
		r.ShortGroup = "databricks"
	})
}
