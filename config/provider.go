/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/viletay/provider-minio/config/iam"
	"github.com/viletay/provider-minio/config/ilm"
	"github.com/viletay/provider-minio/config/kms"
	"github.com/viletay/provider-minio/config/s3"
)

const (
	resourcePrefix = "minio"
	modulePath     = "github.com/viletay/provider-minio"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration

func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("minio.upbound.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		iam.Configure,
		ilm.Configure,
		kms.Configure,
		s3.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
