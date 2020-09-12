# ipseq
#### Command-line interface for module [`github.com/ardnew/ipseq`](https://github.com/ardnew/ipseq)

## Installation

Use the built-in `go get` tool:

```sh
go get -v github.com/ardnew/ipseq/cmd/ipseq
```

## Usage

##### Refer to the [module documentation](https://godoc.org/github.com/ardnew/ipseq) for complete usage details

This utility is a very thin wrapper to the module. It merely passes all arguments on the command-line to `Seq`, and prints the in-range IPs to stdout, one per line.

Using the same syntax and IP ranges given in the [README](https://github.com/ardnew/ipseq#example) and [docs](https://godoc.org/github.com/ardnew/ipseq#ex-package), the following shows that it has equivalent semantics:

```sh
$ ipseq 192.168.1.0/30,10.0.0.254-10.0.1.1 8.8.8.8
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


