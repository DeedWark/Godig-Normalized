package dns

import (
	"flag"
	"fmt"
	"net"
)

const (
	blue   = "\033[34m"
	red    = "\033[91m"
	yellow = "\033[93m"
	green  = "\033[32m"
)

func Dkimfinder(domain string, selector string) {
	dkim, _ := net.LookupTXT(selector + "._domainkey." + domain)
	fmt.Println("")
	fmt.Println(bold + "DKIM key:" + end)
	if flag.Arg(1) == "" {
		fmt.Println("Add a selector (ex: domain.com selector)")
		fmt.Println("Try with " + blue + "G" + red + "o" + yellow + "o" + blue + "g" + green + "l" + red + "e" + end + " as selector:" + "\n")
	}
	if len(dkim) == 0 {
		fmt.Println("No DKIM key found" + "\n")
	} else {
		for _, dkimk := range dkim {
			fmt.Println(dkimk)
		}
	}
}
