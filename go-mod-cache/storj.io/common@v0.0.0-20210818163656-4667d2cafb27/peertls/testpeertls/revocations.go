// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package testpeertls

import (
	"crypto"
	"crypto/x509"
	"crypto/x509/pkix"

	"storj.io/common/identity"
	"storj.io/common/peertls"
	"storj.io/common/peertls/extensions"
	"storj.io/common/peertls/tlsopts"
	"storj.io/common/storj"
)

// RevokeLeaf revokes the leaf certificate in the passed chain and replaces it
// with a "revoking" certificate, which contains a revocation extension recording
// this action.
func RevokeLeaf(caKey crypto.PrivateKey, chain []*x509.Certificate) ([]*x509.Certificate, pkix.Extension, error) {
	if len(chain) < 2 {
		return nil, pkix.Extension{}, extensions.Error.New("revoking leaf implies a CA exists; chain too short")
	}
	ca := &identity.FullCertificateAuthority{
		Key:       caKey,
		Cert:      chain[peertls.CAIndex],
		RestChain: chain[:peertls.CAIndex+1],
	}

	var err error
	ca.ID, err = identity.NodeIDFromKey(ca.Cert.PublicKey, storj.LatestIDVersion())
	if err != nil {
		return nil, pkix.Extension{}, err
	}

	ident := &identity.FullIdentity{
		Leaf:      chain[peertls.LeafIndex],
		CA:        ca.Cert,
		ID:        ca.ID,
		RestChain: ca.RestChain,
	}

	manageableIdent := identity.NewManageableFullIdentity(ident, ca)
	if err := manageableIdent.Revoke(); err != nil {
		return nil, pkix.Extension{}, err
	}

	revokingCert := manageableIdent.Leaf
	revocationExt := findRevocationExt(revokingCert)
	if revocationExt == nil {
		return nil, pkix.Extension{}, extensions.ErrRevocation.New("no revocation extension found")
	}

	return append([]*x509.Certificate{revokingCert}, chain[peertls.CAIndex:]...), *revocationExt, nil
}

func findRevocationExt(revokingCert *x509.Certificate) *pkix.Extension {
	for _, ext := range revokingCert.Extensions {
		if extensions.RevocationExtID.Equal(ext.Id) {
			return &ext
		}
	}
	return nil
}

// RevokeCA revokes the CA certificate in the passed chain and adds a revocation
// extension to that certificate, recording this action.
func RevokeCA(caKey crypto.PrivateKey, chain []*x509.Certificate) ([]*x509.Certificate, pkix.Extension, error) {
	nodeID, err := identity.NodeIDFromCert(chain[peertls.CAIndex])
	if err != nil {
		return nil, pkix.Extension{}, err
	}

	ca := &identity.FullCertificateAuthority{
		ID:        nodeID,
		Cert:      chain[peertls.CAIndex],
		Key:       caKey,
		RestChain: chain[peertls.CAIndex+1:],
	}

	if err = ca.Revoke(); err != nil {
		return nil, pkix.Extension{}, err
	}

	extMap := tlsopts.NewExtensionsMap(ca.Cert)
	revocationExt, ok := extMap[extensions.RevocationExtID.String()]
	if !ok {
		return nil, pkix.Extension{}, extensions.ErrRevocation.New("no revocation extension found")
	}
	return append([]*x509.Certificate{chain[peertls.LeafIndex], ca.Cert}, ca.RestChain...), revocationExt, nil
}

// NewRevokedLeafChain creates a certificate chain (of length 2) with a leaf
// that contains a valid revocation extension.
func NewRevokedLeafChain() ([]crypto.PrivateKey, []*x509.Certificate, pkix.Extension, error) {
	keys, certs, err := NewCertChain(2, storj.LatestIDVersion().Number)
	if err != nil {
		return nil, nil, pkix.Extension{}, err
	}

	newChain, revocation, err := RevokeLeaf(keys[peertls.CAIndex], certs)
	return keys, newChain, revocation, err
}
