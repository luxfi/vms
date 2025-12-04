// Copyright (C) 2019-2025, Lux Industries, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Package vms provides VM factory interfaces.
package vms

import "github.com/luxfi/log"

// Factory creates new instances of a VM.
type Factory interface {
	New(log.Logger) (interface{}, error)
}
