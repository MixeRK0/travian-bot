package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var CommonCookie = "__cmpcc=1; __cmpconsentx17155=CP9AEggP9AEggAfSDBRUAwEgAAAAAEPAAAYgAABBQgJgA4AM-AjwBKoDfAHbAO5AgoBIgCSgEowJaATHAmSBNICfYFBAKDhBQAAA; __cmpcccx17155=aBP9ASYwAAgAzA_gACAAcABgAHgAUABgADgAJwAXABgAD0AIQAiABQADEAGgAQQAmgBeAD2AIcATIAxABlgEFAIWARIAjoBOACeAFPAKuAWYA0IBzAEYgI7gUaBRwCpwG6AN2Ab6BBkCFgENgIkgSlAlmBMACZYFdwLAgWZAuCBcMDHYGPwMjAZ4A68CIgEl4JdATBAm_BRoCoAFRwAoXVQvihlZDpmrIEA;"
var JwtCookie = ""
var Cookie = ""

const HostHeader = "ts100.x10.europe.travian.com"
const Host = "https://ts100.x10.europe.travian.com"

const Username = "777sweety777"
const Password = "qwe123"
const VillageId = 18054

func Login() {
	type Payload struct {
		Name                string `json:"name"`
		Password            string `json:"password"`
		W                   string `json:"w"`
		MobileOptimizations bool   `json:"mobileOptimizations"`
	}

	data := Payload{
		Name:                Username,
		Password:            Password,
		MobileOptimizations: false,
		W:                   "1920:1080",
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", Host+"/api/v1/auth/login", body)
	if err != nil {
		log.Fatal(err)
	}
	req.Host = HostHeader
	req.Header.Set("Cookie", "__cmpcc=1; __cmpconsentx17155=CP9AEggP9AEggAfSDBRUAwEgAAAAAEPAAAYgAABBQgJgA4AM-AjwBKoDfAHbAO5AgoBIgCSgEowJaATHAmSBNICfYFBAKDhBQAAA; __cmpcccx17155=aBP9ASB4AAgAzA_gACAAcABgAHgAUABgADgAJwAXABgAD0AIQAiABQADEAGgAQQAmgBeAD2AIcATIAxABlgEFAIWARIAjoBOACeAFPAKuAWYA0IBzAEYgI7gUaBRwCpwG6AN2Ab6BBkCFgENgIkgSlAlmBMACZYFdwLAgWZAuCBcMDHYGPwMjAZ4A68CIgEl4JdATBAm_BRoCoAFRwAoXVQvihlZDpmrIEA")
	req.Header.Set("Sec-Ch-Ua", "\"Google Chrome\";v=\"123\", \"Not:A-Brand\";v=\"8\", \"Chromium\";v=\"123\"")
	req.Header.Set("X-Version", "2435.8")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Windows\"")
	req.Header.Set("Origin", Host)
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", Host)
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	type RespBody struct {
		Code string `json:"code"`
	}

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var respBody *RespBody
	if err = json.Unmarshal(respBodyBytes, &respBody); err != nil {
		log.Fatal(err)
	}

	ResolveCookie(respBody.Code)
}

func ResolveCookie(code string) {
	req, err := http.NewRequest("GET", fmt.Sprintf(Host+"/api/v1/auth?code=%s", code), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Host = HostHeader
	req.Header.Set("Cookie", CommonCookie)
	req.Header.Set("Sec-Ch-Ua", "\"Google Chrome\";v=\"123\", \"Not:A-Brand\";v=\"8\", \"Chromium\";v=\"123\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Windows\"")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer", Host)
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	JwtCookie = resp.Header.Get("Set-Cookie")
	Cookie = CommonCookie + JwtCookie

	println(Cookie)
}

func TryToUpdateCookieAfterRequest(r *http.Response) {
	if r == nil {
		return
	}

	if r.Header.Get("Set-Cookie") != "" {
		JwtCookie = r.Header.Get("Set-Cookie")
		Cookie = CommonCookie + JwtCookie

		println("Cookie updated")
	}
}
