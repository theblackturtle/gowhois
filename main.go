package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strings"

	whoisparser "github.com/likexian/whois-parser-go"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}

		var domain string
		domain = line
		if strings.HasPrefix(line, "http") {
			u, err := url.Parse(line)
			if err != nil {
				domain = u.Hostname()
			}
		}

		out, err := exec.Command("whois", domain).Output()
		if err != nil {
			printError("Whois error: " + err.Error())
			continue
		}
		result, err := whoisparser.Parse(string(out))
		if err != nil {
			printError("Parse error: " + err.Error())
			continue
		}
		j, err := json.Marshal(result)
		if err != nil {
			printError("Marshal error: " + err.Error())
			continue
		}
		fmt.Println(string(j))
	}
}

func printError(s string) {
	fmt.Fprintf(os.Stderr, s+"\n")
}
