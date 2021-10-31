package main

import cw "gorltemplate/console_wrapper"

type tileStatic struct {
	passabilityForMovementType map[byte]bool

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

var tileStatics = map[string]*tileStatic{
	"FLOOR": {
		passabilityForMovementType: map[byte]bool{MOVEMENT_WALK: true, MOVEMENT_AMPHIBIOUS: true},
		opaque:          false,
		char:            '.',
		bgcolor:         0,
		fgcolor:         cw.WHITE,
	},
	"WALL": {
		passabilityForMovementType: map[byte]bool{MOVEMENT_WALK: false, MOVEMENT_AMPHIBIOUS: false},
		opaque:          true,
		char:            ' ',
		bgcolor:         cw.DARK_RED,
		fgcolor:         0,
	},
	"ENTRYPOINT": {
		passabilityForMovementType: map[byte]bool{MOVEMENT_WALK: true, MOVEMENT_AMPHIBIOUS: true},
		opaque:          false,
		char:            '_',
		bgcolor:         0,
		fgcolor:         cw.WHITE,
	},
	"EXITPOINT": {
		passabilityForMovementType: map[byte]bool{MOVEMENT_WALK: true, MOVEMENT_AMPHIBIOUS: true},
		opaque:          false,
		char:            '>',
		bgcolor:         0,
		fgcolor:         cw.WHITE,
	},
	"WATER": {
		passabilityForMovementType: map[byte]bool{MOVEMENT_WALK: false, MOVEMENT_AMPHIBIOUS: true},
		opaque:          false,
		char:            '~',
		bgcolor:         0,
		fgcolor:         cw.BLUE,
	},

	// don't exlicitly use following codes
	"_DOOR_OPENED": {
		passabilityForMovementType: map[byte]bool{MOVEMENT_WALK: true, MOVEMENT_AMPHIBIOUS: true},
		opaque:          false,
		char:            '\'',
		bgcolor:         0,
		fgcolor:         cw.RED,
	},
	"_DOOR_CLOSED": {
		passabilityForMovementType: map[byte]bool{MOVEMENT_WALK: false, MOVEMENT_AMPHIBIOUS: false},
		opaque:          true,
		char:            '+',
		bgcolor:         0,
		fgcolor:         cw.RED,
	},
	"_DOOR_LOCKED": {
		passabilityForMovementType: map[byte]bool{MOVEMENT_WALK: false, MOVEMENT_AMPHIBIOUS: false},
		opaque:          true,
		char:            '#',
		bgcolor:         0,
		fgcolor:         cw.DARK_BLUE,
	},
	"_SWITCH_OFF": {
		passabilityForMovementType: map[byte]bool{MOVEMENT_WALK: false, MOVEMENT_AMPHIBIOUS: false},
		opaque:          false,
		char:            '*',
		bgcolor:         0,
		fgcolor:         cw.RED,
	},
	"_SWITCH_ON": {
		passabilityForMovementType: map[byte]bool{MOVEMENT_WALK: false, MOVEMENT_AMPHIBIOUS: false},
		opaque:          false,
		char:            '*',
		bgcolor:         0,
		fgcolor:         cw.BLUE,
	},
	"DEFAULT_VALUE": {
		passabilityForMovementType: map[byte]bool{MOVEMENT_WALK: false, MOVEMENT_AMPHIBIOUS: false},
		opaque:          false,
		char:            '?',
		bgcolor:         cw.DARK_MAGENTA,
		fgcolor:         cw.CYAN,
	},
}
