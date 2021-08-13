package main

type tile struct {
	data   *tileStatic
	asDoor *tileDoor
	asSwitch *tileSwitch
	wasSeenPreviously bool
}

type tileDoor struct {
	isOpened  bool
	lockLevel int
}

type tileSwitch struct {
	isActivated bool
	lockLevel int
}

func (t *tile) getStaticData() *tileStatic {
	if t.asSwitch != nil {
		if t.asSwitch.isActivated {
			return getTileStatics("_SWITCH_ON")
		}
		return getTileStatics("_SWITCH_OFF")
	}
	if t.asDoor != nil {
		if t.asDoor.isOpened {
			return getTileStatics("_DOOR_OPENED")
		}
		if t.asDoor.lockLevel > 0 {
			return getTileStatics("_DOOR_LOCKED")
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
