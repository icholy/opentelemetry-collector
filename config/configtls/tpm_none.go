// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build notpm

package configtls // import "go.opentelemetry.io/collector/config/configtls"

import (
	"crypto/tls"
	"errors"
)

var errorNoTPM = errors.New("package was not compiled with TPM support")

func (c TPMConfig) tpmCertificate(keyPem []byte, certPem []byte, openTPM func() (tpmCloser, error)) (tls.Certificate, error) {
	return tls.Certificate{}, errorNoTPM
}

func openTPM(path string) func() (tpmCloser, error) {
	return func() (tpmCloser, error) {
		return nil, errorNoTPM
	}
}
