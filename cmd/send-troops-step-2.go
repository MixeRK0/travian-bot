package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func SendTroopsStep2(x int, y int, count int, villageId int, checksum string, timestamp string) error {
	params := url.Values{}
	params.Add("action", fmt.Sprintf("troopsSend/%d/%s", villageId, timestamp))
	params.Add("eventType", `4`)
	params.Add("villagename", ``)
	params.Add("x", fmt.Sprintf("%d", x))
	params.Add("y", fmt.Sprintf("%d", y))
	params.Add("redeployHero", ``)
	params.Add("checksum", checksum)
	params.Add("troops[0][t1]", fmt.Sprintf("%d", count))
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
	params.Add("troops[0][scoutTarget]", ``)
	params.Add("troops[0][catapultTarget1]", ``)
	params.Add("troops[0][catapultTarget2]", ``)
	params.Add("troops[0][villageId]", fmt.Sprintf("%d", villageId))
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", "https://ts3.x1.international.travian.com/build.php?gid=16&tt=2", body)
	if err != nil {
		return err
	}
	req.Host = "ts3.x1.international.travian.com"
	req.Header.Set("Cookie", "__cmpconsentx17155=CP8ltAgP8ltAgAfSDBRUAvEgAP_AAEPAAAYgg1NX_H__bX9v-X736ft0eY1f99j77uQxBhfJs-4FzLvW_JwX32EzNE36tqYKmRIEu3bBIQNtHJnUTVihaogVrzHsakWchTNKJ-BkiHMRe2dYCF5vm4tj-QKZ5_p_93f52T_9_dv-3dzzz91nv3f9_-f1eLida59tH_v_bROb-_If9_7-_4v0_t_rk2_eT1v_9evv7_--_t______9____7___AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAEQamr_j__tr-3_L979P26PMav--x993IYgwvk2fcC5l3rfk4L77CZmib9W1MFTIkCXbtgkIG2jkzqJqxQtUQK15j2NSLOQpmlE_AyRDmIvbOsBC83zcWx_IFM8_0_-7v87J_-_u3_bu555-6z37v-__P6vFxOtc-2j_3_tonN_fkP-_9_f8X6f2_1ybfvJ63_-vX39__ff2______-____9___gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAACAA; __cmpcccx17155=aBP8oqGSgAgAzBXgACAAcABgAHgAUABgADgAJwAXABgADUAHQAegBAAEUAJAAlABQAC4AGIANAAeQBAAEEAJoAXgA9gCHAEyAMQAZYA3ACCgELAIkARoAjoBOACeAFPAKuAWYAyoBoQDmAIxAR3Ao0CjgFTgN0AbsA30BwQDiQHlgPRAgyBBwCFgENgIhARJAiYBFACUoEswJgATLAruBYECzIFogLggXDAumBjsDH4GRgM1AZ4A68CHQERAJGASXgl0BMECY4E3gJvwUKAo0BUACo4FSwA4XVQvihlZDpmlhJRFSsgQ; __qca=P0-872691519-1712341847479; _ga_XMB8PEEDZK=GS1.2.1712692444.5.1.1712692487.0.0.0; _ga_45619DR32F=GS1.2.1712692825.1.1.1712692899.0.0.0; _ga=GA1.1.958460064.1712341848; active_rallypoint_sub_filters_1=2%2C3; active_rallypoint_sub_filters_2=4; _ga_ZQMM0SFYTB=GS1.1.1712707508.3.0.1712707511.0.0.0; JWT=eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJ1WkNWU083VE5IRDVvdHNWVnlxZWVSRmh1U1F6UGswQSIsImF1ZCI6ImQ2MzNkODAwLWY2ODEtMTFlZS02NDAzLTAwMDAwMDAwMDAwMCIsImV4cCI6MTcxMjgwMDY1NCwicHJvcGVydGllcyI6eyJtb2JpbGVPcHRpbWlzYXRpb25zIjpmYWxzZSwibGFuZ3VhZ2UiOiJlbi1VUyIsInB3IjoiQjBDODM4NlhDR2JGMUw4MU02cDdCcENKVFA3dm9EUHkiLCJoYXNoIjoiMDAwMDAwMDBkMjUwZDI1MFI0VHFWSjJUWVpBa0k3dXAiLCJtb2JpbGVPcHRpbWl6YXRpb25zIjpmYWxzZSwibG9naW5JZCI6ODQ3MCwiZGlkIjoxODQ2OSwidmlsbGFnZVBlcnNwZWN0aXZlIjoicGVyc3BlY3RpdmVSZXNvdXJjZXMifX0.fOKMD-NRUk15KVvBJe7OtSHGsbsZf8D8nIGMxW7e4ZmToAJnnsBTJ6uI6XR4FtWIaAy7N3b7EBimPrMHT8M06vnBLxLjM3QvxBX813VY9vU4Lckb9ln9FkX9v2NOFUnwrl04hi3ny357KmUYpsxjDvvaDCzlD8sYLH-Ha77RnNCCVCcxcDnsU5hgRTJEZN7ZyjyJ5vSnWK9Y5Psi14F83tQQP89eln8aXOXZ-HJgqEA4e5g4iQQF5Hg0O4PYExr39tIIGv0xcBSopLWIMvmlBpPYnNzM21Ut8691lzv4WOxtsram1GXeIpj_6oc27TipApoLznRpt2gSzP2NPovblg")
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
	req.Header.Set("Referer", "https://ts3.x1.international.travian.com/build.php?gid=16&tt=2")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Printf("%s: Sended %d trops to x = %d y = %d\n", time.Now().Format(time.TimeOnly), count, x, y)

	return nil
}
