// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Permissions.
func (mg *Permissions) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterPolicyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterPolicyIDRef,
		Selector:     mg.Spec.ForProvider.ClusterPolicyIDSelector,
		To: reference.To{
			List:    &ClusterPolicyList{},
			Managed: &ClusterPolicy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterPolicyID")
	}
	mg.Spec.ForProvider.ClusterPolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterPolicyIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstancePoolID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.InstancePoolIDRef,
		Selector:     mg.Spec.ForProvider.InstancePoolIDSelector,
		To: reference.To{
			List:    &InstancePoolList{},
			Managed: &InstancePool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstancePoolID")
	}
	mg.Spec.ForProvider.InstancePoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstancePoolIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NotebookID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NotebookIDRef,
		Selector:     mg.Spec.ForProvider.NotebookIDSelector,
		To: reference.To{
			List:    &NotebookList{},
			Managed: &Notebook{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NotebookID")
	}
	mg.Spec.ForProvider.NotebookID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NotebookIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PipelineID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PipelineIDRef,
		Selector:     mg.Spec.ForProvider.PipelineIDSelector,
		To: reference.To{
			List:    &PipelineList{},
			Managed: &Pipeline{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PipelineID")
	}
	mg.Spec.ForProvider.PipelineID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PipelineIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SQLAlertID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SQLAlertIDRef,
		Selector:     mg.Spec.ForProvider.SQLAlertIDSelector,
		To: reference.To{
			List:    &SQLAlertList{},
			Managed: &SQLAlert{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SQLAlertID")
	}
	mg.Spec.ForProvider.SQLAlertID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SQLAlertIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SQLDashboardID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SQLDashboardIDRef,
		Selector:     mg.Spec.ForProvider.SQLDashboardIDSelector,
		To: reference.To{
			List:    &SQLDashboardList{},
			Managed: &SQLDashboard{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SQLDashboardID")
	}
	mg.Spec.ForProvider.SQLDashboardID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SQLDashboardIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SQLEndpointID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SQLEndpointIDRef,
		Selector:     mg.Spec.ForProvider.SQLEndpointIDSelector,
		To: reference.To{
			List:    &SQLEndpointList{},
			Managed: &SQLEndpoint{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SQLEndpointID")
	}
	mg.Spec.ForProvider.SQLEndpointID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SQLEndpointIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SQLQueryID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SQLQueryIDRef,
		Selector:     mg.Spec.ForProvider.SQLQueryIDSelector,
		To: reference.To{
			List:    &SQLQueryList{},
			Managed: &SQLQuery{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SQLQueryID")
	}
	mg.Spec.ForProvider.SQLQueryID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SQLQueryIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ClusterIDRef,
		Selector:     mg.Spec.InitProvider.ClusterIDSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClusterID")
	}
	mg.Spec.InitProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClusterIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClusterPolicyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ClusterPolicyIDRef,
		Selector:     mg.Spec.InitProvider.ClusterPolicyIDSelector,
		To: reference.To{
			List:    &ClusterPolicyList{},
			Managed: &ClusterPolicy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClusterPolicyID")
	}
	mg.Spec.InitProvider.ClusterPolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClusterPolicyIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstancePoolID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.InstancePoolIDRef,
		Selector:     mg.Spec.InitProvider.InstancePoolIDSelector,
		To: reference.To{
			List:    &InstancePoolList{},
			Managed: &InstancePool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstancePoolID")
	}
	mg.Spec.InitProvider.InstancePoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstancePoolIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NotebookID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.NotebookIDRef,
		Selector:     mg.Spec.InitProvider.NotebookIDSelector,
		To: reference.To{
			List:    &NotebookList{},
			Managed: &Notebook{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.NotebookID")
	}
	mg.Spec.InitProvider.NotebookID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NotebookIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PipelineID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.PipelineIDRef,
		Selector:     mg.Spec.InitProvider.PipelineIDSelector,
		To: reference.To{
			List:    &PipelineList{},
			Managed: &Pipeline{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PipelineID")
	}
	mg.Spec.InitProvider.PipelineID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PipelineIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SQLAlertID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.SQLAlertIDRef,
		Selector:     mg.Spec.InitProvider.SQLAlertIDSelector,
		To: reference.To{
			List:    &SQLAlertList{},
			Managed: &SQLAlert{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SQLAlertID")
	}
	mg.Spec.InitProvider.SQLAlertID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SQLAlertIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SQLDashboardID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.SQLDashboardIDRef,
		Selector:     mg.Spec.InitProvider.SQLDashboardIDSelector,
		To: reference.To{
			List:    &SQLDashboardList{},
			Managed: &SQLDashboard{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SQLDashboardID")
	}
	mg.Spec.InitProvider.SQLDashboardID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SQLDashboardIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SQLEndpointID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.SQLEndpointIDRef,
		Selector:     mg.Spec.InitProvider.SQLEndpointIDSelector,
		To: reference.To{
			List:    &SQLEndpointList{},
			Managed: &SQLEndpoint{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SQLEndpointID")
	}
	mg.Spec.InitProvider.SQLEndpointID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SQLEndpointIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SQLQueryID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.SQLQueryIDRef,
		Selector:     mg.Spec.InitProvider.SQLQueryIDSelector,
		To: reference.To{
			List:    &SQLQueryList{},
			Managed: &SQLQuery{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SQLQueryID")
	}
	mg.Spec.InitProvider.SQLQueryID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SQLQueryIDRef = rsp.ResolvedReference

	return nil
}
