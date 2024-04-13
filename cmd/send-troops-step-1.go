package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func SendTroopsStep1(x int, y int, troopType int, troopCount int, logs string) (*string, bool, error) {
	params := url.Values{}
	params.Add(fmt.Sprintf("troop[t%d]", troopType), fmt.Sprintf("%d", troopCount))
	params.Add("troop[t11]", ``)
	params.Add("villagename", ``)
	params.Add("x", fmt.Sprintf("%d", x))
	params.Add("y", fmt.Sprintf("%d", y))
	params.Add("eventType", `4`)
	params.Add("ok", `ok`)
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", "https://ts3.x1.international.travian.com/build.php?gid=16&tt=2", body)
	if err != nil {
		return nil, false, err
	}
	req.Host = "ts3.x1.international.travian.com"
	req.Header.Set("Cookie", cookie)
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Sec-Ch-Ua", "\"Google Chrome\";v=\"123\", \"Not:A-Brand\";v=\"8\", \"Chromium\";v=\"123\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Windows\"")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Origin", "https://ts3.x1.international.travian.com")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	//req.Header.Set("Referer", "https://ts3.x1.international.travian.com/build.php?gid=16&tt=2&eventType=4&targetMapId=58762")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, false, err
	}
	defer resp.Body.Close()

	TryToUpdateCookieAfterRequest(resp)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, false, err
	}

	res := string(respBody)

	if strings.Contains(res, "No troops have been selected") {
		return nil, true, err
	}

	if strings.Contains(res, "value=\"1\"") && logs != "Caesaris" {
		return nil, true, err
	}

	return &res, false, nil
}
