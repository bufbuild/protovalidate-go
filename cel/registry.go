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

package cel

import "github.com/google/cel-go/common/types"

// registry embeds a types.Registry to prevent cel-go from copying it on every
// Extend() call. When cel.Env.Extend() is called, cel-go checks if the provider
// is a *types.Registry and copies it if so. By embedding the registry in a
// different struct type, the type assertion fails and cel-go treats the
// provider as immutable.
//
// See cel-go's env.go Extend method for the type assertion that this avoids.
type registry struct {
	*types.Registry
}

// Copy shadows the embedded Registry's Copy method to return the registry
// itself, preventing any copying.
func (r *registry) Copy() *registry { return r }

var (
	_ types.Provider = (*registry)(nil)
	_ types.Adapter  = (*registry)(nil)
)
