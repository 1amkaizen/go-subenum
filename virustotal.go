package main

import (
	"fmt"
)

func fetchVirusTotal(domain string) ([]string, error) {

	fetchURL := fmt.Sprintf("https://www.virustotal.com/vtapi/v2/domain/report?apikey=235d2ac68f0e00d4cfc18b58250c2b55a4729cdcaf240583fe79b2a3beb45da0&domain=%s", domain)

	wrapper := struct {
		Subdomains []string `json:"subdomains"`
	}{}
	err := fetchJSON(fetchURL, &wrapper)
	return wrapper.Subdomains, err
}
