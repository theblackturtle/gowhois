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
			printError(err)
			continue
		}
		response, err := whois.DefaultClient.Fetch(request)
		if err != nil {
			printError(err)
			continue
		}
		result, err := whoisparser.Parse(response.String())
		if err != nil {
			printError(err)
			continue
		}
		j, err := json.Marshal(result)
		if err != nil {
			printError(err)
			continue
		}
		fmt.Println(string(j))
	}
}

func printError(e error) {
	fmt.Fprintf(os.Stderr, e.Error()+"\n")
}
