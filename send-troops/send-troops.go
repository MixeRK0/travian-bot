package send_troops

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	"travian-bot/common"
)

func SendTroops(targets []common.Target, logs string) {
	checkMap := make(map[int]map[int]bool)
	for _, target := range targets {
		if _, ok := checkMap[target.X][target.Y]; ok {
			fmt.Printf("%s: %s, Find duplicate in oasis and skipped, x = %d y = %d\n", time.Now().Format(time.TimeOnly), logs, target.X, target.Y)
			continue
		}

		if _, ok := checkMap[target.X]; !ok {
			checkMap[target.X] = make(map[int]bool)
		}

		checkMap[target.X][target.Y] = true
	}

	fmt.Printf("%s: %s, Targets count = %d\n", time.Now().Format(time.TimeOnly), logs, len(targets))

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(targets), func(i, j int) { targets[i], targets[j] = targets[j], targets[i] })

	for {
		for i := 0; i < len(targets); {
			target := targets[i]
			time.Sleep((3 + time.Duration(rand.Intn(5))) * time.Second)
			x, y, troopsType, troopsCount := target.X, target.Y, target.TroopsType, target.Count

			if IsEnemyInOasis(x, y) {
				fmt.Printf("%s: %s, Find enemy in oasis and skipped, x = %d y = %d\n", time.Now().Format(time.TimeOnly), logs, x, y)
				i++
				continue
			}
			time.Sleep((1000 + time.Duration(rand.Intn(2000))) * time.Millisecond)

			htmlStep1, isNoTroops, err := SendTroopsStep1(x, y, troopsType, troopsCount, logs)
			if err != nil {
				log.Fatal("step1 gg", err)
			}

			if isNoTroops {
				sleepSeconds := 300 + time.Duration(rand.Intn(900))
				fmt.Printf("%s: %s, No troops for atack, sleep %d seconds, x = %d y = %d\n", time.Now().Format(time.TimeOnly), logs, sleepSeconds, x, y)
				time.Sleep(sleepSeconds * time.Second)
				continue
			}

			checksum, timestamp := ParseStep1Result(*htmlStep1)

			time.Sleep((1000 + time.Duration(rand.Intn(2000))) * time.Millisecond)
			err = SendTroopsStep2(x, y, troopsType, troopsCount, common.VillageId, checksum, timestamp)
			if err != nil {
				log.Fatal("step2 gg", err)
			}

			fmt.Printf("%s: %s, Sended %d trops to x = %d y = %d\n", time.Now().Format(time.TimeOnly), logs, troopsCount, x, y)

			i++
		}
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(targets), func(i, j int) { targets[i], targets[j] = targets[j], targets[i] })
	}
}
