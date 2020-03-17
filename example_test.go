package ipseq_test

import (
	"fmt"

	"github.com/ardnew/ipseq"
)

func ExampleMakeIPSeq() {
	seq := ipseq.MakeIPSeq("192.168.1.0/30,10.0.0.10-10.0.0.12")
	for ip := range seq {
		fmt.Printf("%s\n", ip)
	}
	// Output:
	// 192.168.1.0
	// 192.168.1.1
	// 192.168.1.2
	// 192.168.1.3
	// 10.0.0.10
	// 10.0.0.11
	// 10.0.0.12
}
