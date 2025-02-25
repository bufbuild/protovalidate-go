# [![The Buf logo](../../../.github/buf-logo.svg)][buf] protovalidate-go

# Go conformance executor

This binary is the [conformance testing executor](https://github.com/bufbuild/protovalidate/tree/main/tools/protovalidate-conformance) for the Go implementation. From the root of the project, the Go conformance tests can be executed with make:

```shell
make conformance # runs all conformance tests

make conformance ARGS='-suite uint64' # pass flags to the conformance harness
```

[buf]: https://buf.build
