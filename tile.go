package main

type tile struct {
	data   *tileStatic
	asDoor *tileDoor
	wasSeenPreviously bool
}

type tileDoor struct {
	isOpened  bool
	lockLevel int
}

func (t *tile) getStaticData() *tileStatic {
	if t.asDoor != nil {
		if t.asDoor.isOpened {
			return getTileStatics("_DOOR_OPENED")
		}
		return getTileStatics("_DOOR_CLOSED")
	}
	return t.data
}

//func (t *tile) getTileVisuals() (rune, int, int) {
//	if t.asDoor != nil {
//		if t.asDoor.isOpened {
//			staticData := getTileStatics("_DOOR_OPENED")
//			return staticData.char, staticData.fgcolor, staticData.bgcolor
//		}
//	}
//}

func (t *tile) isPassable() bool {
	return t.getStaticData().passable
}
