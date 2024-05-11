package common

type Target struct {
	X          int
	Y          int
	TroopsType int
	Count      int
}

const LEGIONER = 1
const IMPERIAN = 3
const IMPERATORIS = 5

var ImpTargets = []Target{
	{46, 16, IMPERATORIS, 1}, // distance: 1.414214
	{47, 17, IMPERATORIS, 1}, // distance: 2.000000
	{43, 17, IMPERATORIS, 1}, // distance: 2.000000
	{43, 18, IMPERATORIS, 1}, // distance: 2.236068
	{43, 19, IMPERATORIS, 1}, // distance: 2.828427
	{47, 19, IMPERATORIS, 1}, // distance: 2.828427
	{45, 14, IMPERATORIS, 1}, // distance: 3.000000
	{41, 17, IMPERATORIS, 1}, // distance: 4.000000
	{44, 13, IMPERATORIS, 1}, // distance: 4.123106
	{48, 14, IMPERATORIS, 1}, // distance: 4.242641
	{43, 13, IMPERATORIS, 1}, // distance: 4.472136
	{41, 19, IMPERATORIS, 1}, // distance: 4.472136
	{42, 13, IMPERATORIS, 1}, // distance: 5.000000
	{42, 21, IMPERATORIS, 1}, // distance: 5.000000
	{48, 21, IMPERATORIS, 1}, // distance: 5.000000
	{49, 14, IMPERATORIS, 1}, // distance: 5.000000
	{50, 15, IMPERATORIS, 1}, // distance: 5.385165
	{41, 21, IMPERATORIS, 1}, // distance: 5.656854
	{39, 16, IMPERATORIS, 1}, // distance: 6.082763
	{47, 23, IMPERATORIS, 1}, // distance: 6.324555
	{49, 22, IMPERATORIS, 1}, // distance: 6.403124
	{38, 16, IMPERATORIS, 1}, // distance: 7.071068
	{52, 16, IMPERATORIS, 1}, // distance: 7.071068
	{52, 18, IMPERATORIS, 1}, // distance: 7.071068
	{44, 24, IMPERATORIS, 1}, // distance: 7.071068
	{40, 12, IMPERATORIS, 1}, // distance: 7.071068
	{49, 11, IMPERATORIS, 1}, // distance: 7.211103
	{51, 13, IMPERATORIS, 1}, // distance: 7.211103
	{38, 15, IMPERATORIS, 1}, // distance: 7.280110
	{52, 15, IMPERATORIS, 1}, // distance: 7.280110
	{42, 24, IMPERATORIS, 1}, // distance: 7.615773
	{37, 15, IMPERATORIS, 1}, // distance: 8.246211
	{43, 25, IMPERATORIS, 1}, // distance: 8.246211
	{53, 20, IMPERATORIS, 1}, // distance: 8.544004
	{42, 25, IMPERATORIS, 1}, // distance: 8.544004
	{48, 9, IMPERATORIS, 1},  // distance: 8.544004
	{49, 9, IMPERATORIS, 1},  // distance: 8.944272
	{45, 8, IMPERATORIS, 1},  // distance: 9.000000
	{46, 26, IMPERATORIS, 1}, // distance: 9.055385
	{36, 19, IMPERATORIS, 1}, // distance: 9.219544
	{53, 22, IMPERATORIS, 1}, // distance: 9.433981
	{37, 22, IMPERATORIS, 1}, // distance: 9.433981
	{54, 13, IMPERATORIS, 1}, // distance: 9.848858
	{38, 10, IMPERATORIS, 1}, // distance: 9.899495
	{53, 11, IMPERATORIS, 1}, // distance: 10.000000
	//{40, 26, IMPERATORIS, 1}, // distance: 10.295630
	//{36, 23, IMPERATORIS, 1}, // distance: 10.816654
	//{45, 6, IMPERATORIS, 1},  // distance: 11.000000
	//{54, 24, IMPERATORIS, 1}, // distance: 11.401754
	//{48, 6, IMPERATORIS, 1},  // distance: 11.401754
	//{56, 21, IMPERATORIS, 1}, // distance: 11.704700
	//{47, 5, IMPERATORIS, 1},  // distance: 12.165525
	//{55, 10, IMPERATORIS, 1}, // distance: 12.206556
	//{35, 10, IMPERATORIS, 1}, // distance: 12.206556
	//{39, 28, IMPERATORIS, 1}, // distance: 12.529964
	//{33, 21, IMPERATORIS, 1}, // distance: 12.649111
	//{49, 5, IMPERATORIS, 1},  // distance: 12.649111
	//{54, 26, IMPERATORIS, 1}, // distance: 12.727922
	//{55, 9, IMPERATORIS, 1},  // distance: 12.806248
	//{40, 5, IMPERATORIS, 1},  // distance: 13.000000
	//{32, 17, IMPERATORIS, 1}, // distance: 13.000000
	//{40, 29, IMPERATORIS, 1}, // distance: 13.000000
	//{45, 30, IMPERATORIS, 1}, // distance: 13.000000
	//{33, 22, IMPERATORIS, 1}, // distance: 13.000000
	//{32, 16, IMPERATORIS, 1}, // distance: 13.038405
	//{38, 6, IMPERATORIS, 1},  // distance: 13.038405
	//{39, 29, IMPERATORIS, 1}, // distance: 13.416408
	//{36, 27, IMPERATORIS, 1}, // distance: 13.453624
	//{41, 30, IMPERATORIS, 1}, // distance: 13.601471
	//{32, 21, IMPERATORIS, 1}, // distance: 13.601471
	//{46, 3, IMPERATORIS, 1},  // distance: 14.035669
	//{47, 3, IMPERATORIS, 1},  // distance: 14.142136
	//{39, 30, IMPERATORIS, 1}, // distance: 14.317821
	//{39, 4, IMPERATORIS, 1},  // distance: 14.317821
	//{59, 21, IMPERATORIS, 1}, // distance: 14.560220
	//{41, 3, IMPERATORIS, 1},  // distance: 14.560220
	//{35, 6, IMPERATORIS, 1},  // distance: 14.866069
	//{56, 27, IMPERATORIS, 1}, // distance: 14.866069
	//{60, 17, IMPERATORIS, 1}, // distance: 15.000000
	//{30, 16, IMPERATORIS, 1}, // distance: 15.033296
	//{46, 2, IMPERATORIS, 1},  // distance: 15.033296
	//{60, 16, IMPERATORIS, 1}, // distance: 15.033296
	//{60, 19, IMPERATORIS, 1}, // distance: 15.132746
	//{60, 15, IMPERATORIS, 1}, // distance: 15.132746
	//{39, 3, IMPERATORIS, 1},  // distance: 15.231546
	//{60, 20, IMPERATORIS, 1}, // distance: 15.297059
	//{48, 2, IMPERATORIS, 1},  // distance: 15.297059
	//{60, 13, IMPERATORIS, 1}, // distance: 15.524175
	//{59, 10, IMPERATORIS, 1}, // distance: 15.652476
	//{31, 10, IMPERATORIS, 1}, // distance: 15.652476
	//{58, 26, IMPERATORIS, 1}, // distance: 15.811388
	//{58, 8, IMPERATORIS, 1},  // distance: 15.811388
	//{59, 25, IMPERATORIS, 1}, // distance: 16.124515
	//{37, 3, IMPERATORIS, 1},  // distance: 16.124515
	//{34, 5, IMPERATORIS, 1},  // distance: 16.278821
	//{54, 32, IMPERATORIS, 1}, // distance: 17.492856
	//{33, 30, IMPERATORIS, 1}, // distance: 17.691806
	//{32, 29, IMPERATORIS, 1}, // distance: 17.691806
	//{59, 28, IMPERATORIS, 1}, // distance: 17.804494
	//{60, 28, IMPERATORIS, 1}, // distance: 18.601075
	//{30, 3, IMPERATORIS, 1},  // distance: 20.518285
}
