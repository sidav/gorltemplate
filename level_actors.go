package main

func (l *level) cleanDeadActors() {
	for i := len(l.actors)-1; i >= 0; i-- {
		if l.actors[i].hp <= 0 {
			// switch i-th and last element from slice, remove last element from slice
			l.actors[i] = l.actors[len(l.actors)-1]
			l.actors = l.actors[:len(l.actors)-1]
		}
	}
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

// true if door was opened
func (l *level) tryOpenDoorForActor(a *actor, vx, vy int) bool {
	doorX, doorY := a.x + vx, a.y + vy
	if l.coordsValid(doorX, doorY) {
		if l.tiles[doorX][doorY].asDoor != nil && l.tiles[doorX][doorY].asDoor.lockLevel == 0 && !l.tiles[doorX][doorY].asDoor.isOpened {
			l.tiles[doorX][doorY].asDoor.isOpened = true
			return true
		}
	}
	return false
}

func (l *level) tryActivateSwitchAsActor(a *actor, vx, vy int) bool {
	switchX, switchY := a.x + vx, a.y + vy
	return l.activateSwitchAt(switchX, switchY)
}

func (l *level) tryMoveActorByVector(a *actor, vx, vy int) bool {
	x, y := a.x, a.y
	if l.isTilePassableFor(x+vx, y+vy, a) {
		a.x += vx
		a.y += vy
		return true
	}
	return false
}
