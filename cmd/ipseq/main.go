package main

import (
	"fmt"
	"os"

	"github.com/ardnew/ipseq"
)

func main() {
	// capture the errors, which are discarded by default
	ipseq.Errors = os.Stderr
	for ip := range ipseq.Seq(os.Args[1:]...) {
		fmt.Println(ip)
	}
}
