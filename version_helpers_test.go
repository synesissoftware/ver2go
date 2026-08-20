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

func Test_CombineVersion_convenience_alpha(t *testing.T) {
	test_helpers.EqualInteger(t, 0x0000_0001_0002_4001, CombineVersion(0, 1, 2, Alpha1))
	test_helpers.EqualInteger(t, 0x0000_0001_0002_4002, CombineVersion(0, 1, 2, Alpha2))
	test_helpers.EqualInteger(t, 0x0000_0001_0002_4003, CombineVersion(0, 1, 2, Alpha3))
	test_helpers.EqualInteger(t, 0x0000_0001_0002_4004, CombineVersion(0, 1, 2, Alpha4))
	test_helpers.EqualInteger(t, 0x0000_0001_0002_4005, CombineVersion(0, 1, 2, Alpha5))
}

func Test_CombineVersion_convenience_beta(t *testing.T) {
	test_helpers.EqualInteger(t, 0x0000_0001_0002_8001, CombineVersion(0, 1, 2, Beta1))
	test_helpers.EqualInteger(t, 0x0000_0001_0002_8002, CombineVersion(0, 1, 2, Beta2))
	test_helpers.EqualInteger(t, 0x0000_0001_0002_8003, CombineVersion(0, 1, 2, Beta3))
	test_helpers.EqualInteger(t, 0x0000_0001_0002_8004, CombineVersion(0, 1, 2, Beta4))
	test_helpers.EqualInteger(t, 0x0000_0001_0002_8005, CombineVersion(0, 1, 2, Beta5))
}

func Test_CombineVersion_convenience_rc(t *testing.T) {
	test_helpers.EqualInteger(t, 0x0001_0002_0003_C001, CombineVersion(1, 2, 3, ReleaseCandidate1))
	test_helpers.EqualInteger(t, 0x0001_0002_0003_C002, CombineVersion(1, 2, 3, ReleaseCandidate2))
	test_helpers.EqualInteger(t, 0x0001_0002_0003_C003, CombineVersion(1, 2, 3, ReleaseCandidate3))
	test_helpers.EqualInteger(t, 0x0001_0002_0003_C004, CombineVersion(1, 2, 3, ReleaseCandidate4))
	test_helpers.EqualInteger(t, 0x0001_0002_0003_C005, CombineVersion(1, 2, 3, ReleaseCandidate5))
}

func Test_CombineVersion_convenience_release(t *testing.T) {
	test_helpers.EqualInteger(t, Version(), CombineVersion(VersionMajor, VersionMinor, VersionPatch, Release))
	test_helpers.EqualInteger(t, 0x0000_0001_0004_FFFF, CombineVersion(VersionMajor, VersionMinor, VersionPatch, Release))
}

func Test_CombineVersion_convenience_zero_minor(t *testing.T) {
	test_helpers.EqualInteger(t, 0x0000_0000_0000_4001, CombineVersion(0, 0, 0, Alpha1))
	test_helpers.EqualInteger(t, 0x0000_0000_0000_8002, CombineVersion(0, 0, 0, Beta2))
	test_helpers.EqualInteger(t, 0x0000_0000_0000_C003, CombineVersion(0, 0, 0, ReleaseCandidate3))
}

func Test_CombineVersion_convenience_zero_minor_zero_patch_alpha(t *testing.T) {
	test_helpers.EqualInteger(t, 0x0000_0000_0001_4001, CombineVersion(0, 0, 1, Alpha1))
	test_helpers.EqualInteger(t, 0x0000_0000_0001_4002, CombineVersion(0, 0, 1, Alpha2))
	test_helpers.EqualInteger(t, 0x0000_0000_0001_4003, CombineVersion(0, 0, 1, Alpha3))
	test_helpers.EqualInteger(t, 0x0000_0000_0001_4004, CombineVersion(0, 0, 1, Alpha4))
	test_helpers.EqualInteger(t, 0x0000_0000_0001_4005, CombineVersion(0, 0, 1, Alpha5))
}

func Test_CombineVersion_convenience_zero_minor_zero_patch_beta(t *testing.T) {
	test_helpers.EqualInteger(t, 0x0000_0000_0001_8001, CombineVersion(0, 0, 1, Beta1))
	test_helpers.EqualInteger(t, 0x0000_0000_0001_8002, CombineVersion(0, 0, 1, Beta2))
	test_helpers.EqualInteger(t, 0x0000_0000_0001_8003, CombineVersion(0, 0, 1, Beta3))
	test_helpers.EqualInteger(t, 0x0000_0000_0001_8004, CombineVersion(0, 0, 1, Beta4))
	test_helpers.EqualInteger(t, 0x0000_0000_0001_8005, CombineVersion(0, 0, 1, Beta5))
}

func Test_CombineVersion_convenience_zero_minor_zero_patch_rc(t *testing.T) {
	test_helpers.EqualInteger(t, 0x0000_0000_0001_C001, CombineVersion(0, 0, 1, ReleaseCandidate1))
	test_helpers.EqualInteger(t, 0x0000_0000_0001_C002, CombineVersion(0, 0, 1, ReleaseCandidate2))
	test_helpers.EqualInteger(t, 0x0000_0000_0001_C003, CombineVersion(0, 0, 1, ReleaseCandidate3))
	test_helpers.EqualInteger(t, 0x0000_0000_0001_C004, CombineVersion(0, 0, 1, ReleaseCandidate4))
	test_helpers.EqualInteger(t, 0x0000_0000_0001_C005, CombineVersion(0, 0, 1, ReleaseCandidate5))
}

func Test_CombineVersion_convenience_zero_minor_release(t *testing.T) {
	test_helpers.EqualInteger(t, 0x0000_0000_0004_FFFF, CombineVersion(VersionMajor, 0, VersionPatch, Release))
	test_helpers.EqualInteger(t, 0x0001_0000_0000_FFFF, CombineVersion(1, 0, 0, Release))
	test_helpers.EqualInteger(t, 0x0001_0000_0002_FFFF, CombineVersion(1, 0, 2, Release))
}
