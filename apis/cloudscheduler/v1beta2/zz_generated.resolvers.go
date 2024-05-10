// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Job.
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
)

func (mg *Job) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	if mg.Spec.ForProvider.PubsubTarget != nil {
		{
			m, l, err = apisresolver.GetManagedResource("pubsub.gcp.upbound.io", "v1beta2", "Topic", "TopicList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PubsubTarget.TopicName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.PubsubTarget.TopicNameRef,
				Selector:     mg.Spec.ForProvider.PubsubTarget.TopicNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.PubsubTarget.TopicName")
		}
		mg.Spec.ForProvider.PubsubTarget.TopicName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.PubsubTarget.TopicNameRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.PubsubTarget != nil {
		{
			m, l, err = apisresolver.GetManagedResource("pubsub.gcp.upbound.io", "v1beta2", "Topic", "TopicList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PubsubTarget.TopicName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.PubsubTarget.TopicNameRef,
				Selector:     mg.Spec.InitProvider.PubsubTarget.TopicNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.PubsubTarget.TopicName")
		}
		mg.Spec.InitProvider.PubsubTarget.TopicName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.PubsubTarget.TopicNameRef = rsp.ResolvedReference

	}

	return nil
}
