package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("provide CIDR block (e.g. 192.168.0.1/24) as argument")
		os.Exit(1)
	}
	ip, ipnet, err := net.ParseCIDR(args[0])
	if err != nil {
		fmt.Printf("%s is invalid CIDR block\n", args[0])
		os.Exit(1)
	}
	ipcalc(ip, ipnet)
}

func ipcalc(ip net.IP, ipnet *net.IPNet) {

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
	if broadcast.ToInt() == hostMin.ToInt() {
		hostMax = hostMin
	}

	fmt.Printf("Address:   %-21s %s\n", ipv4.String(), ipv4.BinaryString())
	fmt.Printf("Netmask:   %-21s %s\n", netmaskStr, netmask.BinaryString())
	fmt.Printf("Wildcard:  %-21s %s\n", wildcard.String(), wildcard.BinaryString())
	fmt.Println("=>")
	fmt.Printf("Network:   %-21s %s\n", networkCIDR, network.BinaryString())
	fmt.Printf("Broadcast: %-21s %s\n", broadcast.String(), broadcast.BinaryString())
	fmt.Printf("HostMin:   %-21s %s\n", hostMin.String(), hostMin.BinaryString())
	fmt.Printf("HostMax:   %-21s %s\n", hostMax.String(), hostMax.BinaryString())
	fmt.Printf("Hosts/Net: %d\n", hostMax.ToInt() - hostMin.ToInt())
}

type Byte []byte

func NewByteFromInt(i uint32) Byte {

	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, i)
	return b
}

func (b Byte) ToInt() uint32 {
	return binary.BigEndian.Uint32(b)
}

func (b Byte) BinaryString() string {

	var out []string
	for _, n := range b {
		out = append(out, fmt.Sprintf("%08b", n))
	}
	return strings.Join(out, ".")
}

func (b Byte) String() string {

	var out []string
	for _, n := range b {
		out = append(out, fmt.Sprintf("%d", n))
	}
	return strings.Join(out, ".")
}
