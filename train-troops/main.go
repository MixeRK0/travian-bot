package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var commonCookie = "__cmpcc=1; __cmpconsentx17155=CP9AEggP9AEggAfSDBRUAwEgAAAAAEPAAAYgAABBQgJgA4AM-AjwBKoDfAHbAO5AgoBIgCSgEowJaATHAmSBNICfYFBAKDhBQAAA; __cmpcccx17155=aBP9ASYwAAgAzA_gACAAcABgAHgAUABgADgAJwAXABgAD0AIQAiABQADEAGgAQQAmgBeAD2AIcATIAxABlgEFAIWARIAjoBOACeAFPAKuAWYA0IBzAEYgI7gUaBRwCpwG6AN2Ab6BBkCFgENgIkgSlAlmBMACZYFdwLAgWZAuCBcMDHYGPwMjAZ4A68CIgEl4JdATBAm_BRoCoAFRwAoXVQvihlZDpmrIEA;"
var jwtCookie = ""
var cookie = ""

const hostHeader = "ts3.x1.international.travian.com"
const host = "https://ts3.x1.international.travian.com"

const username = "777McTRAXER777"
const password = "qwe123"
const villageId = "18469"

func main() {
	Login(username, password)

	for {
		sleepSeconds := 300 + time.Duration(rand.Intn(1500))
		if isEnoughResources(550, 440, 320, 100) {
			Train()
			fmt.Printf("%s: Launched train troop, sleep %d seconds\n", time.Now().Format(time.TimeOnly), sleepSeconds)
		} else {
			fmt.Printf("%s: Not enough resourses for train troop, sleep %d seconds\n", time.Now().Format(time.TimeOnly), sleepSeconds)
		}
		time.Sleep(sleepSeconds * time.Second)
	}
}

func Train() {
	req, err := http.NewRequest("GET", host+"/build.php?id=29&gid=20", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Host = hostHeader
	req.Header.Set("Cookie", cookie)
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Ch-Ua", "\"Google Chrome\";v=\"123\", \"Not:A-Brand\";v=\"8\", \"Chromium\";v=\"123\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Windows\"")
	req.Header.Set("Referer", host+"/build.php?id=22&gid=20")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	TryToUpdateCookieAfterRequest(resp)

	res := string(respBody)

	splited := strings.Split(res, "<input type=\"hidden\" name=\"checksum\" value=\"")
	if len(splited) < 2 {
		fmt.Print("Can't split checksum")
		return
	}

	splited1 := strings.Split(splited[1], `"`)
	if len(splited) < 2 {
		fmt.Print("Can't split checksum 2")
		return
	}

	checksum := splited1[0]

	params := url.Values{}
	params.Add("action", `trainTroops`)
	params.Add("checksum", checksum)
	params.Add("s", `1`)
	params.Add("did", villageId)
	params.Add("t5", `1`)
	params.Add("s1", `ok`)
	body := strings.NewReader(params.Encode())

	req, err = http.NewRequest("POST", host+"/build.php?id=29&gid=20", body)
	if err != nil {
		log.Fatal(err)
	}
	req.Host = hostHeader
	req.Header.Set("Cookie", cookie)
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Sec-Ch-Ua", "\"Google Chrome\";v=\"123\", \"Not:A-Brand\";v=\"8\", \"Chromium\";v=\"123\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Windows\"")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Origin", host)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer", host+"?id=29&gid=20")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	TryToUpdateCookieAfterRequest(resp)
}

func isEnoughResources(needWood int, needClay int, needIron int, needCrop int) bool {
	req, err := http.NewRequest("GET", host+"/dorf1.php", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Host = hostHeader
	req.Header.Set("Cookie", cookie)
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
	//req.Header.Set("Referer", "https://ts3.x1.international.travian.com/build.php?id=29&gid=20")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	TryToUpdateCookieAfterRequest(resp)

	res := string(respBody)
	splited := strings.Split(res, "<div id=\"l1\" class=\"value\">&#x202d;")
	if len(splited) < 2 {
		fmt.Print("Can't split wood")
		return false
	}
	woodStr := strings.Replace(strings.Split(splited[1], "&#x202c;</div>")[0], ",", "", -1)
	wood, err := strconv.ParseInt(woodStr, 10, 0)
	if err != nil {
		log.Fatal(err)
	}

	splited = strings.Split(res, "<div id=\"l2\" class=\"value\">&#x202d;")
	if len(splited) < 2 {
		fmt.Print("Can't split wood")
		return false
	}
	clayStr := strings.Replace(strings.Split(splited[1], "&#x202c;</div>")[0], ",", "", -1)
	clay, err := strconv.ParseInt(clayStr, 10, 0)
	if err != nil {
		log.Fatal(err)
	}

	splited = strings.Split(res, "<div id=\"l3\" class=\"value\">&#x202d;")
	if len(splited) < 2 {
		fmt.Print("Can't split wood")
		return false
	}
	ironStr := strings.Replace(strings.Split(splited[1], "&#x202c;</div>")[0], ",", "", -1)
	iron, err := strconv.ParseInt(ironStr, 10, 0)
	if err != nil {
		log.Fatal(err)
	}

	splited = strings.Split(res, "<div id=\"l3\" class=\"value\">&#x202d;")
	if len(splited) < 2 {
		fmt.Print("Can't split wood")
		return false
	}
	cropStr := strings.Replace(strings.Split(splited[1], "&#x202c;</div>")[0], ",", "", -1)
	crop, err := strconv.ParseInt(cropStr, 10, 0)
	if err != nil {
		log.Fatal(err)
	}

	if needWood > int(wood) {
		return false
	}

	if needClay > int(clay) {
		return false
	}

	if needIron > int(iron) {
		return false
	}

	if needCrop > int(crop) {
		return false
	}

	return true
}
