package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
)

func escape(s string) (string, error) {
	return url.QueryEscape(s), nil
}

func trimToLen(s string, maxlen int) string {
	if len(s) <= maxlen || maxlen < 0 {
		return s
	}
	return s[:maxlen]
}

func main() {
	decode := flag.Bool("d", false, "Decode")
	flag.Parse()

	info, err := os.Stdin.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	var args []string
	if info.Mode()&os.ModeCharDevice == 0 {
		in, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not read from stdin\n")
			os.Exit(1)
		}
		args = []string{string(in)}
	} else {
		args = flag.CommandLine.Args()
	}

	exitcode := 0
	var op func(string) (string, error) = escape
	var opLabel = "encode"
	if *decode {
		op = url.QueryUnescape
		opLabel = "decode"
	}

	for _, a := range args {
		v, err := op(a)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not %s %s: %v\n", opLabel, trimToLen(a, 10), err)
			exitcode = 1
			continue
		}
		fmt.Println(v)
	}
	os.Exit(exitcode)
}
