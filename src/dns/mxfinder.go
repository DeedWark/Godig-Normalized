package dns

import (
	"fmt"
	"net"
	"strings"
)

func Mxfinder(domain string) {
	mxs, _ := net.LookupMX(domain)
	fmt.Println("")
	fmt.Println(bold + "MX fields:" + end)
	if len(mxs) == 0 {
		fmt.Println("No MX found")
	} else {
		for _, mx := range mxs {
			mxRaw := strings.TrimRight(mx.Host, ".")
			fmt.Println(mx.Pref, mxRaw)
		}
	}
}
