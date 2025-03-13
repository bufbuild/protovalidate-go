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
	"fmt"
	"net/url"
	"os"
)

func isIP(str string, version int64) bool {
	if version == 6 {
		return true
		// return new Ipv6(str).address();
	}
	if version == 4 {
		return NewIpv4(str).address()
	}
	// if (version === undefined || version == 0) {
	// return new Ipv4(str).address() || new Ipv6(str).address();
	// }
	return false
}

type URI struct {
	str             string
	index           int64
	l               int64
	pctEncodedFound bool
}

// func (u *URI) log(str string) {
// 	if u.str == "https://example.com##" {
// 	}
// }

// URI = scheme ":" hier-part [ "?" query ] [ "#" fragment ].
func (u *URI) uri() bool {
	start := u.index
	if !(u.scheme() && u.take(':') && u.hierPart()) {
		u.index = start
		return false
	}
	if u.take('?') && !u.query() {
		return false
	}
	if u.take('#') && !u.fragment() {
		return false
	}
	if u.index != u.l {
		u.index = start
		return false
	}
	return true
}

func (u *URI) log(s string) {
	fmt.Fprintf(os.Stderr, "%s: index:%d strlen:%d\n", s, u.index, u.l)
}

// hier-part = "//" authority path-abempty
// path-absolute
// path-rootless
// path-empty
func (u *URI) hierPart() bool {
	u.log("hierpart START")
	start := u.index
	if u.take('/') && //nolint:staticcheck
		u.take('/') &&
		u.authority() &&
		u.pathAbempty() {
		return true
	}
	u.log("hierpart done")
	u.index = start
	return u.pathAbsolute() || u.pathRootless() || u.pathEmpty()
}

// URI-reference = URI / relative-ref.
func (u *URI) uriReference() bool {
	return u.uri() || u.relativeRef()
}

// relative-ref = relative-part [ "?" query ] [ "#" fragment ].
func (u *URI) relativeRef() bool {
	start := u.index
	if !u.relativePart() {
		return false
	}
	if u.take('?') && !u.query() {
		u.index = start
		return false
	}
	if u.take('#') && !u.fragment() {
		u.index = start
		return false
	}
	if u.index != u.l {
		u.index = start
		return false
	}
	return true
}

// relative-part = "//" authority path-abempty
// path-absolute
// path-noscheme
// path-empty
func (u *URI) relativePart() bool {
	start := u.index
	if u.take('/') && //nolint:staticcheck
		u.take('/') &&
		u.authority() &&
		u.pathAbempty() {
		return true
	}
	u.index = start
	return u.pathAbsolute() || u.pathNoscheme() || u.pathEmpty()
}

// scheme = ALPHA *( ALPHA / DIGIT / "+" / "-" / "." )
// Terminated by ":".
func (u *URI) scheme() bool {
	start := u.index
	if u.alpha() {
		u.log("scheme loop start")
		for {
			if !u.alpha() && !u.digit() && !u.take('+') && !u.take('-') && !u.take('.') {
				break
			}
		}
		if u.str[u.index] == ':' {
			return true
		}
	}
	u.index = start
	return false
}

// authority = [ userinfo "@" ] host [ ":" port ]
// Lead by double slash ("").
// Terminated by "/", "?", "#", or end of URI.
func (u *URI) authority() bool {
	start := u.index
	u.log("authority start")
	if u.userinfo() {
		if !u.take('@') {
			u.index = start
			return false
		}
	}
	u.log("userinfo block over")
	if !u.host() {
		u.index = start
		return false
	}
	u.log("found a host")
	if u.take(':') {
		u.log("checking port")
		if !u.port() {
			u.log("no port")
			u.index = start
			return false
		}
	}
	if !u.isAuthorityEnd() {
		u.index = start
		return false
	}
	u.log("authority done")
	return true
}

// > The authority component [...] is terminated by the next slash ("/"),
// > question mark ("?"), or number > sign ("#") character, or by the
// > end of the URI.
func (u *URI) isAuthorityEnd() bool {
	return u.index >= u.l ||
		u.str[u.index] == '?' ||
		u.str[u.index] == '#' ||
		u.str[u.index] == '/'
}

// userinfo = *( unreserved / pct-encoded / sub-delims / ":" )
// Terminated by "@" in authority.
// If the end of the string is found before the "@" terminator, false is returned
func (u *URI) userinfo() bool {
	start := u.index
	for {
		if u.unreserved() ||
			u.pctEncoded() ||
			u.subDelims() ||
			u.take(':') {
			continue
		}
		if u.index < u.l {
			if u.str[u.index] == '@' {
				u.log("found an at returning")
				return true
			}
		}
		u.index = start
		return false
	}
}

