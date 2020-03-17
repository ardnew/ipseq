// Package ipseq provides a method to easily parse and iterate over ranges of
// IPv4 addresses.
package ipseq

import (
	"net"
	"strings"
)

// IPSeq is a channel of discrete IPv4 address sequences.
type IPSeq chan net.IP

// MakeIPSeq returns a new IPSeq. The given string s is parsed and the channel
// starts filling with the individual IPv4 addresses in a separate goroutine
// immediately.
// The string s contains comma-separated ranges, and each range can be expressed
// in either CIDR format or as a hyphenated interval.
func MakeIPSeq(s string) IPSeq {
	q := make(IPSeq)
	q.generate(s)
	return q
}

// maskToUint32 converts a 4-byte slice to its corresponding uint32 value.
func maskToUint32(mask []byte) uint32 {
	u := uint32(0)
	for i, b := range mask {
		u |= uint32(b) << (8 * (3 - i))
	}
	return u
}

// IPToUint32 converts an IPv4 address to its corresponding uint32 value.
func IPToUint32(ip net.IP) uint32 {
	return maskToUint32([]byte(ip.To4()))
}

// Uint32ToIP converts a uint32 value to its corresponding IPv4 address.
func Uint32ToIP(u uint32) net.IP {
	return net.IPv4(uint8(u>>24), uint8(u>>16), uint8(u>>8), uint8(u))
}

// generate dispatches a goroutine to parse each range in the comma-separated
// string s.
func (q IPSeq) generate(s string) {
	go func() {
		sub := strings.Split(s, ",")
		for _, r := range sub {
			q.parse(r)
		}
		close(q)
	}()
}

// parse writes to the channel each IPv4 address in sequence in the given range
// parsed from string s.
func (q IPSeq) parse(s string) {
	var lo, hi uint32
	if p := strings.Split(s, "/"); len(p) == 2 {
		// CIDR format
		if ip, sub, err := net.ParseCIDR(s); nil == err {
			_, bits := sub.Mask.Size()
			diff := maskToUint32([]byte(net.CIDRMask(bits, bits))) - maskToUint32([]byte(sub.Mask))
			lo = IPToUint32(ip)
			hi = lo + diff
		}
	} else {
		if p := strings.Split(s, "-"); len(p) == 2 {
			// hyphenated range
			ipLo, ipHi := net.ParseIP(p[0]), net.ParseIP(p[1])
			if nil != ipLo && nil != ipHi {
				lo, hi = IPToUint32(ipLo), IPToUint32(ipHi)
			}
		} else {
			// bare IP
			ipLo := net.ParseIP(s)
			if nil != ipLo {
				lo = IPToUint32(ipLo)
				hi = lo
			}
		}
	}
	if lo != 0 && hi != 0 {
		for i := lo; i <= hi; i++ {
			q <- Uint32ToIP(i)
		}
	}
}
