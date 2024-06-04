package main

import (
	"time"
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

	//go adventure.GoToAdventures()

	for {
		time.Sleep(time.Minute * 100)
	}
}

func launchBuild() {
	time.Sleep(time.Second * 10)
	go build.Build(v1, common.CreateBuildingList(common.BuildingPlan4446...), true)

	time.Sleep(time.Second * 10)
	go build.Build(v1, common.CreateBuildingList(common.MainBuilding, common.MainBuilding, common.MainBuilding, common.MainBuilding, common.Ambar, common.Ambar, common.Wall, common.Wall, common.Wall, common.Wall, common.Wall, common.Wall, common.Cranny, common.Cranny, common.Cranny, common.Cranny, common.Cranny), false)
}
