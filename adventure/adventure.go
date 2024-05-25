package adventure

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
	"travian-bot/common"
)

type Adventure struct {
	MapID             int    `json:"mapId"`
	X                 int    `json:"x"`
	Y                 int    `json:"y"`
	Place             string `json:"place"`
	Difficulty        int    `json:"difficulty"`
	TravelingDuration int    `json:"travelingDuration"`
}

func GoToAdventures() {
	for {
		adventures, err := getAdventuresList()
		if err != nil {
			sleepSeconds := 1200 + time.Duration(rand.Intn(300))
			fmt.Printf("%s: Adventures: error: %s, Sleep %d seconds \n", time.Now().Format(time.TimeOnly), err.Error(), sleepSeconds)
			time.Sleep(sleepSeconds * time.Second)
			continue
		}

		time.Sleep(time.Second * 3)

		if len(adventures) == 0 {
			sleepSeconds := 1200 + time.Duration(rand.Intn(300))
			fmt.Printf("%s: Adventures: error: %s, Sleep %d seconds \n", time.Now().Format(time.TimeOnly), err.Error(), sleepSeconds)
			time.Sleep(sleepSeconds * time.Second)
			continue
		}

		err = startAdventure(adventures[0])
		if err != nil {
			sleepSeconds := 1200 + time.Duration(rand.Intn(300))
			fmt.Printf("%s: Adventures: error: %s, Sleep %d seconds \n", time.Now().Format(time.TimeOnly), err.Error(), sleepSeconds)
			time.Sleep(sleepSeconds * time.Second)
			continue
		}
	}
}

func startAdventure(adventure Adventure) error {
	xNonce, err := startAdventureStep1(adventure)
	if err != nil {
		return err
	}

	err = startAdventureStep2(adventure, *xNonce)
	if err != nil {
		return err
	}

	return nil
}

func startAdventureStep1(adventure Adventure) (*string, error) {
	type Troops struct {
		T11 int `json:"t11"`
	}
	type Payload struct {
		Action      string   `json:"action"`
		TargetMapID int      `json:"targetMapId"`
		EventType   int      `json:"eventType"`
		Troops      []Troops `json:"troops"`
	}

	data := Payload{
		Action:      "troopsSend",
		TargetMapID: adventure.MapID,
		EventType:   50,
		Troops: []Troops{
			{T11: 1},
		},
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(payloadBytes)
	bodyStr := string(payloadBytes)
	_ = bodyStr

	req, err := http.NewRequest("PUT", common.Host+"/api/v1/troop/send", body)
	if err != nil {
		return nil, err
	}
	req.Host = common.HostHeader
	req.Header.Set("Cookie", common.Cookie)
	req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"124\", \"Google Chrome\";v=\"124\", \"Not-A.Brand\";v=\"99\"")
	req.Header.Set("X-Version", "2473.3")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	//req.Header.Set("X-Nonce", "44JeDMy0PF2q2S5e7WxQVrmLBSk20mJcVfX7mClMSIdXiqIgXKhzcWvHz09Teiti") // ???????
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Windows\"")
	req.Header.Set("Origin", common.Host)
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", common.Host+"/hero/adventures")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Priority", "u=1, i")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	xNonce := resp.Header.Get("X-Nonce")

	return &xNonce, nil
}

func startAdventureStep2(adventure Adventure, xNonce string) error {
	type Troops struct {
		T11 int `json:"t11"`
	}
	type Payload struct {
		Action      string   `json:"action"`
		TargetMapID int      `json:"targetMapId"`
		EventType   int      `json:"eventType"`
		Troops      []Troops `json:"troops"`
	}

	data := Payload{
		Action:      "troopsSend",
		TargetMapID: adventure.MapID,
		EventType:   50,
		Troops: []Troops{
			{T11: 1},
		},
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	body := bytes.NewReader(payloadBytes)
	bodyStr := string(payloadBytes)
	_ = bodyStr

	req, err := http.NewRequest("POST", common.Host+"/api/v1/troop/send", body)
	if err != nil {
		return err
	}
	req.Host = common.HostHeader
	req.Header.Set("Cookie", common.Cookie)
	req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"124\", \"Google Chrome\";v=\"124\", \"Not-A.Brand\";v=\"99\"")
	req.Header.Set("X-Version", "2473.3")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("X-Nonce", xNonce)
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Windows\"")
	req.Header.Set("Origin", common.Host)
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", common.Host+"/hero/adventures")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Priority", "u=1, i")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func getAdventuresList() ([]Adventure, error) {
	req, err := http.NewRequest("GET", common.Host+"/hero/adventures", nil)
	if err != nil {
		return nil, err
	}
	req.Host = common.HostHeader
	req.Header.Set("Cookie", common.Cookie)
	req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"124\", \"Google Chrome\";v=\"124\", \"Not-A.Brand\";v=\"99\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Windows\"")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	//req.Header.Set("Referer", "https://ts50.x5.arabics.travian.com/dorf1.php")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Priority", "u=0, i")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	respBodyString := string(respBody)

	if !strings.Contains(respBodyString, "\"adventures\":[") {
		return nil, errors.New("no adventures")
	}
	if strings.Contains(respBodyString, "\"inVillage\":null") {
		return nil, errors.New("in way")
	}

	splited1 := strings.Split(respBodyString, `"adventures":[`)
	adventuresString := "[" + strings.Split(splited1[1], `]`)[0] + "]"
	adventuresArray := make([]Adventure, 0)
	err = json.Unmarshal([]byte(adventuresString), &adventuresArray)
	if err != nil {
		return nil, err
	}

	return adventuresArray, nil
}
