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
	"net/url"
	"strconv"
	"strings"
)

// Returns whether the given string is a valid IP address for the given version.
// If version is 4, it will validate str as an ipv4 address. If version is 6,
// it will validate as ipv6. If version is 0, it will validate that str is
// _either_ ipv4 or ipv6.
func isIP(str string, version int64) bool {
	if version == 6 {
		return NewIpv6(str).address()
	}
	if version == 4 {
		return NewIpv4(str).address()
	}
	if version == 0 {
		return NewIpv4(str).address() || NewIpv6(str).address()
	}
	return false
}

/**
 * Returns true if the string is a valid IP with prefix length, optionally
 * limited to a specific version (v4 or v6), and optionally requiring the host
 * portion to be all zeros.
 *
 * An address prefix divides an IP address into a network portion, and a host
 * portion. The prefix length specifies how many bits the network portion has.
 * For example, the IPv6 prefix "2001:db8:abcd:0012::0/64" designates the
 * left-most 64 bits as the network prefix. The range of the network is 2**64
 * addresses, from 2001:db8:abcd:0012::0 to 2001:db8:abcd:0012:ffff:ffff:ffff:ffff.
 *
 * An address prefix may include a specific host address, for example
 * "2001:db8:abcd:0012::1f/64". With strict = true, this is not permitted. The
 * host portion must be all zeros, as in "2001:db8:abcd:0012::0/64".
 *
 * The same principle applies to IPv4 addresses. "192.168.1.0/24" designates
 * the first 24 bits of the 32-bit IPv4 as the network prefix.
 */
func isIPPrefix(
	str string,
	version int64,
	strict bool,
) bool {
	if version == 6 {
		ip := NewIpv6(str)
		return ip.addressPrefix() && (!strict || ip.isPrefixOnly())
	}
	if version == 4 {
		ip := NewIpv4(str)
		return ip.addressPrefix() && (!strict || ip.isPrefixOnly())
	}
	if version == 0 {
		return isIPPrefix(str, 6, strict) || isIPPrefix(str, 4, strict)
	}
	return false
}

/**
 * Returns true if the string is a valid hostname, for example "foo.example.com".
 *
 * A valid hostname follows the rules below:
 * - The name consists of one or more labels, separated by a dot (".").
 * - Each label can be 1 to 63 alphanumeric characters.
 * - A label can contain hyphens ("-"), but must not start or end with a hyphen.
 * - The right-most label must not be digits only.
 * - The name can have a trailing dot, for example "foo.example.com.".
 * - The name can be 253 characters at most, excluding the optional trailing dot.
 */
func isHostname(val string) bool {
	if len(val) > 253 {
		return false
	}
	var str string
	if strings.HasSuffix(val, ".") {
		str = val[0 : len(val)-1]
	} else {
		str = val
	}

	allDigits := false
	parts := strings.Split(strings.ToLower(str), ".")

	// split hostname on '.' and validate each part
	for _, part := range parts {
		allDigits = true
		// if part is empty, longer than 63 chars, or starts/ends with '-', it is invalid
		l := len(part)
		if l == 0 || l > 63 || strings.HasPrefix(part, "-") || strings.HasSuffix(part, "-") {
			return false
		}
		// for each character in part
		for i := 0; i < len(part); i++ {
			c := part[i]
			// if the character is not a-z, 0-9, or '-', it is invalid
			if (c < 'a' || c > 'z') && (c < '0' || c > '9') && c != '-' {
				return false
			}
			allDigits = allDigits && c >= '0' && c <= '9'
		}
	}
	// the last part cannot be all numbers
	return !allDigits
}

/**
 * Returns true if the string is a valid host/port pair, for example "example.com:8080".
 *
 * If the argument `portRequired` is true, the port is required. If the argument
 * is false, the port is optional.
 *
 * The host can be one of:
 * - An IPv4 address in dotted decimal format, for example "192.168.0.1".
 * - An IPv6 address enclosed in square brackets, for example "[::1]".
 * - A hostname, for example "example.com".
 *
 * The port is separated by a colon. It must be non-empty, with a decimal number
 * in the range of 0-65535, inclusive.
 */
