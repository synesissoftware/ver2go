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
