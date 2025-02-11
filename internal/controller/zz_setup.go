// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	customappintegration "github.com/schuhlisa/provider-databricks/internal/controller/apps/customappintegration"
	cluster "github.com/schuhlisa/provider-databricks/internal/controller/compute/cluster"
	clusterpolicy "github.com/schuhlisa/provider-databricks/internal/controller/compute/clusterpolicy"
	instancepool "github.com/schuhlisa/provider-databricks/internal/controller/compute/instancepool"
	job "github.com/schuhlisa/provider-databricks/internal/controller/compute/job"
	library "github.com/schuhlisa/provider-databricks/internal/controller/compute/library"
	pipeline "github.com/schuhlisa/provider-databricks/internal/controller/compute/pipeline"
	instanceprofile "github.com/schuhlisa/provider-databricks/internal/controller/deployment/instanceprofile"
	mwscredentials "github.com/schuhlisa/provider-databricks/internal/controller/deployment/mwscredentials"
	mwscustomermanagedkeys "github.com/schuhlisa/provider-databricks/internal/controller/deployment/mwscustomermanagedkeys"
	mwsnccbinding "github.com/schuhlisa/provider-databricks/internal/controller/deployment/mwsnccbinding"
	mwsnccprivateendpointrule "github.com/schuhlisa/provider-databricks/internal/controller/deployment/mwsnccprivateendpointrule"
	mwsnetworkconnectivityconfig "github.com/schuhlisa/provider-databricks/internal/controller/deployment/mwsnetworkconnectivityconfig"
	mwsnetworks "github.com/schuhlisa/provider-databricks/internal/controller/deployment/mwsnetworks"
	mwsprivateaccesssettings "github.com/schuhlisa/provider-databricks/internal/controller/deployment/mwsprivateaccesssettings"
	mwsstorageconfigurations "github.com/schuhlisa/provider-databricks/internal/controller/deployment/mwsstorageconfigurations"
	mwsvpcendpoint "github.com/schuhlisa/provider-databricks/internal/controller/deployment/mwsvpcendpoint"
	mwsworkspaces "github.com/schuhlisa/provider-databricks/internal/controller/deployment/mwsworkspaces"
	budget "github.com/schuhlisa/provider-databricks/internal/controller/finops/budget"
	mwslogdelivery "github.com/schuhlisa/provider-databricks/internal/controller/log/mwslogdelivery"
	mlflowexperiment "github.com/schuhlisa/provider-databricks/internal/controller/mlflow/mlflowexperiment"
	mlflowmodel "github.com/schuhlisa/provider-databricks/internal/controller/mlflow/mlflowmodel"
	vectorsearchendpoint "github.com/schuhlisa/provider-databricks/internal/controller/mosaic/vectorsearchendpoint"
	vectorsearchindex "github.com/schuhlisa/provider-databricks/internal/controller/mosaic/vectorsearchindex"
	providerconfig "github.com/schuhlisa/provider-databricks/internal/controller/providerconfig"
	accesscontrolruleset "github.com/schuhlisa/provider-databricks/internal/controller/security/accesscontrolruleset"
	entitlements "github.com/schuhlisa/provider-databricks/internal/controller/security/entitlements"
	group "github.com/schuhlisa/provider-databricks/internal/controller/security/group"
	groupinstanceprofile "github.com/schuhlisa/provider-databricks/internal/controller/security/groupinstanceprofile"
	groupmember "github.com/schuhlisa/provider-databricks/internal/controller/security/groupmember"
	grouprole "github.com/schuhlisa/provider-databricks/internal/controller/security/grouprole"
	ipaccesslist "github.com/schuhlisa/provider-databricks/internal/controller/security/ipaccesslist"
	mwspermissionassignment "github.com/schuhlisa/provider-databricks/internal/controller/security/mwspermissionassignment"
	obotoken "github.com/schuhlisa/provider-databricks/internal/controller/security/obotoken"
	permissionassignment "github.com/schuhlisa/provider-databricks/internal/controller/security/permissionassignment"
	permissions "github.com/schuhlisa/provider-databricks/internal/controller/security/permissions"
	secret "github.com/schuhlisa/provider-databricks/internal/controller/security/secret"
	secretacl "github.com/schuhlisa/provider-databricks/internal/controller/security/secretacl"
	secretscope "github.com/schuhlisa/provider-databricks/internal/controller/security/secretscope"
	serviceprincipal "github.com/schuhlisa/provider-databricks/internal/controller/security/serviceprincipal"
	serviceprincipalrole "github.com/schuhlisa/provider-databricks/internal/controller/security/serviceprincipalrole"
	serviceprincipalsecret "github.com/schuhlisa/provider-databricks/internal/controller/security/serviceprincipalsecret"
	sqlpermissions "github.com/schuhlisa/provider-databricks/internal/controller/security/sqlpermissions"
	token "github.com/schuhlisa/provider-databricks/internal/controller/security/token"
	user "github.com/schuhlisa/provider-databricks/internal/controller/security/user"
	userinstanceprofile "github.com/schuhlisa/provider-databricks/internal/controller/security/userinstanceprofile"
	userrole "github.com/schuhlisa/provider-databricks/internal/controller/security/userrole"
	modelserving "github.com/schuhlisa/provider-databricks/internal/controller/serving/modelserving"
	compliancesecurityprofileworkspacesetting "github.com/schuhlisa/provider-databricks/internal/controller/settings/compliancesecurityprofileworkspacesetting"
	defaultnamespacesetting "github.com/schuhlisa/provider-databricks/internal/controller/settings/defaultnamespacesetting"
	enhancedsecuritymonitoringworkspacesetting "github.com/schuhlisa/provider-databricks/internal/controller/settings/enhancedsecuritymonitoringworkspacesetting"
	restrictworkspaceadminssetting "github.com/schuhlisa/provider-databricks/internal/controller/settings/restrictworkspaceadminssetting"
	provider "github.com/schuhlisa/provider-databricks/internal/controller/sharing/provider"
	recipient "github.com/schuhlisa/provider-databricks/internal/controller/sharing/recipient"
	share "github.com/schuhlisa/provider-databricks/internal/controller/sharing/share"
	alert "github.com/schuhlisa/provider-databricks/internal/controller/sql/alert"
	dashboard "github.com/schuhlisa/provider-databricks/internal/controller/sql/dashboard"
	query "github.com/schuhlisa/provider-databricks/internal/controller/sql/query"
	sqlalert "github.com/schuhlisa/provider-databricks/internal/controller/sql/sqlalert"
	sqldashboard "github.com/schuhlisa/provider-databricks/internal/controller/sql/sqldashboard"
	sqlendpoint "github.com/schuhlisa/provider-databricks/internal/controller/sql/sqlendpoint"
	sqlglobalconfig "github.com/schuhlisa/provider-databricks/internal/controller/sql/sqlglobalconfig"
	sqlquery "github.com/schuhlisa/provider-databricks/internal/controller/sql/sqlquery"
	sqlvisualization "github.com/schuhlisa/provider-databricks/internal/controller/sql/sqlvisualization"
	sqlwidget "github.com/schuhlisa/provider-databricks/internal/controller/sql/sqlwidget"
	dbfsfile "github.com/schuhlisa/provider-databricks/internal/controller/storage/dbfsfile"
	file "github.com/schuhlisa/provider-databricks/internal/controller/storage/file"
	mount "github.com/schuhlisa/provider-databricks/internal/controller/storage/mount"
	artifactallowlist "github.com/schuhlisa/provider-databricks/internal/controller/unity/artifactallowlist"
	catalog "github.com/schuhlisa/provider-databricks/internal/controller/unity/catalog"
	catalogworkspacebinding "github.com/schuhlisa/provider-databricks/internal/controller/unity/catalogworkspacebinding"
	connection "github.com/schuhlisa/provider-databricks/internal/controller/unity/connection"
	credential "github.com/schuhlisa/provider-databricks/internal/controller/unity/credential"
	externallocation "github.com/schuhlisa/provider-databricks/internal/controller/unity/externallocation"
	grant "github.com/schuhlisa/provider-databricks/internal/controller/unity/grant"
	grantmap "github.com/schuhlisa/provider-databricks/internal/controller/unity/grantmap"
	lakehousemonitor "github.com/schuhlisa/provider-databricks/internal/controller/unity/lakehousemonitor"
	metastore "github.com/schuhlisa/provider-databricks/internal/controller/unity/metastore"
	metastoreassignment "github.com/schuhlisa/provider-databricks/internal/controller/unity/metastoreassignment"
	metastoredataaccess "github.com/schuhlisa/provider-databricks/internal/controller/unity/metastoredataaccess"
	onlinetable "github.com/schuhlisa/provider-databricks/internal/controller/unity/onlinetable"
	qualitymonitor "github.com/schuhlisa/provider-databricks/internal/controller/unity/qualitymonitor"
	registeredmodel "github.com/schuhlisa/provider-databricks/internal/controller/unity/registeredmodel"
	schema "github.com/schuhlisa/provider-databricks/internal/controller/unity/schema"
	sqltable "github.com/schuhlisa/provider-databricks/internal/controller/unity/sqltable"
	storagecredential "github.com/schuhlisa/provider-databricks/internal/controller/unity/storagecredential"
	systemschema "github.com/schuhlisa/provider-databricks/internal/controller/unity/systemschema"
	volume "github.com/schuhlisa/provider-databricks/internal/controller/unity/volume"
	workspacebinding "github.com/schuhlisa/provider-databricks/internal/controller/unity/workspacebinding"
	directory "github.com/schuhlisa/provider-databricks/internal/controller/workspace/directory"
	gitcredential "github.com/schuhlisa/provider-databricks/internal/controller/workspace/gitcredential"
	globalinitscript "github.com/schuhlisa/provider-databricks/internal/controller/workspace/globalinitscript"
	notebook "github.com/schuhlisa/provider-databricks/internal/controller/workspace/notebook"
	notificationdestination "github.com/schuhlisa/provider-databricks/internal/controller/workspace/notificationdestination"
	repo "github.com/schuhlisa/provider-databricks/internal/controller/workspace/repo"
	workspaceconf "github.com/schuhlisa/provider-databricks/internal/controller/workspace/workspaceconf"
	workspacefile "github.com/schuhlisa/provider-databricks/internal/controller/workspace/workspacefile"
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
