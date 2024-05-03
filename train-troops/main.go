package train_troops

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
	"travian-bot/common"
)

func TrainTroops() {
	for {
		//common.Login()
		sleepMins := 30 + time.Duration(rand.Intn(60))
		isEnough, count := isEnoughResources(550, 440, 320, 100)
		if isEnough {
			Train(count)
			fmt.Printf("%s: Launched train troop, sleep %d minutes\n", time.Now().Format(time.TimeOnly), sleepMins)
		} else {
			fmt.Printf("%s: Not enough resourses for train troop, sleep %d minutes\n", time.Now().Format(time.TimeOnly), sleepMins)
		}
		time.Sleep(sleepMins * time.Minute)
	}
}

func Train(count int) {
	println(count)
	req, err := http.NewRequest("GET", common.Host+"/build.php?gid=20", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Host = common.HostHeader
	req.Header.Set("Cookie", common.Cookie)
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
	req.Header.Set("Referer", common.Host+"/dorf1.php")
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

	common.TryToUpdateCookieAfterRequest(resp)

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
	params.Add("did", fmt.Sprintf("%d", common.VillageId))
	params.Add("t5", fmt.Sprintf("%d", count))
	params.Add("s1", `ok`)
	body := strings.NewReader(params.Encode())

	req, err = http.NewRequest("POST", common.Host+"/build.php?id=29&gid=20", body)
	if err != nil {
		log.Fatal(err)
	}
	req.Host = common.HostHeader
	req.Header.Set("Cookie", common.Cookie)
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Sec-Ch-Ua", "\"Google Chrome\";v=\"123\", \"Not:A-Brand\";v=\"8\", \"Chromium\";v=\"123\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Windows\"")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Origin", common.Host)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer", common.Host+"?gid=20")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	common.TryToUpdateCookieAfterRequest(resp)
}

func isEnoughResources(needWood int, needClay int, needIron int, needCrop int) (bool, int) {
	req, err := http.NewRequest("GET", common.Host+"/dorf1.php", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Host = common.HostHeader
	req.Header.Set("Cookie", common.Cookie)
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

	common.TryToUpdateCookieAfterRequest(resp)

	res := string(respBody)
	splited := strings.Split(res, "<div id=\"l1\" class=\"value\">&#x202d;")
	if len(splited) < 2 {
		fmt.Print("Can't split wood")
		return false, 0
	}
	woodStr := strings.Replace(strings.Split(splited[1], "&#x202c;</div>")[0], " ", "", -1)
	wood, err := strconv.ParseInt(woodStr, 10, 0)
	if err != nil {
		log.Fatal(err)
	}

	splited = strings.Split(res, "<div id=\"l2\" class=\"value\">&#x202d;")
	if len(splited) < 2 {
		fmt.Print("Can't split wood")
		return false, 0
	}
	clayStr := strings.Replace(strings.Split(splited[1], "&#x202c;</div>")[0], " ", "", -1)
	clay, err := strconv.ParseInt(clayStr, 10, 0)
	if err != nil {
		log.Fatal(err)
	}

	splited = strings.Split(res, "<div id=\"l3\" class=\"value\">&#x202d;")
	if len(splited) < 2 {
		fmt.Print("Can't split wood")
		return false, 0
	}
	ironStr := strings.Replace(strings.Split(splited[1], "&#x202c;</div>")[0], " ", "", -1)
	iron, err := strconv.ParseInt(ironStr, 10, 0)
	if err != nil {
		log.Fatal(err)
	}

	splited = strings.Split(res, "<div id=\"l3\" class=\"value\">&#x202d;")
	if len(splited) < 2 {
		fmt.Print("Can't split wood")
		return false, 0
	}
	cropStr := strings.Replace(strings.Split(splited[1], "&#x202c;</div>")[0], " ", "", -1)
	crop, err := strconv.ParseInt(cropStr, 10, 0)
	if err != nil {
		log.Fatal(err)
	}

	if needWood > int(wood) {
		return false, 0
	}

	if needClay > int(clay) {
		return false, 0
	}

	if needIron > int(iron) {
		return false, 0
	}

	if needCrop > int(crop) {
		return false, 0
	}

	woodCount := int(wood) / needWood
	clayCount := int(clay) / needClay
	ironCount := int(iron) / needIron
	cropCount := int(crop) / needCrop

	minCount := woodCount
	if clayCount < minCount {
		minCount = clayCount
	}
	if ironCount < minCount {
		minCount = ironCount
	}
	if cropCount < minCount {
		minCount = cropCount
	}

	return true, minCount
}
