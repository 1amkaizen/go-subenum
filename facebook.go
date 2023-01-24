package main

import "fmt"

func fetchFacebook(domain string) ([]string, error) {
	out := make([]string, 0)

	fetchURL := fmt.Sprintf("https://graph.facebook.com/certificates?fields=domains&access_token=EAAJQBRLEG30BAOeUZBooVdIQWHQBZB40fLlem2f61oZCEurxYg5P9ZBKl3t5l1QvUqMUhalG0w6ZCRYqZBeptMv9ssfBg3jljkr8IOMnZBGZAbVpz7t9IOFF8OuZAMpyCPM2zEbqQm3sar9To5d4yaEMSCTuU4XCPU7QlremvNZCtszIZA8NKj4RUBULUxLpdIK7TqH6ThgZCQqjekRDujpKIZBE2m98E4N4QyRoN4ji6YpyUZAPOT3GDVYZBxs&query=*.%s", domain)
	for {
		wrapper := struct {
			Data []struct {
				Domains []string `json:"domains"`
			} `json:"data"`
			Paging struct {
				Next string `json:"next"`
			} `json:"paging"`
		}{}
		err := fetchJSON(fetchURL, &wrapper)
		if err != nil {
			return out, err
		}

		for _, data := range wrapper.Data {
			for _, d := range data.Domains {
				out = append(out, d)
			}

		}
		fetchURL = wrapper.Paging.Next
		if fetchURL == "" {
			break
		}
	}
	return out, nil
}
