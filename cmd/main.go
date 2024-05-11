package main

import (
	"time"
	"travian-bot/common"
	send_troops "travian-bot/send-troops"
)

func main() {
	common.Login()

	go send_troops.SendTroops(common.ImpTargets, "Imperatoris")
	//time.Sleep(time.Second * 333)
	//go send_troops.SendTroops(common.LegsTargets, "Legs")
	//time.Sleep(time.Second * 111)
	//go train_troops.TrainTroops()

	for {
		time.Sleep(time.Minute * 100)
	}
}
