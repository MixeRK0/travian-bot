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
)

func main() {
	common.Login()

	go train_troops.TrainTroops(v1, train_troops.ImperatorisTrainConfig)

	go launchBuild()

	//go adventure.GoToAdventures()

	for {
		time.Sleep(time.Minute * 100)
	}
}

func launchBuild() {
	time.Sleep(time.Second * 10)
	go build.Build(v2, common.CreateBuildingList(common.BuildingPlan4446...), true)

	time.Sleep(time.Second * 10)
	go build.Build(v2, common.CreateBuildingList(
		common.Ambar,
		common.Ambar,
		common.Marketplace,
		common.Marketplace,
		common.Marketplace,
		common.Marketplace,
		common.Marketplace,
		common.MainBuilding,
		common.MainBuilding,
		common.MainBuilding,
		common.MainBuilding,
		common.MainBuilding,
		common.Casarm,
		common.Casarm,
		common.Casarm,
		common.Casarm,
		common.Casarm,
		common.Casarm,
	), false)

	time.Sleep(time.Second * 10)
	go build.Build(v3, common.CreateBuildingList(common.BuildingPlan5346...), true)

	time.Sleep(time.Second * 10)
	go build.Build(v3, common.CreateBuildingList(
		common.Warehouse,
		common.Warehouse,
		common.Warehouse,
		common.Ambar,
		common.Ambar,
		common.MainBuilding,
		common.MainBuilding,
		common.MainBuilding,
		common.MainBuilding,
		common.MainBuilding,
		common.Warehouse,
		common.Ambar,
		common.Ambar,
		common.Warehouse,
		common.Warehouse,
		common.Ambar,
		common.Ambar,
		common.MainBuilding,
		common.MainBuilding,
		common.MainBuilding,
		common.MainBuilding,
		common.Warehouse,
		common.Warehouse,
		common.Ambar,
		common.Ambar,
		common.MainBuilding,
		common.MainBuilding,
		common.MainBuilding,
		common.MainBuilding,
	), false)
}
