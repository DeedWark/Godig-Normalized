package dns

import (
	"fmt"
	"net"
)

const (
	bold = "\033[1m"
	end  = "\033[00m"
)

func Afinder(domain string) {
	a, _ := net.LookupIP(domain)
	fmt.Println(bold + "DNS (A):" + end)
	if len(a) == 0 {
		fmt.Println("No DNS found")
	} else {
		for _, ip := range a {
			fmt.Println(ip.String())
		}
	}
}
