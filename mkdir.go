package main

import (
	"fmt"
	"log"
	"os"
)

func Create(domain string) {
	create := fmt.Sprintf("/home/kali/recon/%s", domain)
	if err := os.MkdirAll(create, os.ModePerm); err != nil {
		log.Fatal(err)
	}

}
