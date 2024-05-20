package switch_village

import (
	"fmt"
	"net/http"
	"travian-bot/common"
)

var curVillage = 0

func Switch(villageId int) error {
	if curVillage == villageId {
		return nil
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/dorf1.php?newdid=%d&", common.Host, villageId), nil)
	if err != nil {
		return err
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
	req.Header.Set("Referer", common.Host+"/dorf1.php")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Priority", "u=0, i")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	common.TryToUpdateCookieAfterRequest(resp)

	curVillage = villageId

	return nil
}
