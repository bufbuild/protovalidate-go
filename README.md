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
                          │ .tmp/bench/2026-04-24:17:56:31.bench.txt │ .tmp/bench/2026-04-24:18:00:25.bench.txt │
                          │                  sec/op                  │      sec/op       vs base                │
Scalar-10                                               165.90n ± 1%        68.80n ± 1%  -58.53% (p=0.000 n=10)
Repeated/Scalar-10                                      279.80n ± 0%        95.84n ± 1%  -65.75% (p=0.000 n=10)
Repeated/Message-10                                      685.5n ± 1%        288.8n ± 1%  -57.86% (p=0.000 n=10)
Repeated/Unique/Scalar-10                               1248.0n ± 1%        311.3n ± 0%  -75.05% (p=0.000 n=10)
Repeated/Unique/Bytes-10                                2491.0n ± 2%        448.3n ± 1%  -82.00% (p=0.000 n=10)
Map-10                                                   284.2n ± 1%        100.8n ± 0%  -64.52% (p=0.000 n=10)
ComplexSchema-10                                         40.50µ ± 1%        13.47µ ± 0%  -66.74% (p=0.000 n=10)
Int32GT-10                                              2744.5n ± 0%        774.6n ± 0%  -71.78% (p=0.000 n=10)
TestByteMatching-10                                     1334.0n ± 1%        208.2n ± 1%  -84.39% (p=0.000 n=10)
StringMatching-10                                       1893.0n ± 1%        906.5n ± 1%  -52.11% (p=0.000 n=10)
WrapperTesting-10                                        3.044µ ± 1%        1.140µ ± 1%  -62.57% (p=0.000 n=10)
Compile-10                                               8.137m ± 1%        1.983m ± 1%  -75.63% (p=0.000 n=10)
CompileInt32GT-10                                        5.775m ± 1%        1.639m ± 2%  -71.61% (p=0.000 n=10)
MultiRuleError-10                                        783.4n ± 1%        686.8n ± 2%  -12.33% (p=0.000 n=10)
MultiRuleNoError-10                                     234.80n ± 1%        69.11n ± 0%  -70.57% (p=0.000 n=10)
geomean                                                  3.585µ             1.160µ       -67.64%

                          │ .tmp/bench/2026-04-24:17:56:31.bench.txt │ .tmp/bench/2026-04-24:18:00:25.bench.txt  │
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
Compile-10                                            6.826Mi ± 0%     2.039Mi ± 0%   -70.13% (p=0.000 n=10)
CompileInt32GT-10                                     5.210Mi ± 0%     1.842Mi ± 0%   -64.65% (p=0.000 n=10)
MultiRuleError-10                                       784.0 ± 0%       808.0 ± 0%    +3.06% (p=0.000 n=10)
MultiRuleNoError-10                                     0.000 ± 0%       0.000 ± 0%         ~ (p=1.000 n=10) ¹
geomean                                                            ²                 ?                       ² ³
¹ all samples are equal
² summaries must be >0 to compute geomean
³ ratios must be >0 to compute geomean

                          │ .tmp/bench/2026-04-24:17:56:31.bench.txt │ .tmp/bench/2026-04-24:18:00:25.bench.txt │
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
Compile-10                                             96.37k ± 0%     21.58k ± 0%   -77.61% (p=0.000 n=10)
CompileInt32GT-10                                      64.73k ± 0%     19.62k ± 0%   -69.69% (p=0.000 n=10)
MultiRuleError-10                                       21.00 ± 0%      22.00 ± 0%    +4.76% (p=0.000 n=10)
MultiRuleNoError-10                                     0.000 ± 0%      0.000 ± 0%         ~ (p=1.000 n=10) ¹
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
