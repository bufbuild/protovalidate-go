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
	"slices"
	"strconv"
)

type Ipv6 struct {
	str             string
	index           int
	strLen          int
	pieces          []int64 // 16-bit pieces found
	doubleColonAt   int     // number of 16-bit pieces found when double colon was found
	doubleColonSeen bool
	dottedRaw       string // dotted notation for right-most 32 bits
	dottedAddr      *Ipv4  // dotted notation successfully parsed as IPv4
	zoneIDFound     bool
	prefixLen       int64 // 0 - 128
}

// Return the 128-bit value of an address parsed through address() or addressPrefix(),
// as a 4-tuple of 32-bit values.
// Return [0,0,0,0] if no address was parsed successfully.
func (i *Ipv6) getBits() [4]int64 {
	p16 := i.pieces
	// handle dotted decimal, add to p16
	if i.dottedAddr != nil {
		dotted32 := i.dottedAddr.getBits() // right-most 32 bits
		p16 = append(p16, dotted32>>16)    // high 16 bits
		p16 = append(p16, dotted32)        // low 16 bits
	}
	// handle double colon, fill pieces with 0
	if i.doubleColonSeen {
		for {
			if len(p16) >= 8 {
				break
			}
			// delete 0 entries at pos, insert a 0
			p16 = slices.Insert(p16, i.doubleColonAt, 0x00000000)
		}
	}
	if len(p16) != 8 {
		return [4]int64{0, 0, 0, 0}
	}
	return [4]int64{
		((p16[0] << 16) | p16[1]),
		((p16[2] << 16) | p16[3]),
		((p16[4] << 16) | p16[5]),
		((p16[6] << 16) | p16[7]),
	}
}

// Return true if all bits to the right of the prefix-length are all zeros.
// Behavior is undefined if addressPrefix() has not been called before, or has
// returned false.
func (i *Ipv6) isPrefixOnly() bool {
	// For each 32-bit piece of the address, require that values to the right of the prefix are zero
	for idx, p32 := range i.getBits() {
		size := i.prefixLen - 32*int64(idx)
		var mask int64
		if size >= 32 { //nolint:gocritic
			mask = 0xffffffff
		} else if size < 0 {
			mask = 0x00000000
		} else {
			mask = ^(0xffffffff >> size)
		}
		masked := (p32 & mask)
		if p32 != masked {
			return false
		}
	}
	return true
}

// Parse IPv6 Address following RFC 4291, with optional zone id following RFC 4007.
func (i *Ipv6) address() bool {
	return i.addressPart() && i.index == i.strLen
}

// Parse IPv6 Address Prefix following RFC 4291. Zone id is not permitted.
func (i *Ipv6) addressPrefix() bool {
	return i.addressPart() &&
		!i.zoneIDFound &&
		i.take('/') &&
		i.prefixLength() &&
		i.index == i.strLen
}

// Determines the prefix length. Stores value in `prefixLen`.
func (i *Ipv6) prefixLength() bool {
	start := i.index
	for {
		if i.index >= i.strLen || !i.digit() {
			break
		}
		if i.index-start > 3 {
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
	if value > 128 {
		// max 128 bits
		return false
	}
	i.prefixLen = value
	return true
}

// Stores dotted notation for right-most 32 bits in `dottedRaw` / `dottedAddr` if found.
func (i *Ipv6) addressPart() bool {
	for {
		if i.index >= i.strLen {
			break
		}
		// dotted notation for right-most 32 bits, e.g. 0:0:0:0:0:ffff:192.1.56.10
		if (i.doubleColonSeen || len(i.pieces) == 6) && i.dotted() {
			dotted := NewIpv4(i.dottedRaw)
			if dotted.address() {
				i.dottedAddr = dotted
				return true
			}
			return false
		}
		if i.h16() {
			continue
		}
		if i.take(':') { //nolint:nestif
			if i.take(':') {
				if i.doubleColonSeen {
					return false
				}
				i.doubleColonSeen = true
				i.doubleColonAt = len(i.pieces)
				if i.take(':') {
					return false
				}
			}
			continue
		}
		if i.str[i.index] == '%' && !i.zoneID() {
			return false
		}
		break
	}
	return i.doubleColonSeen || len(i.pieces) == 8
}

// There is no definition for the character set allowed in the zone
// identifier. RFC 4007 permits basically any non-null string.
//
// RFC 6874 defines it as ZoneID = 1*( unreserved / pct-encoded ).
func (i *Ipv6) zoneID() bool {
	start := i.index
	if i.take('%') {
		if i.strLen-i.index > 0 {
			// permit any non-null string
			i.index = i.strLen
			i.zoneIDFound = true
			return true
		}
	}
	i.index = start
	i.zoneIDFound = false
	return false
}

// Parses str from the current index to determine if it contains a valid
// dotted sequence defined as:
//
// 1*3DIGIT "." 1*3DIGIT "." 1*3DIGIT "." 1*3DIGIT
// Stores match in `dottedRaw`.
func (i *Ipv6) dotted() bool {
	start := i.index
	i.dottedRaw = ""
	for {
		if i.index < i.strLen && (i.digit() || i.take('.')) {
			continue
		}
		break
	}
	if i.index-start >= 7 {
		i.dottedRaw = i.str[start:i.index]
		return true
	}
	i.index = start
	return false
}

// h16 = 1*4HEXDIG
// Stores 16-bit value in `pieces`.
func (i *Ipv6) h16() bool {
	start := i.index
	for {
		if i.index >= i.strLen || !i.hexdig() {
			break
		}
	}
	str := i.str[start:i.index]
	if len(str) == 0 {
		// too short
		return false
	}
	if len(str) > 4 {
		// too long
		return false
	}

	value, err := strconv.ParseInt(str, 16, 32)
	if err != nil {
		return false
	}
	i.pieces = append(i.pieces, value)
	return true
}

// HEXDIG =  DIGIT / "A" / "B" / "C" / "D" / "E" / "F".
func (i *Ipv6) hexdig() bool {
	c := i.str[i.index]
	if ('0' <= c && c <= '9') ||
		('a' <= c && c <= 'f') ||
		('A' <= c && c <= 'F') {
		i.index++
		return true
	}
	return false
}

// DIGIT = %x30-39  ; 0-9.
func (i *Ipv6) digit() bool {
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
func (i *Ipv6) take(char byte) bool {
	if i.index >= i.strLen {
		return false
	}
	if i.str[i.index] == char {
		i.index++
		return true
	}
	return false
}

// NewIpv6 creates a new Ipv6 based on str.
func NewIpv6(str string) *Ipv6 {
	return &Ipv6{
		str:           str,
		strLen:        len(str),
		pieces:        make([]int64, 0),
		doubleColonAt: -1,
		dottedAddr:    nil,
	}
}
