package main

import (
	"time"
	"travian-bot/adventure"
	"travian-bot/build"
	"travian-bot/common"
)

const (
	v1 = 20215
)

func main() {
	common.Login()

	//go train_troops.TrainTroops(v1, train_troops.PretsTrainConfig)
	//time.Sleep(time.Second * 10)

	go launchBuild()

	go adventure.GoToAdventures()

	for {
		time.Sleep(time.Minute * 100)
	}
}

func launchBuild() {
	time.Sleep(time.Second * 10)
	go build.Build(v1, common.CreateBuildingList(common.BuildingPlan4446...), true)

	time.Sleep(time.Second * 10)
	go build.Build(v1, common.CreateBuildingList(common.Warehouse, common.Warehouse, common.Warehouse, common.Ambar, common.Ambar, common.Ambar, common.MainBuilding, common.MainBuilding, common.MainBuilding, common.MainBuilding, common.MainBuilding, common.Warehouse, common.Warehouse, common.Ambar, common.Ambar, common.Marketplace, common.Marketplace, common.Marketplace, common.Marketplace, common.Academy, common.Academy, common.Academy, common.Cranny, common.Cranny, common.Cranny, common.Cranny, common.Cranny, common.Cranny, common.Cranny, common.Cranny, common.Wall, common.Wall, common.Wall, common.Wall, common.Wall, common.Wall, common.Wall, common.Warehouse, common.Warehouse, common.Ambar, common.Ambar), false)
}
