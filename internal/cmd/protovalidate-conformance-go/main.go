// Copyright 2023-2025 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"buf.build/go/hyperpb"
	"buf.build/go/protovalidate"
	"buf.build/go/protovalidate/internal/gen/buf/validate/conformance/harness"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/dynamicpb"
	"google.golang.org/protobuf/types/known/anypb"
)

type Config struct {
	HyperPB bool
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("[protovalidate-go] ")

	config := Config{
		HyperPB: os.Getenv("HYPERPB") != "",
	}

	req := &harness.TestConformanceRequest{}
	if data, err := io.ReadAll(os.Stdin); err != nil {
		log.Fatalf("failed to read input from stdin: %v", err)
	} else if err = proto.Unmarshal(data, req); err != nil {
		log.Fatalf("failed to unmarshal conformance request: %v", err)
	}

	resp, err := TestConformance(req, config)
	if err != nil {
		log.Fatalf("unable to test conformance: %v", err)
	} else if data, err := proto.Marshal(resp); err != nil {
		log.Fatalf("unable to marshal conformance response: %v", err)
	} else if _, err = os.Stdout.Write(data); err != nil {
		log.Fatalf("unable to write output to stdout: %v", err)
	}
}

func TestConformance(req *harness.TestConformanceRequest, config Config) (*harness.TestConformanceResponse, error) {
	files, err := protodesc.NewFiles(req.GetFdset())
	if err != nil {
		err = fmt.Errorf("failed to parse file descriptors: %w", err)
		return nil, err
	}
	registry := &protoregistry.Types{}
	files.RangeFiles(func(file protoreflect.FileDescriptor) bool {
		for i := range file.Extensions().Len() {
			if err = registry.RegisterExtension(
				dynamicpb.NewExtensionType(file.Extensions().Get(i)),
			); err != nil {
				return false
			}
		}
		return err == nil
	})
	if err != nil {
		return nil, err
	}
	val, err := protovalidate.New(protovalidate.WithExtensionTypeResolver(registry))
	if err != nil {
		err = fmt.Errorf("failed to initialize validator: %w", err)
		return nil, err
	}
	resp := harness.TestConformanceResponse_builder{Results: map[string]*harness.TestResult{}}
	for caseName, testCase := range req.GetCases() {
		resp.Results[caseName] = TestCase(val, files, testCase, config.HyperPB)
	}
	return resp.Build(), nil
}

func TestCase(val protovalidate.Validator, files *protoregistry.Files, testCase *anypb.Any, useHyperPB bool) *harness.TestResult {
	urlParts := strings.Split(testCase.GetTypeUrl(), "/")
	fullName := protoreflect.FullName(urlParts[len(urlParts)-1])
	desc, err := files.FindDescriptorByName(fullName)
	if err != nil {
		return unexpectedErrorResult("unable to find descriptor: %v", err)
	}
	msgDesc, ok := desc.(protoreflect.MessageDescriptor)
	if !ok {
		return unexpectedErrorResult("expected message descriptor, got %T", desc)
	}

	var dyn proto.Message
	if useHyperPB {
		dyn = hyperpb.NewMessage(hyperpb.CompileMessageDescriptor(msgDesc))
	} else {
		dyn = dynamicpb.NewMessage(msgDesc)
	}
	if err = anypb.UnmarshalTo(testCase, dyn, proto.UnmarshalOptions{}); err != nil {
		return unexpectedErrorResult("unable to unmarshal test case: %v", err)
	}

	err = val.Validate(dyn)
	if err == nil {
		return harness.TestResult_builder{
			Success: proto.Bool(true),
		}.Build()
	}
	switch res := err.(type) {
	case *protovalidate.ValidationError:
		return harness.TestResult_builder{
			ValidationError: res.ToProto(),
		}.Build()
	case *protovalidate.RuntimeError:
		return harness.TestResult_builder{
			RuntimeError: proto.String(res.Error()),
		}.Build()
	case *protovalidate.CompilationError:
		return harness.TestResult_builder{
			CompilationError: proto.String(res.Error()),
		}.Build()
	default:
		return unexpectedErrorResult("unknown error: %v", err)
	}
}

func unexpectedErrorResult(format string, args ...any) *harness.TestResult {
	return harness.TestResult_builder{
		UnexpectedError: proto.String(fmt.Sprintf(format, args...)),
	}.Build()
}
