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

	"github.com/schuhlisa/provider-databricks/config/access_control_rule_set"
	"github.com/schuhlisa/provider-databricks/config/alert"
	"github.com/schuhlisa/provider-databricks/config/artifact_allowlist"
	"github.com/schuhlisa/provider-databricks/config/budget"
	"github.com/schuhlisa/provider-databricks/config/catalog"
	"github.com/schuhlisa/provider-databricks/config/catalog_workspace_binding"
	"github.com/schuhlisa/provider-databricks/config/cluster"
	"github.com/schuhlisa/provider-databricks/config/cluster_policy"
	"github.com/schuhlisa/provider-databricks/config/compliance_security_profile_workspace_setting"
	"github.com/schuhlisa/provider-databricks/config/connection"
	"github.com/schuhlisa/provider-databricks/config/credential"
	"github.com/schuhlisa/provider-databricks/config/custom_app_integration"
	"github.com/schuhlisa/provider-databricks/config/dashboard"
	"github.com/schuhlisa/provider-databricks/config/dbfs_file"
	"github.com/schuhlisa/provider-databricks/config/default_namespace_setting"
	"github.com/schuhlisa/provider-databricks/config/directory"
	"github.com/schuhlisa/provider-databricks/config/enhanced_security_monitoring_workspace_setting"
	"github.com/schuhlisa/provider-databricks/config/entitlements"
	"github.com/schuhlisa/provider-databricks/config/external_location"
	"github.com/schuhlisa/provider-databricks/config/file"
	"github.com/schuhlisa/provider-databricks/config/git_credential"
	"github.com/schuhlisa/provider-databricks/config/global_init_script"
	"github.com/schuhlisa/provider-databricks/config/grant"
	"github.com/schuhlisa/provider-databricks/config/grants"
	"github.com/schuhlisa/provider-databricks/config/group"
	"github.com/schuhlisa/provider-databricks/config/group_instance_profile"
	"github.com/schuhlisa/provider-databricks/config/group_member"
	"github.com/schuhlisa/provider-databricks/config/group_role"
	"github.com/schuhlisa/provider-databricks/config/instance_pool"
	"github.com/schuhlisa/provider-databricks/config/instance_profile"
	"github.com/schuhlisa/provider-databricks/config/ip_access_list"
	"github.com/schuhlisa/provider-databricks/config/job"
	"github.com/schuhlisa/provider-databricks/config/lakehouse_monitor"
	"github.com/schuhlisa/provider-databricks/config/library"
	"github.com/schuhlisa/provider-databricks/config/metastore"
	"github.com/schuhlisa/provider-databricks/config/metastore_assignment"
	"github.com/schuhlisa/provider-databricks/config/metastore_data_access"
	"github.com/schuhlisa/provider-databricks/config/mlflow_experiment"
	"github.com/schuhlisa/provider-databricks/config/mlflow_model"
	"github.com/schuhlisa/provider-databricks/config/mlflow_webhook"
	"github.com/schuhlisa/provider-databricks/config/model_serving"
	"github.com/schuhlisa/provider-databricks/config/mount"
	"github.com/schuhlisa/provider-databricks/config/mws_credentials"
	"github.com/schuhlisa/provider-databricks/config/mws_customer_managed_keys"
	"github.com/schuhlisa/provider-databricks/config/mws_log_delivery"
	"github.com/schuhlisa/provider-databricks/config/mws_ncc_binding"
	"github.com/schuhlisa/provider-databricks/config/mws_ncc_private_endpoint_rule"
	"github.com/schuhlisa/provider-databricks/config/mws_network_connectivity_config"
	"github.com/schuhlisa/provider-databricks/config/mws_networks"
	"github.com/schuhlisa/provider-databricks/config/mws_permission_assignment"
	"github.com/schuhlisa/provider-databricks/config/mws_private_access_settings"
	"github.com/schuhlisa/provider-databricks/config/mws_storage_configurations"
	"github.com/schuhlisa/provider-databricks/config/mws_vpc_endpoint"
	"github.com/schuhlisa/provider-databricks/config/mws_workspaces"
	"github.com/schuhlisa/provider-databricks/config/notebook"
	"github.com/schuhlisa/provider-databricks/config/notification_destination"
	"github.com/schuhlisa/provider-databricks/config/obo_token"
	"github.com/schuhlisa/provider-databricks/config/online_table"
	"github.com/schuhlisa/provider-databricks/config/permission_assignment"
	"github.com/schuhlisa/provider-databricks/config/permissions"
	"github.com/schuhlisa/provider-databricks/config/pipeline"
	"github.com/schuhlisa/provider-databricks/config/provider"
	"github.com/schuhlisa/provider-databricks/config/quality_monitor"
	"github.com/schuhlisa/provider-databricks/config/query"
	"github.com/schuhlisa/provider-databricks/config/recipient"
	"github.com/schuhlisa/provider-databricks/config/registered_model"
	"github.com/schuhlisa/provider-databricks/config/repo"
	"github.com/schuhlisa/provider-databricks/config/restrict_workspace_admins_setting"
	"github.com/schuhlisa/provider-databricks/config/schema"
	"github.com/schuhlisa/provider-databricks/config/secret"
	"github.com/schuhlisa/provider-databricks/config/secret_acl"
	"github.com/schuhlisa/provider-databricks/config/secret_scope"
	"github.com/schuhlisa/provider-databricks/config/service_principal"
	"github.com/schuhlisa/provider-databricks/config/service_principal_role"
	"github.com/schuhlisa/provider-databricks/config/service_principal_secret"
	"github.com/schuhlisa/provider-databricks/config/share"
	"github.com/schuhlisa/provider-databricks/config/sql_alert"
	"github.com/schuhlisa/provider-databricks/config/sql_dashboard"
	"github.com/schuhlisa/provider-databricks/config/sql_endpoint"
	"github.com/schuhlisa/provider-databricks/config/sql_global_config"
	"github.com/schuhlisa/provider-databricks/config/sql_permissions"
	"github.com/schuhlisa/provider-databricks/config/sql_query"
	"github.com/schuhlisa/provider-databricks/config/sql_table"
	"github.com/schuhlisa/provider-databricks/config/sql_visualization"
	"github.com/schuhlisa/provider-databricks/config/sql_widget"
	"github.com/schuhlisa/provider-databricks/config/storage_credential"
	"github.com/schuhlisa/provider-databricks/config/system_schema"
	"github.com/schuhlisa/provider-databricks/config/token"
	"github.com/schuhlisa/provider-databricks/config/user"
	"github.com/schuhlisa/provider-databricks/config/user_instance_profile"
	"github.com/schuhlisa/provider-databricks/config/user_role"
	"github.com/schuhlisa/provider-databricks/config/vector_search_endpoint"
	"github.com/schuhlisa/provider-databricks/config/vector_search_index"
	"github.com/schuhlisa/provider-databricks/config/volume"
	"github.com/schuhlisa/provider-databricks/config/workspace_binding"
	"github.com/schuhlisa/provider-databricks/config/workspace_conf"
	"github.com/schuhlisa/provider-databricks/config/workspace_file"
)

