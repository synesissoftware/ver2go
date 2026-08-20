/*
 * Copyright (c) 2025-2026 Matt Wilson and Synesis Information Systems
 *
 * Distributed under the 3-Clause BSD License (aka "New BSD-3 License"). See
 * accompanying file LICENSE file for details.
 */

/*
 * Created: 13th February 2025
 * Updated: 20th August 2026
 */

package ver2go

const (
	VersionMajor = 0
	VersionMinor = 2
	VersionPatch = 1
	VersionAB    = Beta1
)

var (
	version              = CombineVersion(VersionMajor, VersionMinor, VersionPatch, VersionAB)
	versionString string = CalcVersionString(VersionMajor, VersionMinor, VersionPatch, VersionAB)
)

// Version returns this library's version as a packed 64-bit integer, formed
// by CombineVersion from VersionMajor, VersionMinor, VersionPatch, and
// VersionAB. The result is suitable for numeric comparison: a later release
// has a strictly greater value than an earlier one that uses the same
// packing.
func Version() uint64 {
	return version
}

// VersionString returns this library's version as a human-readable string,
// formed by CalcVersionString from VersionMajor, VersionMinor,
// VersionPatch, and VersionAB. For a final (non-prerelease) version the
// result is of the form "MAJOR.MINOR.PATCH", e.g. "0.2.1".
func VersionString() string {
	return versionString
}
