package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {
	var d bool
	flag.BoolVar(&d, "d", false, "hanya mencari subdomain dari domin")
	flag.Parse()

	domain := flag.Arg(0)
	if domain == "" {
		fmt.Println("Tidak ada domain\n   Usage : -d domain.com")
		return
	}
	domain = strings.ToLower(domain)

	sources := []fetchFn{
		fetchCertspotter,
		fetchVirusTotal,
		fetchCrtSh,
		fetchHackerTarget,
		fetchShodan,
		fetchFacebook,
		fetchAlienVault,
		fetchSecurityTails,
	}
	out := make(chan string)
	var wg sync.WaitGroup

	for _, source := range sources {
		wg.Add(1)
		fn := source
		go func() {

			defer wg.Done()

			names, err := fn(domain)
			if err != nil {
				fmt.Fprintf(os.Stderr, "err: %s\n", err)
				return
			}
			for _, n := range names {
				out <- n
			}

		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	printed := make(map[string]bool)
	for n := range out {
		n = cleanDomain(n)
		if _, ok := printed[n]; ok {
			continue
		}

		if d && !strings.HasSuffix(n, domain) {
			continue
		}
		fmt.Println(n)
		printed[n] = true
		//

		//file := fmt.Sprintf("/home/kali/recon/%s/domains", domain)
		//		f, err := os.Create(file)
		//		if err != nil {
		//			log.Fatal(err)
		//		}
		//		defer f.Close()
		//		_, err2 := f.WriteString(n + "\n")
		//		if err2 != nil {
		//			log.Fatal(err2)
		//		}
		//

	}
}

type fetchFn func(string) ([]string, error)

func httpGet(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}

	raw, err := ioutil.ReadAll(res.Body)

	res.Body.Close()
	if err != nil {
		return []byte{}, err
	}
	return raw, nil

}

func cleanDomain(CD string) string {
	CD = strings.ToLower(CD)

	if len(CD) < 2 {
		return CD
	}
	if CD[0] == '*' || CD[0] == '%' {
		CD = CD[1:]
	}
	if CD[0] == '.' {
		CD = CD[1:]
	}
	return CD
}

func fetchJSON(url string, wrapper interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)

	return dec.Decode(wrapper)
}
