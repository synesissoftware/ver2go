/*
 * Copyright (c) 2026 Matt Wilson and Synesis Information Systems
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

func Test_CombineVersion_zero(t *testing.T) {
	var expected uint64 = 0

	test_helpers.EqualInteger(t, expected, CombineVersion(0, 0, 0, 0))
}

func Test_CombineVersion_example(t *testing.T) {
	var expected uint64 = 0x0000_0001_0004_FFFF

	test_helpers.EqualInteger(t, expected, CombineVersion(0, 1, 4, 0xFFFF))
}

func Test_CombineVersion_fields(t *testing.T) {
	{
		var expected uint64 = 0x0001_0000_0000_0000

		test_helpers.EqualInteger(t, expected, CombineVersion(1, 0, 0, 0))
	}

	{
		var expected uint64 = 0x0000_0001_0000_0000

		test_helpers.EqualInteger(t, expected, CombineVersion(0, 1, 0, 0))
	}

	{
		var expected uint64 = 0x0000_0000_0001_0000

		test_helpers.EqualInteger(t, expected, CombineVersion(0, 0, 1, 0))
	}

	{
		var expected uint64 = 0x0000_0000_0000_0001

		test_helpers.EqualInteger(t, expected, CombineVersion(0, 0, 0, 1))
	}
}

func Test_CombineVersion_max(t *testing.T) {
	{
		var expected uint64 = 0xFFFF_0000_0000_0000

		test_helpers.EqualInteger(t, expected, CombineVersion(0xFFFF, 0, 0, 0))
	}

	{
		var expected uint64 = 0x0000_FFFF_0000_0000

		test_helpers.EqualInteger(t, expected, CombineVersion(0, 0xFFFF, 0, 0))
	}

	{
		var expected uint64 = 0x0000_0000_FFFF_0000

		test_helpers.EqualInteger(t, expected, CombineVersion(0, 0, 0xFFFF, 0))
	}

	{
		var expected uint64 = 0x0000_0000_0000_FFFF

		test_helpers.EqualInteger(t, expected, CombineVersion(0, 0, 0, 0xFFFF))
	}

	{
		var expected uint64 = 0xFFFF_FFFF_FFFF_FFFF

		test_helpers.EqualInteger(t, expected, CombineVersion(0xFFFF, 0xFFFF, 0xFFFF, 0xFFFF))
	}
}

func Test_CombineVersion_mixed(t *testing.T) {
	{
		var expected uint64 = 0x007B_01C8_0315_1234

		test_helpers.EqualInteger(t, expected, CombineVersion(123, 456, 789, 0x1234))
	}

	{
		var expected uint64 = 0x0001_0002_0003_C007

		test_helpers.EqualInteger(t, expected, CombineVersion(1, 2, 3, 0xC007))
	}
}

func Test_CombineVersion_ordering(t *testing.T) {
	if CombineVersion(1, 0, 0, 0) <= CombineVersion(0, 0xFFFF, 0xFFFF, 0xFFFF) {
		t.Errorf("major increment must outrank lower fields")
	}

	if CombineVersion(0, 1, 0, 0) <= CombineVersion(0, 0, 0xFFFF, 0xFFFF) {
		t.Errorf("minor increment must outrank lower fields")
	}

	if CombineVersion(0, 0, 1, 0) <= CombineVersion(0, 0, 0, 0xFFFF) {
		t.Errorf("patch increment must outrank αβ-designator")
	}
}

func Test_CombineVersion_library(t *testing.T) {
	test_helpers.EqualInteger(t, Version(), CombineVersion(VersionMajor, VersionMinor, VersionPatch, VersionAB))
}
