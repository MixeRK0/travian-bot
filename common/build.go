package common

type BuildingId struct {
	Id  int
	Gid int
}

var BuildingPlan = []BuildingId{
	// wood
	{1, 1},
	{3, 1},
	{14, 1},
	{17, 1},

	// clay
	{4, 2}, // iron
	{5, 2},
	{6, 2},
	{16, 2},
	{18, 2},

	// iron
	//{10, 3},
	{7, 3},
	{11, 3},

	// crop
	{2, 4},
	{8, 4},
	{9, 4},
	{12, 4},
	{13, 4},
	{15, 4},
}
