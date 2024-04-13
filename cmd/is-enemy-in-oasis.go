package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func IsEnemyInOasis(x int, y int) bool {
	type Payload struct {
		X int `json:"x"`
		Y int `json:"y"`
	}

	data := Payload{
		X: x,
		Y: y,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://ts3.x1.international.travian.com/api/v1/map/tile-details", body)
	if err != nil {
		log.Fatal(err)
	}
	req.Host = "ts3.x1.international.travian.com"
	req.Header.Set("Cookie", "__cmpcc=1; __cmpconsentx17155=CP9AEggP9AEggAfSDBRUAwEgAAAAAEPAAAYgAABBQgJgA4AM-AjwBKoDfAHbAO5AgoBIgCSgEowJaATHAmSBNICfYFBAKDhBQAAA; __cmpcccx17155=aBP9ASYwAAgAzA_gACAAcABgAHgAUABgADgAJwAXABgAD0AIQAiABQADEAGgAQQAmgBeAD2AIcATIAxABlgEFAIWARIAjoBOACeAFPAKuAWYA0IBzAEYgI7gUaBRwCpwG6AN2Ab6BBkCFgENgIkgSlAlmBMACZYFdwLAgWZAuCBcMDHYGPwMjAZ4A68CIgEl4JdATBAm_BRoCoAFRwAoXVQvihlZDpmrIEA; active_rallypoint_sub_filters_2=4; JWT=eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJHTHFPeHh0UG0xbFU2dVRjWmlpVmwxMzJUZVJieHdPMSIsImF1ZCI6ImQ2MzNkODAwLWY2ODEtMTFlZS02NDAzLTAwMDAwMDAwMDAwMCIsImV4cCI6MTcxMjk3MDIwNSwicHJvcGVydGllcyI6eyJwdyI6IjBKMzk2TXBrNGY5SHJ5cjhiSndaUEY5RkhIM2dwb3VzIiwiaGFzaCI6IjAwMDAwMDAwMDAwMGQyNTBpdlA1MU8wYjhiOWd5ME00IiwibW9iaWxlT3B0aW1pemF0aW9ucyI6ZmFsc2UsImxvZ2luSWQiOjI5NzIwLCJkaWQiOjE4NDY5LCJsYW5ndWFnZSI6ImVuLVVTIiwidmlsbGFnZVBlcnNwZWN0aXZlIjoicGVyc3BlY3RpdmVSZXNvdXJjZXMifX0.K08lTlU2AbutW0GPGEyflAKS1YoGELMuBbJMZsGAnNjO5xxb2Scml2cyyjsH_ezr05vJYpqTDQwn9Jj8ssSQE6tQxcu3Hp-toZk1VkwBvA8Vr2Z-JkVzMdc4i2e0JOaNA4HbHTL40ilY2v5oGDGqEWRLxU1pOx3xBQq_A2AC_PO6_GonJWfjlzJJfI5wSiYVNv6JuPfUdsM22lp0w0HxDSHMapBkvxTa_MN0Sf4808_gX4SWesjkbhG7Qy6ABf7-reGPHsGA7WzL4rzvRIh-79YiS7-mwBPmNGKNtujS7m3WyCXAtAYPjmOPLm0S1E9u2U-4I_YLLzirnYE0PTMWlw")
	req.Header.Set("Sec-Ch-Ua", "\"Google Chrome\";v=\"123\", \"Not:A-Brand\";v=\"8\", \"Chromium\";v=\"123\"")
	req.Header.Set("X-Version", "2435.8")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Windows\"")
	req.Header.Set("Origin", "https://ts3.x1.international.travian.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://ts3.x1.international.travian.com/karte.php?x=15&y=55&zoom=1")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	respBodyString := string(respBodyBytes)
	if strings.Contains(respBodyString, "none") {
		return false
	}

	return true
}
