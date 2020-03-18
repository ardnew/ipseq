// Package ipseq provides a method to easily parse and iterate over ranges of
// IPv4 addresses.
package ipseq

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"strings"
)

// Errors contains the io.Writer where errors will be written.
var Errors io.Writer = ioutil.Discard

// IPSeq is a channel of discrete IPv4 address sequences.
type IPSeq chan net.IP

// Seq returns a channel being populated with IPv4 addresses from a separate
// goroutine. The channel is closed after all addresses have been added.
// Each string s contains comma-separated ranges, and each range can be expressed
// in either CIDR format or as a hyphenated interval.
func Seq(s ...string) IPSeq {
	q := make(IPSeq)
	q.gen(s...)
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

// gen creates a single goroutine to parse all ranges in each comma-separated
// string s, closing the channel when complete.
func (q IPSeq) gen(s ...string) {
	go func() {
		for _, t := range s {
			sub := strings.Split(t, ",")
			for _, r := range sub {
				q.parse(r)
			}
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
	} else {
		fmt.Fprintf(Errors, "invalid range: %q", s)
	}
}
