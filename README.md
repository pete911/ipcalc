# ipcalc
Go implementation of ipcalc command.

## build
`go build` or `go install`

## run
ipcalc takes one argument which is CIDR block e.g.:

```
ipcalc 192.168.0.2/30

Address:   192.168.0.2           11000000.10101000.00000000.00000010
Netmask:   255.255.255.252 = 30  11111111.11111111.11111111.11111100
Wildcard:  0.0.0.3               00000000.00000000.00000000.00000011
=>
Network:   192.168.0.0/30        11000000.10101000.00000000.00000000
Broadcast: 192.168.0.3           11000000.10101000.00000000.00000011
HostMin:   192.168.0.1           11000000.10101000.00000000.00000001
HostMax:   192.168.0.2           11000000.10101000.00000000.00000010
Hosts/Net: 2
```
