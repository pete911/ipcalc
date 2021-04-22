package main

import (
	"fmt"
	"github.com/pete911/ipcalc/internal"
	"os"
	"strconv"
)

var Version = "dev"

func main() {

	args := os.Args[1:]
	cidr, err := GetCIDR(args)
	if err != nil {
		fmt.Printf("version %s\n", Version)
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	cidr.PrettyPrint()
}

func GetCIDR(args []string) (internal.CIDR, error) {

	var (
		cidr internal.CIDR
		err  error
	)

	switch len(args) {
	case 0:
		err = fmt.Errorf("provide <CIDR> (e.g. 192.168.0.1/24), <CIDR> <hostnum> or <CIDR> <newbits> <netnum> as argument")
	case 1:
		cidr, err = internal.IPCalc(args[0])
	case 2:
		cidr, err = internal.CIDRHost(args[0], ToInt(args[1]))
	case 3:
		cidr, err = internal.CIDRSubnet(args[0], ToInt(args[1]), ToInt(args[2]))
	default:
		err = fmt.Errorf("received %d arguments, provide 1 (CIDR), 2 (CIDR, hostnum) or 3 (CIDR, newbits, netmum)", len(args))
	}
	return cidr, err
}

func ToInt(s string) int {

	out, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("cannot convert %s to int\n", s)
		os.Exit(1)
	}
	return out
}
