package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
	"travian-bot/common"
)

const xMin = 22
const xMax = 25
const yMin = 48
const yMax = 50

func main() {
	common.Login()

	result := make([][]int, 0)
	for x := xMin; x <= xMax; x++ {
		for y := yMin; y <= yMax; y++ {
			isNeedToAdd := isOasis(x, y)
			if isNeedToAdd {
				result = append(result, []int{x, y})
				fmt.Printf("%s: Finded oasis, x = %d y = %d\n", time.Now().Format(time.TimeOnly), x, y)
			}
			time.Sleep((1000 + time.Duration(rand.Intn(2000))) * time.Millisecond)
		}
	}

	for _, item := range result {
		fmt.Println(item)
	}
}

func isOasis(x int, y int) bool {
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

	req, err := http.NewRequest("POST", common.Host+"/api/v1/map/tile-details", body)
	if err != nil {
		log.Fatal(err)
	}
	req.Host = common.HostHeader
	req.Header.Set("Cookie", common.Cookie)
	req.Header.Set("Sec-Ch-Ua", "\"Google Chrome\";v=\"123\", \"Not:A-Brand\";v=\"8\", \"Chromium\";v=\"123\"")
	req.Header.Set("X-Version", "2435.8")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Windows\"")
	req.Header.Set("Origin", common.Host)
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	//req.Header.Set("Referer", "https://ts3.x1.international.travian.com/karte.php?x=15&y=55&zoom=1")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	type RespBody struct {
		Html string `json:"html"`
	}

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var respBody *RespBody
	if err = json.Unmarshal(respBodyBytes, &respBody); err != nil {
		log.Fatal(err)
	}

	respBodyString := respBody.Html
	if strings.Contains(respBodyString, "Unoccupied oasis") || strings.Contains(respBodyString, "Свободный оазис") {
		return true
	}

	return false
}
