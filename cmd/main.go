package main

type TargetCoords struct {
	x int
	y int
}

var coords = []TargetCoords{
	//{13, 56},
	//{13, 57},
	//{15, 54},
	//{12, 52},
	//{10, 53},
	//{9, 53},
	//{9, 55},
	//{11, 56},
	{7, 56},
	//{15, 57},
	//{19, 54},
}

const count = 3

func main() {
	SendTroops(coords, count)

	return
}
