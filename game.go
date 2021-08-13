package main

import "gorltemplate/fibrandom"

const (
	TicksInTurn = 10
)

var (
	GAMEISRUNNING    = true
	CURRENTLEVEL     = level{}
	RENDERER         renderer
	PLAYERCONTROLLER playerController
	GAMETICK         int

	seededRnd fibrandom.FibRandom // for seed-based recreation
	rnd       fibrandom.FibRandom // for everything else
)

func gameLoop() {
	rnd.InitDefault()
	seededRnd.InitDefault()

	initLevel()

	for GAMEISRUNNING {
		if GAMETICK%TicksInTurn == 0 {
			PLAYERCONTROLLER.playerTurn()
		}

		for i := range CURRENTLEVEL.actors {
			CURRENTLEVEL.actors[i].decide()
		}

		for i := range CURRENTLEVEL.actors {
			CURRENTLEVEL.executeActorsIntent(CURRENTLEVEL.actors[i])
		}

		GAMETICK++
	}
}
