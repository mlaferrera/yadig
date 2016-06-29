## Overview

yadig (pronounced "you dig") allows for DNS queries from the command line using Google's HTTPS DNS service.

## Install *

- Make sure golang is installed
- ```go get github.com/mlaferrera/yadig```
- ```go build yadig.go```

    \* Eventually I'll release precompiled binaries

## Usage

```
Usage of yadig:
-q string
    Host to conduct a query on
-t string
    DNS record type (default "A")
```

## Examples

* Lookup the ```A``` record for ```cnn.com```

```
./yadig -q cnn.com
Query: cnn.com., DNSSEC: false, Type: A, TTL: 291, Response: 157.166.226.26
Query: cnn.com., DNSSEC: false, Type: A, TTL: 291, Response: 157.166.226.25
```

* Lookup the ```MX``` record for ```cnn.com```

```
./yadig -q cnn.com -t MX
Query: cnn.com., DNSSEC: false, Type: MX, TTL: 299, Response: 10 ppsprmsf.turner.com.
Query: cnn.com., DNSSEC: false, Type: MX, TTL: 299, Response: 10 ppsprmse.turner.com.
Query: cnn.com., DNSSEC: false, Type: MX, TTL: 299, Response: 10 ppsprmsa.turner.com.
Query: cnn.com., DNSSEC: false, Type: MX, TTL: 299, Response: 10 ppsprmsb.turner.com.
Query: cnn.com., DNSSEC: false, Type: MX, TTL: 299, Response: 10 ppsprmsc.turner.com.
Query: cnn.com., DNSSEC: false, Type: MX, TTL: 299, Response: 10 ppsprmsd.turner.com.
Query: cnn.com., DNSSEC: false, Type: MX, TTL: 299, Response: 10 ppsprmsg.turner.com.
Query: cnn.com., DNSSEC: false, Type: MX, TTL: 299, Response: 10 ppsprmsh.turner.com.
```

* Conduct a mass lookup of ```A``` records for domains in a file

```
cat domains.txt | ./yadig
Query: cnn.com., DNSSEC: false, Type: A, TTL: 299, Response: 157.166.226.26
Query: cnn.com., DNSSEC: false, Type: A, TTL: 299, Response: 157.166.226.25
Query: google.com., DNSSEC: false, Type: A, TTL: 299, Response: 74.125.29.102
Query: google.com., DNSSEC: false, Type: A, TTL: 299, Response: 74.125.29.139
Query: google.com., DNSSEC: false, Type: A, TTL: 299, Response: 74.125.29.113
Query: google.com., DNSSEC: false, Type: A, TTL: 299, Response: 74.125.29.100
Query: google.com., DNSSEC: false, Type: A, TTL: 299, Response: 74.125.29.138
Query: google.com., DNSSEC: false, Type: A, TTL: 299, Response: 74.125.29.101
Query: abc.com., DNSSEC: false, Type: A, TTL: 125, Response: 199.181.132.250
Query: msnbc.com., DNSSEC: false, Type: A, TTL: 19, Response: 69.192.61.76
```

* Conduct a mass lookup of ```NS``` records for domains in a file

```
cat domains.txt | ./yadig -t NS
Query: cnn.com., DNSSEC: false, Type: NS, TTL: 21599, Response: ns2.p42.dynect.net.
Query: cnn.com., DNSSEC: false, Type: NS, TTL: 21599, Response: ns1.timewarner.net.
Query: cnn.com., DNSSEC: false, Type: NS, TTL: 21599, Response: ns1.p42.dynect.net.
Query: cnn.com., DNSSEC: false, Type: NS, TTL: 21599, Response: ns3.timewarner.net.
Query: google.com., DNSSEC: false, Type: NS, TTL: 19983, Response: ns4.google.com.
Query: google.com., DNSSEC: false, Type: NS, TTL: 19983, Response: ns2.google.com.
Query: google.com., DNSSEC: false, Type: NS, TTL: 19983, Response: ns1.google.com.
Query: google.com., DNSSEC: false, Type: NS, TTL: 19983, Response: ns3.google.com.
Query: abc.com., DNSSEC: false, Type: NS, TTL: 299, Response: orns01.dig.com.
Query: abc.com., DNSSEC: false, Type: NS, TTL: 299, Response: orns02.dig.com.
Query: abc.com., DNSSEC: false, Type: NS, TTL: 299, Response: sens01.dig.com.
Query: abc.com., DNSSEC: false, Type: NS, TTL: 299, Response: sens02.dig.com.
Query: msnbc.com., DNSSEC: false, Type: NS, TTL: 299, Response: ns1-102.akam.net.
Query: msnbc.com., DNSSEC: false, Type: NS, TTL: 299, Response: eur3.akam.net.
Query: msnbc.com., DNSSEC: false, Type: NS, TTL: 299, Response: asia3.akam.net.
Query: msnbc.com., DNSSEC: false, Type: NS, TTL: 299, Response: eur4.akam.net.
Query: msnbc.com., DNSSEC: false, Type: NS, TTL: 299, Response: usc1.akam.net.
Query: msnbc.com., DNSSEC: false, Type: NS, TTL: 299, Response: aus1.akam.net.
Query: msnbc.com., DNSSEC: false, Type: NS, TTL: 299, Response: ns1-161.akam.net.
Query: msnbc.com., DNSSEC: false, Type: NS, TTL: 299, Response: use3.akam.net.
Query: msnbc.com., DNSSEC: false, Type: NS, TTL: 299, Response: usw1.akam.net.
```

* Conduct a query one at a time via Stdin

```
./yadig
cnn.com
Query: cnn.com., DNSSEC: false, Type: A, TTL: 297, Response: 157.166.226.26
Query: cnn.com., DNSSEC: false, Type: A, TTL: 297, Response: 157.166.226.25
abc.com
Query: abc.com., DNSSEC: false, Type: A, TTL: 68, Response: 199.181.132.250
google.com
Query: google.com., DNSSEC: false, Type: A, TTL: 299, Response: 209.85.144.138
Query: google.com., DNSSEC: false, Type: A, TTL: 299, Response: 209.85.144.100
Query: google.com., DNSSEC: false, Type: A, TTL: 299, Response: 209.85.144.102
Query: google.com., DNSSEC: false, Type: A, TTL: 299, Response: 209.85.144.101
Query: google.com., DNSSEC: false, Type: A, TTL: 299, Response: 209.85.144.113
Query: google.com., DNSSEC: false, Type: A, TTL: 299, Response: 209.85.144.139
```
