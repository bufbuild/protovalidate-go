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
and can be disabled by compiling with the build tag `cel_rules`:

```
go build -tags="cel_rules" ...
```

We continue to validate that the native rules and the CEL rules produce identical results. 
The `compliance`, `test`, and `test-opaque` Makefile targets have been updated to run twice, 
once with the native rules enabled, and once with the CEL rules enabled.

Performance improvements on the included benchmarks:

```
$ benchstat 2026-04-21:17:17:52.bench.txt 2026-04-21:17:20:31.bench.txt
goos: darwin
goarch: arm64
pkg: buf.build/go/protovalidate
cpu: Apple M1 Max
                          │ 2026-04-21:17:17:52.bench.txt │    2026-04-21:17:20:31.bench.txt    │
                          │            sec/op             │   sec/op     vs base                │
Scalar-10                                    167.35n ± 1%   70.97n ± 0%  -57.59% (p=0.000 n=10)
Repeated/Scalar-10                           285.60n ± 1%   99.10n ± 1%  -65.30% (p=0.000 n=10)
Repeated/Message-10                           691.3n ± 1%   296.7n ± 0%  -57.09% (p=0.000 n=10)
Repeated/Unique/Scalar-10                    1249.0n ± 0%   556.8n ± 0%  -55.42% (p=0.000 n=10)
Repeated/Unique/Bytes-10                     2514.5n ± 0%   967.0n ± 1%  -61.54% (p=0.000 n=10)
Map-10                                        289.9n ± 1%   103.2n ± 1%  -64.39% (p=0.000 n=10)
ComplexSchema-10                              40.93µ ± 0%   14.27µ ± 1%  -65.14% (p=0.000 n=10)
Int32GT-10                                   2748.0n ± 0%   845.6n ± 0%  -69.23% (p=0.000 n=10)
TestByteMatching-10                          1348.0n ± 0%   206.2n ± 0%  -84.70% (p=0.000 n=10)
Compile-10                                    8.262m ± 0%   1.357m ± 1%  -83.58% (p=0.000 n=10)
CompileInt32GT-10                             5.887m ± 1%   1.241m ± 0%  -78.91% (p=0.000 n=10)
geomean                                       5.738µ        1.755µ       -69.42%

                          │ 2026-04-21:17:17:52.bench.txt │       2026-04-21:17:20:31.bench.txt       │
                          │             B/op              │     B/op      vs base                     │
Scalar-10                                    0.000 ± 0%       0.000 ± 0%         ~ (p=1.000 n=10) ¹
Repeated/Scalar-10                          120.00 ± 0%       48.00 ± 0%   -60.00% (p=0.000 n=10)
Repeated/Message-10                         120.00 ± 0%       48.00 ± 0%   -60.00% (p=0.000 n=10)
Repeated/Unique/Scalar-10                    536.0 ± 0%       272.0 ± 0%   -49.25% (p=0.000 n=10)
Repeated/Unique/Bytes-10                    1784.0 ± 0%       832.0 ± 0%   -53.36% (p=0.000 n=10)
Map-10                                      128.00 ± 0%       64.00 ± 0%   -50.00% (p=0.000 n=10)
ComplexSchema-10                          10.552Ki ± 0%     4.383Ki ± 0%   -58.46% (p=0.000 n=10)
Int32GT-10                                   0.000 ± 0%       0.000 ± 0%         ~ (p=1.000 n=10) ¹
TestByteMatching-10                          408.0 ± 0%         0.0 ± 0%  -100.00% (p=0.000 n=10)
Compile-10                                 6.811Mi ± 0%     1.671Mi ± 0%   -75.47% (p=0.000 n=10)
CompileInt32GT-10                          5.199Mi ± 0%     1.606Mi ± 0%   -69.11% (p=0.000 n=10)
geomean                                                 ²                 ?                       ² ³
¹ all samples are equal
² summaries must be >0 to compute geomean
³ ratios must be >0 to compute geomean

                          │ 2026-04-21:17:17:52.bench.txt │      2026-04-21:17:20:31.bench.txt       │
                          │           allocs/op           │  allocs/op   vs base                     │
Scalar-10                                    0.000 ± 0%      0.000 ± 0%         ~ (p=1.000 n=10) ¹
Repeated/Scalar-10                           3.000 ± 0%      1.000 ± 0%   -66.67% (p=0.000 n=10)
Repeated/Message-10                          3.000 ± 0%      1.000 ± 0%   -66.67% (p=0.000 n=10)
Repeated/Unique/Scalar-10                    34.00 ± 0%      14.00 ± 0%   -58.82% (p=0.000 n=10)
Repeated/Unique/Bytes-10                     73.00 ± 0%      24.00 ± 0%   -67.12% (p=0.000 n=10)
Map-10                                       2.000 ± 0%      1.000 ± 0%   -50.00% (p=0.000 n=10)
ComplexSchema-10                             419.0 ± 0%      131.0 ± 0%   -68.74% (p=0.000 n=10)
Int32GT-10                                   0.000 ± 0%      0.000 ± 0%         ~ (p=1.000 n=10) ¹
TestByteMatching-10                          17.00 ± 0%       0.00 ± 0%  -100.00% (p=0.000 n=10)
Compile-10                                  96.14k ± 0%     18.35k ± 0%   -80.91% (p=0.000 n=10)
CompileInt32GT-10                           64.56k ± 0%     17.53k ± 0%   -72.84% (p=0.000 n=10)
geomean                                                 ²                ?                       ² ³
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
