package main

import (
	"fmt"
	"net"
	"strings"
)

func fetchShodan(domain string) ([]string, error) {
	host, _ := net.LookupHost(domain)
	h := fmt.Sprintf("%s", host)
	c := strings.Trim(h, "[]")
	fetchURL := fmt.Sprintf("https://api.shodan.io/shodan/host/%s?key=jH6y3CEE4Ddp7s2N8RtRDGucC4mOvlWe", c)

	wrapper := struct {
		Subdomains []string `json:"hostnames"`
	}{}
	err := fetchJSON(fetchURL, &wrapper)
	return wrapper.Subdomains, err
}
