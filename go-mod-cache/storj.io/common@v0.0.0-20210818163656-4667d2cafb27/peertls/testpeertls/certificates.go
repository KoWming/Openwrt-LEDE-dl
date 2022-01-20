// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package testpeertls

import (
	"crypto"
	"crypto/x509"

	"storj.io/common/peertls"
	"storj.io/common/peertls/extensions"
	"storj.io/common/pkcrypto"
	"storj.io/common/storj"
)

// NewCertChain creates a valid peertls certificate chain (and respective keys) of the desired length.
// NB: keys are in the reverse order compared to certs (i.e. first key belongs to last cert)!
func NewCertChain(length int, versionNumber storj.IDVersionNumber) (keys []crypto.PrivateKey, certs []*x509.Certificate, _ error) {
	version, err := storj.GetIDVersion(versionNumber)
	if err != nil {
		return nil, nil, err
	}

	for i := 0; i < length; i++ {
		key, err := pkcrypto.GeneratePrivateKey()
		if err != nil {
			return nil, nil, err
		}
		keys = append([]crypto.PrivateKey{key}, keys...)

		var template *x509.Certificate
		if i != length-1 {
			template, err = peertls.CATemplate()
			if err != nil {
				return nil, nil, err
			}
			if err := extensions.AddExtraExtension(template, storj.NewVersionExt(version)); err != nil {
				return nil, nil, err
			}
		} else {
			template, err = peertls.LeafTemplate()
		}
		if err != nil {
			return nil, nil, err
		}

		var cert *x509.Certificate
		if i == 0 {
			cert, err = peertls.CreateSelfSignedCertificate(key, template)
		} else {
			var pubKey crypto.PublicKey
			pubKey, err = pkcrypto.PublicKeyFromPrivate(key)
			if err == nil {
				// NB: 	`keys[1]`: key has already been prepended; parent key is at first index
				// 		`certs[0]`: cert hasn't been prepended yet; parent cert is at zeroth index
				cert, err = peertls.CreateCertificate(pubKey, keys[1], template, certs[0])
			}
		}
		if err != nil {
			return nil, nil, err
		}

		certs = append([]*x509.Certificate{cert}, certs...)
	}
	return keys, certs, nil
}