const (
	resourcePrefix = "databricks"
	modulePath     = "github.com/schuhlisa/provider-databricks"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		// ujconfig.WithShortName("databricks"),
		ujconfig.WithRootGroup("databricks.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// app.Configure,
		cluster.Configure,
		cluster_policy.Configure,
		credential.Configure,
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
		sql_table.Configure,
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
		access_control_rule_set.Configure,
		artifact_allowlist.Configure,
		catalog_workspace_binding.Configure,
		compliance_security_profile_workspace_setting.Configure,
		custom_app_integration.Configure,
		dashboard.Configure,
		dbfs_file.Configure,
		default_namespace_setting.Configure,
		directory.Configure,
		enhanced_security_monitoring_workspace_setting.Configure,
		file.Configure,
		global_init_script.Configure,
		grant.Configure,
		group_instance_profile.Configure,
		instance_profile.Configure,
		lakehouse_monitor.Configure,
		metastore.Configure,
		metastore_assignment.Configure,
		metastore_data_access.Configure,
		mount.Configure,
		mws_credentials.Configure,
		mws_customer_managed_keys.Configure,
		mws_log_delivery.Configure,
		mws_ncc_binding.Configure,
		mws_ncc_private_endpoint_rule.Configure,
		mws_network_connectivity_config.Configure,
		mws_networks.Configure,
		mws_permission_assignment.Configure,
		mws_private_access_settings.Configure,
		mws_storage_configurations.Configure,
		mws_vpc_endpoint.Configure,
		mws_workspaces.Configure,
		notification_destination.Configure,
		obo_token.Configure,
		online_table.Configure,
		quality_monitor.Configure,
		recipient.Configure,
		registered_model.Configure,
		repo.Configure,
		restrict_workspace_admins_setting.Configure,
		secret_acl.Configure,
		service_principal_secret.Configure,
		share.Configure,
		storage_credential.Configure,
		system_schema.Configure,
		user.Configure,
		user_instance_profile.Configure,
		user_role.Configure,
		vector_search_endpoint.Configure,
		vector_search_index.Configure,
		volume.Configure,
		workspace_binding.Configure,
		workspace_conf.Configure,
		workspace_file.Configure,
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
