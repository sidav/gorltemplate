package main

type level struct {
	tiles  [][]tile
	actors []*actor
}

func (l *level) isAnyActorPresentAt(x, y int) bool {
	for i := range l.actors {
		if l.actors[i].x == x && l.actors[i].y == y {
			return true
		}
	}
	return false
}

func (l *level) getFirstActorAtCoords(x, y int) *actor {
	for i := range l.actors {
		if l.actors[i].x == x && l.actors[i].y == y {
			return l.actors[i]
		}
	}
	return nil
}

func (l *level) getAllActorsAtCoords(x, y int) []*actor {
	var actors []*actor
	for i := range l.actors {
		if l.actors[i].x == x && l.actors[i].y == y {
			actors = append(actors, l.actors[i])
		}
	}
	return actors
}

func (l *level) moveActorByVector(a *actor, vx, vy int) {
	x, y := a.x, a.y
	if l.coordsValid(x+vx, y+vy) && l.tiles[x+vx][y+vy].isPassable() {
		a.x += vx
		a.y += vy
	}
}

func (l *level) coordsValid(x, y int) bool {
	return x >= 0 && y >= 0 && x < len(l.tiles) && y < len(l.tiles[0])
}
