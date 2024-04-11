package main

import (
	"log"
	"math/rand"
	"time"
)

func SendTroops(targets []TargetCoords, count int) {

	//for {
	for _, target := range targets {
		time.Sleep((3 + time.Duration(rand.Intn(2))) * time.Second)
		x, y := target.x, target.y
		htmlStep1, err := SendTroopsStep1(x, y, count)
		if err != nil {
			log.Fatal("step1 gg", err)
		}

		checksum, timestamp := ParseStep1Result(*htmlStep1)

		err = SendTroopsStep2(x, y, count, 18469, checksum, timestamp)
		if err != nil {
			log.Fatal("step2 gg", err)
		}
	}
	//}
}
