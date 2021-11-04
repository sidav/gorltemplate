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

	seededRnd fibrandom.FibRandom // for seed-based recreation
	rnd       fibrandom.FibRandom // for everything else
)

func gameLoop() {
	log.Init(2)
	rnd.InitDefault()
	seededRnd.InitDefault()

	initLevel()

	for GAMEISRUNNING {
		if GAMETICK%TicksInTurn == 0 {
			for GAMEISRUNNING && !CURRENTLEVEL.tryExecuteActorsIntent(PLAYERCONTROLLER.player) {
				PLAYERCONTROLLER.playerTurn()
			}
		}

		for i := range CURRENTLEVEL.actors {
			CURRENTLEVEL.actors[i].aiAct()
		}

		for i := range CURRENTLEVEL.actors {
			CURRENTLEVEL.tryExecuteActorsIntent(CURRENTLEVEL.actors[i])
		}

		GAMETICK++
	}
}
