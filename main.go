package main

import "anotherroguelike/console_wrapper"

func main() {
	console_wrapper.Init_console()
	defer console_wrapper.Close_console()
	initLevel()
	gameLoop()
}
