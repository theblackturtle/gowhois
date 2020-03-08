package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/domainr/whois"
	whoisparser "github.com/likexian/whois-parser-go"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		request, err := whois.NewRequest(line)
		if err != nil {
			printError("Request error: " + err.Error())
			continue
		}
		response, err := whois.DefaultClient.Fetch(request)
		if err != nil {
			printError("Fetch error: " + err.Error())
			continue
		}
		result, err := whoisparser.Parse(response.String())
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
