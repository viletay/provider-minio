// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	group "github.com/viletay/provider-minio/internal/controller/iam/group"
	groupmembership "github.com/viletay/provider-minio/internal/controller/iam/groupmembership"
	grouppolicy "github.com/viletay/provider-minio/internal/controller/iam/grouppolicy"
	grouppolicyattachment "github.com/viletay/provider-minio/internal/controller/iam/grouppolicyattachment"
	groupuserattachment "github.com/viletay/provider-minio/internal/controller/iam/groupuserattachment"
	policy "github.com/viletay/provider-minio/internal/controller/iam/policy"
	serviceaccount "github.com/viletay/provider-minio/internal/controller/iam/serviceaccount"
	user "github.com/viletay/provider-minio/internal/controller/iam/user"
	userpolicyattachment "github.com/viletay/provider-minio/internal/controller/iam/userpolicyattachment"
	policyilm "github.com/viletay/provider-minio/internal/controller/ilm/policy"
	key "github.com/viletay/provider-minio/internal/controller/kms/key"
	providerconfig "github.com/viletay/provider-minio/internal/controller/providerconfig"
	bucket "github.com/viletay/provider-minio/internal/controller/s3/bucket"
	bucketnotification "github.com/viletay/provider-minio/internal/controller/s3/bucketnotification"
	bucketpolicy "github.com/viletay/provider-minio/internal/controller/s3/bucketpolicy"
	bucketreplication "github.com/viletay/provider-minio/internal/controller/s3/bucketreplication"
	bucketserversideencryption "github.com/viletay/provider-minio/internal/controller/s3/bucketserversideencryption"
	bucketversioning "github.com/viletay/provider-minio/internal/controller/s3/bucketversioning"
	object "github.com/viletay/provider-minio/internal/controller/s3/object"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		group.Setup,
		groupmembership.Setup,
		grouppolicy.Setup,
		grouppolicyattachment.Setup,
		groupuserattachment.Setup,
		policy.Setup,
		serviceaccount.Setup,
		user.Setup,
		userpolicyattachment.Setup,
		policyilm.Setup,
		key.Setup,
		providerconfig.Setup,
		bucket.Setup,
		bucketnotification.Setup,
		bucketpolicy.Setup,
		bucketreplication.Setup,
		bucketserversideencryption.Setup,
		bucketversioning.Setup,
		object.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
