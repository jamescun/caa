# CAA Validator

CAA Validator is a Go package and command line utility for fetching and validating CAA Certificate Authority Authorization records from DNS.


## Getting Started

### Setup

Grab the latest binary from the [releases](https://github.com/jamescun/caa/releases) page.

To build from source or hack on CAA Validator, you can install via `go get`:

```sh
go get -u github.com/jamescun/caa/cmd
```


### Usage

To fetch all CAA records associated with the domain `www.digicert.com`, run:

```sh
$ caa www.digicert.com
CAA Records for www.digicert.com
  digicert.com.
    issuewild  digicert.com
    issue      digicert.com
```

This will fetch all CAA records associated with `www.digicert.com`, `digicert.com` and `com` using Google DNS.

To use a resolver other than Google DNS (inheriting system resolver is TODO), such as OpenDNS, apply the `--resolver` option:

```sh
$ caa --resolver=208.67.222.222 www.digicert.com
CAA Records for www.digicert.com
  digicert.com.
    issuewild  digicert.com
    issue      digicert.com
```

To output JSON as opposed to a human readable output, apply the `--json` option:

```sh
$ caa --json www.digicert.com
{
  "records":[
    {
      "name":"digicert.com.",
      "issuers":[
        {
          "name":"digicert.com",
          "wildcard":false,
          "critical":false
        },
        {
          "name":"digicert.com",
          "wildcard":true,
          "critical":false
        }
      ]
    }
  ]
}
```
