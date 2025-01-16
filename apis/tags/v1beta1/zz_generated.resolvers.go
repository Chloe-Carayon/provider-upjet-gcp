// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	errors "github.com/pkg/errors"
	common "github.com/upbound/provider-gcp/config/common"
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *LocationTagBinding) ResolveReferences( // ResolveReferences of this LocationTagBinding.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("tags.gcp.upbound.io", "v1beta1", "TagValue", "TagValueList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TagValue),
			Extract:      common.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.TagValueRef,
			Selector:     mg.Spec.ForProvider.TagValueSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TagValue")
	}
	mg.Spec.ForProvider.TagValue = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TagValueRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("tags.gcp.upbound.io", "v1beta1", "TagValue", "TagValueList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.TagValue),
			Extract:      common.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.TagValueRef,
			Selector:     mg.Spec.InitProvider.TagValueSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.TagValue")
	}
	mg.Spec.InitProvider.TagValue = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TagValueRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TagBinding.
func (mg *TagBinding) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("tags.gcp.upbound.io", "v1beta1", "TagValue", "TagValueList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TagValue),
			Extract:      common.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.TagValueRef,
			Selector:     mg.Spec.ForProvider.TagValueSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TagValue")
	}
	mg.Spec.ForProvider.TagValue = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TagValueRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("tags.gcp.upbound.io", "v1beta1", "TagValue", "TagValueList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.TagValue),
			Extract:      common.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.TagValueRef,
			Selector:     mg.Spec.InitProvider.TagValueSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.TagValue")
	}
	mg.Spec.InitProvider.TagValue = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TagValueRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TagValue.
func (mg *TagValue) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("tags.gcp.upbound.io", "v1beta1", "TagKey", "TagKeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Parent),
			Extract:      common.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.ParentRef,
			Selector:     mg.Spec.ForProvider.ParentSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Parent")
	}
	mg.Spec.ForProvider.Parent = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ParentRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("tags.gcp.upbound.io", "v1beta1", "TagKey", "TagKeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Parent),
			Extract:      common.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.ParentRef,
			Selector:     mg.Spec.InitProvider.ParentSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Parent")
	}
	mg.Spec.InitProvider.Parent = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ParentRef = rsp.ResolvedReference

	return nil
}
