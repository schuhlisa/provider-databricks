package mws_customer_managed_keys

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_mws_customer_managed_keys", func(r *config.Resource) {
		r.ShortGroup = "databricks"
	})
}
