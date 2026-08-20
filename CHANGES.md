# ver2go - Changes <!-- omit in toc -->


## 0.2.1-beta1 - 20th August 2026

* enforced Synesis Go import order via **gci** (**.golangci.yml**, **examples/.golangci.yml**);
* restructured **examples/libver** into **examples/libver/main.go** so `go test ./...` no longer collides on multiple `main`s;
* updated **examples/libver** to use **VersionString()**;
* version string updated for the 0.2.1-beta1 release;


## 0.2.0-beta1 - 20th August 2026

* added **CombineVersion()** for packing four 16-bit version components into a **uint64**;
* added **Version()** (replacing the **Version** constant) and documented **VersionString()**;
* added αβ convenience constants (**Alpha1**–**Alpha5**, **Beta1**–**Beta5**, **ReleaseCandidate1**–**ReleaseCandidate5**, **Release**);
* expanded unit-test coverage for **CalcVersionString()** and **CombineVersion()**;
* CI reliability fix (golangci-lint config verification disabled in CI);
* macOS test reliability fix in **run_all_unit_tests.sh** (`dyld` / **LC_UUID** mitigation);
* removed retired Go Report Card badge from **README.md**;
* version string updated for the 0.2.0-beta1 release;


## 0.1.4 - 20th August 2026

* CI modernisation (matrix + lint);
* boilerplate additions (scripts, markdown docs, project identity);
* version string updated for the 0.1.4 release;


## 0.1.3 - 18th March 2026

* reducing dependencies;
* tidying;


## 0.1.2 - 18th August 2025

* GitHub Actions;
* boilerplate;
* documentation;


## 0.1.1 - 13th August 2025

* distro fix;


## 0.1.0 - 13th August 2025

* release;


## 0.0.1 - 3rd April 2025

* release preparation;


## 0.0.0-alpha3 - 1st March 2025

* fix;


## 0.0.0-alpha2 - 26th February 2025

* tidying;


## 0.0.0 - 13th February 2025

FIRST PUBLIC RELEASE


<!-- ########################### end of file ########################### -->
