package internal

import (
	"fmt"
	"net"
)

func CIDRHost(cidr string, hostnum int) (CIDR, error) {

	ip, _, err := net.ParseCIDR(cidr)
	if err != nil {
		return CIDR{}, fmt.Errorf("%s is invalid CIDR block\n", cidr)
	}

	ipv4 := Byte(ip.To4())
	host := NewByteFromInt(ipv4.ToInt() + uint32(hostnum))

	hostCIDR := fmt.Sprintf("%s/32", host.String())
	return IPCalc(hostCIDR)
}
