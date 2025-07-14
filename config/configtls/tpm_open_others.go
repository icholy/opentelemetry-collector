// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !notpm && !linux && !windows

package configtls // import "go.opentelemetry.io/collector/config/configtls"

import (
	"errors"
)

// for testing
var tpmSimulator tpmCloser

func openTPM(path string) func() (tpmCloser, error) {
	return func() (tpmCloser, error) {
		if path == "simulator" {
			return tpmSimulator, nil
		}
		return nil, errors.New("TPM is not supported on this platform")
	}
}
