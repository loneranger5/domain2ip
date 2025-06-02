/*********

Concurrent domain-to-IP resolver
Author: github.com/loneranger5

**********/

package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/sourcegraph/conc"

	"strings"
)

func Resolve(domains []string, out_file string) {

	output_file, err := os.Create(out_file)

	if err != nil {
		panic("Unable to create output.")
	}

	var wg conc.WaitGroup

	for _, domain := range domains {

		wg.Go(func() {
			ips, err := net.LookupIP(domain)
			if err != nil {
				// fmt.Printf("Error resolving %s: %v\n", domain, err)
				return
			}
			for _, ip := range ips {
				if ipv4 := ip.To4(); ipv4 != nil {
					fmt.Printf(" %s: %s\n", domain, ipv4.String())
					fmt.Fprintf(output_file, "%s -> %s\n", domain, ipv4.String())

				}

			}
		})

	}

	wg.Wait()
}

func main() {
	var DomainFileArg string
	var OutFileArg string
	flag.StringVar(&DomainFileArg, "domain", "", "Provide domain to resolve.")
	flag.StringVar(&OutFileArg, "output", "output.txt", "Filename to output results.")

	flag.Parse()

	content, err := os.ReadFile(DomainFileArg)
	if err != nil {
		fmt.Println("Unable to read ", DomainFileArg)
		panic("")
	}

	if len(content) <= 0 {
		panic("File is empty")
	}

	domains := strings.Split(string(content), "\n")
	Resolve(domains, OutFileArg)

}
