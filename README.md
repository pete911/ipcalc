# ipcalc

[![pipeline](https://github.com/pete911/ipcalc/actions/workflows/pipeline.yml/badge.svg)](https://github.com/pete911/ipcalc/actions/workflows/pipeline.yml)

Go implementation of ipcalc command.

## build

`go build` or `go install`

## download

- [binary](https://github.com/pete911/ipcalc/releases)

## build/install

### brew

- add tap `brew tap pete911/tap`
- install `brew install pete911/tap/ipcalc`

### go

[go](https://golang.org/dl/) has to be installed.
- build `go build` (or with `Taskfile.yml` - `task`)

## release

Releases are published when the new tag is created e.g.
`git tag -m "add super cool feature" v1.0.0 && git push --follow-tags`

## run

ipcalc takes either:
- one argument: `<CIDR>` ([ipcalc](http://jodies.de/ipcalc))
- two arguments: `<CIDR> <hostnum>` ([terraform cidrhost](https://www.terraform.io/docs/language/functions/cidrhost.html))
- three arguments: `<CIDR> <newbits> <netnum>` ([terraform cidrsubnet](https://www.terraform.io/docs/language/functions/cidrsubnet.html))

### ipcalc example

```
ipcalc 192.168.0.2/30

Address:    192.168.0.2           11000000.10101000.00000000.00000010
Netmask:    255.255.255.252 = 30  11111111.11111111.11111111.11111100
Wildcard:   0.0.0.3               00000000.00000000.00000000.00000011
=>
Network:    192.168.0.0/30        11000000.10101000.00000000.00000000
Broadcast:  192.168.0.3           11000000.10101000.00000000.00000011
HostMin:    192.168.0.1           11000000.10101000.00000000.00000001
HostMax:    192.168.0.2           11000000.10101000.00000000.00000010
Hosts/Net:  2
TotalHosts: 4
```

### terraform cidrhost example

```
ipcalc 10.12.127.0/20 268

Address:    10.12.128.12          00001010.00001100.10000000.00001100
Netmask:    255.255.255.255 = 32  11111111.11111111.11111111.11111111
Wildcard:   0.0.0.0               00000000.00000000.00000000.00000000
=>
Network:    10.12.128.12/32       00001010.00001100.10000000.00001100
Broadcast:  N/A                   N/A
HostMin:    N/A                   N/A
HostMax:    N/A                   N/A
Hosts/Net:  0
TotalHosts: 1
```

### terraform cidrsubnet example

```
ipcalc 10.1.2.0/24 4 15

Address:    10.1.2.240            00001010.00000001.00000010.11110000
Netmask:    255.255.255.240 = 28  11111111.11111111.11111111.11110000
Wildcard:   0.0.0.15              00000000.00000000.00000000.00001111
=>
Network:    10.1.2.240/28         00001010.00000001.00000010.11110000
Broadcast:  10.1.2.255            00001010.00000001.00000010.11111111
HostMin:    10.1.2.241            00001010.00000001.00000010.11110001
HostMax:    10.1.2.254            00001010.00000001.00000010.11111110
Hosts/Net:  14
TotalHosts: 16
```
