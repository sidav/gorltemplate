package main

import "gorltemplate/console_wrapper"

func main() {
	console_wrapper.Init_console()
	defer console_wrapper.Close_console()
	initLevel()
	gameLoop()
}
