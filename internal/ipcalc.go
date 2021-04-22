package internal

import (
	"fmt"
	"net"
)

type CIDR struct {
	ipv4        Byte
	netmask     Byte
	netmaskStr  string
	wildcard    Byte
	networkCIDR string
	network     Byte
	broadcast   Byte
	hostMin     Byte
	hostMax     Byte
	hosts       uint32
	totalHosts  uint32
}

func (c CIDR) PrettyPrint() {

	fmt.Printf("Address:    %-21s %s\n", c.ipv4.String(), c.ipv4.BinaryString())
	fmt.Printf("Netmask:    %-21s %s\n", c.netmaskStr, c.netmask.BinaryString())
	fmt.Printf("Wildcard:   %-21s %s\n", c.wildcard.String(), c.wildcard.BinaryString())
	fmt.Println("=>")
	fmt.Printf("Network:    %-21s %s\n", c.networkCIDR, c.network.BinaryString())
	fmt.Printf("Broadcast:  %-21s %s\n", c.broadcast.String(), c.broadcast.BinaryString())
	fmt.Printf("HostMin:    %-21s %s\n", c.hostMin.String(), c.hostMin.BinaryString())
	fmt.Printf("HostMax:    %-21s %s\n", c.hostMax.String(), c.hostMax.BinaryString())
	fmt.Printf("Hosts/Net:  %d\n", c.hosts)
	fmt.Printf("TotalHosts: %d\n", c.totalHosts)
}

func IPCalc(cidrNotation string) (CIDR, error) {

	ip, ipnet, err := net.ParseCIDR(cidrNotation)
	if err != nil {
		return CIDR{}, fmt.Errorf("%s is invalid CIDR notation\n", cidrNotation)
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

	return CIDR{
		ipv4:        ipv4,
		netmask:     netmask,
		netmaskStr:  netmaskStr,
		wildcard:    wildcard,
		networkCIDR: networkCIDR,
		network:     network,
		broadcast:   broadcast,
		hostMin:     hostMin,
		hostMax:     hostMax,
		hosts:       hosts,
		totalHosts:  totalHosts,
	}, nil
}
