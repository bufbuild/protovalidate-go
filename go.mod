module github.com/bufbuild/protovalidate-go

go 1.21.1

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.35.2-20240920164238-5a7b106cbb87.1
	github.com/envoyproxy/protoc-gen-validate v1.1.0
	github.com/google/cel-go v0.22.0
	github.com/stretchr/testify v1.9.0
	google.golang.org/protobuf v1.35.2
)

require (
	cel.dev/expr v0.18.0 // indirect
	github.com/antlr4-go/antlr/v4 v4.13.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stoewer/go-strcase v1.3.0 // indirect
	golang.org/x/exp v0.0.0-20240325151524-a685a6edb6d8 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240826202546-f6391c0de4c7 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240826202546-f6391c0de4c7 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// DO NOT MERGE!
replace buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go => buf.build/gen/go/jchadwick-buf/protovalidate/protocolbuffers/go v1.35.1-20241107000346-6ecee89ee0c9.1
