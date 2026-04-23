[![The Buf logo](.github/buf-logo.svg)][buf]

# protovalidate-go

[![CI](https://github.com/bufbuild/protovalidate-go/actions/workflows/ci.yaml/badge.svg)](https://github.com/bufbuild/protovalidate-go/actions/workflows/ci.yaml)
[![Conformance](https://github.com/bufbuild/protovalidate-go/actions/workflows/conformance.yaml/badge.svg)](https://github.com/bufbuild/protovalidate-go/actions/workflows/conformance.yaml)
[![Report Card](https://goreportcard.com/badge/buf.build/go/protovalidate)](https://goreportcard.com/report/buf.build/go/protovalidate)
[![GoDoc](https://pkg.go.dev/badge/buf.build/go/protovalidate.svg)](https://pkg.go.dev/buf.build/go/protovalidate)
[![BSR](https://img.shields.io/badge/BSR-Module-0C65EC)][buf-mod]

[Protovalidate][protovalidate] is the semantic validation library for Protobuf. It provides standard annotations to validate common rules on messages and fields, as well as the ability to use [CEL][cel] to write custom rules. It's the next generation of [protoc-gen-validate][protoc-gen-validate].

With Protovalidate, you can annotate your Protobuf messages with both standard and custom validation rules:

```protobuf
syntax = "proto3";

package acme.user.v1;

import "buf/validate/validate.proto";

message User {
  string id = 1 [(buf.validate.field).string.uuid = true];
  uint32 age = 2 [(buf.validate.field).uint32.lte = 150]; // We can only hope.
  string email = 3 [(buf.validate.field).string.email = true];
  string first_name = 4 [(buf.validate.field).string.max_len = 64];
  string last_name = 5 [(buf.validate.field).string.max_len = 64];

  option (buf.validate.message).cel = {
    id: "first_name_requires_last_name"
    message: "last_name must be present if first_name is present"
    expression: "!has(this.first_name) || has(this.last_name)"
  };
}
```

Once you've added `protovalidate-go` to your project, validation is idiomatic Go:

```go
if err = protovalidate.Validate(moneyTransfer); err != nil {
    // Handle failure.
}
```

## Installation

> [!TIP]
> The easiest way to get started with Protovalidate for RPC APIs are the quickstarts in Buf's documentation. They're available for both [Connect][connect-go] and [gRPC][grpc-go].

To install the package, use `go get` from within your Go module:

```shell
go get buf.build/go/protovalidate
```

## Documentation

Comprehensive documentation for Protovalidate is available at [protovalidate.com][protovalidate]. 

Highlights for Go developers include:

* The [developer quickstart][quickstart]
* Comprehensive RPC quickstarts for [Connect][connect-go] and [gRPC][grpc-go]
* A [migration guide for protoc-gen-validate][migration-guide] users

API documentation for Go is available on [pkg.go.dev][pkg-go].

### Native standard validation rules
This release provides native support for standard validation rule processing. They are enabled by default 
and can be disabled by setting the ValidatorOption `WithDisableNativeRules`.

We continue to validate that the native rules and the CEL rules produce identical results. 
The `compliance` Makefile target has been updated to run twice, 
once with the native rules enabled, and once with the CEL rules enabled.

Performance improvements on the included benchmarks:

```
goos: darwin
goarch: arm64
pkg: buf.build/go/protovalidate
cpu: Apple M1 Max
                          │ .tmp/bench/2026-04-23:18:10:09.bench.txt │ .tmp/bench/2026-04-23:18:13:13.bench.txt │
                          │                  sec/op                  │      sec/op       vs base                │
Scalar-10                                               174.30n ± 1%        78.03n ± 1%  -55.24% (p=0.000 n=10)
Repeated/Scalar-10                                      291.70n ± 1%        98.94n ± 0%  -66.08% (p=0.000 n=10)
Repeated/Message-10                                      711.0n ± 1%        323.9n ± 0%  -54.44% (p=0.000 n=10)
Repeated/Unique/Scalar-10                               1275.5n ± 0%        322.2n ± 0%  -74.74% (p=0.000 n=10)
Repeated/Unique/Bytes-10                                2560.5n ± 0%        462.2n ± 0%  -81.95% (p=0.000 n=10)
Map-10                                                   298.0n ± 1%        104.3n ± 0%  -64.99% (p=0.000 n=10)
ComplexSchema-10                                         41.19µ ± 1%        14.89µ ± 0%  -63.84% (p=0.000 n=10)
Int32GT-10                                               2.807µ ± 1%        1.015µ ± 0%  -63.86% (p=0.000 n=10)
TestByteMatching-10                                     1368.0n ± 0%        213.5n ± 1%  -84.39% (p=0.000 n=10)
StringMatching-10                                       1937.0n ± 1%        936.3n ± 0%  -51.66% (p=0.000 n=10)
WrapperTesting-10                                        3.133µ ± 0%        1.293µ ± 0%  -58.73% (p=0.000 n=10)
Compile-10                                               8.315m ± 2%        1.371m ± 1%  -83.52% (p=0.000 n=10)
CompileInt32GT-10                                        5.925m ± 1%        1.296m ± 3%  -78.13% (p=0.000 n=10)
geomean                                                  5.120µ             1.538µ       -69.96%

                          │ .tmp/bench/2026-04-23:18:10:09.bench.txt │ .tmp/bench/2026-04-23:18:13:13.bench.txt  │
                          │                   B/op                   │     B/op      vs base                     │
Scalar-10                                               0.000 ± 0%       0.000 ± 0%         ~ (p=1.000 n=10) ¹
Repeated/Scalar-10                                     120.00 ± 0%       48.00 ± 0%   -60.00% (p=0.000 n=10)
Repeated/Message-10                                    120.00 ± 0%       48.00 ± 0%   -60.00% (p=0.000 n=10)
Repeated/Unique/Scalar-10                               536.0 ± 0%       132.0 ± 0%   -75.37% (p=0.000 n=10)
Repeated/Unique/Bytes-10                               1784.0 ± 0%       264.0 ± 0%   -85.20% (p=0.000 n=10)
Map-10                                                 128.00 ± 0%       64.00 ± 0%   -50.00% (p=0.000 n=10)
ComplexSchema-10                                     10.552Ki ± 0%     4.250Ki ± 0%   -59.72% (p=0.000 n=10)
Int32GT-10                                              0.000 ± 0%       0.000 ± 0%         ~ (p=1.000 n=10) ¹
TestByteMatching-10                                     408.0 ± 0%         0.0 ± 0%  -100.00% (p=0.000 n=10)
StringMatching-10                                       387.0 ± 0%         0.0 ± 0%  -100.00% (p=0.000 n=10)
WrapperTesting-10                                       192.0 ± 0%         0.0 ± 0%  -100.00% (p=0.000 n=10)
Compile-10                                            6.816Mi ± 0%     1.675Mi ± 0%   -75.42% (p=0.000 n=10)
CompileInt32GT-10                                     5.204Mi ± 0%     1.611Mi ± 0%   -69.05% (p=0.000 n=10)
geomean                                                            ²                 ?                       ² ³
¹ all samples are equal
² summaries must be >0 to compute geomean
³ ratios must be >0 to compute geomean

                          │ .tmp/bench/2026-04-23:18:10:09.bench.txt │ .tmp/bench/2026-04-23:18:13:13.bench.txt │
                          │                allocs/op                 │  allocs/op   vs base                     │
Scalar-10                                               0.000 ± 0%      0.000 ± 0%         ~ (p=1.000 n=10) ¹
Repeated/Scalar-10                                      3.000 ± 0%      1.000 ± 0%   -66.67% (p=0.000 n=10)
Repeated/Message-10                                     3.000 ± 0%      1.000 ± 0%   -66.67% (p=0.000 n=10)
Repeated/Unique/Scalar-10                               34.00 ± 0%      11.00 ± 0%   -67.65% (p=0.000 n=10)
Repeated/Unique/Bytes-10                                73.00 ± 0%      10.00 ± 0%   -86.30% (p=0.000 n=10)
Map-10                                                  2.000 ± 0%      1.000 ± 0%   -50.00% (p=0.000 n=10)
ComplexSchema-10                                        419.0 ± 0%      121.0 ± 0%   -71.12% (p=0.000 n=10)
Int32GT-10                                              0.000 ± 0%      0.000 ± 0%         ~ (p=1.000 n=10) ¹
TestByteMatching-10                                     17.00 ± 0%       0.00 ± 0%  -100.00% (p=0.000 n=10)
StringMatching-10                                       23.00 ± 0%       0.00 ± 0%  -100.00% (p=0.000 n=10)
WrapperTesting-10                                       17.00 ± 0%       0.00 ± 0%  -100.00% (p=0.000 n=10)
Compile-10                                             96.21k ± 0%     18.41k ± 0%   -80.86% (p=0.000 n=10)
CompileInt32GT-10                                      64.63k ± 0%     17.60k ± 0%   -72.77% (p=0.000 n=10)
geomean                                                            ²                ?                       ² ³
¹ all samples are equal
² summaries must be >0 to compute geomean
³ ratios must be >0 to compute geomean
```

Use `make bench` to generate benchmarks with native rules and `make bench-cel` to generate benchmarks with CEL rules.

## Additional languages and repositories

Protovalidate isn't just for Go! You might be interested in sibling repositories for other languages: 

- [`protovalidate-java`][pv-java] (Java)
- [`protovalidate-python`][pv-python] (Python)
- [`protovalidate-cc`][pv-cc] (C++)
- [`protovalidate-es`][pv-es] (TypeScript and JavaScript)

Additionally, [protovalidate's core repository](https://github.com/bufbuild/protovalidate) provides:

- [Protovalidate's Protobuf API][validate-proto]
- [Conformance testing utilities][conformance] for acceptance testing of `protovalidate` implementations

## Contributing

We genuinely appreciate any help! If you'd like to contribute, check out these resources:

- [Contributing Guidelines][contributing]: Guidelines to make your contribution process straightforward and meaningful
- [Conformance testing utilities](https://github.com/bufbuild/protovalidate/tree/main/docs/conformance.md): Utilities providing acceptance testing of `protovalidate` implementations
- [Go conformance executor][conformance-executable]: Conformance testing executor for `protovalidate-go`

## Legal

Offered under the [Apache 2 license][license].

[buf]: https://buf.build
[cel]: https://cel.dev

[pv-go]: https://github.com/bufbuild/protovalidate-go
[pv-java]: https://github.com/bufbuild/protovalidate-java
[pv-python]: https://github.com/bufbuild/protovalidate-python
[pv-cc]: https://github.com/bufbuild/protovalidate-cc
[pv-es]: https://github.com/bufbuild/protovalidate-es

[buf-mod]: https://buf.build/bufbuild/protovalidate
[license]: LICENSE
[contributing]: .github/CONTRIBUTING.md

[protoc-gen-validate]: https://github.com/bufbuild/protoc-gen-validate

[protovalidate]: https://protovalidate.com/
[quickstart]: https://protovalidate.com/quickstart/
[connect-go]: https://protovalidate.com/quickstart/connect-go/
[grpc-go]: https://protovalidate.com/quickstart/grpc-go/
[grpc-java]: https://protovalidate.com/quickstart/grpc-java/
[grpc-python]: https://protovalidate.com/quickstart/grpc-python/
[migration-guide]: https://protovalidate.com/migration-guides/migrate-from-protoc-gen-validate/
[conformance-executable]: ./internal/cmd/protovalidate-conformance-go/README.md
[pkg-go]: https://pkg.go.dev/buf.build/go/protovalidate

[validate-proto]: https://buf.build/bufbuild/protovalidate/docs/main:buf.validate
[conformance]: https://github.com/bufbuild/protovalidate/blob/main/docs/conformance.md
[examples]: https://github.com/bufbuild/protovalidate/tree/main/examples
[migrate]: https://protovalidate.com/migration-guides/migrate-from-protoc-gen-validate/
