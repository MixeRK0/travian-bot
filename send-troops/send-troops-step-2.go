package send_troops

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"travian-bot/common"
)

func SendTroopsStep2(x int, y int, troopType int, troopCount int, villageId int, checksum string, timestamp string) error {
	params := url.Values{}
	params.Add("action", fmt.Sprintf("troopsSend/%d/%s", villageId, timestamp))
	params.Add("eventType", `4`)
	params.Add("villagename", ``)
	params.Add("x", fmt.Sprintf("%d", x))
	params.Add("y", fmt.Sprintf("%d", y))
	params.Add("redeployHero", ``)
	params.Add("checksum", checksum)
	params.Add("troops[0][t1]", `0`)
	params.Add("troops[0][t2]", `0`)
	params.Add("troops[0][t3]", `0`)
	params.Add("troops[0][t4]", `0`)
	params.Add("troops[0][t5]", `0`)
	params.Add("troops[0][t6]", `0`)
	params.Add("troops[0][t7]", `0`)
	params.Add("troops[0][t8]", `0`)
	params.Add("troops[0][t9]", `0`)
	params.Add("troops[0][t10]", `0`)
	params.Add("troops[0][t11]", `0`)
	params.Add(fmt.Sprintf("troops[0][t%d]", troopType), fmt.Sprintf("%d", troopCount))
	params.Add("troops[0][scoutTarget]", ``)
	params.Add("troops[0][catapultTarget1]", ``)
	params.Add("troops[0][catapultTarget2]", ``)
	params.Add("troops[0][villageId]", fmt.Sprintf("%d", villageId))
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", common.Host+"/build.php?gid=16&tt=2", body)
	if err != nil {
		return err
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
	req.Header.Set("Referer", common.Host+"/build.php?gid=16&tt=2")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	common.TryToUpdateCookieAfterRequest(resp)

	return nil
}
