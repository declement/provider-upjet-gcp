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
	v1beta1 "github.com/upbound/provider-gcp/apis/compute/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Spoke.
func (mg *Spoke) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Hub),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.HubRef,
		Selector:     mg.Spec.ForProvider.HubSelector,
		To: reference.To{
			List:    &HubList{},
			Managed: &Hub{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Hub")
	}
	mg.Spec.ForProvider.Hub = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.HubRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.LinkedRouterApplianceInstances); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.LinkedRouterApplianceInstances[i3].Instances); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LinkedRouterApplianceInstances[i3].Instances[i4].VirtualMachine),
				Extract:      resource.ExtractParamPath("self_link", true),
				Reference:    mg.Spec.ForProvider.LinkedRouterApplianceInstances[i3].Instances[i4].VirtualMachineRef,
				Selector:     mg.Spec.ForProvider.LinkedRouterApplianceInstances[i3].Instances[i4].VirtualMachineSelector,
				To: reference.To{
					List:    &v1beta1.InstanceList{},
					Managed: &v1beta1.Instance{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.LinkedRouterApplianceInstances[i3].Instances[i4].VirtualMachine")
			}
			mg.Spec.ForProvider.LinkedRouterApplianceInstances[i3].Instances[i4].VirtualMachine = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.LinkedRouterApplianceInstances[i3].Instances[i4].VirtualMachineRef = rsp.ResolvedReference

		}
	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Hub),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.HubRef,
		Selector:     mg.Spec.InitProvider.HubSelector,
		To: reference.To{
			List:    &HubList{},
			Managed: &Hub{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Hub")
	}
	mg.Spec.InitProvider.Hub = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.HubRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.LinkedRouterApplianceInstances); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.LinkedRouterApplianceInstances[i3].Instances); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LinkedRouterApplianceInstances[i3].Instances[i4].VirtualMachine),
				Extract:      resource.ExtractParamPath("self_link", true),
				Reference:    mg.Spec.InitProvider.LinkedRouterApplianceInstances[i3].Instances[i4].VirtualMachineRef,
				Selector:     mg.Spec.InitProvider.LinkedRouterApplianceInstances[i3].Instances[i4].VirtualMachineSelector,
				To: reference.To{
					List:    &v1beta1.InstanceList{},
					Managed: &v1beta1.Instance{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.LinkedRouterApplianceInstances[i3].Instances[i4].VirtualMachine")
			}
			mg.Spec.InitProvider.LinkedRouterApplianceInstances[i3].Instances[i4].VirtualMachine = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.LinkedRouterApplianceInstances[i3].Instances[i4].VirtualMachineRef = rsp.ResolvedReference

		}
	}

	return nil
}
