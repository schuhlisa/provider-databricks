package vector_search_index

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_vector_search_index", func(r *config.Resource) {
		r.ShortGroup = "mosaic"
	})
}
