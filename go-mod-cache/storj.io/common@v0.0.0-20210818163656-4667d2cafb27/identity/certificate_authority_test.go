// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package identity_test

import (
	"context"
	"crypto/x509/pkix"
	"encoding/asn1"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"storj.io/common/identity"
	"storj.io/common/identity/testidentity"
	"storj.io/common/peertls/extensions"
	"storj.io/common/peertls/tlsopts"
	"storj.io/common/storj"
	"storj.io/common/testcontext"
	"storj.io/common/testrand"
)

func TestNewCA(t *testing.T) {
	const expectedDifficulty = 4

	for _, version := range storj.IDVersions {
		ca, err := identity.NewCA(context.Background(), identity.NewCAOptions{
			VersionNumber: version.Number,
			Difficulty:    expectedDifficulty,
			Concurrency:   4,
		})
		require.NoError(t, err)
		require.NotEmpty(t, ca)

		assert.Equal(t, version.Number, ca.ID.Version().Number)

		caVersion, err := ca.Version()
		require.NoError(t, err)
		assert.Equal(t, version.Number, caVersion.Number)

		actualDifficulty, err := ca.ID.Difficulty()
		require.NoError(t, err)
		assert.True(t, actualDifficulty >= expectedDifficulty)
	}
}

func TestFullCertificateAuthority_NewIdentity(t *testing.T) {
	ctx := testcontext.New(t)

	ca, err := identity.NewCA(ctx, identity.NewCAOptions{
		Difficulty:  12,
		Concurrency: 4,
	})
	require.NoError(t, err)
	require.NotNil(t, ca)

	fi, err := ca.NewIdentity()
	require.NoError(t, err)
	require.NotNil(t, fi)

	assert.Equal(t, ca.Cert, fi.CA)
	assert.Equal(t, ca.ID, fi.ID)
	assert.NotEqual(t, ca.Key, fi.Key)
	assert.NotEqual(t, ca.Cert, fi.Leaf)

	err = fi.Leaf.CheckSignatureFrom(ca.Cert)
	assert.NoError(t, err)
}

func TestFullCertificateAuthority_Sign(t *testing.T) {
	ctx := testcontext.New(t)

	caOpts := identity.NewCAOptions{
		Difficulty:  12,
		Concurrency: 4,
	}

	ca, err := identity.NewCA(ctx, caOpts)
	require.NoError(t, err)
	require.NotNil(t, ca)

	toSign, err := identity.NewCA(ctx, caOpts)
	require.NoError(t, err)
	require.NotNil(t, toSign)

	signed, err := ca.Sign(toSign.Cert)
	require.NoError(t, err)
	require.NotNil(t, signed)

	assert.Equal(t, toSign.Cert.RawTBSCertificate, signed.RawTBSCertificate)
	assert.NotEqual(t, toSign.Cert.Signature, signed.Signature)
	assert.NotEqual(t, toSign.Cert.Raw, signed.Raw)

	err = signed.CheckSignatureFrom(ca.Cert)
	assert.NoError(t, err)
}

func TestFullCAConfig_Save(t *testing.T) {
	// TODO(bryanchriswhite): test with both
	// TODO(bryanchriswhite): test with only cert path
	// TODO(bryanchriswhite): test with only key path
	t.SkipNow()
}

func TestFullCAConfig_Load_extensions(t *testing.T) {
	ctx := testcontext.New(t)

	for versionNumber, version := range storj.IDVersions {
		caCfg := identity.CASetupConfig{
			VersionNumber: uint(versionNumber),
			CertPath:      ctx.File("ca.cert"),
			KeyPath:       ctx.File("ca.key"),
		}

		{
			ca, err := caCfg.Create(ctx, nil)
			require.NoError(t, err)

			caVersion, err := ca.Version()
			require.NoError(t, err)
			require.Equal(t, version.Number, caVersion.Number)
		}

		{
			ca, err := caCfg.FullConfig().Load()
			require.NoError(t, err)
			caVersion, err := ca.Version()
			require.NoError(t, err)
			assert.Equal(t, version.Number, caVersion.Number)
		}

	}
}

func BenchmarkNewCA(b *testing.B) {
	ctx := context.Background()
	for _, difficulty := range []uint16{8, 12} {
		testDifficulty := difficulty
		for _, testConcurrency := range []uint{1, 2, 5, 10} {
			concurrency := testConcurrency
			test := fmt.Sprintf("%d/%d", testDifficulty, concurrency)
			b.Run(test, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = identity.NewCA(ctx, identity.NewCAOptions{
						Difficulty:  testDifficulty,
						Concurrency: concurrency,
					})
				}
			})
		}
	}
}

func TestFullCertificateAuthority_AddExtension(t *testing.T) {
	ctx := testcontext.New(t)

	ca, err := testidentity.NewTestCA(ctx)
	require.NoError(t, err)

	oldCert := ca.Cert
	assert.Len(t, ca.Cert.ExtraExtensions, 0)

	randBytes := testrand.Bytes(10)
	randExt := pkix.Extension{
		Id:    asn1.ObjectIdentifier{2, 999, int(randBytes[0])},
		Value: randBytes,
	}

	err = ca.AddExtension(randExt)
	require.NoError(t, err)

	assert.Len(t, ca.Cert.ExtraExtensions, 0)
	assert.Len(t, ca.Cert.Extensions, len(oldCert.Extensions)+1)

	assert.Equal(t, oldCert.SerialNumber, ca.Cert.SerialNumber)
	assert.Equal(t, oldCert.IsCA, ca.Cert.IsCA)
	assert.Equal(t, oldCert.PublicKey, ca.Cert.PublicKey)
	assert.Equal(t, randExt, tlsopts.NewExtensionsMap(ca.Cert)[randExt.Id.String()])

	assert.NotEqual(t, oldCert.Raw, ca.Cert.Raw)
	assert.NotEqual(t, oldCert.RawTBSCertificate, ca.Cert.RawTBSCertificate)
	assert.NotEqual(t, oldCert.Signature, ca.Cert.Signature)
}

func TestFullCertificateAuthority_Revoke(t *testing.T) {
	ctx := testcontext.New(t)

	ca, err := testidentity.NewTestCA(ctx)
	require.NoError(t, err)

	oldCert := ca.Cert
	assert.Len(t, ca.Cert.ExtraExtensions, 0)

	err = ca.Revoke()
	require.NoError(t, err)

	assert.Len(t, ca.Cert.ExtraExtensions, 0)
	assert.Len(t, ca.Cert.Extensions, len(oldCert.Extensions)+1)

	assert.Equal(t, oldCert.SerialNumber, ca.Cert.SerialNumber)
	assert.Equal(t, oldCert.IsCA, ca.Cert.IsCA)
	assert.Equal(t, oldCert.PublicKey, ca.Cert.PublicKey)

	assert.NotEqual(t, oldCert.Raw, ca.Cert.Raw)
	assert.NotEqual(t, oldCert.RawTBSCertificate, ca.Cert.RawTBSCertificate)
	assert.NotEqual(t, oldCert.Signature, ca.Cert.Signature)

	revocationExt := tlsopts.NewExtensionsMap(ca.Cert)[extensions.RevocationExtID.String()]
	assert.True(t, extensions.RevocationExtID.Equal(revocationExt.Id))

	var rev extensions.Revocation
	err = rev.Unmarshal(revocationExt.Value)
	require.NoError(t, err)

	err = rev.Verify(ca.Cert)
	assert.NoError(t, err)
}
