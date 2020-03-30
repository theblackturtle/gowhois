package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	whoisparser "github.com/likexian/whois-parser-go"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func main() {
	flag.Parse()
	var domains []string
	if flag.NArg() > 0 {
		domains = append(domains, flag.Args()...)
	} else {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			line := strings.TrimSpace(sc.Text())
			if line == "" {
				continue
			}
			domains = append(domains, line)
		}
	}

	if len(domains) == 0 {
		printError("No domains")
		os.Exit(1)
	}

	for _, domain := range domains {
		whois(domain)
	}
}

func whois(domain string) {
	if strings.HasPrefix(domain, "http") {
		u, err := url.Parse(domain)
		if err != nil {
			printError("Parse URL error: " + err.Error())
			return
		}
		domain = u.Hostname()
	}
	out, err := exec.Command("whois", domain).Output()
	if err != nil {
		printError("Exec command error: " + err.Error())
		return
	}
	result, err := whoisparser.Parse(string(out))
	if err != nil {
		printError("Parse error: " + err.Error())
		return
	}
	j, err := json.Marshal(result)
	if err != nil {
		printError("Marshal error: " + err.Error())
		return
	}
	fmt.Println(string(j))
}

func printError(s string) {
	fmt.Fprintf(os.Stderr, s+"\n")
}
