package dijkstra_maps_test

import (
	cw "gorltemplate/console_wrapper"
	"strconv"
	"gorltemplate/fibrandom"
)

var dmap [][]int // -1 means unpassable
var rnd fibrandom.FibRandom
var (
	px     = SIZE /2
	py     = SIZE /2
	babaiX = SIZE/2+1
	babaiY = SIZE/2+1
)

const SIZE = 25

func getNonZerosSumNear(xx, yy int) int {
	sum := 0
	for x := xx-1; x <= xx+1; x++ {
		for y := yy-1; y <= yy+1; y++ {
			if x != xx && y != yy {
				continue
			}
			if x >= 0 && x < len(dmap) && y >= 0 && y < len(dmap[x]) && dmap[x][y] > 0 {
				sum += dmap[x][y]
			}
		}
	}
	return sum
}

func getMinValueNear(xx, yy int) int {
	minValue := 0
	for x := xx-1; x <= xx+1; x++ {
		for y := yy-1; y <= yy+1; y++ {
			if x != xx && y != yy {
				continue
			}
			if x >= 0 && x < len(dmap) && y >= 0 && y < len(dmap[x]) {
				if dmap[x][y] > 0&& (minValue == 0 || dmap[x][y] < minValue) {
					minValue = dmap[x][y]
				}
			}
		}
	}
	return minValue
}

func getBestStepFrom(xx, yy int) (int, int) {
	minValue := 0
	vx, vy := 0, 0
	for x := xx-1; x <= xx+1; x++ {
		for y := yy-1; y <= yy+1; y++ {
			if x >= 0 && x < len(dmap) && y >= 0 && y < len(dmap[x]) {
				if x == xx && y == yy || dmap[x][y] < 0 {
					continue
				}
				if minValue == 0 || (dmap[x][y] > 0 && dmap[x][y] <= minValue) {
					if dmap[x][y] == minValue && rnd.OneChanceFrom(3) {
						break
					}
					minValue = dmap[x][y]
					vx = x-xx
					vy = y-yy
				}
			}
		}
	}
	if minValue == 0 {
		return 0, 0
	}
	return vx, vy
}

func initDijkstraMap(targetCoords, wallCoords [][2]int) {
	dmap = make([][]int, SIZE)
	for i := range dmap {
		dmap[i] = make([]int, SIZE)
	}
	// minimum is 1, 0 is count as non-set!
	for i := range targetCoords {
		dmap[targetCoords[i][0]][targetCoords[i][1]] = 1
	}
	for i := range wallCoords {
		dmap[wallCoords[i][0]][wallCoords[i][1]] = -1
	}
	dmap[babaiX][babaiY] = 9999999999
	// fill
	atLeastOneChanged := true
	for atLeastOneChanged {
		listOfCoordsToUpd := make([][3]int, 0)
		atLeastOneChanged = false
		for x := range dmap {
			for y := range dmap[x] {
				if dmap[x][y] == 0 {
					minv := getNonZerosSumNear(x, y)
					//minv = getMinValueNear(x, y)
					if minv > 0 {
						atLeastOneChanged = true
						listOfCoordsToUpd = append(listOfCoordsToUpd, [3]int{x, y, minv + 1})
					}
				}
			}
		}
		for i := range listOfCoordsToUpd {
			x := listOfCoordsToUpd[i][0]
			y := listOfCoordsToUpd[i][1]
			dmap[x][y] = listOfCoordsToUpd[i][2]
		}
	}
}

func Simulate() {
	rnd.InitDefault()
	tcoords := make([][2]int, 0)
	for i := 0; i < 10; i++ {
		tcoords = append(tcoords, [2]int{rnd.Rand(15), rnd.Rand(15)})
	}
	walls := make([][2]int, 0)
	for i := 0; i < 40; i++ {
		walls = append(walls, [2]int{rnd.Rand(15), rnd.Rand(15)})
	}
	initDijkstraMap(tcoords, walls)
	drawmap()
	for cw.ReadKey() != "ESCAPE" {
		vx, vy := getBestStepFrom(px, py)
		px += vx
		py += vy
		if dmap[px][py] == 1 {
			for i := range tcoords {
				if tcoords[i][0] == px && tcoords[i][1] == py {
					tcoords = append(tcoords[:i], tcoords[i+1:]...)
					break
				}
			}
		}
		babaiX += rnd.RandInRange(-1, 1)
		babaiY += rnd.RandInRange(-1, 1)
		if babaiX < 0 {
			babaiX = 0
		}
		if babaiY < 0 {
			babaiY = 0
		}
		if babaiX >= SIZE {
			babaiX = SIZE -1
		}
		if babaiY >= SIZE {
			babaiY = SIZE -1
		}
		initDijkstraMap(tcoords, walls)
		drawmap()
	}
}

func drawmap() {
	for x := range dmap {
		for y := range dmap[x] {
			if dmap[x][y] == -1 {
				cw.SetColor(cw.BLACK, cw.RED)
				cw.PutChar(' ', x, y)
			} else if dmap[x][y] == 1 {
				cw.SetColor(cw.MAGENTA, cw.BLACK)
				cw.PutChar('*', x, y)
			} else if dmap[x][y] < 10 {
				cw.SetColor(cw.WHITE, cw.BLACK)
				cw.PutChar(rune(strconv.Itoa(dmap[x][y])[0]), x, y)
			} else {
				cw.SetColor(cw.WHITE, cw.BLACK)
				cw.PutChar(' ', x, y)
			}
		}
	}
	cw.SetColor(cw.BLACK, cw.WHITE)
	cw.PutChar('@', px, py)
	cw.SetColor(cw.DARK_RED, cw.WHITE)
	cw.PutChar('B', babaiX, babaiY)
	cw.Flush_console()
}
