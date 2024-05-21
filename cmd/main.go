package main

import (
	"time"
	"travian-bot/common"
	train_troops "travian-bot/train-troops"
)

func main() {
	common.Login()

	go train_troops.TrainTroops(18054, train_troops.PretsTrainConfig)
	time.Sleep(time.Second * 10)
	go train_troops.TrainTroops(20342, train_troops.LegsTrainConfig)
	time.Sleep(time.Second * 10)
	go train_troops.TrainTroops(24271, train_troops.LegsTrainConfig)

	for {
		time.Sleep(time.Minute * 100)
	}
}
