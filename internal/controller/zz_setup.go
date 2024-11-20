// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	accesscontrolruleset "github.com/glalanne/provider-databricks/internal/controller/databricks/accesscontrolruleset"
	alert "github.com/glalanne/provider-databricks/internal/controller/databricks/alert"
	artifactallowlist "github.com/glalanne/provider-databricks/internal/controller/databricks/artifactallowlist"
	automaticclusterupdateworkspacesetting "github.com/glalanne/provider-databricks/internal/controller/databricks/automaticclusterupdateworkspacesetting"
	awss3mount "github.com/glalanne/provider-databricks/internal/controller/databricks/awss3mount"
	azureadlsgen1mount "github.com/glalanne/provider-databricks/internal/controller/databricks/azureadlsgen1mount"
	azureadlsgen2mount "github.com/glalanne/provider-databricks/internal/controller/databricks/azureadlsgen2mount"
	azureblobmount "github.com/glalanne/provider-databricks/internal/controller/databricks/azureblobmount"
	budget "github.com/glalanne/provider-databricks/internal/controller/databricks/budget"
	catalog "github.com/glalanne/provider-databricks/internal/controller/databricks/catalog"
	catalogworkspacebinding "github.com/glalanne/provider-databricks/internal/controller/databricks/catalogworkspacebinding"
	cluster "github.com/glalanne/provider-databricks/internal/controller/databricks/cluster"
	clusterpolicy "github.com/glalanne/provider-databricks/internal/controller/databricks/clusterpolicy"
	compliancesecurityprofileworkspacesetting "github.com/glalanne/provider-databricks/internal/controller/databricks/compliancesecurityprofileworkspacesetting"
	connection "github.com/glalanne/provider-databricks/internal/controller/databricks/connection"
	customappintegration "github.com/glalanne/provider-databricks/internal/controller/databricks/customappintegration"
	dashboard "github.com/glalanne/provider-databricks/internal/controller/databricks/dashboard"
	dbfsfile "github.com/glalanne/provider-databricks/internal/controller/databricks/dbfsfile"
	defaultnamespacesetting "github.com/glalanne/provider-databricks/internal/controller/databricks/defaultnamespacesetting"
	directory "github.com/glalanne/provider-databricks/internal/controller/databricks/directory"
	enhancedsecuritymonitoringworkspacesetting "github.com/glalanne/provider-databricks/internal/controller/databricks/enhancedsecuritymonitoringworkspacesetting"
	entitlements "github.com/glalanne/provider-databricks/internal/controller/databricks/entitlements"
	externallocation "github.com/glalanne/provider-databricks/internal/controller/databricks/externallocation"
	file "github.com/glalanne/provider-databricks/internal/controller/databricks/file"
	gitcredential "github.com/glalanne/provider-databricks/internal/controller/databricks/gitcredential"
	globalinitscript "github.com/glalanne/provider-databricks/internal/controller/databricks/globalinitscript"
	grant "github.com/glalanne/provider-databricks/internal/controller/databricks/grant"
	grants "github.com/glalanne/provider-databricks/internal/controller/databricks/grants"
	group "github.com/glalanne/provider-databricks/internal/controller/databricks/group"
	groupinstanceprofile "github.com/glalanne/provider-databricks/internal/controller/databricks/groupinstanceprofile"
	groupmember "github.com/glalanne/provider-databricks/internal/controller/databricks/groupmember"
	grouprole "github.com/glalanne/provider-databricks/internal/controller/databricks/grouprole"
	instancepool "github.com/glalanne/provider-databricks/internal/controller/databricks/instancepool"
	instanceprofile "github.com/glalanne/provider-databricks/internal/controller/databricks/instanceprofile"
	ipaccesslist "github.com/glalanne/provider-databricks/internal/controller/databricks/ipaccesslist"
	job "github.com/glalanne/provider-databricks/internal/controller/databricks/job"
	lakehousemonitor "github.com/glalanne/provider-databricks/internal/controller/databricks/lakehousemonitor"
	library "github.com/glalanne/provider-databricks/internal/controller/databricks/library"
	metastore "github.com/glalanne/provider-databricks/internal/controller/databricks/metastore"
	metastoreassignment "github.com/glalanne/provider-databricks/internal/controller/databricks/metastoreassignment"
	metastoredataaccess "github.com/glalanne/provider-databricks/internal/controller/databricks/metastoredataaccess"
	mlflowexperiment "github.com/glalanne/provider-databricks/internal/controller/databricks/mlflowexperiment"
	mlflowmodel "github.com/glalanne/provider-databricks/internal/controller/databricks/mlflowmodel"
	modelserving "github.com/glalanne/provider-databricks/internal/controller/databricks/modelserving"
	mount "github.com/glalanne/provider-databricks/internal/controller/databricks/mount"
	mwscredentials "github.com/glalanne/provider-databricks/internal/controller/databricks/mwscredentials"
	mwscustomermanagedkeys "github.com/glalanne/provider-databricks/internal/controller/databricks/mwscustomermanagedkeys"
	mwslogdelivery "github.com/glalanne/provider-databricks/internal/controller/databricks/mwslogdelivery"
	mwsnccbinding "github.com/glalanne/provider-databricks/internal/controller/databricks/mwsnccbinding"
	mwsnccprivateendpointrule "github.com/glalanne/provider-databricks/internal/controller/databricks/mwsnccprivateendpointrule"
	mwsnetworkconnectivityconfig "github.com/glalanne/provider-databricks/internal/controller/databricks/mwsnetworkconnectivityconfig"
	mwsnetworks "github.com/glalanne/provider-databricks/internal/controller/databricks/mwsnetworks"
	mwspermissionassignment "github.com/glalanne/provider-databricks/internal/controller/databricks/mwspermissionassignment"
	mwsprivateaccesssettings "github.com/glalanne/provider-databricks/internal/controller/databricks/mwsprivateaccesssettings"
	mwsstorageconfigurations "github.com/glalanne/provider-databricks/internal/controller/databricks/mwsstorageconfigurations"
	mwsvpcendpoint "github.com/glalanne/provider-databricks/internal/controller/databricks/mwsvpcendpoint"
	mwsworkspaces "github.com/glalanne/provider-databricks/internal/controller/databricks/mwsworkspaces"
	notebook "github.com/glalanne/provider-databricks/internal/controller/databricks/notebook"
	notificationdestination "github.com/glalanne/provider-databricks/internal/controller/databricks/notificationdestination"
	obotoken "github.com/glalanne/provider-databricks/internal/controller/databricks/obotoken"
	onlinetable "github.com/glalanne/provider-databricks/internal/controller/databricks/onlinetable"
	permissionassignment "github.com/glalanne/provider-databricks/internal/controller/databricks/permissionassignment"
	permissions "github.com/glalanne/provider-databricks/internal/controller/databricks/permissions"
	pipeline "github.com/glalanne/provider-databricks/internal/controller/databricks/pipeline"
	provider "github.com/glalanne/provider-databricks/internal/controller/databricks/provider"
	qualitymonitor "github.com/glalanne/provider-databricks/internal/controller/databricks/qualitymonitor"
	query "github.com/glalanne/provider-databricks/internal/controller/databricks/query"
	recipient "github.com/glalanne/provider-databricks/internal/controller/databricks/recipient"
	registeredmodel "github.com/glalanne/provider-databricks/internal/controller/databricks/registeredmodel"
	repo "github.com/glalanne/provider-databricks/internal/controller/databricks/repo"
	restrictworkspaceadminssetting "github.com/glalanne/provider-databricks/internal/controller/databricks/restrictworkspaceadminssetting"
	schema "github.com/glalanne/provider-databricks/internal/controller/databricks/schema"
	secret "github.com/glalanne/provider-databricks/internal/controller/databricks/secret"
	secretacl "github.com/glalanne/provider-databricks/internal/controller/databricks/secretacl"
	secretscope "github.com/glalanne/provider-databricks/internal/controller/databricks/secretscope"
	serviceprincipal "github.com/glalanne/provider-databricks/internal/controller/databricks/serviceprincipal"
	serviceprincipalrole "github.com/glalanne/provider-databricks/internal/controller/databricks/serviceprincipalrole"
	serviceprincipalsecret "github.com/glalanne/provider-databricks/internal/controller/databricks/serviceprincipalsecret"
	share "github.com/glalanne/provider-databricks/internal/controller/databricks/share"
	sharepluginframework "github.com/glalanne/provider-databricks/internal/controller/databricks/sharepluginframework"
	sqlalert "github.com/glalanne/provider-databricks/internal/controller/databricks/sqlalert"
	sqldashboard "github.com/glalanne/provider-databricks/internal/controller/databricks/sqldashboard"
	sqlendpoint "github.com/glalanne/provider-databricks/internal/controller/databricks/sqlendpoint"
	sqlglobalconfig "github.com/glalanne/provider-databricks/internal/controller/databricks/sqlglobalconfig"
	sqlpermissions "github.com/glalanne/provider-databricks/internal/controller/databricks/sqlpermissions"
	sqlquery "github.com/glalanne/provider-databricks/internal/controller/databricks/sqlquery"
	sqltable "github.com/glalanne/provider-databricks/internal/controller/databricks/sqltable"
	sqlvisualization "github.com/glalanne/provider-databricks/internal/controller/databricks/sqlvisualization"
	sqlwidget "github.com/glalanne/provider-databricks/internal/controller/databricks/sqlwidget"
	storagecredential "github.com/glalanne/provider-databricks/internal/controller/databricks/storagecredential"
	systemschema "github.com/glalanne/provider-databricks/internal/controller/databricks/systemschema"
	table "github.com/glalanne/provider-databricks/internal/controller/databricks/table"
	token "github.com/glalanne/provider-databricks/internal/controller/databricks/token"
	user "github.com/glalanne/provider-databricks/internal/controller/databricks/user"
	userinstanceprofile "github.com/glalanne/provider-databricks/internal/controller/databricks/userinstanceprofile"
	userrole "github.com/glalanne/provider-databricks/internal/controller/databricks/userrole"
	vectorsearchendpoint "github.com/glalanne/provider-databricks/internal/controller/databricks/vectorsearchendpoint"
	vectorsearchindex "github.com/glalanne/provider-databricks/internal/controller/databricks/vectorsearchindex"
	volume "github.com/glalanne/provider-databricks/internal/controller/databricks/volume"
	workspacebinding "github.com/glalanne/provider-databricks/internal/controller/databricks/workspacebinding"
	workspaceconf "github.com/glalanne/provider-databricks/internal/controller/databricks/workspaceconf"
	workspacefile "github.com/glalanne/provider-databricks/internal/controller/databricks/workspacefile"
	providerconfig "github.com/glalanne/provider-databricks/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesscontrolruleset.Setup,
		alert.Setup,
		artifactallowlist.Setup,
		automaticclusterupdateworkspacesetting.Setup,
		awss3mount.Setup,
		azureadlsgen1mount.Setup,
		azureadlsgen2mount.Setup,
		azureblobmount.Setup,
		budget.Setup,
		catalog.Setup,
		catalogworkspacebinding.Setup,
		cluster.Setup,
		clusterpolicy.Setup,
		compliancesecurityprofileworkspacesetting.Setup,
		connection.Setup,
		customappintegration.Setup,
		dashboard.Setup,
		dbfsfile.Setup,
		defaultnamespacesetting.Setup,
		directory.Setup,
		enhancedsecuritymonitoringworkspacesetting.Setup,
		entitlements.Setup,
		externallocation.Setup,
		file.Setup,
		gitcredential.Setup,
		globalinitscript.Setup,
		grant.Setup,
		grants.Setup,
		group.Setup,
		groupinstanceprofile.Setup,
		groupmember.Setup,
		grouprole.Setup,
		instancepool.Setup,
		instanceprofile.Setup,
		ipaccesslist.Setup,
		job.Setup,
		lakehousemonitor.Setup,
		library.Setup,
		metastore.Setup,
		metastoreassignment.Setup,
		metastoredataaccess.Setup,
		mlflowexperiment.Setup,
		mlflowmodel.Setup,
		modelserving.Setup,
		mount.Setup,
		mwscredentials.Setup,
		mwscustomermanagedkeys.Setup,
		mwslogdelivery.Setup,
		mwsnccbinding.Setup,
		mwsnccprivateendpointrule.Setup,
		mwsnetworkconnectivityconfig.Setup,
		mwsnetworks.Setup,
		mwspermissionassignment.Setup,
		mwsprivateaccesssettings.Setup,
		mwsstorageconfigurations.Setup,
		mwsvpcendpoint.Setup,
		mwsworkspaces.Setup,
		notebook.Setup,
		notificationdestination.Setup,
		obotoken.Setup,
		onlinetable.Setup,
		permissionassignment.Setup,
		permissions.Setup,
		pipeline.Setup,
		provider.Setup,
		qualitymonitor.Setup,
		query.Setup,
		recipient.Setup,
		registeredmodel.Setup,
		repo.Setup,
		restrictworkspaceadminssetting.Setup,
		schema.Setup,
		secret.Setup,
		secretacl.Setup,
		secretscope.Setup,
		serviceprincipal.Setup,
		serviceprincipalrole.Setup,
		serviceprincipalsecret.Setup,
		share.Setup,
		sharepluginframework.Setup,
		sqlalert.Setup,
		sqldashboard.Setup,
		sqlendpoint.Setup,
		sqlglobalconfig.Setup,
		sqlpermissions.Setup,
		sqlquery.Setup,
		sqltable.Setup,
		sqlvisualization.Setup,
		sqlwidget.Setup,
		storagecredential.Setup,
		systemschema.Setup,
		table.Setup,
		token.Setup,
		user.Setup,
		userinstanceprofile.Setup,
		userrole.Setup,
		vectorsearchendpoint.Setup,
		vectorsearchindex.Setup,
		volume.Setup,
		workspacebinding.Setup,
		workspaceconf.Setup,
		workspacefile.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
