package main

import (
	"time"
	"travian-bot/common"
	send_troops "travian-bot/send-troops"
	train_troops "travian-bot/train-troops"
)

func main() {
	common.Login()

	go send_troops.SendTroops(common.LegsTargets, "Legs")
	go func() {
		println("Sleep before send")
		time.Sleep(time.Hour * 3)
		send_troops.SendTroops(common.ImperatorisTargets, "Imperatoris")
	}()
	go func() {
		println("Sleep before train")
		time.Sleep(time.Hour * 2)
		train_troops.TrainTroops()
	}()

	for {
		time.Sleep(time.Minute * 100)
	}
}
