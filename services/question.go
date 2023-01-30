package services

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Question1get(client *http.Client) string {
	url := host + "/question/1"
	method := "GET"

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9,ru;q=0.8")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Connection", "keep-alive")
	// req.Header.Add("Cookie", "SID=nar1W5eac9j8edLvyjJs3YmB65fN4XmA")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Referer", host+"/")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", userAgent)

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	rule := ruleParser(string(body))
	if !rule.Validate() {
		log.Fatalln("unexpected rule")
	}

	form := formParser(string(body))
	if !form.Validate() {
		log.Fatalln("unexpected form")
	}

	return fill(rule, form)
}

func QuestionPost(client *http.Client, stepNumber uint, data string) (bool, string) {
	//count := 5

	//wg := sync.WaitGroup{}
	//for i := 0; i < count; i++ {
	//	wg.Add(1)

	//	go func(x int) (bool, string) {
	//		defer wg.Done()
	url := host + fmt.Sprintf("/question/%d", stepNumber)
	method := "POST"

	payload := strings.NewReader(data)

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9,ru;q=0.8")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Add("Cookie", "SID=nar1W5eac9j8edLvyjJs3YmB65fN4XmA")
	req.Header.Add("Origin", host)
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Referer", host+fmt.Sprintf("/question/%d", stepNumber))
	req.Header.Add("Upgrade-Insecure-Requests", fmt.Sprintf("%d", stepNumber))
	req.Header.Add("User-Agent", userAgent)

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if strings.Contains(string(body), "Test successfully passed") {
		return true, ""
	}

	rule := ruleParser(string(body))
	if !rule.Validate() {
		log.Fatalln("unexpected rule")
	}

	form := formParser(string(body))
	if !form.Validate() {
		log.Fatalln("unexpected form")
	}
	return false, fill(rule, form)
	//	}(i)

	//}
}
