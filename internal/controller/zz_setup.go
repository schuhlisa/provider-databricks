// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	cluster "github.com/glalanne/provider-databricks/internal/controller/databricks/cluster"
	clusterpolicy "github.com/glalanne/provider-databricks/internal/controller/databricks/clusterpolicy"
	entitlements "github.com/glalanne/provider-databricks/internal/controller/databricks/entitlements"
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
	secret "github.com/glalanne/provider-databricks/internal/controller/databricks/secret"
	secretscope "github.com/glalanne/provider-databricks/internal/controller/databricks/secretscope"
	serviceprincipal "github.com/glalanne/provider-databricks/internal/controller/databricks/serviceprincipal"
	serviceprincipalrole "github.com/glalanne/provider-databricks/internal/controller/databricks/serviceprincipalrole"
	sqlendpoint "github.com/glalanne/provider-databricks/internal/controller/databricks/sqlendpoint"
	sqlpermissions "github.com/glalanne/provider-databricks/internal/controller/databricks/sqlpermissions"
	token "github.com/glalanne/provider-databricks/internal/controller/databricks/token"
	providerconfig "github.com/glalanne/provider-databricks/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		clusterpolicy.Setup,
		entitlements.Setup,
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
		secret.Setup,
		secretscope.Setup,
		serviceprincipal.Setup,
		serviceprincipalrole.Setup,
		sqlendpoint.Setup,
		sqlpermissions.Setup,
		token.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
