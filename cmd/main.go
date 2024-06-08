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
	v5 = 26261
)

func main() {
	common.Login()
	time.Sleep(time.Second)

	go train_troops.TrainTroops(v1, train_troops.ImperatorisTrainConfig)
	time.Sleep(time.Second * 5)
	go func() {
		time.Sleep(time.Minute * 30)
		train_troops.TrainTroops(v1, train_troops.ImpsTrainConfig)
	}()

	go train_troops.TrainTroops(v2, train_troops.PretsTrainConfig)
	time.Sleep(time.Second * 5)
	go train_troops.TrainTroops(v3, train_troops.PretsTrainConfig)
	time.Sleep(time.Second * 5)
	go train_troops.TrainTroops(v4, train_troops.LegsTrainConfig)
	time.Sleep(time.Second * 5)

	go launchBuild()

	//go adventure.GoToAdventures()

	for {
		time.Sleep(time.Minute * 100)
	}
}

func launchBuild() {
	go build.Build(v5, common.BuildingPlan4446, true)

	time.Sleep(time.Second * 10)
	go build.Build(v1, common.CreateBuildingList(
		common.WarehouseSecond,
		common.WarehouseSecond,
		common.WarehouseSecond,
		common.WarehouseSecond,
		common.WarehouseSecond,
		common.Academy,
		common.Academy,
		common.Casarm,
		common.Casarm,
		common.Casarm,
		common.Casarm,
		common.Casarm,
		common.Casarm,
		common.Casarm,
		common.Casarm,
		common.Marketplace,
		common.Marketplace,
		common.Marketplace,
		common.Marketplace,
		common.Marketplace,
		common.Marketplace,
	), false)

	time.Sleep(time.Second * 10)
	go build.Build(v2, common.CreateBuildingList(
		common.Ratushe,
		common.Ratushe,
		common.Ratushe,
		common.Ratushe,
		common.Ratushe,
		common.Marketplace,
		common.Marketplace,
		common.Marketplace,
		common.Marketplace,
		common.Marketplace,
	), false)

	time.Sleep(time.Second * 10)
	go build.Build(v3, common.CreateBuildingList(
		common.Ratushe,
		common.Ratushe,
		common.Ratushe,
		common.MainBuilding,
		common.MainBuilding,
		common.MainBuilding,
		common.MainBuilding,
		common.MainBuilding,
		common.Warehouse,
		common.Ambar,
		common.Ambar,
		common.Ambar,
		common.Marketplace,
		common.Marketplace,
		common.Marketplace,
		common.Marketplace,
		common.Marketplace,
	), false)

	time.Sleep(time.Second * 10)
	go build.Build(v4, common.CreateBuildingList(
		common.MainBuilding,
		common.Marketplace,
		common.Marketplace,
		common.Marketplace,
		common.Marketplace,
		common.Marketplace,
		common.Marketplace,
		common.Marketplace,
		common.Marketplace,
		common.Marketplace,
		common.Marketplace,
		common.Ambar,
		common.Ambar,
		common.Warehouse,
	), false)
	time.Sleep(time.Second * 10)
	go build.Build(v5, common.CreateBuildingList(
		common.Academy,
		common.Academy,
		common.Academy,
		common.Academy,
		common.Academy,
		common.Academy,
		common.Academy,
		common.Academy,
		common.Academy,
		common.MainBuilding,
		common.Residence,
		common.Residence,
		common.Residence,
		common.Residence,
		common.Residence,
		common.Residence,
		common.Residence,
		common.Residence,
		common.Residence,
		common.Casarm,
		common.Casarm,
		common.Casarm,
		common.Casarm,
		common.Casarm,
		common.Casarm,
		common.Casarm,
		common.Casarm,
		common.Ambar,
		common.Ambar,
		common.Ambar,
		common.Warehouse,
		common.Warehouse,
		common.Warehouse,
	), false)
}
