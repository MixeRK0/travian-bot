package main

import (
	"time"
	"travian-bot/common"
	send_troops "travian-bot/send-troops"
	train_troops "travian-bot/train-troops"
)

func main() {
	common.Login()

	go send_troops.SendTroops(common.ImperatorisTargets1, "Imperatoris")
	time.Sleep(time.Second * 333)
	go train_troops.TrainTroops()

	for {
		time.Sleep(time.Minute * 100)
	}
}
