/*
 * Copyright (c) 2026 Matt Wilson and Synesis Information Systems
 *
 * Distributed under the 3-Clause BSD License (aka "New BSD-3 License"). See
 * accompanying file LICENSE file for details.
 */

/*
 * Created: 20th August 2026
 * Updated: 20th August 2026
 */

package ver2go


// CombineVersion packs four 16-bit version components into a single uint64
// for compact storage and comparison. Each component occupies a 16-bit
// field, in order from the most-significant bits:
//
//	bits 63-48: versionMajor
//	bits 47-32: versionMinor
//	bits 31-16: versionPatch
//	bits 15-0:  versionAB (the αβ-designator)
//
// For example, CombineVersion(0, 1, 4, 0xFFFF) yields 0x000000010004FFFF.
// The same four values are the inputs to CalcVersionString, which produces
// the corresponding human-readable form.
func CombineVersion(
	versionMajor uint16,
	versionMinor uint16,
	versionPatch uint16,
	versionAB uint16,
) uint64 {

	return 0 +
		uint64(versionMajor&0xFFFF)<<48 +
		uint64(versionMinor&0xFFFF)<<32 +
		uint64(versionPatch&0xFFFF)<<16 +
		uint64(versionAB&0xFFFF)<<0 +
		0
}
