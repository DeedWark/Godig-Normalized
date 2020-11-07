package dns

import (
	"fmt"
	"net"
)

func Dmarcfinder(domain string) {
	dmarc, _ := net.LookupTXT("_dmarc." + domain)
	fmt.Println("")
	fmt.Println(bold + "DMARC key:" + end)
	if len(dmarc) == 0 {
		fmt.Println("No DMARC key found")
	} else {
		for _, dmkey := range dmarc {
			fmt.Println(dmkey)
		}
	}
}
