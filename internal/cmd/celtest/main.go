// Copyright 2023-2024 Buf Technologies, Inc.
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
	"log"

	pb "github.com/bufbuild/protovalidate-go/internal/gen/tests/example/v1"
)

// Copyright 2023-2024 Buf Technologies, Inc.
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

func main() {

	validator, err := New()
	testPb := &pb.SimpleMap{
		Val: map[int32]int32{
			42: 1,
		},
	}

	mInt := map[any]any{
		42: 1,
	}
	kInt := 42
	fmt.Println(mInt[kInt])

	mStr := map[any]any{
		"42": 1,
	}
	fmt.Println(mStr["42"])

	out, err := validator.Validate(testPb)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(out)
}
