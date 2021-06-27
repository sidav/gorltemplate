package main

import cw "gorltemplate/console_wrapper"

type tiletypecode uint8

const (
	TILE_FLOOR tiletypecode = iota
	TILE_WALL  tiletypecode = iota
)

type tile struct {
	data   *tileStatic
	asDoor *tileDoor
}

type tileDoor struct {
	isOpened  bool
	lockLevel int
}

type tileStatic struct {
	passable         bool
	opaque           bool
	char             rune
	bgcolor, fgcolor int
}

var tileStatics = map[tiletypecode]*tileStatic{
	TILE_FLOOR: {
		passable: true,
		opaque:   true,
		char:     '.',
		bgcolor:  0,
		fgcolor:  cw.WHITE,
	},
	TILE_WALL: {
		passable: true,
		opaque:   true,
		char:     ' ',
		bgcolor:  cw.RED,
		fgcolor:  0,
	},
}
