/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	"strings"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/types/name"

	"github.com/glalanne/provider-databricks/config/alert"
	"github.com/glalanne/provider-databricks/config/budget"
	"github.com/glalanne/provider-databricks/config/catalog"
	"github.com/glalanne/provider-databricks/config/cluster"
	"github.com/glalanne/provider-databricks/config/cluster_policy"
	"github.com/glalanne/provider-databricks/config/connection"
	"github.com/glalanne/provider-databricks/config/entitlements"
	"github.com/glalanne/provider-databricks/config/external_location"
	"github.com/glalanne/provider-databricks/config/git_credential"
	"github.com/glalanne/provider-databricks/config/grants"
	"github.com/glalanne/provider-databricks/config/group"
	"github.com/glalanne/provider-databricks/config/group_member"
	"github.com/glalanne/provider-databricks/config/group_role"
	"github.com/glalanne/provider-databricks/config/instance_pool"
	"github.com/glalanne/provider-databricks/config/ip_access_list"
	"github.com/glalanne/provider-databricks/config/job"
	"github.com/glalanne/provider-databricks/config/library"
	"github.com/glalanne/provider-databricks/config/mlflow_experiment"
	"github.com/glalanne/provider-databricks/config/mlflow_model"
	"github.com/glalanne/provider-databricks/config/mlflow_webhook"
	"github.com/glalanne/provider-databricks/config/model_serving"
	"github.com/glalanne/provider-databricks/config/notebook"
	"github.com/glalanne/provider-databricks/config/permission_assignment"
	"github.com/glalanne/provider-databricks/config/permissions"
	"github.com/glalanne/provider-databricks/config/pipeline"
	"github.com/glalanne/provider-databricks/config/provider"
	"github.com/glalanne/provider-databricks/config/query"
	"github.com/glalanne/provider-databricks/config/schema"
	"github.com/glalanne/provider-databricks/config/secret"
	"github.com/glalanne/provider-databricks/config/secret_scope"
	"github.com/glalanne/provider-databricks/config/service_principal"
	"github.com/glalanne/provider-databricks/config/service_principal_role"
	"github.com/glalanne/provider-databricks/config/sql_alert"
	"github.com/glalanne/provider-databricks/config/sql_dashboard"
	"github.com/glalanne/provider-databricks/config/sql_endpoint"
	"github.com/glalanne/provider-databricks/config/sql_global_config"
	"github.com/glalanne/provider-databricks/config/sql_permissions"
	"github.com/glalanne/provider-databricks/config/sql_query"
	"github.com/glalanne/provider-databricks/config/sql_visualization"
	"github.com/glalanne/provider-databricks/config/sql_widget"
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
		ujconfig.WithShortName("databricks"),
		ujconfig.WithRootGroup("crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		cluster.Configure,
		cluster_policy.Configure,
		token.Configure,
		secret.Configure,
		secret_scope.Configure,
		notebook.Configure,
		job.Configure,
		instance_pool.Configure,
		permissions.Configure,
		sql_endpoint.Configure,
		entitlements.Configure,
		group.Configure,
		group_member.Configure,
		group_role.Configure,
		ip_access_list.Configure,
		permission_assignment.Configure,
		service_principal.Configure,
		service_principal_role.Configure,
		sql_permissions.Configure,
		grants.Configure,
		pipeline.Configure,
		alert.Configure,
		query.Configure,
		sql_alert.Configure,
		sql_dashboard.Configure,
		sql_global_config.Configure,
		sql_query.Configure,
		budget.Configure,
		git_credential.Configure,
		catalog.Configure,
		connection.Configure,
		external_location.Configure,
		schema.Configure,
		library.Configure,
		sql_visualization.Configure,
		sql_widget.Configure,
		provider.Configure,
		mlflow_experiment.Configure,
		mlflow_model.Configure,
		mlflow_webhook.Configure,
		model_serving.Configure,
	} {
		configure(pc)
	}

	// Rename resources to make it more pleasing to the eye
	for _, r := range pc.Resources {

		parts := strings.Split(r.Name, "_")
		if len(parts) > 1 {
			r.ShortGroup = resourcePrefix
			r.Kind = name.NewFromSnake(strings.Join(parts[1:], "_")).Camel
		}

	}

	pc.ConfigureResources()
	return pc
}
