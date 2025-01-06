// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	customappintegration "github.com/glalanne/provider-databricks/internal/controller/apps/customappintegration"
	cluster "github.com/glalanne/provider-databricks/internal/controller/compute/cluster"
	clusterpolicy "github.com/glalanne/provider-databricks/internal/controller/compute/clusterpolicy"
	instancepool "github.com/glalanne/provider-databricks/internal/controller/compute/instancepool"
	job "github.com/glalanne/provider-databricks/internal/controller/compute/job"
	library "github.com/glalanne/provider-databricks/internal/controller/compute/library"
	pipeline "github.com/glalanne/provider-databricks/internal/controller/compute/pipeline"
	instanceprofile "github.com/glalanne/provider-databricks/internal/controller/deployment/instanceprofile"
	mwscredentials "github.com/glalanne/provider-databricks/internal/controller/deployment/mwscredentials"
	mwscustomermanagedkeys "github.com/glalanne/provider-databricks/internal/controller/deployment/mwscustomermanagedkeys"
	mwsnccbinding "github.com/glalanne/provider-databricks/internal/controller/deployment/mwsnccbinding"
	mwsnccprivateendpointrule "github.com/glalanne/provider-databricks/internal/controller/deployment/mwsnccprivateendpointrule"
	mwsnetworkconnectivityconfig "github.com/glalanne/provider-databricks/internal/controller/deployment/mwsnetworkconnectivityconfig"
	mwsnetworks "github.com/glalanne/provider-databricks/internal/controller/deployment/mwsnetworks"
	mwsprivateaccesssettings "github.com/glalanne/provider-databricks/internal/controller/deployment/mwsprivateaccesssettings"
	mwsstorageconfigurations "github.com/glalanne/provider-databricks/internal/controller/deployment/mwsstorageconfigurations"
	mwsvpcendpoint "github.com/glalanne/provider-databricks/internal/controller/deployment/mwsvpcendpoint"
	mwsworkspaces "github.com/glalanne/provider-databricks/internal/controller/deployment/mwsworkspaces"
	budget "github.com/glalanne/provider-databricks/internal/controller/finops/budget"
	mwslogdelivery "github.com/glalanne/provider-databricks/internal/controller/log/mwslogdelivery"
	mlflowexperiment "github.com/glalanne/provider-databricks/internal/controller/mlflow/mlflowexperiment"
	mlflowmodel "github.com/glalanne/provider-databricks/internal/controller/mlflow/mlflowmodel"
	vectorsearchendpoint "github.com/glalanne/provider-databricks/internal/controller/mosaic/vectorsearchendpoint"
	vectorsearchindex "github.com/glalanne/provider-databricks/internal/controller/mosaic/vectorsearchindex"
	providerconfig "github.com/glalanne/provider-databricks/internal/controller/providerconfig"
	accesscontrolruleset "github.com/glalanne/provider-databricks/internal/controller/security/accesscontrolruleset"
	entitlements "github.com/glalanne/provider-databricks/internal/controller/security/entitlements"
	group "github.com/glalanne/provider-databricks/internal/controller/security/group"
	groupinstanceprofile "github.com/glalanne/provider-databricks/internal/controller/security/groupinstanceprofile"
	groupmember "github.com/glalanne/provider-databricks/internal/controller/security/groupmember"
	grouprole "github.com/glalanne/provider-databricks/internal/controller/security/grouprole"
	ipaccesslist "github.com/glalanne/provider-databricks/internal/controller/security/ipaccesslist"
	mwspermissionassignment "github.com/glalanne/provider-databricks/internal/controller/security/mwspermissionassignment"
	obotoken "github.com/glalanne/provider-databricks/internal/controller/security/obotoken"
	permissionassignment "github.com/glalanne/provider-databricks/internal/controller/security/permissionassignment"
	permissions "github.com/glalanne/provider-databricks/internal/controller/security/permissions"
	secret "github.com/glalanne/provider-databricks/internal/controller/security/secret"
	secretacl "github.com/glalanne/provider-databricks/internal/controller/security/secretacl"
	secretscope "github.com/glalanne/provider-databricks/internal/controller/security/secretscope"
	serviceprincipal "github.com/glalanne/provider-databricks/internal/controller/security/serviceprincipal"
	serviceprincipalrole "github.com/glalanne/provider-databricks/internal/controller/security/serviceprincipalrole"
	serviceprincipalsecret "github.com/glalanne/provider-databricks/internal/controller/security/serviceprincipalsecret"
	sqlpermissions "github.com/glalanne/provider-databricks/internal/controller/security/sqlpermissions"
	token "github.com/glalanne/provider-databricks/internal/controller/security/token"
	user "github.com/glalanne/provider-databricks/internal/controller/security/user"
	userinstanceprofile "github.com/glalanne/provider-databricks/internal/controller/security/userinstanceprofile"
	userrole "github.com/glalanne/provider-databricks/internal/controller/security/userrole"
	modelserving "github.com/glalanne/provider-databricks/internal/controller/serving/modelserving"
	compliancesecurityprofileworkspacesetting "github.com/glalanne/provider-databricks/internal/controller/settings/compliancesecurityprofileworkspacesetting"
	defaultnamespacesetting "github.com/glalanne/provider-databricks/internal/controller/settings/defaultnamespacesetting"
	enhancedsecuritymonitoringworkspacesetting "github.com/glalanne/provider-databricks/internal/controller/settings/enhancedsecuritymonitoringworkspacesetting"
	restrictworkspaceadminssetting "github.com/glalanne/provider-databricks/internal/controller/settings/restrictworkspaceadminssetting"
	provider "github.com/glalanne/provider-databricks/internal/controller/sharing/provider"
	recipient "github.com/glalanne/provider-databricks/internal/controller/sharing/recipient"
	share "github.com/glalanne/provider-databricks/internal/controller/sharing/share"
	alert "github.com/glalanne/provider-databricks/internal/controller/sql/alert"
	dashboard "github.com/glalanne/provider-databricks/internal/controller/sql/dashboard"
	query "github.com/glalanne/provider-databricks/internal/controller/sql/query"
	sqlalert "github.com/glalanne/provider-databricks/internal/controller/sql/sqlalert"
	sqldashboard "github.com/glalanne/provider-databricks/internal/controller/sql/sqldashboard"
	sqlendpoint "github.com/glalanne/provider-databricks/internal/controller/sql/sqlendpoint"
	sqlglobalconfig "github.com/glalanne/provider-databricks/internal/controller/sql/sqlglobalconfig"
	sqlquery "github.com/glalanne/provider-databricks/internal/controller/sql/sqlquery"
	sqlvisualization "github.com/glalanne/provider-databricks/internal/controller/sql/sqlvisualization"
	sqlwidget "github.com/glalanne/provider-databricks/internal/controller/sql/sqlwidget"
	dbfsfile "github.com/glalanne/provider-databricks/internal/controller/storage/dbfsfile"
	file "github.com/glalanne/provider-databricks/internal/controller/storage/file"
	mount "github.com/glalanne/provider-databricks/internal/controller/storage/mount"
	artifactallowlist "github.com/glalanne/provider-databricks/internal/controller/unity/artifactallowlist"
	catalog "github.com/glalanne/provider-databricks/internal/controller/unity/catalog"
	catalogworkspacebinding "github.com/glalanne/provider-databricks/internal/controller/unity/catalogworkspacebinding"
	connection "github.com/glalanne/provider-databricks/internal/controller/unity/connection"
	credential "github.com/glalanne/provider-databricks/internal/controller/unity/credential"
	externallocation "github.com/glalanne/provider-databricks/internal/controller/unity/externallocation"
	grant "github.com/glalanne/provider-databricks/internal/controller/unity/grant"
	grantmap "github.com/glalanne/provider-databricks/internal/controller/unity/grantmap"
	lakehousemonitor "github.com/glalanne/provider-databricks/internal/controller/unity/lakehousemonitor"
	metastore "github.com/glalanne/provider-databricks/internal/controller/unity/metastore"
	metastoreassignment "github.com/glalanne/provider-databricks/internal/controller/unity/metastoreassignment"
	metastoredataaccess "github.com/glalanne/provider-databricks/internal/controller/unity/metastoredataaccess"
	onlinetable "github.com/glalanne/provider-databricks/internal/controller/unity/onlinetable"
	qualitymonitor "github.com/glalanne/provider-databricks/internal/controller/unity/qualitymonitor"
	registeredmodel "github.com/glalanne/provider-databricks/internal/controller/unity/registeredmodel"
	schema "github.com/glalanne/provider-databricks/internal/controller/unity/schema"
	sqltable "github.com/glalanne/provider-databricks/internal/controller/unity/sqltable"
	storagecredential "github.com/glalanne/provider-databricks/internal/controller/unity/storagecredential"
	systemschema "github.com/glalanne/provider-databricks/internal/controller/unity/systemschema"
	volume "github.com/glalanne/provider-databricks/internal/controller/unity/volume"
	workspacebinding "github.com/glalanne/provider-databricks/internal/controller/unity/workspacebinding"
	directory "github.com/glalanne/provider-databricks/internal/controller/workspace/directory"
	gitcredential "github.com/glalanne/provider-databricks/internal/controller/workspace/gitcredential"
	globalinitscript "github.com/glalanne/provider-databricks/internal/controller/workspace/globalinitscript"
	notebook "github.com/glalanne/provider-databricks/internal/controller/workspace/notebook"
	notificationdestination "github.com/glalanne/provider-databricks/internal/controller/workspace/notificationdestination"
	repo "github.com/glalanne/provider-databricks/internal/controller/workspace/repo"
	workspaceconf "github.com/glalanne/provider-databricks/internal/controller/workspace/workspaceconf"
	workspacefile "github.com/glalanne/provider-databricks/internal/controller/workspace/workspacefile"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		customappintegration.Setup,
		cluster.Setup,
		clusterpolicy.Setup,
		instancepool.Setup,
		job.Setup,
		library.Setup,
		pipeline.Setup,
		instanceprofile.Setup,
		mwscredentials.Setup,
		mwscustomermanagedkeys.Setup,
		mwsnccbinding.Setup,
		mwsnccprivateendpointrule.Setup,
		mwsnetworkconnectivityconfig.Setup,
		mwsnetworks.Setup,
		mwsprivateaccesssettings.Setup,
		mwsstorageconfigurations.Setup,
		mwsvpcendpoint.Setup,
		mwsworkspaces.Setup,
		budget.Setup,
		mwslogdelivery.Setup,
		mlflowexperiment.Setup,
		mlflowmodel.Setup,
		vectorsearchendpoint.Setup,
		vectorsearchindex.Setup,
		providerconfig.Setup,
		accesscontrolruleset.Setup,
		entitlements.Setup,
		group.Setup,
		groupinstanceprofile.Setup,
		groupmember.Setup,
		grouprole.Setup,
		ipaccesslist.Setup,
		mwspermissionassignment.Setup,
		obotoken.Setup,
		permissionassignment.Setup,
		permissions.Setup,
		secret.Setup,
		secretacl.Setup,
		secretscope.Setup,
		serviceprincipal.Setup,
		serviceprincipalrole.Setup,
		serviceprincipalsecret.Setup,
		sqlpermissions.Setup,
		token.Setup,
		user.Setup,
		userinstanceprofile.Setup,
		userrole.Setup,
		modelserving.Setup,
		compliancesecurityprofileworkspacesetting.Setup,
		defaultnamespacesetting.Setup,
		enhancedsecuritymonitoringworkspacesetting.Setup,
		restrictworkspaceadminssetting.Setup,
		provider.Setup,
		recipient.Setup,
		share.Setup,
		alert.Setup,
		dashboard.Setup,
		query.Setup,
		sqlalert.Setup,
		sqldashboard.Setup,
		sqlendpoint.Setup,
		sqlglobalconfig.Setup,
		sqlquery.Setup,
		sqlvisualization.Setup,
		sqlwidget.Setup,
		dbfsfile.Setup,
		file.Setup,
		mount.Setup,
		artifactallowlist.Setup,
		catalog.Setup,
		catalogworkspacebinding.Setup,
		connection.Setup,
		credential.Setup,
		externallocation.Setup,
		grant.Setup,
		grantmap.Setup,
		lakehousemonitor.Setup,
		metastore.Setup,
		metastoreassignment.Setup,
		metastoredataaccess.Setup,
		onlinetable.Setup,
		qualitymonitor.Setup,
		registeredmodel.Setup,
		schema.Setup,
		sqltable.Setup,
		storagecredential.Setup,
		systemschema.Setup,
		volume.Setup,
		workspacebinding.Setup,
		directory.Setup,
		gitcredential.Setup,
		globalinitscript.Setup,
		notebook.Setup,
		notificationdestination.Setup,
		repo.Setup,
		workspaceconf.Setup,
		workspacefile.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
