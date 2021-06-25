package main

import cw "anotherroguelike/console_wrapper"

func initLevel() {
	// temporary
	CURRENTLEVEL.tiles = make([][]tile, 15)
	for i := range CURRENTLEVEL.tiles {
		CURRENTLEVEL.tiles[i] = make([]tile, 10)
	}
	for x := 0; x < len(CURRENTLEVEL.tiles); x++ {
		for y := 0; y < len(CURRENTLEVEL.tiles[x]); y++ {
			CURRENTLEVEL.tiles[x][y].data = tileStatics[TILE_FLOOR]
		}
	}

	plr := &actor{
		x:  0,
		y:  0,
		hp: 5,
		data: &actorStatic{
			maxhp: 5,
			char:  '@',
			color: cw.WHITE,
		},
	}
	CURRENTLEVEL.actors = append(CURRENTLEVEL.actors, plr)
	PLAYERCONTROLLER.player = plr
}
