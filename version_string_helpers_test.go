/*
 * Copyright (c) 2025 Matt Wilson and Synesis Information Systems
 *
 * Distributed under the 3-Clause BSD License (aka "New BSD-3 License"). See
 * accompanying file LICENSE file for details.
 */

package ver2go

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Version_String_J_N_P(t *testing.T) {
	{
		var major uint16 = 0
		var minor uint16 = 0
		var patch uint16 = 0
		var ab uint16 = 0

		require.Equal(t, "0.0.0", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 1
		var minor uint16 = 2
		var patch uint16 = 3
		var ab uint16 = 0xffff

		require.Equal(t, "1.2.3", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 123
		var minor uint16 = 456
		var patch uint16 = 789
		var ab uint16 = 0xffff

		require.Equal(t, "123.456.789", CalcVersionString(major, minor, patch, ab))
	}
}

func Test_Version_String_J_N_P_N(t *testing.T) {
	{
		var major uint16 = 0
		var minor uint16 = 0
		var patch uint16 = 1
		var ab uint16 = 0x1234

		require.Equal(t, "0.0.1.4660", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 0
		var minor uint16 = 0
		var patch uint16 = 1
		var ab uint16 = 0x3fff

		require.Equal(t, "0.0.1.16383", CalcVersionString(major, minor, patch, ab))
	}
}

func Test_Version_String_J_N_P_alpha(t *testing.T) {

	{
		var major uint16 = 0
		var minor uint16 = 1
		var patch uint16 = 2
		var ab uint16 = 0x4321

		require.Equal(t, "0.1.2-alpha801", CalcVersionString(major, minor, patch, ab))
	}

}

func Test_Version_String_J_N_P_beta(t *testing.T) {
	{
		var major uint16 = 0
		var minor uint16 = 1
		var patch uint16 = 0
		var ab uint16 = 0x8765

		require.Equal(t, "0.1.0-beta1893", CalcVersionString(major, minor, patch, ab))
	}
}

func Test_Version_String_J_N_P_rc(t *testing.T) {
	{
		var major uint16 = 1
		var minor uint16 = 2
		var patch uint16 = 3
		var ab uint16 = 0xC123

		require.Equal(t, "1.2.3-rc291", CalcVersionString(major, minor, patch, ab))
	}

	{
		var major uint16 = 1
		var minor uint16 = 2
		var patch uint16 = 3
		var ab uint16 = 0xfffe

		require.Equal(t, "1.2.3-rc16382", CalcVersionString(major, minor, patch, ab))
	}
}
