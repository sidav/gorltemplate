package main

type level struct {
	tiles  [][]tile
	actors []*actor
}

func (l *level) coordsValid(x, y int) bool {
	return x >= 0 && y >= 0 && x < len(l.tiles) && y < len(l.tiles[0])
}
