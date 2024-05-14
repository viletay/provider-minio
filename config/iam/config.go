package iam

import ujconfig "github.com/crossplane/upjet/pkg/config"

const (
	// Group is the short group for this provider.
	Group = "iam"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("minio_iam_group", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.Kind = "Group"
		r.ExternalName = ujconfig.IdentifierFromProvider
	})
	p.AddResourceConfigurator("minio_iam_group_membership", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.Kind = "GroupMembership"
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["group"] = ujconfig.Reference{
			Type: "Group",
		}
	})
	p.AddResourceConfigurator("minio_iam_group_policy", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.Kind = "GroupPolicy"
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["group"] = ujconfig.Reference{
			Type: "Group",
		}
	})
	p.AddResourceConfigurator("minio_iam_group_policy_attachment", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.Kind = "GroupPolicyAttachment"
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["group_name"] = ujconfig.Reference{
			Type: "Group",
		}
		r.References["policy_name"] = ujconfig.Reference{
			Type: "GroupPolicy",
		}
	})
	p.AddResourceConfigurator("minio_iam_group_user_attachment", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.Kind = "GroupUserAttachment"
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["group_name"] = ujconfig.Reference{
			Type: "Group",
		}
		r.References["user_name"] = ujconfig.Reference{
			Type: "User",
		}
	})
	p.AddResourceConfigurator("minio_iam_policy", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.Kind = "Policy"
		r.ExternalName = ujconfig.IdentifierFromProvider
	})
	p.AddResourceConfigurator("minio_iam_service_account", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.Kind = "ServiceAccount"
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["target_name"] = ujconfig.Reference{
			Type: "User",
		}
	})
	p.AddResourceConfigurator("minio_iam_user", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.Kind = "User"
		r.ExternalName = ujconfig.IdentifierFromProvider
	})
	p.AddResourceConfigurator("minio_iam_user_policy_attachment", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.Kind = "UserPolicyAttachment"
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["user_name"] = ujconfig.Reference{
			Type: "User",
		}
		r.References["policy_name"] = ujconfig.Reference{
			Type: "User",
		}
	})
}
