// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	alert "github.com/glalanne/provider-databricks/internal/controller/databricks/alert"
	budget "github.com/glalanne/provider-databricks/internal/controller/databricks/budget"
	catalog "github.com/glalanne/provider-databricks/internal/controller/databricks/catalog"
	cluster "github.com/glalanne/provider-databricks/internal/controller/databricks/cluster"
	clusterpolicy "github.com/glalanne/provider-databricks/internal/controller/databricks/clusterpolicy"
	connection "github.com/glalanne/provider-databricks/internal/controller/databricks/connection"
	entitlements "github.com/glalanne/provider-databricks/internal/controller/databricks/entitlements"
	externallocation "github.com/glalanne/provider-databricks/internal/controller/databricks/externallocation"
	gitcredential "github.com/glalanne/provider-databricks/internal/controller/databricks/gitcredential"
	grants "github.com/glalanne/provider-databricks/internal/controller/databricks/grants"
	group "github.com/glalanne/provider-databricks/internal/controller/databricks/group"
	groupmember "github.com/glalanne/provider-databricks/internal/controller/databricks/groupmember"
	grouprole "github.com/glalanne/provider-databricks/internal/controller/databricks/grouprole"
	instancepool "github.com/glalanne/provider-databricks/internal/controller/databricks/instancepool"
	ipaccesslist "github.com/glalanne/provider-databricks/internal/controller/databricks/ipaccesslist"
	job "github.com/glalanne/provider-databricks/internal/controller/databricks/job"
	notebook "github.com/glalanne/provider-databricks/internal/controller/databricks/notebook"
	permissionassignment "github.com/glalanne/provider-databricks/internal/controller/databricks/permissionassignment"
	permissions "github.com/glalanne/provider-databricks/internal/controller/databricks/permissions"
	pipeline "github.com/glalanne/provider-databricks/internal/controller/databricks/pipeline"
	query "github.com/glalanne/provider-databricks/internal/controller/databricks/query"
	schema "github.com/glalanne/provider-databricks/internal/controller/databricks/schema"
	secret "github.com/glalanne/provider-databricks/internal/controller/databricks/secret"
	secretscope "github.com/glalanne/provider-databricks/internal/controller/databricks/secretscope"
	serviceprincipal "github.com/glalanne/provider-databricks/internal/controller/databricks/serviceprincipal"
	serviceprincipalrole "github.com/glalanne/provider-databricks/internal/controller/databricks/serviceprincipalrole"
	sqlalert "github.com/glalanne/provider-databricks/internal/controller/databricks/sqlalert"
	sqldashboard "github.com/glalanne/provider-databricks/internal/controller/databricks/sqldashboard"
	sqlendpoint "github.com/glalanne/provider-databricks/internal/controller/databricks/sqlendpoint"
	sqlglobalconfig "github.com/glalanne/provider-databricks/internal/controller/databricks/sqlglobalconfig"
	sqlpermissions "github.com/glalanne/provider-databricks/internal/controller/databricks/sqlpermissions"
	sqlquery "github.com/glalanne/provider-databricks/internal/controller/databricks/sqlquery"
	sqltable "github.com/glalanne/provider-databricks/internal/controller/databricks/sqltable"
	token "github.com/glalanne/provider-databricks/internal/controller/databricks/token"
	providerconfig "github.com/glalanne/provider-databricks/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alert.Setup,
		budget.Setup,
		catalog.Setup,
		cluster.Setup,
		clusterpolicy.Setup,
		connection.Setup,
		entitlements.Setup,
		externallocation.Setup,
		gitcredential.Setup,
		grants.Setup,
		group.Setup,
		groupmember.Setup,
		grouprole.Setup,
		instancepool.Setup,
		ipaccesslist.Setup,
		job.Setup,
		notebook.Setup,
		permissionassignment.Setup,
		permissions.Setup,
		pipeline.Setup,
		query.Setup,
		schema.Setup,
		secret.Setup,
		secretscope.Setup,
		serviceprincipal.Setup,
		serviceprincipalrole.Setup,
		sqlalert.Setup,
		sqldashboard.Setup,
		sqlendpoint.Setup,
		sqlglobalconfig.Setup,
		sqlpermissions.Setup,
		sqlquery.Setup,
		sqltable.Setup,
		token.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
