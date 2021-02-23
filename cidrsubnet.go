package main

import (
	"fmt"
	"net"
)

func CIDRSubnet(ip net.IP, ipnet *net.IPNet, newbits, netnum int) error {

	ipv4 := Byte(ip.To4())
	maskOnes, _ := ipnet.Mask.Size()

	networkPrefix := maskOnes + newbits
	shift := 32 - networkPrefix
	if shift < 0 {
		return fmt.Errorf("newbits %d plus network prefix /%d cannot be more than 32", newbits, maskOnes)
	}

	mask := NewByteFromInt(uint32(netnum << shift))
	subnet := NewByteFromInt(ipv4.ToInt() | mask.ToInt())

	cidr := fmt.Sprintf("%s/%d", subnet.String(), networkPrefix)
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return fmt.Errorf("cannot parse CIDR %s: %w", cidr, err)
	}

	IPCalc(ip, ipnet)
	return nil
}
