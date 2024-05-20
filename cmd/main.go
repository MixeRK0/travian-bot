package main

import (
	"time"
	"travian-bot/build"
	"travian-bot/common"
	train_troops "travian-bot/train-troops"
)

func main() {
	common.Login()

	go build.Build(24271)
	time.Sleep(time.Second * 30)
	go train_troops.TrainTroops(18054)

	for {
		time.Sleep(time.Minute * 100)
	}
}
