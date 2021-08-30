package main

import cw "gorltemplate/console_wrapper"

type actorStatic struct {
	maxhp int
	name string
	char  rune
	color int
}

var allActorStatics = map[string]*actorStatic {
	"TESTENEMY": {
		maxhp: 10,
		name: "Test enemy",
		char: 'e',
		color: cw.RED,
	},
}
