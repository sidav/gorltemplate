package main

import cw "gorltemplate/console_wrapper"

type tileStatic struct {
	passable         bool
	opaque           bool
	char             rune
	bgcolor, fgcolor int
}

func getTileStatics(code string) *tileStatic {
	ts := tileStatics[code]
	if ts == nil {
		return tileStatics["DEFAULT_VALUE"]
	}
	return ts
}

func (ts *tileStatic) getUnseenColors() (int, int) {
	if ts.fgcolor == 0 {
		return 0, cw.DARK_GRAY
	}
	return cw.DARK_GRAY, 0
}

var tileStatics = map[string] *tileStatic {
	"FLOOR": {
		passable: true,
		opaque:   false,
		char:     '.',
		bgcolor:  0,
		fgcolor:  cw.WHITE,
	},
	"WALL": {
		passable: false,
		opaque:   true,
		char:     ' ',
		bgcolor:  cw.DARK_RED,
		fgcolor:  0,
	},
	"ENTRYPOINT": {
		passable: true,
		opaque:   false,
		char:     '_',
		bgcolor:  0,
		fgcolor:  cw.WHITE,
	},
	"EXITPOINT": {
		passable: true,
		opaque:   false,
		char:     '>',
		bgcolor:  0,
		fgcolor:  cw.WHITE,
	},

	// don't exlicitly use following codes
	"_DOOR_OPENED": {
		passable: true,
		opaque:   false,
		char:     '\'',
		bgcolor:  0,
		fgcolor:  cw.RED,
	},
	"_DOOR_CLOSED": {
		passable: false,
		opaque:   true,
		char:     '+',
		bgcolor:  0,
		fgcolor:  cw.RED,
	},
	"_DOOR_LOCKED": {
		passable: false,
		opaque:   true,
		char:     '#',
		bgcolor:  0,
		fgcolor:  cw.DARK_BLUE,
	},
	"_SWITCH_OFF": {
		passable: false,
		opaque:   false,
		char:     '*',
		bgcolor:  0,
		fgcolor:  cw.RED,
	},
	"_SWITCH_ON": {
		passable: false,
		opaque:   false,
		char:     '*',
		bgcolor:  0,
		fgcolor:  cw.BLUE,
	},
	"DEFAULT_VALUE": {
		passable: false,
		opaque:   false,
		char: '?',
		bgcolor:  cw.DARK_MAGENTA,
		fgcolor:  cw.CYAN,
	},
}
