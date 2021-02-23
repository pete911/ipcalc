# ipcalc

Go implementation of ipcalc command.

## build

`go build` or `go install`

## run

ipcalc takes either:
 - one argument: `<CIDR>` ([ipcalc](http://jodies.de/ipcalc))
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
