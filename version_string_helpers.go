/*
 * Copyright (c) 2025 Matt Wilson and Synesis Information Systems
 *
 * Distributed under the 3-Clause BSD License (aka "New BSD-3 License"). See
 * accompanying file LICENSE file for details.
 */

/*
 * Created: 13th February 2025
 * Updated: 1st March 2025
 */

package ver2go

import "fmt"

// Calculates a version string, which is conventionally a dotted string of
// the form "<MAJOR>.<MINOR>.<PATCH>", where each of <MAJOR>, <MINOR>,
// <PATCH> is a number, e.g. "1.2.3". This is the form presented when either
// of the following two conditions is true:
//
// 1. All of `verMajor`, `verMinor`, `verPatch` have the value 0; or
// 2. `VersionAB` has the value 0xffff.
//
// In all other cases, the version string takes the form of
// "<MAJOR>.<MINOR>.<PATCH>.<AB_DGNTR>", where <AB_DGNTR> - called the
// αβ-designator - has a variety of forms as described by the following
// rules (listed in order of precedence):
//
//  1. If both `verMajor` and `verMinor` have the value 0, then the
//     αβ-designator is a simple number, as in "0.0.13.5432";
//  2. If `verAB` is >= 0xC000, then αβ-designator designates a "release
//     candidate" where the RC number is verAB-0xC000, e.g. "0.1.2-rc1" (for
//     a `verAB` of 0xC002);
//  2. If `verAB` is >= 0x8000, then αβ-designator designates a "beta" where
//     the RC number is verAB-0x800), e.g. "0.1.2-beta13" (for a `verAB` of
//     0x800D);
//  3. If `verAB` is >= 0x4000, then αβ-designator designates an "alpha"
//     where the RC number is verAB-0x800), e.g. "0.1.2-alpha7" (for a
//     `verAB` of 0x8007);
//  4. If `verAB` is > 0, then αβ-designator designates something sub-alpha,
//     perhaps an experimental release, and is given a plain number form,
//     e.g. "0.1.2.4660" (for a `verAB` of 0x1234).
func CalcVersionString(verMajor, verMinor, verPatch, verAB uint16) string {
	if verMajor == 0 && verMinor == 0 && verPatch == 0 {
		return fmt.Sprintf("%d.%d.%d", verMajor, verMinor, verPatch)
	}

	if verAB == 0xffff {
		return fmt.Sprintf("%d.%d.%d", verMajor, verMinor, verPatch)
	}

	if verMajor == 0 && verMinor == 0 && verPatch != 0 {
		return fmt.Sprintf("%d.%d.%d.%d", verMajor, verMinor, verPatch, verAB)
	}

	if verAB >= 0xC000 {
		return fmt.Sprintf("%d.%d.%d-rc%d", verMajor, verMinor, verPatch, verAB-0xC000)
	}

	if verAB >= 0x8000 {
		return fmt.Sprintf("%d.%d.%d-beta%d", verMajor, verMinor, verPatch, verAB-0x8000)
	}

	if verAB >= 0x4000 {
		return fmt.Sprintf("%d.%d.%d-alpha%d", verMajor, verMinor, verPatch, verAB-0x4000)
	}

	return fmt.Sprintf("%d.%d.%d.%d", verMajor, verMinor, verPatch, verAB)
}

