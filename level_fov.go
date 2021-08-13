package main

import fov "gorltemplate/fov"

func (l *level) getFovMapFrom(x, y, radius int) [][]bool{
	return fov.GetFovMapFrom(
		x, y, radius,
		len(l.tiles), len(l.tiles[0]),
		func(x, y int) bool {
			return l.tiles[x][y].getStaticData().opaque
		},
		)
}

func (l *level) updateWasSeenFromFovMap(fovMap [][]bool) {
	for x := range fovMap {
		for y := range fovMap[x] {
			l.tiles[x][y].wasSeenPreviously = l.tiles[x][y].wasSeenPreviously || fovMap[x][y]
		}
	}
}