func isHostAndPort(str string, portRequired bool) bool {
	if len(str) == 0 {
		return false
	}
	splitIdx := strings.LastIndex(str, ":")
	if str[0] == '[' {
		end := strings.LastIndex(str, "]")
		switch end + 1 {
		case len(str): // no port
			return !portRequired && isIP(str[1:end], 6)
		case splitIdx: // port
			return isIP(str[1:end], 6) && isPort(str[splitIdx+1:])
		default: // malformed
			return false
		}
	}
	if splitIdx < 0 {
		return !portRequired && (isHostname(str) || isIP(str, 4))
	}
	host := str[0:splitIdx]
	port := str[splitIdx+1:]
	return (isHostname(host) || isIP(host, 4)) && isPort(port)
}

// Returns true if the string is a valid port.
func isPort(str string) bool {
	if len(str) == 0 {
		return false
	}
	for i := 0; i < len(str); i++ {
		c := str[i]
		if '0' <= c && c <= '9' {
			continue
		}
		return false
	}
	val, err := strconv.ParseInt(str, 0, 32)
	if err != nil {
		return false
	}
	return val <= 65535
}

type URI struct {
	str             string
	index           int64
	strLen          int64
	pctEncodedFound bool
}

// Returns true if the string used for instantiation is a valid URI, i.e. it
// follows the syntax URI = scheme ":" hier-part [ "?" query ] [ "#" fragment ].
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
	if u.index != u.strLen {
		u.index = start
		return false
	}
	return true
}

// Returns true if the string used for instantiation is a valid URI reference,
// i.e., it follows the syntax URI-reference = URI / relative-ref.
func (u *URI) uriReference() bool {
	return u.uri() || u.relativeRef()
}

// Parses str from the current index to determine if it contains a valid
// hier-part defined as:
//
// hier-part = "//" authority path-abempty
// / path-absolute
// / path-rootless
// / path-empty.
func (u *URI) hierPart() bool {
	start := u.index
	if u.take('/') && //nolint:staticcheck
		u.take('/') &&
		u.authority() &&
		u.pathAbempty() {
		return true
	}
	u.index = start
	return u.pathAbsolute() || u.pathRootless() || u.pathEmpty()
}

// Parses str from the current index to determine if it contains a valid
// relative-ref defined as:
//
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
	if u.index != u.strLen {
		u.index = start
		return false
	}
	return true
}

// Parses str from the current index to determine if it contains a valid
// relative-part defined as:
//
// relative-part = "//" authority path-abempty.
// path-absolute.
// path-noscheme.
// path-empty.
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

