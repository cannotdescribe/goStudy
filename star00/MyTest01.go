package main

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

func (ipAddr IPAddr) String() string{
	var result string;
	for _,v := range ipAddr{
		result += fmt.Sprintf("%v.", v);
	}
	index := strings.LastIndex(result, ".");
	return string([]byte(result)[:index])
}

// TODO: Add a "String() string" method to IPAddr.

//func main() {
//	hosts := map[string]IPAddr{
//		"loopback":  {127, 0, 0, 1},
//		"googleDNS": {8, 8, 8, 8},
//	}
//	for name, ip := range hosts {
//		fmt.Printf("%v: %v\n", name, ip)
//	}
//}
