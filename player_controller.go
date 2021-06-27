package main

import "gorltemplate/console_wrapper"

type playerController struct {
	player *actor
}

func (pc *playerController) playerTurn() {
	RENDERER.render(pc.player.x, pc.player.y)

	key := console_wrapper.ReadKey()
	switch key {
	case "ESCAPE": GAMEISRUNNING = false
	case "UP": CURRENTLEVEL.moveActorByVector(pc.player, 0, -1)
	case "DOWN": CURRENTLEVEL.moveActorByVector(pc.player, 0, 1)
	case "LEFT": CURRENTLEVEL.moveActorByVector(pc.player, -1,0)
	case "RIGHT": CURRENTLEVEL.moveActorByVector(pc.player, 1, 0)
	}
}
