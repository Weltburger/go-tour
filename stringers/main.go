package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (i IPAddr) String() string {
	str := make([]string, len(i))
	for i, val := range i {
		str[i] = strconv.Itoa(int(val))
	}
	res := strings.Join(str, ".")

	return res
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
