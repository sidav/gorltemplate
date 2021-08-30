package main

import (
	"gorltemplate/console_wrapper"
	"time"
)

type playerController struct {
	player                               *actor
	isRunning                            bool
	runX, runY, prevPlayerX, prevPlayerY int
}

func (pc *playerController) playerTurn() {
	var key string
	fovMap := CURRENTLEVEL.getFovMapFrom(pc.player.x, pc.player.y, 10)
	CURRENTLEVEL.updateWasSeenFromFovMap(fovMap)
	RENDERER.render(pc.player.x, pc.player.y, fovMap)

	if pc.isRunning {
		pc.run()
		time.Sleep(25*time.Millisecond)
		return
	}

	key = console_wrapper.ReadKey()
	switch key {
	case "ESCAPE": GAMEISRUNNING = false
	// simple movement/action
	case "UP", "k": pc.player.setIntent(INTENT_MOVE_OR_OPEN_DOOR, 0, -1, 10)
	case "y": pc.player.setIntent(INTENT_MOVE_OR_OPEN_DOOR, -1, -1, 10)
	case "DOWN", "j": pc.player.setIntent(INTENT_MOVE_OR_OPEN_DOOR, 0, 1, 10)
	case "u": pc.player.setIntent(INTENT_MOVE_OR_OPEN_DOOR, 1, -1, 10)
	case "LEFT", "h": pc.player.setIntent(INTENT_MOVE_OR_OPEN_DOOR, -1,0, 10)
	case "b": pc.player.setIntent(INTENT_MOVE_OR_OPEN_DOOR, -1, 1, 10)
	case "RIGHT", "l": pc.player.setIntent(INTENT_MOVE_OR_OPEN_DOOR, 1, 0, 10)
	case "n": pc.player.setIntent(INTENT_MOVE_OR_OPEN_DOOR, 1, 1, 10)
	// automove
	case "H", "J", "K", "L":
		pc.prevPlayerX, pc.prevPlayerY = -1, -1
		pc.runX, pc.runY = pc.keyToDirectionVector(key)
		pc.isRunning = true
	}
}

// run until something is encountered
func (pc *playerController) run() {
	if pc.prevPlayerX == pc.player.x && pc.prevPlayerY == pc.player.y {
		pc.isRunning = false
		return
	}
	pc.prevPlayerX, pc.prevPlayerY = pc.player.getCoords()
	pc.player.setIntent(INTENT_MOVE_ONLY, pc.runX, pc.runY, 10)
}

func (pc *playerController) keyToDirectionVector(key string) (int, int) {
	switch key {
	case "UP", "k", "K": return 0, -1
	case "DOWN", "j", "J": return 0, 1
	case "LEFT", "h", "H": return -1, 0
	case "RIGHT", "l", "L": return 1, 0
	}
	return 0, 0
}
