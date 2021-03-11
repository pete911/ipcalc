package main

import (
	"fmt"
	"os"
	"strconv"
)

var Version = "dev"

func main() {

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Printf("version %s\n", Version)
		fmt.Println("provide <CIDR> (e.g. 192.168.0.1/24), <CIDR> <hostnum> or <CIDR> <newbits> <netnum> as argument")
		os.Exit(1)
	}

	// ipcalc
	if len(args) == 1 {
		if err := IPCalc(args[0]); err != nil {
			fmt.Printf("ipcalc: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// cidrhost
	if len(args) == 2 {
		if err := CIDRHost(args[0], ToInt(args[1])); err != nil {
			fmt.Printf("cidrhost: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// cidrsubnet
	if len(args) == 3 {
		if err := CIDRSubnet(args[0], ToInt(args[1]), ToInt(args[2])); err != nil {
			fmt.Printf("cidrsubnet: %v\n", err)
			os.Exit(1)
		}
		return
	}

	fmt.Printf("received %d arguments, provide 1 (CIDR) or 3 (CIDR, newbits, netmum)\n", len(args))
	os.Exit(1)
}

func ToInt(s string) int {

	out, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("cannot convert %s to int\n", s)
		os.Exit(1)
	}
	return out
}
