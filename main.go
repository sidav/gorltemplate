package main

import (
	"gorltemplate/console_wrapper"
)

func main() {
	console_wrapper.Init_console()
	defer console_wrapper.Close_console()

	rnd.InitDefault()
	// dijkstra_maps_test.Simulate()

	gameLoop()
}