// host = IP-literal / IPv4address / reg-name
func (u *URI) host() bool {
	if u.index >= u.l {
		return false
	}
	start := u.index
	u.pctEncodedFound = false
	// Note: IPv4address is a subset of reg-name
	if (u.str[u.index] == '[' && u.ipLiteral()) || u.regName() {
		if u.pctEncodedFound {
			rawHost := u.str[start:u.index]
			// RFC 3986:
			// > URI producing applications must not use percent-encoding in host
			// > unless it is used to represent a UTF-8 character sequence.
			if _, err := url.PathUnescape(rawHost); err != nil {
				return false
			}
			// decodeURIComponent() throws an error if a pct-encoded escape
			// sequence does not encode a valid UTF-8 character.
			// Other implementations may have to implement this check themselves.
			// For example:
			// - Decode pct-encoded rawHost
			//   - Allocate an octet array
			//   - For every octet in rawHost
			//     - For "%", percent-decode the following two hex digits to an
			//       octet, add it to the octet array
			//     - For every other octet, add it to the octet array
			// - Check that the octet array is valid UTF-8
			// decodeURIComponent(rawHost)
		}
		return true
	}
	u.log("not a valid host--")
	return false
}

// port = *DIGIT
// Terminated by end of authority.
func (u *URI) port() bool {
	start := u.index
	for {
		if u.digit() {
			continue
		}
		if u.isAuthorityEnd() {
			return true
		}
		u.index = start
		return false
	}
}

// RFC 6874:
// IP-literal = "[" ( IPv6address / IPv6addrz / IPvFuture  ) "]"
func (u *URI) ipLiteral() bool {
	start := u.index
	if u.take('[') {
		currIdx := u.index
		if u.ipv6Address() && u.take(']') {
			return true
		}
		u.index = currIdx
		if u.ipv6addrz() && u.take(']') {
			return true
		}
		u.index = currIdx
		if u.ipvFuture() && u.take(']') {
			return true
		}
	}
	u.index = start
	return false
}

// IPv6address
// Relies on the implementation of isIP6() to match the RFC 3986 grammar.
func (u *URI) ipv6Address() bool {
	start := u.index
	for {
		if !u.hexdig() && !u.take(':') {
			break
		}
	}
	if isIP(u.str[start:u.index], 6) {
		return true
	}
	u.index = start
	return false
}

// RFC 6874:
// IPv6addrz = IPv6address "%25" ZoneID
func (u *URI) ipv6addrz() bool {
	start := u.index
	if u.ipv6Address() &&
		u.take('%') &&
		u.take('2') &&
		u.take('5') &&
		u.zoneID() {
		return true
	}
	u.index = start
	return false
}

// RFC 6874:
// ZoneID = 1*( unreserved / pct-encoded )
func (u *URI) zoneID() bool {
	start := u.index
	for {
		if !u.unreserved() && !u.pctEncoded() {
			break
		}
	}
	if u.index-start > 0 {
		return true
	}
	u.index = start
	return false
}

// IPvFuture  = "v" 1*HEXDIG "." 1*( unreserved / sub-delims / ":" )
func (u *URI) ipvFuture() bool {
	start := u.index
	if u.take('v') && u.hexdig() { //nolint:nestif
		for {
			if !u.hexdig() {
				break
			}
		}
		if u.take('.') {
			counter := 0
			for {
				if !u.unreserved() && !u.subDelims() && !u.take(':') {
					break
				}
				counter++
			}
			if counter >= 1 {
				return true
			}
		}
	}
	u.index = start
	return false
}

// reg-name = *( unreserved / pct-encoded / sub-delims )
// Terminates on start of port (":") or end of authority.
func (u *URI) regName() bool {
	u.log("regname")
	start := u.index
	for {
		if u.unreserved() || u.pctEncoded() || u.subDelims() {
			continue
		}
		if u.isAuthorityEnd() {
			// End of authority
			return true
		}
		if u.str[u.index] == ':' {
			return true
		}
		u.index = start
		u.log("regname END")
		return false
	}
}

// > The path is terminated by the first question mark ("?") or
// > number sign ("#") character, or by the end of the URI.
func (u *URI) isPathEnd() bool {
	u.log("pathend")
	return u.index >= u.l || u.str[u.index] == '?' || u.str[u.index] == '#'
}

// path-abempty = *( "/" segment )
// Terminated by end of path: "?", "#", or end of URI.
func (u *URI) pathAbempty() bool {
	start := u.index
	for {
		if !u.take('/') || !u.segment() {
			break
		}
	}
	if u.isPathEnd() {
		return true
	}
	u.index = start
	return false
}

// path-absolute = "/" [ segment-nz *( "/" segment ) ]
// Terminated by end of path: "?", "#", or end of URI.
func (u *URI) pathAbsolute() bool {
	u.log("pathAbsolute")
	start := u.index
	if u.take('/') {
		if u.segmentNz() {
			for {
				if !u.take('/') || !u.segment() {
					break
				}
			}
		}
		if u.isPathEnd() {
			return true
		}
	}
	u.index = start
	return false
}

// path-noscheme = segment-nz-nc *( "/" segment )
// Terminated by end of path: "?", "#", or end of URI.
func (u *URI) pathNoscheme() bool {
	start := u.index
	if u.segmentNzNc() {
		for {
			if !u.take('/') || !u.segment() {
				break
			}
		}
		if u.isPathEnd() {
			return true
		}
	}
	u.index = start
	return false
}