// Parses str from the current index to determine if it contains a valid
// scheme defined as:
//
// scheme = ALPHA *( ALPHA / DIGIT / "+" / "-" / "." )
//
// Terminated by ":".
func (u *URI) scheme() bool {
	start := u.index
	if u.alpha() {
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

// Parses str from the current index to determine if it contains a valid
// authority defined as:
//
// authority = [ userinfo "@" ] host [ ":" port ]
//
// Lead by double slash ("") and terminated by "/", "?", "#", or end of URI.
func (u *URI) authority() bool {
	start := u.index
	if u.userinfo() {
		if !u.take('@') {
			u.index = start
			return false
		}
	}
	if !u.host() {
		u.index = start
		return false
	}
	if u.take(':') {
		if !u.port() {
			u.index = start
			return false
		}
	}
	if !u.isAuthorityEnd() {
		u.index = start
		return false
	}
	return true
}

// The authority component [...] is terminated by the next slash ("/"),
// question mark ("?"), or number > sign ("#") character, or by the
// end of the URI.
func (u *URI) isAuthorityEnd() bool {
	return u.index >= u.strLen ||
		u.str[u.index] == '?' ||
		u.str[u.index] == '#' ||
		u.str[u.index] == '/'
}

// Parses str from the current index to determine if it contains a valid
// userinfo defined as:
//
// userinfo = *( unreserved / pct-encoded / sub-delims / ":" ).
//
// Terminated by "@" in authority.
// If the end of the string is found before the "@" terminator, false is returned.
func (u *URI) userinfo() bool {
	start := u.index
	for {
		if u.unreserved() ||
			u.pctEncoded() ||
			u.subDelims() ||
			u.take(':') {
			continue
		}
		if u.index < u.strLen {
			if u.str[u.index] == '@' {
				return true
			}
		}
		u.index = start
		return false
	}
}

/* TODO - JavaScript's implementation of decodeURIComponent() throws an error if
 * a pct-encoded escape sequence does not encode a valid UTF-8 character.
 *
 * Go does not have an equivalent function and the closest it has is
 * url.PathUnescape, which is what we use below. However, this is not
 * consistent with JavaScript's stricter implementation so we will have to
 * implement our own.
 * For example:
 * - Decode pct-encoded rawHost
 *   - Allocate an octet array
 *   - For every octet in rawHost
 *     - For "%", percent-decode the following two hex digits to an
 *       octet, add it to the octet array
 *     - For every other octet, add it to the octet array
 * - Check that the octet array is valid UTF-8.
 */
func (u *URI) decodeURIComponent(str string) bool {
	if _, err := url.PathUnescape(str); err != nil {
		return false
	}
	return true
}

// Parses str from the current index to determine if it contains a valid
// host defined as:
//
// host = IP-literal / IPv4address / reg-name.
func (u *URI) host() bool {
	if u.index >= u.strLen {
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
			if !u.decodeURIComponent(rawHost) {
				return false
			}
		}
		return true
	}
	return false
}

// Parses str from the current index to determine if it contains a valid
// port defined as:
//
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

// Parses str from the current index to determine if it contains a valid
// IP-literal defined by RFC-6874 as:
//
// IP-literal = "[" ( IPv6address / IPv6addrz / IPvFuture  ) "]".
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

// Parses str from the current index to determine if it contains a valid
// IPv6 address. Relies on the implementation of isIP(str, 6) to match RFC 3986
// grammar.
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

// Parses str from the current index to determine if it contains a valid
// IPv6addrz defined by RFC 6874 as:
//
// IPv6addrz = IPv6address "%25" ZoneID.
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

// Parses str from the current index to determine if it contains a valid
// ZoneID defined by RFC 6874 as:
//
// ZoneID = 1*( unreserved / pct-encoded ).
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

// Parses str from the current index to determine if it contains a valid
// IPvFuture defined as:
//
// IPvFuture  = "v" 1*HEXDIG "." 1*( unreserved / sub-delims / ":" ).
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

// Parses str from the current index to determine if it contains a valid
// reg-name defined as:
//
// reg-name = *( unreserved / pct-encoded / sub-delims ).
//
// Terminates on start of port (":") or end of authority.
func (u *URI) regName() bool {
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
		return false
	}
}

// The path is terminated by the first question mark ("?") or
// number sign ("#") character, or by the end of the URI.
func (u *URI) isPathEnd() bool {
	return u.index >= u.strLen || u.str[u.index] == '?' || u.str[u.index] == '#'
}

// Parses str from the current index to determine if it contains a valid
// path-abempty defined as:
//
// path-abempty = *( "/" segment )
//
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

