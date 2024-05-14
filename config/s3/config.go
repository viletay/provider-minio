package s3

import ujconfig "github.com/crossplane/upjet/pkg/config"

const (
	// Group is the short group for this provider.
	Group = "s3"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("minio_s3_bucket", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.Kind = "Bucket"
		r.ExternalName = ujconfig.IdentifierFromProvider
	})
	p.AddResourceConfigurator("minio_s3_bucket_notification", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.Kind = "BucketNotification"
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["bucket"] = ujconfig.Reference{
			Type: "Bucket",
		}
	})
	p.AddResourceConfigurator("minio_s3_bucket_policy", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.Kind = "BucketPolicy"
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["bucket"] = ujconfig.Reference{
			Type: "Bucket",
		}
		r.References["policy"] = ujconfig.Reference{
			Type: "github.com/viletay/provider-minio/apis/iam/v1alpha1.Policy",
		}
	})
	p.AddResourceConfigurator("minio_s3_bucket_replication", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.Kind = "BucketReplication"
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["bucket"] = ujconfig.Reference{
			Type: "Bucket",
		}
	})
	p.AddResourceConfigurator("minio_s3_bucket_server_side_encryption", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.Kind = "BucketServerSideEncryption"
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["bucket"] = ujconfig.Reference{
			Type: "Bucket",
		}
		r.References["kms_key_id"] = ujconfig.Reference{
			Type: "github.com/viletay/provider-minio/apis/kms/v1alpha1.Key",
		}
	})
	p.AddResourceConfigurator("minio_s3_bucket_versioning", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.Kind = "BucketVersioning"
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["bucket"] = ujconfig.Reference{
			Type: "Bucket",
		}
	})
	p.AddResourceConfigurator("minio_s3_object", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.Kind = "Object"
		r.ExternalName = ujconfig.IdentifierFromProvider
		r.References["bucket_name"] = ujconfig.Reference{
			Type: "Bucket",
		}
	})
}
