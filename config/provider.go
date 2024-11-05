/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/glalanne/provider-databricks/config/cluster"
	"github.com/glalanne/provider-databricks/config/cluster_policy"
	"github.com/glalanne/provider-databricks/config/instance_pool"
	"github.com/glalanne/provider-databricks/config/job"
	"github.com/glalanne/provider-databricks/config/notebook"
	"github.com/glalanne/provider-databricks/config/secret"
	"github.com/glalanne/provider-databricks/config/secret_scope"
	"github.com/glalanne/provider-databricks/config/token"
)

const (
	resourcePrefix = "databricks"
	modulePath     = "github.com/glalanne/provider-databricks"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("databricks.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		// add custom config functions
		cluster.Configure,
		cluster_policy.Configure,
		token.Configure,
		secret.Configure,
		secret_scope.Configure,
		notebook.Configure,
		job.Configure,
		instance_pool.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
