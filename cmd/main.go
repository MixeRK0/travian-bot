package main

import (
	"time"
	"travian-bot/build"
	"travian-bot/common"
	train_troops "travian-bot/train-troops"
)

const (
	v1 = 20215
	v2 = 23298
	v3 = 24004
	v4 = 25232
)

func main() {
	common.Login()
	time.Sleep(time.Second)

	go train_troops.TrainTroops(v1, train_troops.ImperatorisTrainConfig)
	time.Sleep(time.Second * 5)
	go train_troops.TrainTroops(v2, train_troops.LegsTrainConfig)
	time.Sleep(time.Second * 5)
	go train_troops.TrainTroops(v3, train_troops.LegsTrainConfig)
	time.Sleep(time.Second * 5)

	go launchBuild()

	//go adventure.GoToAdventures()

	for {
		time.Sleep(time.Minute * 100)
	}
}

func launchBuild() {

	go build.Build(v4, common.BuildingPlan4446, true)

	time.Sleep(time.Second * 10)
	go build.Build(v4, common.CreateBuildingList(
		common.Warehouse,
		common.Ambar,
		common.Warehouse,
		common.Ambar,
		common.Warehouse,
		common.Ambar,
		common.Warehouse,
		common.Ambar,
		common.MainBuilding,
		common.MainBuilding,
		common.MainBuilding,
		common.MainBuilding,
		common.MainBuilding,
		common.MainBuilding,
		common.MainBuilding,
		common.MainBuilding,
		common.MainBuilding,
		common.Warehouse,
		common.Ambar,
		common.Warehouse,
		common.Ambar,
		common.MainBuilding,
		common.MainBuilding,
		common.MainBuilding,
		common.MainBuilding,
	), false)

	time.Sleep(time.Second * 10)
	go build.Build(v3, common.CreateBuildingList(
		common.Casarm,
		common.Casarm,
		common.Casarm,
		common.Casarm,
		common.Casarm,
		common.Casarm,
		common.Casarm,
		common.Casarm,
		common.Casarm,
	), false)
}