// path-rootless = segment-nz *( "/" segment )
// Terminated by end of path: "?", "#", or end of URI.
func (u *URI) pathRootless() bool {
	u.log("pathRootless")
	start := u.index
	if u.segmentNz() {
		for {
			if !u.take('/') || !u.segment() {
				break
			}
		}
		if u.isPathEnd() {
			return true
		}
	}
	u.index = start
	return false
}

// path-empty = 0<pchar>
// Terminated by end of path: "?", "#", or end of URI.
func (u *URI) pathEmpty() bool {
	u.log("pathEmpty")
	return u.isPathEnd()
}

// segment = *pchar
func (u *URI) segment() bool {
	for {
		if !u.pchar() {
			break
		}
	}
	return true
}

// segment-nz = 1*pchar
func (u *URI) segmentNz() bool {
	start := u.index
	if u.pchar() {
		return u.segment()
	}
	u.index = start
	return false
}

// segment-nz-nc = 1*( unreserved / pct-encoded / sub-delims / "@" )
//               ; non-zero-length segment without any colon ":"
func (u *URI) segmentNzNc() bool {
	start := u.index
	for {
		if !u.unreserved() ||
			!u.pctEncoded() ||
			!u.subDelims() ||
			!u.take('@') {
			break
		}
	}
	if u.index-start > 0 {
		return true
	}
	u.index = start
	return false
}

// pchar = unreserved / pct-encoded / sub-delims / ":" / "@"
func (u *URI) pchar() bool {
	return (u.unreserved() ||
		u.pctEncoded() ||
		u.subDelims() ||
		u.take(':') ||
		u.take('@'))
}

// query = *( pchar / "/" / "?" )
// Terminated by "#" or end of URI.
func (u *URI) query() bool {
	start := u.index
	for {
		if u.pchar() || u.take('/') || u.take('?') {
			continue
		}
		if u.index == u.l || u.str[u.index] == '#' {
			return true
		}
		u.index = start
		return false
	}
}

// fragment = *( pchar / "/" / "?" )
// Terminated by end of URI.
func (u *URI) fragment() bool {
	start := u.index
	for {
		if u.pchar() || u.take('/') || u.take('?') {
			continue
		}
		if u.index == u.l {
			return true
		}
		u.index = start
		return false
	}
}

// pct-encoded = "%"+HEXDIG+HEXDIG
// Sets `pctEncodedFound` to true if a valid triplet was found
func (u *URI) pctEncoded() bool {
	start := u.index
	if u.take('%') && u.hexdig() && u.hexdig() {
		u.pctEncodedFound = true
		return true
	}
	u.index = start
	return false
}

// unreserved = ALPHA / DIGIT / "-" / "." / "_" / "~"
func (u *URI) unreserved() bool {
	u.log("unreserved")
	return (u.alpha() ||
		u.digit() ||
		u.take('-') ||
		u.take('_') ||
		u.take('.') ||
		u.take('~'))
}

// sub-delims  = "!" / "$" / "&" / "'" / "(" / ")"
//   / "*" / "+" / "," / ";" / "="
func (u *URI) subDelims() bool {
	u.log("subdelims")
	return (u.take('!') ||
		u.take('$') ||
		u.take('&') ||
		u.take('\'') ||
		u.take('(') ||
		u.take(')') ||
		u.take('*') ||
		u.take('+') ||
		u.take(',') ||
		u.take(';') ||
		u.take('='))
}

// ALPHA =  %x41-5A / %x61-7A ; A-Z / a-z
// Terminated by the end of the URI.
func (u *URI) alpha() bool {
	if u.index >= u.l {
		return false
	}
	c := u.str[u.index]
	if ('A' <= c && c <= 'Z') || ('a' <= c && c <= 'z') {
		u.index++
		return true
	}
	return false
}

// digit returns whether the byte at the current index is a digit (i.e. in the
// %x30-39  ; 0-9). If true, it increments the index.
// Terminated by the end of the URI.
func (u *URI) digit() bool {
	if u.index >= u.l {
		return false
	}
	c := u.str[u.index]
	if '0' <= c && c <= '9' {
		u.index++
		return true
	}
	return false
}

// HEXDIG =  DIGIT / "A" / "B" / "C" / "D" / "E" / "F"
func (u *URI) hexdig() bool {
	if u.index >= u.l {
		return false
	}
	c := u.str[u.index]
	if ('0' <= c && c <= '9') ||
		('a' <= c && c <= 'f') ||
		('A' <= c && c <= 'F') {
		u.index++
		return true
	}
	return false
}

// If char is at the current index, return true and increment the index.
// Otherwise return false
// Terminated by the end of the URI.
func (u *URI) take(char byte) bool {
	if u.index >= u.l {
		return false
	}
	if u.str[u.index] == char {
		u.index++
		return true
	}
	return false
}

func NewURI(str string) *URI {
	return &URI{
		str:             str,
		index:           0,
		l:               int64(len(str)),
		pctEncodedFound: false,
	}
}
