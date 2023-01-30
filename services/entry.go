package services

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Entry(client *http.Client) {
	url := host + "/"
	method := "GET"

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9,ru;q=0.8")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", userAgent)

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	/*body*/
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
}
