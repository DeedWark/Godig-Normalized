package main

import (
	"flag"
	"fmt"
	"os"

	"tuto.go/src/dns"
)

func main() {

	name := os.Args[0]
	help := "GODIG - Domain DNS Resolver in Golang" + "\r\n" +
		"        Usage:   " + name + " [domain] [selector]" + "\r\n\r\n" +
		"        Example: " + name + " domain.com" + "\r\n" +
		"                 " + name + " domain.com mailjet" + "\r\n\r\n" +
		"Use [" + name + " help] to show this message"

	flag.Parse()
	domain := flag.Arg(0)
	if domain == "" || domain == "help" || domain == "-h" || domain == "--help" {
		fmt.Println(help)
		os.Exit(1)
	}

	selector := flag.Arg(1)
	if selector == "" {
		selector = "google"
	}

	dns.Afinder(domain)
	dns.Mxfinder(domain)
	dns.Txtfinder(domain)
	dns.Dmarcfinder(domain)
	dns.Dkimfinder(domain, selector)
}
