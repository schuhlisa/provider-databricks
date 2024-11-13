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
	group "github.com/glalanne/provider-databricks/internal/controller/databricks/group"
	groupmember "github.com/glalanne/provider-databricks/internal/controller/databricks/groupmember"
	instancepool "github.com/glalanne/provider-databricks/internal/controller/databricks/instancepool"
	job "github.com/glalanne/provider-databricks/internal/controller/databricks/job"
	notebook "github.com/glalanne/provider-databricks/internal/controller/databricks/notebook"
	permissions "github.com/glalanne/provider-databricks/internal/controller/databricks/permissions"
	secret "github.com/glalanne/provider-databricks/internal/controller/databricks/secret"
	secretscope "github.com/glalanne/provider-databricks/internal/controller/databricks/secretscope"
	sqlendpoint "github.com/glalanne/provider-databricks/internal/controller/databricks/sqlendpoint"
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
		group.Setup,
		groupmember.Setup,
		instancepool.Setup,
		job.Setup,
		notebook.Setup,
		permissions.Setup,
		secret.Setup,
		secretscope.Setup,
		sqlendpoint.Setup,
		token.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
