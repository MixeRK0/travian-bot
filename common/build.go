package common

type BuildingId struct {
	Id   int
	Gid  int
	Dorf int
}

func CreateBuildingList(items ...BuildingId) []BuildingId {
	res := make([]BuildingId, 0)
	for _, item := range items {
		res = append(res, item)
	}

	return res
}

var MainBuilding = BuildingId{26, 15, 2}
var Warehouse = BuildingId{20, 10, 2}
var Ambar = BuildingId{21, 11, 2}
var Casarm = BuildingId{31, 19, 2}
var Stable = BuildingId{29, 20, 2}
var Academy = BuildingId{30, 22, 2}
var Wall = BuildingId{40, 31, 2}
var RallyPoint = BuildingId{39, 16, 2}
var Cranny = BuildingId{35, 23, 2}
var Tavern = BuildingId{37, 37, 2}
var Ratushe = BuildingId{32, 24, 2}
var Marketplace = BuildingId{22, 17, 2}

var BuildingPlan4536 = []BuildingId{
	// wood
	{1, 1, 1},
	{3, 1, 1},
	{14, 1, 1},
	{17, 1, 1},

	// clay
	{4, 2, 1}, // iron
	{5, 2, 1},
	{6, 2, 1},
	{16, 2, 1},
	{18, 2, 1},

	// iron
	//{10, 3},
	{7, 3, 1},
	{11, 3, 1},

	// crop
	{2, 4, 1},
	{8, 4, 1},
	{9, 4, 1},
	{12, 4, 1},
	{13, 4, 1},
	{15, 4, 1},
}

var BuildingPlan5346 = []BuildingId{
	// wood
	{1, 1, 1},
	{3, 1, 1},
	{5, 1, 1},
	{14, 1, 1},
	{17, 1, 1},

	// clay
	{6, 2, 1},
	{16, 2, 1},
	{18, 2, 1},

	// iron
	{4, 3, 1},
	{7, 3, 1},
	{10, 3, 1},
	{11, 3, 1},

	// crop
	{2, 4, 1},
	{8, 4, 1},
	{9, 4, 1},
	{12, 4, 1},
	{13, 4, 1},
	{15, 4, 1},
}

var BuildingPlan4446 = []BuildingId{
	// wood
	//{1, 1, 1},
	{3, 1, 1},
	//{14, 1, 1},
	{17, 1, 1},

	// clay
	{5, 2, 1},
	//{6, 2, 1},
	{16, 2, 1},
	{18, 2, 1},

	// iron
	{4, 3, 1},
	{7, 3, 1},
	{10, 3, 1},
	{11, 3, 1},

	// crop
	{2, 4, 1},
	{8, 4, 1},
	{9, 4, 1},
	{12, 4, 1},
	{13, 4, 1},
	{15, 4, 1},
}
