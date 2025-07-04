# Change Log

All notable changes to this project will be documented in this file.
This project adheres to [Semantic Versioning](http://semver.org/).

## [Unreleased]

## [1.6.11] - 2025-07-04

### Changed

- Update dependencies in [#77](https://github.com/cybozu-go/etcdutil/pull/77)
    - Update etcd to v3.6.1
    - Update Golang to 1.24.4

## [1.6.10] - 2025-03-07

### Changed

- Update dependencies in [#75](https://github.com/cybozu-go/etcdutil/pull/75)
    - Update etcd to v3.5.19
    - Update Golang to 1.24.1
    - Update Ubuntu to 24.04

## [1.6.9] - 2024-11-13

### Changed

- Update dependencies in [#73](https://github.com/cybozu-go/etcdutil/pull/73)
    - Update etcd to v3.5.16
    - Update Golang to 1.23.3

## [1.6.8] - 2024-07-12

### Changed
- Bump golang.org/x/net from 0.22.0 to 0.23.0 [#69](https://github.com/cybozu-go/etcdutil/pull/69)
- Update dependencies in [#70](https://github.com/cybozu-go/etcdutil/pull/70)
- Update release procedure to use gh command [#71](https://github.com/cybozu-go/etcdutil/pull/71)

## [1.6.7] - 2024-03-19

### Changed
- Update dependencies in [#67](https://github.com/cybozu-go/etcdutil/pull/67)
- refactor test in [#65](https://github.com/cybozu-go/etcdutil/pull/65)
- Bump google.golang.org/protobuf from 1.31.0 to 1.33.0 [#66](https://github.com/cybozu-go/etcdutil/pull/66)

## [1.6.6] - 2023-11-15

### Changed
- Update dependencies in [#62](https://github.com/cybozu-go/etcdutil/pull/62), [#63](https://github.com/cybozu-go/etcdutil/pull/63)
  - Update etcd to v3.5.10
  - Update Golang to 1.21
- Bump golang.org/x/net from 0.12.0 to 0.17.0 [#60](https://github.com/cybozu-go/etcdutil/pull/60)
- Bump google.golang.org/grpc from 1.56.2 to 1.56.3 [#61](https://github.com/cybozu-go/etcdutil/pull/61)

## [1.6.5] - 2023-07-14

### Changed
- Update dependencies in [#58](https://github.com/cybozu-go/etcdutil/pull/58)
  - Update etcd to v3.5.9
  - Update Golang to 1.20

## [1.6.4] - 2023-02-20

### Changed
- Update dependencies in [#54](https://github.com/cybozu-go/etcdutil/pull/54)
    - Update etcd to v3.5.7

## [1.6.3] - 2023-01-19

### Changed
- Update dependencies in [#52](https://github.com/cybozu-go/etcdutil/pull/52)
    - Update etcd to v3.5.6
    - Update Golang to 1.19

## [1.6.2] - 2022-11-08

### Changed
- Update etcd to v3.5.5 ([#50](https://github.com/cybozu-go/etcdutil/pull/50))

## [1.6.1] - 2022-08-23

### Changed
- Update dependencies ([#48](https://github.com/cybozu-go/etcdutil/pull/48))
    - Update etcd to v3.5.4
    - Update Golang to 1.18

## [1.6.0] - 2022-04-14

### Changed
- update etcd to v3.5.3 and some dependencies ([#44](https://github.com/cybozu-go/etcdutil/pull/44))

## [1.5.0] - 2021-12-20

### Changed
- update etcd to v3.5.1 ([#42](https://github.com/cybozu-go/etcdutil/pull/42))

## [1.4.1] - 2021-10-04

### Changed
- Update etcd to v3.4.17 ([#40](https://github.com/cybozu-go/etcdutil/pull/40)).

## [1.4.0] - 2021-05-19

### Changed
- Update for etcd 3.4.16 ([#37](https://github.com/cybozu-go/etcdutil/pull/37)).

## [1.3.7] - 2021-05-07
### Fixed
- Fix release CI workflow ([#35](https://github.com/cybozu-go/etcdutil/pull/35)).

## [1.3.6] - 2021-05-06
### Changed
- Enable keepalive checks ([#34](https://github.com/cybozu-go/etcdutil/pull/34)).

## [1.3.5] - 2020-09-10
### Changed
- Update etcd client library as of [etcd-3.3.25](https://github.com/etcd-io/etcd/releases/tag/v3.3.25).

### Fixed
- Fix NewConfig() to return copy of default object ([#28](https://github.com/cybozu-go/etcdutil/pull/28)).

## [1.3.4] - 2019-10-24
### Changed
- Update golang 1.13.3 ([#22](https://github.com/cybozu-go/etcdutil/pull/22))

## [1.3.3] - 2019-08-20
### Changed
- Update etcd client library as of [etcd-3.3.15](https://github.com/etcd-io/etcd/releases/tag/v3.3.15).

## [1.3.2] - 2019-08-19
### Changed
- Update etcd client library as of [etcd-3.3.14](https://github.com/etcd-io/etcd/releases/tag/v3.3.14).
- Revert [#11](https://github.com/cybozu-go/etcdutil/pull/11) "Workaround for [etcd bug #9949](https://github.com/etcd-io/etcd/issues/9949)".

## [1.3.1] - 2018-11-19
### Changed
- Workaround for [etcd bug #9949](https://github.com/etcd-io/etcd/issues/9949).

## [1.3.0] - 2018-10-15
### Added
- AddPFlags method for github.com/spf13/pflag package.

## [1.2.2] - 2018-10-10
### Changed
- Update Go module dependencies ([#9](https://github.com/cybozu-go/etcdutil/pull/9)).

## [1.2.1] - 2018-10-10
### Changed
- Remove http://127.0.0.1:4001 from the default endpoints ([#8](https://github.com/cybozu-go/etcdutil/pull/8)).

## [1.2.0] - 2018-10-09
### Added
- Common command-line flags ([#7](https://github.com/cybozu-go/etcdutil/pull/7)).

## [1.1.0] - 2018-09-14
### Added
- Opt in to [Go modules](https://github.com/golang/go/wiki/Modules).

## 1.0.0 - 2018-09-03

This is the first release.

[Unreleased]: https://github.com/cybozu-go/etcdutil/compare/v1.6.11...HEAD
[1.6.11]: https://github.com/cybozu-go/etcdutil/compare/v1.6.10...v1.6.11
[1.6.10]: https://github.com/cybozu-go/etcdutil/compare/v1.6.9...v1.6.10
[1.6.9]: https://github.com/cybozu-go/etcdutil/compare/v1.6.8...v1.6.9
[1.6.8]: https://github.com/cybozu-go/etcdutil/compare/v1.6.7...v1.6.8
[1.6.7]: https://github.com/cybozu-go/etcdutil/compare/v1.6.6...v1.6.7
[1.6.6]: https://github.com/cybozu-go/etcdutil/compare/v1.6.5...v1.6.6
[1.6.5]: https://github.com/cybozu-go/etcdutil/compare/v1.6.4...v1.6.5
[1.6.4]: https://github.com/cybozu-go/etcdutil/compare/v1.6.3...v1.6.4
[1.6.3]: https://github.com/cybozu-go/etcdutil/compare/v1.6.2...v1.6.3
[1.6.2]: https://github.com/cybozu-go/etcdutil/compare/v1.6.1...v1.6.2
[1.6.1]: https://github.com/cybozu-go/etcdutil/compare/v1.6.0...v1.6.1
[1.6.0]: https://github.com/cybozu-go/etcdutil/compare/v1.5.0...v1.6.0
[1.5.0]: https://github.com/cybozu-go/etcdutil/compare/v1.4.1...v1.5.0
[1.4.1]: https://github.com/cybozu-go/etcdutil/compare/v1.4.0...v1.4.1
[1.4.0]: https://github.com/cybozu-go/etcdutil/compare/v1.3.7...v1.4.0
[1.3.7]: https://github.com/cybozu-go/etcdutil/compare/v1.3.6...v1.3.7
[1.3.6]: https://github.com/cybozu-go/etcdutil/compare/v1.3.5...v1.3.6
[1.3.5]: https://github.com/cybozu-go/etcdutil/compare/v1.3.4...v1.3.5
[1.3.4]: https://github.com/cybozu-go/etcdutil/compare/v1.3.3...v1.3.4
[1.3.3]: https://github.com/cybozu-go/etcdutil/compare/v1.3.2...v1.3.3
[1.3.2]: https://github.com/cybozu-go/etcdutil/compare/v1.3.1...v1.3.2
[1.3.1]: https://github.com/cybozu-go/etcdutil/compare/v1.3.0...v1.3.1
[1.3.0]: https://github.com/cybozu-go/etcdutil/compare/v1.2.2...v1.3.0
[1.2.2]: https://github.com/cybozu-go/etcdutil/compare/v1.2.1...v1.2.2
[1.2.1]: https://github.com/cybozu-go/etcdutil/compare/v1.2.0...v1.2.1
[1.2.0]: https://github.com/cybozu-go/etcdutil/compare/v1.1.0...v1.2.0
[1.1.0]: https://github.com/cybozu-go/etcdutil/compare/v1.0.0...v1.1.0
