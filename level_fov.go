package main

import fov "gorltemplate/fov"

func (l *level) getFovMapFrom(x, y, radius int) [][]bool{
	return fov.GetFovMapFrom(
		x, y, radius,
		len(l.tiles), len(l.tiles[0]),
		func(x, y int) bool {
			return l.tiles[x][y].data.opaque
		},
		)
}
