package main

import (
	"time"
	"travian-bot/adventure"
	"travian-bot/build"
	"travian-bot/common"
	train_troops "travian-bot/train-troops"
)

const (
	v1 = 18054
	v2 = 20342
	v3 = 24271
	v4 = 25945
)

func main() {
	common.Login()

	go train_troops.TrainTroops(v1, train_troops.PretsTrainConfig)
	time.Sleep(time.Second * 10)
	go train_troops.TrainTroops(v2, train_troops.LegsTrainConfig)
	time.Sleep(time.Second * 10)
	go train_troops.TrainTroops(v3, train_troops.LegsTrainConfig)

	go launchBuild()

	time.Sleep(time.Second * 10)

	go adventure.GoToAdventures()

	for {
		time.Sleep(time.Minute * 100)
	}
}

func launchBuild() {
	//go build.Build(v1, common.CreateBuildingList(common.Tavern))
	//time.Sleep(time.Second * 10)
	//
	//go build.Build(v2, common.CreateBuildingList(common.Ambar))
	//time.Sleep(time.Second * 10)
	//
	//go build.Build(v3, common.CreateBuildingList(common.MainBuilding))
	//time.Sleep(time.Second * 10)
	//go build.Build(v3, common.CreateBuildingList(common.BuildingPlan4536...))
	//time.Sleep(time.Second * 10)

	//go build.Build(v4, common.CreateBuildingList(common.Ambar))
	time.Sleep(time.Second * 10)
	go build.Build(v4, common.CreateBuildingList(common.BuildingPlan5346...))
}
