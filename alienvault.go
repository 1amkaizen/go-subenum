package main

import (
	"fmt"
)

func fetchAlienVault(domain string) ([]string, error) {

	fetchURL := fmt.Sprintf("https://otx.alienvault.com/api/v1/indicators/domain/%s/passive_dns", domain)

	wrapper := struct {
		Subdomains []string `json:"hostname"`
	}{}
	err := fetchJSON(fetchURL, &wrapper)
	return wrapper.Subdomains, err
}
