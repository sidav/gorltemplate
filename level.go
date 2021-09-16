package main

type level struct {
	tiles  [][]tile
	actors []*actor
}

func (l *level) coordsValid(x, y int) bool {
	return x >= 0 && y >= 0 && x < len(l.tiles) && y < len(l.tiles[0])
}

func (l *level) countTilesAround(x, y int, includeCenter, reverseCount bool) int {
	return 0 // TODO
}

func (l *level) isTilePassable(x, y int) bool {
	return l.coordsValid(x, y) && l.tiles[x][y].isPassable() && !l.isAnyActorPresentAt(x, y)
}

func (l *level) isTilePotentiallyPassable(x, y int) bool {
	return l.coordsValid(x, y) && (l.tiles[x][y].isPassable() || l.tiles[x][y].asDoor != nil)
}

func (l *level) activateSwitchAt(sx, sy int) bool {
	if l.coordsValid(sx, sy) {
		if l.tiles[sx][sy].asSwitch != nil {
			newState := !l.tiles[sx][sy].asSwitch.isActivated
			l.tiles[sx][sy].asSwitch.isActivated = newState
			lockLevel := l.tiles[sx][sy].asSwitch.lockLevel

			for x := 0; x < len(l.tiles); x++ {
				for y := 0; y < len(l.tiles[x]); y++ {
					if l.tiles[x][y].asDoor != nil && l.tiles[x][y].asDoor.lockLevel == lockLevel {
						l.tiles[x][y].asDoor.isOpened = newState
					}
				}
			}
			return true
		}
	}
	return false
}
