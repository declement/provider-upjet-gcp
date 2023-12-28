/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Entry.
func (mg *Entry) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EntryGroup),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.EntryGroupRef,
		Selector:     mg.Spec.ForProvider.EntryGroupSelector,
		To: reference.To{
			List:    &EntryGroupList{},
			Managed: &EntryGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EntryGroup")
	}
	mg.Spec.ForProvider.EntryGroup = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EntryGroupRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EntryGroup),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.EntryGroupRef,
		Selector:     mg.Spec.InitProvider.EntryGroupSelector,
		To: reference.To{
			List:    &EntryGroupList{},
			Managed: &EntryGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.EntryGroup")
	}
	mg.Spec.InitProvider.EntryGroup = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.EntryGroupRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Tag.
func (mg *Tag) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Parent),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ParentRef,
		Selector:     mg.Spec.ForProvider.ParentSelector,
		To: reference.To{
			List:    &EntryList{},
			Managed: &Entry{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Parent")
	}
	mg.Spec.ForProvider.Parent = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ParentRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Template),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.TemplateRef,
		Selector:     mg.Spec.ForProvider.TemplateSelector,
		To: reference.To{
			List:    &TagTemplateList{},
			Managed: &TagTemplate{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Template")
	}
	mg.Spec.ForProvider.Template = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TemplateRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Parent),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.ParentRef,
		Selector:     mg.Spec.InitProvider.ParentSelector,
		To: reference.To{
			List:    &EntryList{},
			Managed: &Entry{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Parent")
	}
	mg.Spec.InitProvider.Parent = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ParentRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Template),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.TemplateRef,
		Selector:     mg.Spec.InitProvider.TemplateSelector,
		To: reference.To{
			List:    &TagTemplateList{},
			Managed: &TagTemplate{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Template")
	}
	mg.Spec.InitProvider.Template = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TemplateRef = rsp.ResolvedReference

	return nil
}
