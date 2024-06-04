package build

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
	"travian-bot/common"
	switch_village "travian-bot/switch-village"
)

func Build(villageId int, targets []common.BuildingId, isRepeat bool) {
	for {
		i := 0
		for i < len(targets) {
			err := switch_village.Switch(villageId)
			if err != nil {
				println(err.Error())
				i++
				continue
			}

			err = build(targets[i].Id, targets[i].Gid, targets[i].Dorf)
			if err != nil {
				sleepSeconds := 300 + time.Duration(rand.Intn(100))
				fmt.Printf("%s: Build: villageId=%d, id = %d, gid = %d, Sleep %d seconds \n", time.Now().Format(time.TimeOnly), villageId, targets[i].Id, targets[i].Gid, sleepSeconds)
				time.Sleep(sleepSeconds * time.Second)
				continue
			} else {
				fmt.Printf("%s: Build: villageId=%d, id = %d, gid = %d, Launched building \n", time.Now().Format(time.TimeOnly), villageId, targets[i].Id, targets[i].Gid)
			}

			i++
		}

		if isRepeat == false {
			return
		}

		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(targets), func(i, j int) { targets[i], targets[j] = targets[j], targets[i] })
	}
}

func build(id int, gid int, dorf int) error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/build.php?id=%d&gid=%d", common.Host, id, gid), nil)
	if err != nil {
		return err
	}
	req.Host = common.HostHeader
	req.Header.Set("Cookie", common.Cookie)
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"124\", \"Google Chrome\";v=\"124\", \"Not-A.Brand\";v=\"99\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Windows\"")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Priority", "u=0, i")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	common.TryToUpdateCookieAfterRequest(resp)

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	bodyStr := string(respBodyBytes)

	if !strings.Contains(bodyStr, "checksum=") {
		return errors.New("architect")
	}

	splited1 := strings.Split(bodyStr, `checksum=`)
	checksum := strings.Split(splited1[1], `'`)[0]

	if strings.Contains(checksum, "&amp") {
		return errors.New("architect")
	}

	req, err = http.NewRequest("GET", fmt.Sprintf("%s/dorf%d.php?id=%d&gid=%d&action=build&checksum=%s", common.Host, dorf, id, gid, checksum), nil)
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
	req.Header.Set("Referer", fmt.Sprintf("%s/build.php?id=%d&gid=%d", common.Host, id, gid))
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Priority", "u=0, i")

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	common.TryToUpdateCookieAfterRequest(resp)

	respBodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}
