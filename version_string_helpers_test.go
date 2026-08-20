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

func Test_Version_String_J_N_P(t *testing.T) {
	{
		var major uint16 = 0
		var minor uint16 = 0
		var patch uint16 = 0
		var ab uint16 = 0

		test_helpers.EqualString(t, "0.0.0", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 0
		var patch uint16 = 0
		var ab uint16 = 1

		test_helpers.EqualString(t, "0.0.0", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 0
		var patch uint16 = 0
		var ab uint16 = 0x4000

		test_helpers.EqualString(t, "0.0.0", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 0
		var patch uint16 = 0
		var ab uint16 = 0x8000

		test_helpers.EqualString(t, "0.0.0", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 0
		var patch uint16 = 0
		var ab uint16 = 0xC000

		test_helpers.EqualString(t, "0.0.0", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 0
		var patch uint16 = 0
		var ab uint16 = 0xFFFE

		test_helpers.EqualString(t, "0.0.0", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 0
		var patch uint16 = 0
		var ab uint16 = 0xFFFF

		test_helpers.EqualString(t, "0.0.0", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 1
		var minor uint16 = 2
		var patch uint16 = 3
		var ab uint16 = 0xFFFF

		test_helpers.EqualString(t, "1.2.3", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 0
		var patch uint16 = 1
		var ab uint16 = 0xFFFF

		test_helpers.EqualString(t, "0.0.1", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 1
		var patch uint16 = 0
		var ab uint16 = 0xFFFF

		test_helpers.EqualString(t, "0.1.0", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 1
		var minor uint16 = 0
		var patch uint16 = 0
		var ab uint16 = 0xFFFF

		test_helpers.EqualString(t, "1.0.0", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 1
		var patch uint16 = 4
		var ab uint16 = 0xFFFF

		test_helpers.EqualString(t, "0.1.4", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 123
		var minor uint16 = 456
		var patch uint16 = 789
		var ab uint16 = 0xFFFF

		test_helpers.EqualString(t, "123.456.789", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 65535
		var minor uint16 = 65535
		var patch uint16 = 65535
		var ab uint16 = 0xFFFF

		test_helpers.EqualString(t, "65535.65535.65535", CalcVersionString(major, minor, patch, ab))
	}
}

func Test_Version_String_J_N_P_N(t *testing.T) {
	{
		var major uint16 = 0
		var minor uint16 = 0
		var patch uint16 = 1
		var ab uint16 = 0

		test_helpers.EqualString(t, "0.0.1.0", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 0
		var patch uint16 = 1
		var ab uint16 = 1

		test_helpers.EqualString(t, "0.0.1.1", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 0
		var patch uint16 = 1
		var ab uint16 = 0x1234

		test_helpers.EqualString(t, "0.0.1.4660", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 0
		var patch uint16 = 1
		var ab uint16 = 0x3FFF

		test_helpers.EqualString(t, "0.0.1.16383", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 0
		var patch uint16 = 13
		var ab uint16 = 5432

		test_helpers.EqualString(t, "0.0.13.5432", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 0
		var patch uint16 = 1
		var ab uint16 = 0x4000

		test_helpers.EqualString(t, "0.0.1.16384", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 0
		var patch uint16 = 1
		var ab uint16 = 0x8000

		test_helpers.EqualString(t, "0.0.1.32768", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 0
		var patch uint16 = 1
		var ab uint16 = 0xC000

		test_helpers.EqualString(t, "0.0.1.49152", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 0
		var patch uint16 = 1
		var ab uint16 = 0xC002

		test_helpers.EqualString(t, "0.0.1.49154", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 0
		var patch uint16 = 1
		var ab uint16 = 0xFFFE

		test_helpers.EqualString(t, "0.0.1.65534", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 0
		var patch uint16 = 65535
		var ab uint16 = 0xABCD

		test_helpers.EqualString(t, "0.0.65535.43981", CalcVersionString(major, minor, patch, ab))
	}
}

func Test_Version_String_J_N_P_experimental(t *testing.T) {
	{
		var major uint16 = 0
		var minor uint16 = 1
		var patch uint16 = 2
		var ab uint16 = 0

		test_helpers.EqualString(t, "0.1.2.0", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 1
		var minor uint16 = 0
		var patch uint16 = 0
		var ab uint16 = 0

		test_helpers.EqualString(t, "1.0.0.0", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 1
		var patch uint16 = 2
		var ab uint16 = 1

		test_helpers.EqualString(t, "0.1.2.1", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 1
		var patch uint16 = 2
		var ab uint16 = 0x1234

		test_helpers.EqualString(t, "0.1.2.4660", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 1
		var patch uint16 = 2
		var ab uint16 = 0x3FFF

		test_helpers.EqualString(t, "0.1.2.16383", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 1
		var minor uint16 = 2
		var patch uint16 = 3
		var ab uint16 = 0x1234

		test_helpers.EqualString(t, "1.2.3.4660", CalcVersionString(major, minor, patch, ab))
	}
}

func Test_Version_String_J_N_P_alpha(t *testing.T) {
	{
		var major uint16 = 0
		var minor uint16 = 1
		var patch uint16 = 2
		var ab uint16 = 0x4000

		test_helpers.EqualString(t, "0.1.2-alpha0", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 1
		var patch uint16 = 2
		var ab uint16 = 0x4001

		test_helpers.EqualString(t, "0.1.2-alpha1", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 1
		var patch uint16 = 2
		var ab uint16 = 0x4007

		test_helpers.EqualString(t, "0.1.2-alpha7", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 1
		var patch uint16 = 2
		var ab uint16 = 0x4321

		test_helpers.EqualString(t, "0.1.2-alpha801", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 1
		var minor uint16 = 0
		var patch uint16 = 0
		var ab uint16 = 0x4000

		test_helpers.EqualString(t, "1.0.0-alpha0", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 1
		var patch uint16 = 0
		var ab uint16 = 0x7FFF

		test_helpers.EqualString(t, "0.1.0-alpha16383", CalcVersionString(major, minor, patch, ab))
	}
}

func Test_Version_String_J_N_P_beta(t *testing.T) {
	{
		var major uint16 = 0
		var minor uint16 = 1
		var patch uint16 = 2
		var ab uint16 = 0x8000

		test_helpers.EqualString(t, "0.1.2-beta0", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 1
		var patch uint16 = 2
		var ab uint16 = 0x8001

		test_helpers.EqualString(t, "0.1.2-beta1", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 1
		var patch uint16 = 2
		var ab uint16 = 0x800D

		test_helpers.EqualString(t, "0.1.2-beta13", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 1
		var patch uint16 = 0
		var ab uint16 = 0x8765

		test_helpers.EqualString(t, "0.1.0-beta1893", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 1
		var minor uint16 = 2
		var patch uint16 = 3
		var ab uint16 = 0x8000

		test_helpers.EqualString(t, "1.2.3-beta0", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 1
		var patch uint16 = 0
		var ab uint16 = 0xBFFF

		test_helpers.EqualString(t, "0.1.0-beta16383", CalcVersionString(major, minor, patch, ab))
	}
}

func Test_Version_String_J_N_P_rc(t *testing.T) {
	{
		var major uint16 = 0
		var minor uint16 = 1
		var patch uint16 = 2
		var ab uint16 = 0xC000

		test_helpers.EqualString(t, "0.1.2-rc0", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 1
		var patch uint16 = 2
		var ab uint16 = 0xC001

		test_helpers.EqualString(t, "0.1.2-rc1", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 1
		var patch uint16 = 2
		var ab uint16 = 0xC002

		test_helpers.EqualString(t, "0.1.2-rc2", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 1
		var minor uint16 = 2
		var patch uint16 = 3
		var ab uint16 = 0xC007

		test_helpers.EqualString(t, "1.2.3-rc7", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 1
		var minor uint16 = 2
		var patch uint16 = 3
		var ab uint16 = 0xC123

		test_helpers.EqualString(t, "1.2.3-rc291", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 1
		var minor uint16 = 0
		var patch uint16 = 0
		var ab uint16 = 0xC000

		test_helpers.EqualString(t, "1.0.0-rc0", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 1
		var minor uint16 = 2
		var patch uint16 = 3
		var ab uint16 = 0xFFFE

		test_helpers.EqualString(t, "1.2.3-rc16382", CalcVersionString(major, minor, patch, ab))
	}
}

func Test_Version_String_thresholds(t *testing.T) {
	{
		var major uint16 = 0
		var minor uint16 = 1
		var patch uint16 = 0
		var ab uint16 = 0x3FFF

		test_helpers.EqualString(t, "0.1.0.16383", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 1
		var patch uint16 = 0
		var ab uint16 = 0x4000

		test_helpers.EqualString(t, "0.1.0-alpha0", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 1
		var patch uint16 = 0
		var ab uint16 = 0x7FFF

		test_helpers.EqualString(t, "0.1.0-alpha16383", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 1
		var patch uint16 = 0
		var ab uint16 = 0x8000

		test_helpers.EqualString(t, "0.1.0-beta0", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 1
		var patch uint16 = 0
		var ab uint16 = 0xBFFF

		test_helpers.EqualString(t, "0.1.0-beta16383", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 1
		var patch uint16 = 0
		var ab uint16 = 0xC000

		test_helpers.EqualString(t, "0.1.0-rc0", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 1
		var patch uint16 = 0
		var ab uint16 = 0xFFFE

		test_helpers.EqualString(t, "0.1.0-rc16382", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 1
		var patch uint16 = 0
		var ab uint16 = 0xFFFF

		test_helpers.EqualString(t, "0.1.0", CalcVersionString(major, minor, patch, ab))
	}
}

func Test_Version_String_readme_examples(t *testing.T) {
	test_helpers.EqualString(t, "0.0.0", CalcVersionString(0, 0, 0, 0))
	test_helpers.EqualString(t, "1.2.3", CalcVersionString(1, 2, 3, 0xFFFF))
	test_helpers.EqualString(t, "0.0.1.4660", CalcVersionString(0, 0, 1, 0x1234))
	test_helpers.EqualString(t, "0.1.2-alpha801", CalcVersionString(0, 1, 2, 0x4321))
	test_helpers.EqualString(t, "0.1.0-beta1893", CalcVersionString(0, 1, 0, 0x8765))
	test_helpers.EqualString(t, "1.2.3-rc7", CalcVersionString(1, 2, 3, 0xC007))
}

func Test_CalcVersionString_convenience_alpha(t *testing.T) {
	test_helpers.EqualString(t, "0.1.2-alpha1", CalcVersionString(0, 1, 2, Alpha1))
	test_helpers.EqualString(t, "0.1.2-alpha2", CalcVersionString(0, 1, 2, Alpha2))
	test_helpers.EqualString(t, "0.1.2-alpha3", CalcVersionString(0, 1, 2, Alpha3))
	test_helpers.EqualString(t, "0.1.2-alpha4", CalcVersionString(0, 1, 2, Alpha4))
	test_helpers.EqualString(t, "0.1.2-alpha5", CalcVersionString(0, 1, 2, Alpha5))
}

func Test_CalcVersionString_convenience_beta(t *testing.T) {
	test_helpers.EqualString(t, "0.1.2-beta1", CalcVersionString(0, 1, 2, Beta1))
	test_helpers.EqualString(t, "0.1.2-beta2", CalcVersionString(0, 1, 2, Beta2))
	test_helpers.EqualString(t, "0.1.2-beta3", CalcVersionString(0, 1, 2, Beta3))
	test_helpers.EqualString(t, "0.1.2-beta4", CalcVersionString(0, 1, 2, Beta4))
	test_helpers.EqualString(t, "0.1.2-beta5", CalcVersionString(0, 1, 2, Beta5))
}

func Test_CalcVersionString_convenience_rc(t *testing.T) {
	test_helpers.EqualString(t, "1.2.3-rc1", CalcVersionString(1, 2, 3, ReleaseCandidate1))
	test_helpers.EqualString(t, "1.2.3-rc2", CalcVersionString(1, 2, 3, ReleaseCandidate2))
	test_helpers.EqualString(t, "1.2.3-rc3", CalcVersionString(1, 2, 3, ReleaseCandidate3))
	test_helpers.EqualString(t, "1.2.3-rc4", CalcVersionString(1, 2, 3, ReleaseCandidate4))
	test_helpers.EqualString(t, "1.2.3-rc5", CalcVersionString(1, 2, 3, ReleaseCandidate5))
}

func Test_CalcVersionString_convenience_release(t *testing.T) {
	test_helpers.EqualString(t, "1.2.3", CalcVersionString(1, 2, 3, Release))
}

func Test_CalcVersionString_convenience_zero_minor(t *testing.T) {
	test_helpers.EqualString(t, "0.0.0", CalcVersionString(0, 0, 0, Alpha1))
	test_helpers.EqualString(t, "0.0.0", CalcVersionString(0, 0, 0, Beta2))
	test_helpers.EqualString(t, "0.0.0", CalcVersionString(0, 0, 0, ReleaseCandidate3))
}

func Test_CalcVersionString_convenience_zero_minor_zero_patch_alpha(t *testing.T) {
	test_helpers.EqualString(t, "0.0.1.16385", CalcVersionString(0, 0, 1, Alpha1))
	test_helpers.EqualString(t, "0.0.1.16386", CalcVersionString(0, 0, 1, Alpha2))
	test_helpers.EqualString(t, "0.0.1.16387", CalcVersionString(0, 0, 1, Alpha3))
	test_helpers.EqualString(t, "0.0.1.16388", CalcVersionString(0, 0, 1, Alpha4))
	test_helpers.EqualString(t, "0.0.1.16389", CalcVersionString(0, 0, 1, Alpha5))
}

func Test_CalcVersionString_convenience_zero_minor_zero_patch_beta(t *testing.T) {
	test_helpers.EqualString(t, "0.0.1.32769", CalcVersionString(0, 0, 1, Beta1))
	test_helpers.EqualString(t, "0.0.1.32770", CalcVersionString(0, 0, 1, Beta2))
	test_helpers.EqualString(t, "0.0.1.32771", CalcVersionString(0, 0, 1, Beta3))
	test_helpers.EqualString(t, "0.0.1.32772", CalcVersionString(0, 0, 1, Beta4))
	test_helpers.EqualString(t, "0.0.1.32773", CalcVersionString(0, 0, 1, Beta5))
}

func Test_CalcVersionString_convenience_zero_minor_zero_patch_rc(t *testing.T) {
	test_helpers.EqualString(t, "0.0.1.49153", CalcVersionString(0, 0, 1, ReleaseCandidate1))
	test_helpers.EqualString(t, "0.0.1.49154", CalcVersionString(0, 0, 1, ReleaseCandidate2))
	test_helpers.EqualString(t, "0.0.1.49155", CalcVersionString(0, 0, 1, ReleaseCandidate3))
	test_helpers.EqualString(t, "0.0.1.49156", CalcVersionString(0, 0, 1, ReleaseCandidate4))
	test_helpers.EqualString(t, "0.0.1.49157", CalcVersionString(0, 0, 1, ReleaseCandidate5))
}

func Test_CalcVersionString_convenience_zero_minor_release(t *testing.T) {
	test_helpers.EqualString(t, "0.0.4", CalcVersionString(0, 0, 4, Release))
	test_helpers.EqualString(t, "1.0.0", CalcVersionString(1, 0, 0, Release))
	test_helpers.EqualString(t, "1.0.2", CalcVersionString(1, 0, 2, Release))
}
