package main

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strconv"
	"time"
)
// your auth token here. must find in browser tools by inspecting requests
// neither id or token really change from what ive seen.
// MUST ENTER TO RUN CODE. THIS ESSENTIALLY LOGS YOU IN TO YOUR ACCOUNT.
var authToken string = ""
// also found in browser requests. must have to run code.
var device_id string = ""

func main() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	Client := &http.Client{
		Jar: jar,
	}
	respCode := 200
	arrIdx := 0

	randomBytesForLabel := make([]byte, 10)
	
	// to go through rand domains
	domainsArrList := []string{
		"ipriva.com",
		"mymaskedmail.com",
		"maskedmails.com",
		"dontrackme.com",
		"blurfamily.com",
	}
	for (respCode == 200) {
	// creates random label for email.
	_, err = rand.Read(randomBytesForLabel)
	if err != nil {
		panic(err)
	}
	label := base64.URLEncoding.EncodeToString(randomBytesForLabel)

	jsonBuffer := bytes.NewBuffer([]byte(`{"source":"webapp","label":"` + label[:10] + `","custom":"","domain":"` + domainsArrList[arrIdx % 5] + `","auth_token":"` + authToken + `","client_application":{"version":"web-8.2.9999","tag":"www_abine_com","product":"Webapp","browser":"Safari","user_agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.3 Safari/605.1.15","device_id":"` + device_id + `"}}`))
	arrIdx++;

	req, err := http.NewRequest("POST", "https://emails.abine.com/api/v4/disposables", jsonBuffer)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.3 Safari/605.1.15")
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}
	resp, err := Client.Do(req)
	if err != nil {
		panic(err)
	}
	
	if (resp.StatusCode != 200) {
		arrIdx--;
		respCode = resp.StatusCode
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(body))
	} else {
		// to avoid a rate limit it provides space a "buffer" between requests
		fmt.Println("sleeping until next request in 60s")
		time.Sleep(60 * time.Second)
	}
	}
	fmt.Println("emails created = " + strconv.Itoa(arrIdx))
}