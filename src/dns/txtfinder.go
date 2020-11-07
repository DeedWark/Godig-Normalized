package dns

import (
	"fmt"
	"net"
)

func Txtfinder(domain string) {
	txts, _ := net.LookupTXT(domain)
	fmt.Println("")
	fmt.Println(bold + "TXT records:" + end)
	if len(txts) == 0 {
		fmt.Println("No TXT found")
	} else {
		for _, txt := range txts {
			fmt.Println(txt)
		}
	}
}
