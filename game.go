package main

var (
	GAMEISRUNNING = true
	CURRENTLEVEL = level{}
	RENDERER renderer
	PLAYERCONTROLLER playerController
)

func gameLoop() {
	for GAMEISRUNNING {
		PLAYERCONTROLLER.playerTurn()
	}
}
