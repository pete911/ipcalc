package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("provide <CIDR> (e.g. 192.168.0.1/24), or <CIDR> <newbits> <netnum> as argument")
		os.Exit(1)
	}

	ip, ipnet, err := net.ParseCIDR(args[0])
	if err != nil {
		fmt.Printf("%s is invalid CIDR block\n", args[0])
		os.Exit(1)
	}

	// ipcalc
	if len(args) == 1 {
		IPCalc(ip, ipnet)
		return
	}

	// cidrsubnet
	if len(args) == 3 {
		if err := CIDRSubnet(ip, ipnet, toInt(args[1]), toInt(args[2])); err != nil {
			fmt.Printf("cidr subnet: %v\n", err)
			os.Exit(1)
		}
		return
	}

	fmt.Printf("received %d arguments, provide 1 (CIDR) or 3 (CIDR, newbits, netmum)\n", len(args))
	os.Exit(1)
}

func toInt(s string) int {

	out, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("cannot convert %s to int\n", s)
		os.Exit(1)
	}
	return out
}
