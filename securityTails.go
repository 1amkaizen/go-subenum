package main

import (
	"fmt"
)

func fetchSecurityTails(domain string) ([]string, error) {

	fetchURL := fmt.Sprintf("https://api.securitytrails.com/v1/domain/%s/subdomains?apikey=FhrfCfk2AK01dVZYEx83Q2rADUfiUdTV", domain)

	wrapper := struct {
		Subdomains []string `json:"subdomains"`
	}{}
	err := fetchJSON(fetchURL, &wrapper)
	return wrapper.Subdomains, err

}
