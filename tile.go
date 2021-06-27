package main

type tile struct {
	data   *tileStatic
	asDoor *tileDoor
}

type tileDoor struct {
	isOpened  bool
	lockLevel int
}

func (t *tile) isPassable() bool {
	if t.asDoor != nil {
		return t.asDoor.isOpened
	}
	return t.data.passable
}
