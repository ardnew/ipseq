package ipseq_test

import (
	"fmt"

	"github.com/ardnew/ipseq"
)

func ExampleMakeIPSeq() {
	for ip := range ipseq.Seq("192.168.1.0/30,10.0.0.254-10.0.1.1", "8.8.8.8") {
		fmt.Printf("%s\n", ip)
	}
	// Output:
	// 192.168.1.0
	// 192.168.1.1
	// 192.168.1.2
	// 192.168.1.3
	// 10.0.0.254
	// 10.0.0.255
	// 10.0.1.0
	// 10.0.1.1
	// 8.8.8.8
}
