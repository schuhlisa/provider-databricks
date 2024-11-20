package metastore_data_access

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_metastore_data_access", func(r *config.Resource) {
		r.ShortGroup = "databricks"
	})
}
