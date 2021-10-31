package main

import cw "gorltemplate/console_wrapper"

const (
	MOVEMENT_WALK byte = iota
	MOVEMENT_AMPHIBIOUS
)

type actorStatic struct {
	maxhp int
	name string
	char  rune
	color int
	movementType byte
}

func getActorStaticByCode(code string) *actorStatic {
	as := allActorStatics[code]
	if as == nil {
		panic("NO ACTOR OF THAT TYPE")
	}
	return as
}

var allActorStatics = map[string]*actorStatic {
	"TESTENEMY1": {
		maxhp: 10,
		name: "Test enemy",
		char: 'e',
		color: cw.RED,
		movementType: MOVEMENT_WALK,
	},
	"TESTENEMY2": {
		maxhp: 10,
		name: "Test enemy",
		char: 'w',
		color: cw.RED,
		movementType: MOVEMENT_WALK,
	},
	"TESTENEMY3": {
		maxhp: 10,
		name: "Test enemy",
		char: 'k',
		color: cw.RED,
		movementType: MOVEMENT_WALK,
	},
}
