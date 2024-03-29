package main

import (
	"gorltemplate/fibrandom"
	"gorltemplate/game_log"
)

const (
	TicksInTurn = 10
)

var (
	GAMEISRUNNING    = true
	CURRENTLEVEL     = level{}
	RENDERER         renderer
	PLAYERCONTROLLER playerController
	GAMETICK         int

	log game_log.GameLog

	seededRnd    fibrandom.FibRandom // for seed-based recreation
	rnd          fibrandom.FibRandom // for everything else
	playerFovMap [][]bool
)

func gameLoop() {
	log.Init(2)
	rnd.InitDefault()
	seededRnd.InitDefault()

	initLevel()

	for GAMEISRUNNING {
		if GAMETICK%TicksInTurn == 0 {
			for GAMEISRUNNING && !CURRENTLEVEL.tryExecuteActorsIntent(PLAYERCONTROLLER.player) {
				playerFovMap = CURRENTLEVEL.getFovMapFrom(PLAYERCONTROLLER.player.x, PLAYERCONTROLLER.player.y, 10)
				CURRENTLEVEL.updateWasSeenFromFovMap(playerFovMap)
				PLAYERCONTROLLER.playerTurn()
			}
		}

		cleanupNeeded := false
		for i := range CURRENTLEVEL.actors {
			if CURRENTLEVEL.actors[i].hp <= 0 {
				cleanupNeeded = true
			} else {
				CURRENTLEVEL.actors[i].aiAct()
			}
		}
		if cleanupNeeded {
			CURRENTLEVEL.cleanDeadActors()
		}

		for i := range CURRENTLEVEL.actors {
			CURRENTLEVEL.tryExecuteActorsIntent(CURRENTLEVEL.actors[i])
		}

		GAMETICK++
	}
}