// Parses str from the current index to determine if it contains a valid
// path-absolute defined as:
//
// path-absolute = "/" [ segment-nz *( "/" segment ) ]
//
// Terminated by end of path: "?", "#", or end of URI.
func (u *URI) pathAbsolute() bool {
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

// Parses str from the current index to determine if it contains a valid
// path-noscheme defined as:
//
// path-noscheme = segment-nz-nc *( "/" segment )
//
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

// Parses str from the current index to determine if it contains a valid
// path-rootless defined as:
//
// path-rootless = segment-nz *( "/" segment )
//
// Terminated by end of path: "?", "#", or end of URI.
func (u *URI) pathRootless() bool {
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

// Parses str from the current index to determine if it contains a valid
// path-empty defined as:
//
// path-empty = 0<pchar>
//
// Terminated by end of path: "?", "#", or end of URI.
func (u *URI) pathEmpty() bool {
	return u.isPathEnd()
}

// Parses str from the current index to determine if it contains a valid
// segment defined as:
//
// segment = *pchar.
func (u *URI) segment() bool {
	for {
		if !u.pchar() {
			break
		}
	}
	return true
}

// Parses str from the current index to determine if it contains a valid
// segment-nz defined as:
//
// segment-nz = 1*pchar.
func (u *URI) segmentNz() bool {
	start := u.index
	if u.pchar() {
		return u.segment()
	}
	u.index = start
	return false
}

// Parses str from the current index to determine if it contains a valid
// segment-nz-nc defined as:
//
// segment-nz-nc = 1*( unreserved / pct-encoded / sub-delims / "@" );
// non-zero-length segment without any colon ":".
func (u *URI) segmentNzNc() bool {
	start := u.index
	for {
		if !u.unreserved() &&
			!u.pctEncoded() &&
			!u.subDelims() &&
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

// Parses str from the current index to determine if it contains a valid
// pchar defined as:
//
// pchar = unreserved / pct-encoded / sub-delims / ":" / "@".
func (u *URI) pchar() bool {
	return (u.unreserved() ||
		u.pctEncoded() ||
		u.subDelims() ||
		u.take(':') ||
		u.take('@'))
}

// Parses str from the current index to determine if it contains a valid
// query defined as:
//
// query = *( pchar / "/" / "?" )
// Terminated by "#" or end of URI.
func (u *URI) query() bool {
	start := u.index
	for {
		if u.pchar() || u.take('/') || u.take('?') {
			continue
		}
		if u.index == u.strLen || u.str[u.index] == '#' {
			return true
		}
		u.index = start
		return false
	}
}

// Parses str from the current index to determine if it contains a valid
// fragment defined as:
//
// fragment = *( pchar / "/" / "?" )
// Terminated by end of URI.
func (u *URI) fragment() bool {
	start := u.index
	for {
		if u.pchar() || u.take('/') || u.take('?') {
			continue
		}
		if u.index == u.strLen {
			return true
		}
		u.index = start
		return false
	}
}

// Parses str from the current index to determine if it contains a valid
// pct encoding defined as:
//
// pct-encoded = "%"+HEXDIG+HEXDIG.
// Sets `pctEncodedFound` to true if a valid triplet was found.
func (u *URI) pctEncoded() bool {
	start := u.index
	if u.take('%') && u.hexdig() && u.hexdig() {
		u.pctEncodedFound = true
		return true
	}
	u.index = start
	return false
}

// Returns whether the byte at the current index is an unreserved character
// defined as:
//
// unreserved = ALPHA / DIGIT / "-" / "." / "_" / "~".
func (u *URI) unreserved() bool {
	return (u.alpha() ||
		u.digit() ||
		u.take('-') ||
		u.take('_') ||
		u.take('.') ||
		u.take('~'))
}

// Returns whether the byte at the current index is a subdelim defined as:
//
// sub-delims  = "!" / "$" / "&" / "'" / "(" / ")".
// / "*" / "+" / "," / ";" / "=".
func (u *URI) subDelims() bool {
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

// Returns whether the byte at the current index is an alpha character (defined
// as ALPHA =  %x41-5A / %x61-7A ; A-Z / a-z.
// If true, it increments the index.
func (u *URI) alpha() bool {
	if u.index >= u.strLen {
		return false
	}
	c := u.str[u.index]
	if ('A' <= c && c <= 'Z') || ('a' <= c && c <= 'z') {
		u.index++
		return true
	}
	return false
}

// Returns whether the byte at the current index is a digit (defined as
// %x30-39  ; 0-9). If true, it increments the index.
func (u *URI) digit() bool {
	if u.index >= u.strLen {
		return false
	}
	c := u.str[u.index]
	if '0' <= c && c <= '9' {
		u.index++
		return true
	}
	return false
}

// Returns whether the byte at the current index is a hexadecimal digit (defined
// as HEXDIG =  DIGIT / "A" / "B" / "C" / "D" / "E" / "F". If true, it
// increments the index.
func (u *URI) hexdig() bool {
	if u.index >= u.strLen {
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
// If char is not at the current index or the end of str has been reached,
// return false.
func (u *URI) take(char byte) bool {
	if u.index >= u.strLen {
		return false
	}
	if u.str[u.index] == char {
		u.index++
		return true
	}
	return false
}

// NewURI creates a new URI based on str.
func NewURI(str string) *URI {
	return &URI{
		str:             str,
		index:           0,
		strLen:          int64(len(str)),
		pctEncodedFound: false,
	}
}
