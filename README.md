# **ver2go** <!-- omit in toc -->

[![GitHub release](https://img.shields.io/github/v/release/synesissoftware/ver2go.svg)](https://github.com/synesissoftware/ver2go/releases/latest)
[![Go Reference](https://pkg.go.dev/badge/github.com/synesissoftware/ver2go.svg)](https://pkg.go.dev/github.com/synesissoftware/ver2go)
[![Go Report Card](https://goreportcard.com/badge/github.com/synesissoftware/ver2go)](https://goreportcard.com/report/github.com/synesissoftware/ver2go)


**Ver**sion utilities for **Go**


## Introduction

Provides basic version library for use in Go projects.


## Table of Contents <!-- omit in toc -->

- [Introduction](#introduction)
- [Installation](#installation)
- [Components](#components)
- [Examples](#examples)
- [Project Information](#project-information)
	- [Where to get help](#where-to-get-help)
	- [Contribution guidelines](#contribution-guidelines)
	- [Dependencies](#dependencies)
		- [Development/Testing Dependencies](#developmenttesting-dependencies)
	- [Dependents](#dependents)
	- [License](#license)


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

The current version of the library consists of the single API function `CalcVersionString()`:

```Go
func CalcVersionString(verMajor, verMinor, verPatch, verAB uint16) string
```

The meaning of the four parameters are:
* `verMajor` - the major version;
* `verMinor` - the minor version;
* `verPatch` - the patch version;
* `verAB` - the alpha/beta version;

See the function documentation for details of the rules, but the following examples are illustrative:

```Go
ver2go.CalcVersionString(0, 0, 0, 0) // => "0.0.0"
```

```Go
ver2go.CalcVersionString(1, 2, 3, 0xFFFF) // => "1.2.3"
```

```Go
ver2go.CalcVersionString(0, 0, 1, 0x1234) // => "0.0.1.4660"
```

```Go
ver2go.CalcVersionString(0, 1, 2, 0x4321) // => "0.1.2-alpha801"
```

```Go
ver2go.CalcVersionString(0, 1, 0, 0x8765) // => "0.1.0-beta1893"
```

```Go
ver2go.CalcVersionString(1, 2, 3, 0xC007) // => "1.2.3-rc7"
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
* [**recls.Go**](https://github.com/synesissoftware/recls.Go/);
* [**shwild.Go**](https://github.com/synesissoftware/shwild.Go/);
* [**STEGoL**](https://github.com/synesissoftware/STEGoL/);
* [**to-be.Go**](https://github.com/synesissoftware/to-be.Go/);


### License

**ver2go** is released under the 3-clause BSD license. See [LICENSE](./LICENSE) for details.


<!-- ########################### end of file ########################### -->

