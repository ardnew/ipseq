package ipseq_test

import (
	"fmt"
	"os"

	"github.com/ardnew/ipseq"
)

func ExampleErrors() {
	ipseq.Errors = os.Stdout
	for ip := range ipseq.Seq("123") {
		fmt.Println(ip)
	}
	// Output:
	// invalid range: "123"
}
