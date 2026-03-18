/*
 * Copyright (c) 2025-2026 Matt Wilson and Synesis Information Systems
 *
 * Distributed under the 3-Clause BSD License (aka "New BSD-3 License"). See
 * accompanying file LICENSE file for details.
 */

package ver2go_test

import (
	. "github.com/synesissoftware/ver2go"
	test_helpers "github.com/synesissoftware/ver2go/internal"

	"testing"
)

const (
	Expected_VersionMajor uint16 = 0
	Expected_VersionMinor uint16 = 1
	Expected_VersionPatch uint16 = 3
	Expected_VersionAB    uint16 = 0xFFFF
)

func Test_Version_Elements(t *testing.T) {
	test_helpers.EqualInteger(t, Expected_VersionMajor, VersionMajor)
	test_helpers.EqualInteger(t, Expected_VersionMinor, VersionMinor)
	test_helpers.EqualInteger(t, Expected_VersionPatch, VersionPatch)
	test_helpers.EqualInteger(t, Expected_VersionAB, VersionAB)
}

func Test_Version(t *testing.T) {
	test_helpers.EqualInteger(t, 0x0000_0001_0003_FFFF, Version)
}

func Test_Version_String(t *testing.T) {
	test_helpers.EqualString(t, "0.1.3", VersionString())
}
