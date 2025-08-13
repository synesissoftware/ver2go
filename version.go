/*
 * Copyright (c) 2025 Matt Wilson and Synesis Information Systems
 *
 * Distributed under the 3-Clause BSD License (aka "New BSD-3 License"). See
 * accompanying file LICENSE file for details.
 */

/*
 * Created: 13th February 2025
 * Updated: 13th August 2025
 */

package ver2go

const (
	VersionMajor uint16 = 0
	VersionMinor uint16 = 1
	VersionPatch uint16 = 1
	VersionAB    uint16 = 0xFFFF
	Version      uint64 = (uint64(VersionMajor) << 48) + (uint64(VersionMinor) << 32) + (uint64(VersionPatch) << 16) + (uint64(VersionAB) << 0)
)

// func VersionString

var (
	versionString string = CalcVersionString(VersionMajor, VersionMinor, VersionPatch, VersionAB)
)

func VersionString() string {
	return versionString
}
