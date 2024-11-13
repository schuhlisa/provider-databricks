package git_credential

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_git_credential", func(r *config.Resource) {
		r.ShortGroup = "databricks"
	})
}
