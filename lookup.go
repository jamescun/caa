package caa

import (
	"context"
	"strings"

	"github.com/miekg/dns"
)

// Resolver is the default DNS resolver for CAA Lookups.
var Resolver = "8.8.8.8:53"

// Record is the contents of all CAA related records on a domain.
type Record struct {
	Name string `json:"name"`

	Issuers []Issuer `json:"issuers,omitempty"`

	Reports []string `json:"reports,omitempty"`
}

// Issuer is a certificate authority that is authorized to
// issue certificates for a domain.
type Issuer struct {
	Name string `json:"name"`

	Wildcard bool `json:"wildcard"`

	Critical bool `json:"critical"`
}

// Lookup will traverse the given addr from name to root, appending any
// CAA records found to the result.
func Lookup(c *dns.Client, addr string) ([]*Record, error) {
	return LookupContext(context.Background(), c, addr)
}

// LookupContext will traverse the given addr from name to root, appending any
// CAA records found to the result.
func LookupContext(ctx context.Context, c *dns.Client, addr string) ([]*Record, error) {
	addr = dns.Fqdn(addr)

	var records []*Record

	for len(addr) > 0 {
		if addr == "." {
			break
		}

		record, err := lookupCAA(ctx, c, addr)
		if err != nil {
			return nil, err
		}

		if len(record.Issuers) > 0 || len(record.Reports) > 0 {
			records = append(records, record)
		}

		_, addr = nextInHierarchy(addr)
	}

	return records, nil
}

func lookupCAA(ctx context.Context, c *dns.Client, addr string) (*Record, error) {
	req := new(dns.Msg)
	req.Id = dns.Id()
	req.RecursionDesired = true
	req.Question = []dns.Question{
		dns.Question{addr, dns.TypeCAA, dns.ClassINET},
	}

	res, _, err := c.ExchangeContext(ctx, req, Resolver)
	if err != nil {
		return nil, err
	}

	record := &Record{
		Name: addr,
	}

	for _, answer := range res.Answer {
		if caa, ok := answer.(*dns.CAA); ok {
			switch caa.Tag {
			case "issue":
				record.Issuers = append(record.Issuers, Issuer{
					Name:     caa.Value,
					Critical: caa.Flag > 0,
				})

			case "issuewild":
				record.Issuers = append(record.Issuers, Issuer{
					Name:     caa.Value,
					Wildcard: true,
					Critical: caa.Flag > 0,
				})

			case "iodef":
				record.Reports = append(record.Reports, caa.Value)
			}
		}
	}

	return record, nil
}

// nextInHierarchy strips the left-most name from addr and returns it
// with the rest of the domain hierarchy.
func nextInHierarchy(addr string) (name, rest string) {
	if addr == "." {
		name = "."
		return
	}

	i := strings.IndexByte(addr, '.')
	if i > -1 {
		name = addr[:i]

		if addr[i:] == "." {
			rest = "."
		} else {
			rest = addr[i+1:]
		}
	} else {
		name = addr
	}

	return
}
