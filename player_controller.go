package main

import "gorltemplate/console_wrapper"

type playerController struct {
	player *actor
}

func (pc *playerController) playerTurn() {
	fovMap := CURRENTLEVEL.getFovMapFrom(pc.player.x, pc.player.y, 10)
	CURRENTLEVEL.updateWasSeenFromFovMap(fovMap)
	RENDERER.render(pc.player.x, pc.player.y, fovMap)

	key := console_wrapper.ReadKey()
	switch key {
	case "ESCAPE": GAMEISRUNNING = false
	case "UP", "k": pc.player.setIntent(INTENT_MOVE_OR_OPEN_DOOR, 0, -1, 10)
	case "y": pc.player.setIntent(INTENT_MOVE_OR_OPEN_DOOR, -1, -1, 10)
	case "DOWN", "j": pc.player.setIntent(INTENT_MOVE_OR_OPEN_DOOR, 0, 1, 10)
	case "u": pc.player.setIntent(INTENT_MOVE_OR_OPEN_DOOR, 1, -1, 10)
	case "LEFT", "h": pc.player.setIntent(INTENT_MOVE_OR_OPEN_DOOR, -1,0, 10)
	case "b": pc.player.setIntent(INTENT_MOVE_OR_OPEN_DOOR, -1, 1, 10)
	case "RIGHT", "l": pc.player.setIntent(INTENT_MOVE_OR_OPEN_DOOR, 1, 0, 10)
	case "n": pc.player.setIntent(INTENT_MOVE_OR_OPEN_DOOR, 1, 1, 10)
	}
}
