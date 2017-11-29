package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/jamescun/caa"
	"github.com/miekg/dns"
)

// command line usage
const Usage = `CAA Record Validator
Usage: caa [options] <domain>

Options:
  --resolver  <ip[:port]>  define custom resolver (defaults to Google DNS)
  --json                   output JSON rather than human readable output
`

// command line options
var (
	Resolver = flag.String("resolver", "8.8.8.8:53", "define custom resolver (defaults to system)")

	JSON = flag.Bool("json", false, "output JSON rather than human readable output")
)

type Result struct {
	Records []*caa.Record `json:"records"`
}

func main() {
	flag.Usage = func() { fmt.Fprintln(os.Stderr, Usage) }
	flag.Parse()

	caa.Resolver = appendPort(*Resolver, "53")

	result, err := run(flag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
		return
	}

	if *JSON {
		json.NewEncoder(os.Stdout).Encode(result)
	} else {
		fmt.Printf("CAA Records for %s\n", flag.Arg(0))

		for _, record := range result.Records {
			fmt.Printf("  %s\n", record.Name)

			for _, issuer := range record.Issuers {
				if issuer.Wildcard {
					fmt.Printf("    issuewild  %s\n", issuer.Name)
				} else {
					fmt.Printf("    issue      %s\n", issuer.Name)
				}
			}

			for _, report := range record.Reports {
				fmt.Printf("    iodef      %s\n", report)
			}
		}
	}
}

func run(addr string) (res Result, err error) {
	res.Records, err = caa.Lookup(new(dns.Client), addr)
	if err != nil {
		return
	}

	return
}

func appendPort(addr, defaultPort string) string {
	_, port, _ := net.SplitHostPort(addr)
	if port == "" {
		return addr + defaultPort
	}

	return addr
}
