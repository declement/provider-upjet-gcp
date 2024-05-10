// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	common "github.com/upbound/provider-gcp/config/common"
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *SecretIAMMember) ResolveReferences( // ResolveReferences of this SecretIAMMember.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("secretmanager.gcp.upbound.io", "v1beta1", "Secret", "SecretList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SecretID),
			Extract:      common.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.SecretIDRef,
			Selector:     mg.Spec.ForProvider.SecretIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecretID")
	}
	mg.Spec.ForProvider.SecretID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SecretIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("secretmanager.gcp.upbound.io", "v1beta1", "Secret", "SecretList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SecretID),
			Extract:      common.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.SecretIDRef,
			Selector:     mg.Spec.InitProvider.SecretIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SecretID")
	}
	mg.Spec.InitProvider.SecretID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SecretIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecretVersion.
func (mg *SecretVersion) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("secretmanager.gcp.upbound.io", "v1beta2", "Secret", "SecretList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Secret),
			Extract:      common.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.SecretRef,
			Selector:     mg.Spec.ForProvider.SecretSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Secret")
	}
	mg.Spec.ForProvider.Secret = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SecretRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("secretmanager.gcp.upbound.io", "v1beta2", "Secret", "SecretList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Secret),
			Extract:      common.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.SecretRef,
			Selector:     mg.Spec.InitProvider.SecretSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Secret")
	}
	mg.Spec.InitProvider.Secret = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SecretRef = rsp.ResolvedReference

	return nil
}
