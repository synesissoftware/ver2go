# ver2go <!-- omit in toc -->

**Ver**sion utilities for **Go**

![Language](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)
[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)
[![GitHub release](https://img.shields.io/github/v/release/synesissoftware/ver2go.svg)](https://github.com/synesissoftware/ver2go/releases/latest)
[![Last Commit](https://img.shields.io/github/last-commit/synesissoftware/ver2go)](https://github.com/synesissoftware/ver2go/commits/master)
[![Go](https://github.com/synesissoftware/ver2go/actions/workflows/go.yml/badge.svg)](https://github.com/synesissoftware/ver2go/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/synesissoftware/ver2go.svg)](https://pkg.go.dev/github.com/synesissoftware/ver2go)


## Table of Contents <!-- omit in toc -->

- [Introduction](#introduction)
- [Installation](#installation)
- [Components](#components)
  - [Constants](#constants)
  - [Functions](#functions)
    - [CalcVersionString()](#calcversionstring)
    - [CombineVersion()](#combineversion)
    - [Version()](#version)
    - [VersionString()](#versionstring)
  - [Illustrative calls](#illustrative-calls)
- [Examples](#examples)
- [Project Information](#project-information)
  - [Where to get help](#where-to-get-help)
  - [Contribution guidelines](#contribution-guidelines)
  - [Dependencies](#dependencies)
    - [Development/Testing Dependencies](#developmenttesting-dependencies)
  - [Dependents](#dependents)
  - [License](#license)


## Introduction

**ver2go** provides version packing and string formatting for Go libraries. It is intended to be used by *consumer* libraries that declare their own version components (`VersionMajor`, `VersionMinor`, `VersionPatch`, `VersionAB`) and expose **`Version()`** and **`VersionString()`** entry points, both implemented simply in terms of **ver2go**'s facilities:

```Go
package mylib

import "github.com/synesissoftware/ver2go"

const (
    VersionMajor uint16 = 1
    VersionMinor uint16 = 2
    VersionPatch uint16 = 3
    VersionAB    uint16 = ver2go.Beta5
)

var (
    version              = ver2go.CombineVersion(VersionMajor, VersionMinor, VersionPatch, VersionAB)
    versionString string = ver2go.CalcVersionString(VersionMajor, VersionMinor, VersionPatch, VersionAB)
)

func Version() uint64 { // would obtain 0x0001_0002_0003_8005
    return version
}

func VersionString() string { // would obtain "1.2.3-beta5"
    return versionString
}
```

With `VersionAB` set to **ver2go.Beta1**, **`VersionString()`** would yield `"1.0.0-beta1"`; with **ver2go.Release**, it yields `"1.0.0"`. **`Version()`** returns the packed **uint64** from **CombineVersion()**, suitable for numeric comparison between releases.


## Installation

Install via `go get`, as in:

```bash
go get "github.com/synesissoftware/ver2go"
```

and then import as:

```Go
import ver2go "github.com/synesissoftware/ver2go"
```

or, simply, as:

```Go
import "github.com/synesissoftware/ver2go"
```


## Components

**ver2go** supplies αβ convenience constants, a version-string formatter, and a version packer. Consumer libraries combine these to implement their own version constants and **`Version()`** / **`VersionString()`** API (see [Introduction](#introduction)); **ver2go** itself follows that same pattern.


### Constants

αβ convenience constants for use as `verAB` / `versionAB` in **CalcVersionString()** and **CombineVersion()**:

* **Alpha1**–**Alpha5** — convenience constants representing `-alpha1` … `-alpha5` (`0x4001` … `0x4005`);
* **Beta1**–**Beta5** — convenience constants representing `-beta1` … `-beta5` (`0x8001` … `0x8005`);
* **ReleaseCandidate1**–**ReleaseCandidate5** — convenience constants representing `-rc1` … `-rc5` (`0xC001` … `0xC005`);
* **Release** — convenience constant representing a final release (`0xFFFF`);


### Functions


#### CalcVersionString()

```Go
func CalcVersionString(verMajor, verMinor, verPatch, verAB uint16) string
```

Calculates a human-readable version string. When all of `verMajor`, `verMinor`, and `verPatch` are zero, or when `verAB` is **Release** (`0xFFFF`), the result is `"MAJOR.MINOR.PATCH"` (e.g. `"1.2.3"`). Otherwise the αβ-designator is appended, using `-rcN`, `-betaN`, or `-alphaN` forms where `verAB` falls in the RC, beta, or alpha ranges; see the function documentation for the full precedence rules and special cases (including the `0.0.PATCH` form).


#### CombineVersion()

```Go
func CombineVersion(versionMajor, versionMinor, versionPatch, versionAB uint16) uint64
```

Packs four 16-bit version components into a single **uint64** for compact storage and comparison: major in bits 63–48, minor in 47–32, patch in 31–16, and `versionAB` in 15–0. For example, `CombineVersion(0, 1, 4, Release)` yields `0x000000010004FFFF`. The same four values are the inputs to **CalcVersionString()**, which produces the corresponding human-readable form.


#### Version()

```Go
func Version() uint64
```

Returns the **calling library**'s version as a packed 64-bit integer — conventionally formed by **CombineVersion()** from that library's `VersionMajor`, `VersionMinor`, `VersionPatch`, and `VersionAB`. The result is suitable for numeric comparison: a later release has a strictly greater value than an earlier one that uses the same packing.


#### VersionString()

```Go
func VersionString() string
```

Returns the **calling library**'s version as a human-readable string — conventionally formed by **CalcVersionString()** from that library's version components. For a final (non-prerelease) version the result is of the form `"MAJOR.MINOR.PATCH"`, e.g. `"1.0.0"`.


### Illustrative calls

```Go
ver2go.CalcVersionString(0, 0, 0, 0) // => "0.0.0"
```

```Go
ver2go.CalcVersionString(1, 2, 3, ver2go.Release) // => "1.2.3"
```

```Go
ver2go.CalcVersionString(0, 0, 1, 0x1234) // => "0.0.1.4660"
```

```Go
ver2go.CalcVersionString(0, 1, 2, ver2go.Alpha1) // => "0.1.2-alpha1"
```

```Go
ver2go.CalcVersionString(0, 1, 0, ver2go.Beta1) // => "0.1.0-beta1"
```

```Go
ver2go.CalcVersionString(1, 2, 3, ver2go.ReleaseCandidate1) // => "1.2.3-rc1"
```

```Go
ver2go.CombineVersion(0, 1, 4, ver2go.Release) // => 0x0000_0001_0004_FFFF
```


## Examples

Examples are provided in the ```examples``` directory, along with a markdown description for each. A detailed list TOC of them is provided in [EXAMPLES.md](./EXAMPLES.md).


## Project Information


### Where to get help

[GitHub Page](https://github.com/synesissoftware/ver2go "GitHub Page")


### Contribution guidelines

Defect reports, feature requests, and pull requests are welcome on https://github.com/synesissoftware/ver2go.


### Dependencies

None


#### Development/Testing Dependencies

* [**require**]("github.com/stretchr/testify/require");


### Dependents

* [**ANGoLS**](https://github.com/synesissoftware/ANGoLS/);
* [**CLASP.Go**](https://github.com/synesissoftware/CLASP.Go/);
* [**CLiC4.Go**](https://github.com/synesissoftware/CLiC4.Go/);
* [**Diagnosticism.Go**](https://github.com/synesissoftware/Diagnosticism.Go/);
* [**libpath.Go**](https://github.com/synesissoftware/libpath.Go/);
* [**libCLImate.Go**](https://github.com/synesissoftware/libCLImate.Go);
* [**p99.Go**](https://github.com/synesissoftware/p99.Go);
* [**recls.Go**](https://github.com/synesissoftware/recls.Go/);
* [**shwild.Go**](https://github.com/synesissoftware/shwild.Go/);
* [**STEGoL**](https://github.com/synesissoftware/STEGoL/);
* [**syngo**](https://github.com/synesissoftware/syngo/);
* [**to-be.Go**](https://github.com/synesissoftware/to-be.Go/);
* [**woad.Go**](https://github.com/synesissoftware/woad.Go/);


### License

**ver2go** is released under the 3-clause BSD license. See [LICENSE](./LICENSE) for details.


<!-- ########################### end of file ########################### -->
