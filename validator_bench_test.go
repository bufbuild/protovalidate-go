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

package protovalidate

import (
	"testing"

	pb "buf.build/go/protovalidate/internal/gen/tests/example/v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/rodaine/protogofakeit"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func BenchmarkScalar(b *testing.B) {
	benchSuccess(b, &pb.BenchScalar{})
}

func BenchmarkRepeated(b *testing.B) {
	b.Run("Scalar", func(b *testing.B) {
		benchSuccess(b, &pb.BenchRepeatedScalar{})
	})
	b.Run("Message", func(b *testing.B) {
		benchSuccess(b, &pb.BenchRepeatedMessage{})
	})
	b.Run("Unique", func(b *testing.B) {
		b.Run("Scalar", func(b *testing.B) {
			benchSuccess(b, &pb.BenchRepeatedScalarUnique{})
		})
		b.Run("Bytes", func(b *testing.B) {
			benchSuccess(b, &pb.BenchRepeatedBytesUnique{})
		})
	})
}

func benchSuccess(b *testing.B, msg proto.Message) {
	b.Helper()

	faker := protogofakeit.New(gofakeit.New(1))
	require.NoError(b, faker.FakeProto(msg))
	val, err := New(WithMessages(msg), WithDisableLazy())
	require.NoError(b, err)

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := val.Validate(msg)
			require.NoError(b, err)
		}
	})
}
