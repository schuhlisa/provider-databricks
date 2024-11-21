package global_init_script

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_global_init_script", func(r *config.Resource) {
		r.ShortGroup = "workspace"
	})
}
