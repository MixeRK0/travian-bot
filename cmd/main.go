package main

import (
	"time"
	train_troops "travian-bot/train-troops"
)

func main() {
	//common.Login()

	//go send_troops.SendTroops(common.ImpTargets, "Imperatoris")
	//time.Sleep(time.Second * 333)
	//go send_troops.SendTroops(common.LegsTargets, "Legs")
	//time.Sleep(time.Second * 60 * 30)
	go train_troops.TrainTroops()

	for {
		time.Sleep(time.Minute * 100)
	}
}
