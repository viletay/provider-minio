package ilm

import ujconfig "github.com/crossplane/upjet/pkg/config"

const (
	// Group is the short group for this provider.
	Group = "ilm"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("minio_ilm_policy", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.Kind = "Policy"
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["bucket"] = ujconfig.Reference{
			Type: "github.com/viletay/provider-minio/apis/s3/v1alpha1.Bucket",
		}
	})
}
