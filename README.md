[docimg]:https://godoc.org/github.com/ardnew/ipseq?status.svg
[docurl]:https://godoc.org/github.com/ardnew/ipseq

# ipseq
Go module to parse and iterate over IPv4 address intervals.

[![GoDoc][docimg]][docurl]

## Installation

Use the built-in `go get` tool:

```sh
go get -v github.com/ardnew/ipseq
```

## Example

##### See the [documentation][docurl] for complete usage details.

The following example demonstrates various way IP ranges can be provided, and it also shows how rollover behaves when ranging over multiple octets/subnets.

```go
package main

import (
    "fmt"
    "os"

    "github.com/ardnew/ipseq"
)

// can use os.DevNull to hide error messages
var errors = os.Stderr

func main() {
    ipseq.Errors = errors
    
    // the following calls Seq with three different ranges, which may be
    // specified as comma-separated ranges in a single string, or they may be 
    // given as any number of separate arguments.
    //   1. 192.168.1.0/30        : CIDR bitmask range
    //   2. 10.0.0.254-10.0.1.1   : hyphenated IP range
    //   3. 8.8.8.8               : range containing single IP
    for ip := range ipseq.Seq("192.168.1.0/30,10.0.0.254-10.0.1.1", "8.8.8.8") {
        fmt.Printf("%s\n", ip)
    }
}
```

Running this program will produce the following output:

```
192.168.1.0
192.168.1.1
192.168.1.2
192.168.1.3
10.0.0.254
10.0.0.255
10.0.1.0
10.0.1.1
8.8.8.8
```
