package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"sort"
	"strings"
	"time"
	"travian-bot/common"
)

const xCenter = 45
const yCenter = 17

const yRadius = 15
const xRadius = 15

//const yRadius = 1
//const xRadius = 1

type target struct {
	x        int
	y        int
	distance float64
}

func main() {
	common.Login()

	xMin := xCenter - xRadius
	xMax := xCenter + xRadius

	yMin := yCenter - yRadius
	yMax := yCenter + yRadius
	result := make([]target, 0)
	for x := xMin; x <= xMax; x++ {
		for y := yMin; y <= yMax; y++ {
			isNeedToAdd := isOasis(x, y)
			if isNeedToAdd {
				xDiff := math.Pow(float64(xCenter-x), 2)
				yDiff := math.Pow(float64(yCenter-y), 2)
				distance := math.Sqrt(xDiff + yDiff)
				result = append(result, target{x, y, distance})
				fmt.Printf("%s: Finded oasis, x = %d y = %d, d = %f \n", time.Now().Format(time.TimeOnly), x, y, distance)
			}
			time.Sleep((1000 + time.Duration(rand.Intn(2000))) * time.Millisecond)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].distance < result[j].distance
	})

	for _, item := range result {
		fmt.Printf("x: %d, y: %d, // distance: %f \n", item.x, item.y, item.distance)
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
