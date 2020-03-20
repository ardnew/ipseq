package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ardnew/ipseq"
)

func changeLog() {
	ipseq.PrintChangeLog()
}

func version() {
	fmt.Printf("ipseq version %s\n", ipseq.Version())
}

func main() {

	var argChangeLog, argVersion bool
	flag.BoolVar(&argChangeLog, "changelog", false, "display change history")
	flag.BoolVar(&argVersion, "version", false, "display version information")
	flag.Parse()

	if argChangeLog {
		changeLog()
	} else if argVersion {
		version()
	} else {
		// capture the errors, which are discarded by default
		ipseq.Errors = os.Stderr
		for ip := range ipseq.Seq(os.Args[1:]...) {
			fmt.Println(ip)
		}
	}
}
