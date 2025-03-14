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

package cel

import (
	"strconv"
)

type Ipv4 struct {
	str       string
	index     int64
	strLen    int64
	octets    []int64
	prefixLen int64
}

// Return the 32-bit value of an address parsed through address() or addressPrefix().
// Return -1 if no address was parsed successfully.
func (i *Ipv4) getBits() int64 {
	if len(i.octets) != 4 {
		return -1
	}
	return ((i.octets[0] << 24) | (i.octets[1] << 16) | (i.octets[2] << 8) | i.octets[3]>>0)
}

// Return true if all bits to the right of the prefix-length are all zeros.
// Behavior is undefined if addressPrefix() has not been called before, or has
// returned false.
func (i *Ipv4) isPrefixOnly() bool {
	bits := i.getBits()
	var mask int64
	if i.prefixLen == 32 {
		mask = 0xffffffff
	} else {
		mask = ^(0xffffffff >> i.prefixLen) >> 0
	}
	masked := (bits & mask) >> 0
	return bits == masked
}

// Parse Ipv4 Address in dotted decimal notation.
func (i *Ipv4) address() bool {
	return i.addressPart() && i.index == i.strLen
}

// Parse Ipv4 Address prefix.
func (i *Ipv4) addressPrefix() bool {
	return i.addressPart() &&
		i.take('/') &&
		i.prefixLength() &&
		i.index == i.strLen
}

// Determines the prefix length. Stores value in `prefixLen`.
func (i *Ipv4) prefixLength() bool {
	start := i.index
	for {
		if i.index >= i.strLen || !i.digit() {
			break
		}
		if i.index-start > 2 {
			// max prefix-length is 32 bits, so anything more than 2 digits is invalid
			return false
		}
	}
	str := i.str[start:i.index]
	if len(str) == 0 {
		// too short
		return false
	}
	if len(str) > 1 && str[0] == '0' {
		// bad leading 0
		return false
	}
	value, err := strconv.ParseInt(str, 0, 32)
	if err != nil {
		// Error converting to number
		return false
	}
	if value > 32 {
		// max 32 bits
		return false
	}
	i.prefixLen = value
	return true
}

// Parses str from the current index to determine an address part.
func (i *Ipv4) addressPart() bool {
	start := i.index
	if i.decOctet() &&
		i.take('.') &&
		i.decOctet() &&
		i.take('.') &&
		i.decOctet() &&
		i.take('.') &&
		i.decOctet() {
		return true
	}
	i.index = start
	return false
}

// Parses str from the current index to determine a decimal octet.
func (i *Ipv4) decOctet() bool {
	start := i.index
	for {
		if i.index >= i.strLen || !i.digit() {
			break
		}
		if i.index-start > 3 {
			// decimal octet can be three characters at most
			return false
		}
	}
	str := i.str[start:i.index]
	if len(str) == 0 {
		// too short
		return false
	}
	if len(str) > 1 && str[0] == '0' {
		// bad leading 0
		return false
	}
	value, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return false
	}
	if value > 255 {
		return false
	}
	i.octets = append(i.octets, value)
	return true
}

// Returns whether the byte at the current index is a digit (defined as
// %x30-39  ; 0-9). If true, it increments the index.
func (i *Ipv4) digit() bool {
	c := i.str[i.index]
	if '0' <= c && c <= '9' {
		i.index++
		return true
	}
	return false
}

// If char is at the current index, return true and increment the index.
// If char is not at the current index or the end of str has been reached,
// return false.
func (i *Ipv4) take(char byte) bool {
	if i.index >= i.strLen {
		return false
	}
	if i.str[i.index] == char {
		i.index++
		return true
	}
	return false
}

// NewIpv4 creates a new Ipv4 based on str.
func NewIpv4(str string) *Ipv4 {
	return &Ipv4{
		str:       str,
		index:     0,
		strLen:    int64(len(str)),
		octets:    make([]int64, 0),
		prefixLen: 0,
	}
}
