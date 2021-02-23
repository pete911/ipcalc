package main

import (
	"fmt"
	"net"
)

func IPCalc(cidr string) error {

	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return fmt.Errorf("%s is invalid CIDR block\n", cidr)
	}

	ipv4 := Byte(ip.To4())
	netmask := Byte(ipnet.Mask)
	maskOnes, _ := ipnet.Mask.Size()
	netmaskStr := fmt.Sprintf("%s = %d", netmask.String(), maskOnes)
	wildcard := NewByteFromInt(0xFFFFFFFF >> maskOnes)

	network := Byte(ipnet.IP)
	networkCIDR := fmt.Sprintf("%v/%d", network.String(), maskOnes)

	broadcast := NewByteFromInt(ipv4.ToInt() | wildcard.ToInt())
	hostMin := NewByteFromInt(network.ToInt() | 1<<0)      // set last bit on network
	hostMax := NewByteFromInt(broadcast.ToInt() &^ 1 << 0) // clear last bit on broadcast

	totalHosts := (broadcast.ToInt() - network.ToInt()) + 1
	hosts := totalHosts - 2
	if maskOnes > 30 {
		hosts = 0
		hostMin = NAByte()
		hostMax = NAByte()
	}
	if maskOnes == 32 {
		broadcast = NAByte()
	}

	fmt.Printf("Address:    %-21s %s\n", ipv4.String(), ipv4.BinaryString())
	fmt.Printf("Netmask:    %-21s %s\n", netmaskStr, netmask.BinaryString())
	fmt.Printf("Wildcard:   %-21s %s\n", wildcard.String(), wildcard.BinaryString())
	fmt.Println("=>")
	fmt.Printf("Network:    %-21s %s\n", networkCIDR, network.BinaryString())
	fmt.Printf("Broadcast:  %-21s %s\n", broadcast.String(), broadcast.BinaryString())
	fmt.Printf("HostMin:    %-21s %s\n", hostMin.String(), hostMin.BinaryString())
	fmt.Printf("HostMax:    %-21s %s\n", hostMax.String(), hostMax.BinaryString())
	fmt.Printf("Hosts/Net:  %d\n", hosts)
	fmt.Printf("TotalHosts: %d\n", totalHosts)
	return nil
}
