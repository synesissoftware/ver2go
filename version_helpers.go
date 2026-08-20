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

const (
	// Convenience constant represent the "-alpha1".
	Alpha1 = 0x4001
	// Convenience constant represent the "-alpha2".
	Alpha2 = 0x4002
	// Convenience constant represent the "-alpha3".
	Alpha3 = 0x4003
	// Convenience constant represent the "-alpha4".
	Alpha4 = 0x4004
	// Convenience constant represent the "-alpha5".
	Alpha5 = 0x4005

	// Convenience constant represent the "-beta1".
	Beta1 = 0x8001
	// Convenience constant represent the "-beta2".
	Beta2 = 0x8002
	// Convenience constant represent the "-beta3".
	Beta3 = 0x8003
	// Convenience constant represent the "-beta4".
	Beta4 = 0x8004
	// Convenience constant represent the "-beta5".
	Beta5 = 0x8005

	// Convenience constant represent the "-rc1".
	ReleaseCandidate1 = 0xC001
	// Convenience constant represent the "-rc2".
	ReleaseCandidate2 = 0xC002
	// Convenience constant represent the "-rc3".
	ReleaseCandidate3 = 0xC003
	// Convenience constant represent the "-rc4".
	ReleaseCandidate4 = 0xC004
	// Convenience constant represent the "-rc5".
	ReleaseCandidate5 = 0xC005

	// Convenience constant represent the final release version.
	Release = 0xFFFF
)

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
