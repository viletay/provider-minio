package kms

import ujconfig "github.com/crossplane/upjet/pkg/config"

const (
	// Group is the short group for this provider.
	Group = "kms"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("minio_kms_key", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.Kind = "Key"
		r.ExternalName = ujconfig.IdentifierFromProvider
	})
}
