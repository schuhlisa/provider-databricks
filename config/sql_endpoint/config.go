package sql_endpoint

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_sql_endpoint", func(r *config.Resource) {
		r.ShortGroup = "databricks"
		r.Kind = "SqlEndpoint"
		if s, ok := r.TerraformResource.Schema["jdbc_url"]; ok {
			s.Sensitive = true
		}

		if s, ok := r.TerraformResource.Schema["odbc_params"]; ok {
			s.Sensitive = true
		}
	})
}
