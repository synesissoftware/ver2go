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
	expected_VersionMajor uint16 = 0
	expected_VersionMinor uint16 = 2
	expected_VersionPatch uint16 = 0
	expected_VersionAB    uint16 = 0x8001
)

func Test_Version_Elements(t *testing.T) {
	test_helpers.EqualInteger(t, expected_VersionMajor, VersionMajor)
	test_helpers.EqualInteger(t, expected_VersionMinor, VersionMinor)
	test_helpers.EqualInteger(t, expected_VersionPatch, VersionPatch)
	test_helpers.EqualInteger(t, expected_VersionAB, VersionAB)
}

func Test_Version(t *testing.T) {
	test_helpers.EqualInteger(t, 0x0000_0002_0000_8001, Version())
}

func Test_Version_String(t *testing.T) {
	test_helpers.EqualString(t, "0.2.0-beta1", VersionString())
}
