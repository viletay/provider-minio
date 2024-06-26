/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this GroupMembership.
func (mg *GroupMembership) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Group),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.GroupRef,
		Selector:     mg.Spec.ForProvider.GroupSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Group")
	}
	mg.Spec.ForProvider.Group = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GroupRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this GroupPolicy.
func (mg *GroupPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Group),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.GroupRef,
		Selector:     mg.Spec.ForProvider.GroupSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Group")
	}
	mg.Spec.ForProvider.Group = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GroupRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this GroupPolicyAttachment.
func (mg *GroupPolicyAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.GroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.GroupNameRef,
		Selector:     mg.Spec.ForProvider.GroupNameSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GroupName")
	}
	mg.Spec.ForProvider.GroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PolicyName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PolicyNameRef,
		Selector:     mg.Spec.ForProvider.PolicyNameSelector,
		To: reference.To{
			List:    &GroupPolicyList{},
			Managed: &GroupPolicy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PolicyName")
	}
	mg.Spec.ForProvider.PolicyName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PolicyNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this GroupUserAttachment.
func (mg *GroupUserAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.GroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.GroupNameRef,
		Selector:     mg.Spec.ForProvider.GroupNameSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GroupName")
	}
	mg.Spec.ForProvider.GroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.UserName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.UserNameRef,
		Selector:     mg.Spec.ForProvider.UserNameSelector,
		To: reference.To{
			List:    &UserList{},
			Managed: &User{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.UserName")
	}
	mg.Spec.ForProvider.UserName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this UserPolicyAttachment.
func (mg *UserPolicyAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PolicyName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PolicyNameRef,
		Selector:     mg.Spec.ForProvider.PolicyNameSelector,
		To: reference.To{
			List:    &UserList{},
			Managed: &User{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PolicyName")
	}
	mg.Spec.ForProvider.PolicyName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PolicyNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.UserName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.UserNameRef,
		Selector:     mg.Spec.ForProvider.UserNameSelector,
		To: reference.To{
			List:    &UserList{},
			Managed: &User{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.UserName")
	}
	mg.Spec.ForProvider.UserName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserNameRef = rsp.ResolvedReference

	return nil
}
