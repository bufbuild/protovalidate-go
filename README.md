# [![The Buf logo](.github/buf-logo.svg)][buf] protovalidate-go

[![CI](https://github.com/bufbuild/protovalidate-go/actions/workflows/ci.yaml/badge.svg)](https://github.com/bufbuild/protovalidate-go/actions/workflows/ci.yaml)
[![Conformance](https://github.com/bufbuild/protovalidate-go/actions/workflows/conformance.yaml/badge.svg)](https://github.com/bufbuild/protovalidate-go/actions/workflows/conformance.yaml)
[![Report Card](https://goreportcard.com/badge/github.com/bufbuild/protovalidate-go)](https://goreportcard.com/report/github.com/bufbuild/protovalidate-go)
[![GoDoc](https://pkg.go.dev/badge/github.com/bufbuild/protovalidate-go.svg)](https://pkg.go.dev/github.com/bufbuild/protovalidate-go)
[![BSR](https://img.shields.io/badge/BSR-Module-0C65EC)][buf-mod]

`protovalidate-go` is the Go language implementation
of [`protovalidate`](https://github.com/bufbuild/protovalidate) designed
to validate Protobuf messages at runtime based on user-defined validation constraints.
Powered by Google's Common Expression Language ([CEL](https://github.com/google/cel-spec)), it provides a
flexible and efficient foundation for defining and evaluating custom validation
rules.
The primary goal of `protovalidate` is to help developers ensure data
consistency and integrity across the network without requiring generated code.

## The `protovalidate` project

Head over to the core [`protovalidate`](https://github.com/bufbuild/protovalidate/) repository for:

- [The API definition](https://github.com/bufbuild/protovalidate/tree/main/proto/protovalidate/buf/validate/validate.proto): used to describe validation constraints
- [Documentation](https://github.com/bufbuild/protovalidate/tree/main/docs): how to apply `protovalidate` effectively
- [Migration tooling](https://github.com/bufbuild/protovalidate/tree/main/docs/migrate.md): incrementally migrate from `protoc-gen-validate`
- [Conformance testing utilities](https://github.com/bufbuild/protovalidate/tree/main/docs/conformance.md): for acceptance testing of `protovalidate` implementations

Other `protovalidate` runtime implementations:

- C++: [`protovalidate-cc`][pv-cc]
- Java: [`protovalidate-java`][pv-java]
- Python: [`protovalidate-python`][pv-python]

And others coming soon:

- TypeScript: `protovalidate-ts`

## Installation

**Requires the `go` toolchain (≥ v1.18)**

To install the package, use the `go get` command from within your Go module:

```shell
go get github.com/bufbuild/protovalidate-go
```

Import the package into your Go project:

```go
import "github.com/bufbuild/protovalidate-go"
```

Remember to always check for the latest version of `protovalidate-go` on the
project's [GitHub releases page](https://github.com/bufbuild/protovalidate-go/releases)
to ensure you're using the most up-to-date version.

## Usage

### Implementing validation constraints

Validation constraints are defined directly within `.proto` files.
Documentation for adding constraints can be found in the `protovalidate` project
[README](https://github.com/bufbuild/protovalidate) and its [comprehensive docs](https://github.com/bufbuild/protovalidate/tree/main/docs).

```protobuf
syntax = "proto3";

package my.package;

import "google/protobuf/timestamp.proto";
import "buf/validate/validate.proto";

message Transaction {
  uint64 id = 1 [(buf.validate.field).uint64.gt = 999];
  google.protobuf.Timestamp purchase_date = 2;
  google.protobuf.Timestamp delivery_date = 3;
  
  string price = 4 [(buf.validate.field).cel = {
    id: "transaction.price",
    message: "price must be positive and include a valid currency symbol ($ or £)",
    expression: "(this.startsWith('$') || this.startsWith('£')) && double(this.substring(1)) > 0"
  }];

  option (buf.validate.message).cel = {
    id: "transaction.delivery_date",
    message: "delivery date must be after purchase date",
    expression: "this.delivery_date > this.purchase_date"
  };
}
```

#### Buf managed mode

`protovalidate-go` assumes the constraint extensions are imported into
the generated code via `buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go`.

If you are using Buf [managed mode](https://buf.build/docs/generate/managed-mode/) to augment Go code generation, ensure 
that the `protovalidate` module is excluded in your [`buf.gen.yaml`](https://buf.build/docs/configuration/v1/buf-gen-yaml#except):

```yaml
version: v1
# <snip>
managed:
  enabled: true
  go_package_prefix:
    except:
      - buf.build/bufbuild/protovalidate
# <snip>
```

### Example

```go
package main

import (
	"fmt"
	"time"
	
	pb "github.com/path/to/generated/protos"
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	msg := &pb.Transaction{
		Id:           1234,
		Price:        "$5.67",
		PurchaseDate: timestamppb.New(time.Now()),
		DeliveryDate: timestamppb.New(time.Now().Add(time.Hour)),
	}

	v, err := protovalidate.New()
	if err != nil {
		fmt.Println("failed to initialize validator:", err)
	}

	if err = v.Validate(msg); err != nil {
		fmt.Println("validation failed:", err)
	} else {
		fmt.Println("validation succeeded")
	}
}
```

### Lazy mode

`protovalidate-go` defaults to lazily construct validation logic for Protobuf 
message types the first time they are encountered. A validator's internal 
cache can be pre-warmed with the `WithMessages` or `WithDescriptors` options 
during initialization:

```go
validator, err := protovalidate.New(
  protovalidate.WithMessages(
    &pb.MyFoo{}, 
    &pb.MyBar{}, 
  ),
)
```

Lazy mode requires usage of a mutex to keep the validator thread-safe, which 
results in about 50% of CPU time spent obtaining a read lock. While [performance](#performance)
is sub-microsecond, the mutex overhead can be further reduced by disabling lazy 
mode with the `WithDisableLazy` option. Note that all expected messages must be
provided during initialization of the validator:

```go
validator, err := protovalidate.New(
  protovalidate.WithDisableLazy(true),
  protovalidate.WithMessages(
    &pb.MyFoo{},
    &pb.MyBar{},
  ),
)
```

### Support legacy `protoc-gen-validate` constraints

The `protovalidate-go` module comes with a `legacy` package which adds opt-in support
for existing `protoc-gen-validate` constraints. Provide the`legacy.WithLegacySupport` 
option when initializing the validator:

```go
validator, err := protovalidate.New(
  legacy.WithLegacySupport(legacy.ModeMerge),
)
```

`protoc-gen-validate` code generation is **not** used by `protovalidate-go`. The 
`legacy` package assumes the `protoc-gen-validate` extensions are imported into
the generated code via `github.com/envoyproxy/protoc-gen-validate/validate`.

A [migration tool](https://github.com/bufbuild/protovalidate/tree/main/tools/protovalidate-migrate) is also available to incrementally upgrade legacy constraints in `.proto` files.

## Performance

[Benchmarks](validator_bench_test.go) are provided to test a variety of use-cases. Generally, after the 
initial cold start, validation on a message is sub-microsecond 
and only allocates in the event of a validation error.

```
[circa 15 May 2023]
goos: darwin
goarch: arm64
pkg: github.com/bufbuild/protovalidate-go
BenchmarkValidator
BenchmarkValidator/ColdStart
BenchmarkValidator/ColdStart-10         	    4372	    276457 ns/op	  470780 B/op	    9255 allocs/op
BenchmarkValidator/Lazy/Valid
BenchmarkValidator/Lazy/Valid-10        	 9022392	     134.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidator/Lazy/Invalid
BenchmarkValidator/Lazy/Invalid-10      	 3416996	     355.9 ns/op	     632 B/op	      14 allocs/op
BenchmarkValidator/Lazy/FailFast
BenchmarkValidator/Lazy/FailFast-10     	 6751131	     172.6 ns/op	     168 B/op	       3 allocs/op
BenchmarkValidator/PreWarmed/Valid
BenchmarkValidator/PreWarmed/Valid-10   	17557560	     69.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidator/PreWarmed/Invalid
BenchmarkValidator/PreWarmed/Invalid-10 	 3621961	     332.9 ns/op	     632 B/op	      14 allocs/op
BenchmarkValidator/PreWarmed/FailFast
BenchmarkValidator/PreWarmed/FailFast-10	13960359	     92.22 ns/op	     168 B/op	       3 allocs/op
PASS
```

### Ecosystem

- [`protovalidate`](https://github.com/bufbuild/protovalidate) core repository
- [Buf][buf]
- [CEL Go][cel-go]
- [CEL Spec][cel-spec]

## Legal

Offered under the [Apache 2 license][license].

[license]: LICENSE
[buf]: https://buf.build
[buf-mod]: https://buf.build/bufbuild/protovalidate
[cel-go]: https://github.com/google/cel-go
[cel-spec]: https://github.com/google/cel-spec
[pv-cc]: https://github.com/bufbuild/protovalidate-cc
[pv-java]: https://github.com/bufbuild/protovalidate-java
[pv-python]: https://github.com/bufbuild/protovalidate-python