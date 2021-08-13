package main

import (
	"github.com/sidav/cyclicdungeongenerator/generator"
	"github.com/sidav/cyclicdungeongenerator/generator/layout_tiler"
	cw "gorltemplate/console_wrapper"
)

func initLevel() {
	cdg := generator.InitGeneratorsWrapper()
	ptnName := cdg.ListPatternFilenamesInPath("assets/")[0]
	generatedMap := cdg.GenerateTiledMapFromPattern(ptnName,5, 5, 6, 5, "", -1)
	// temporary
	CURRENTLEVEL.tiles = make([][]tile, len(generatedMap))
	for i := range CURRENTLEVEL.tiles {
		CURRENTLEVEL.tiles[i] = make([]tile, len(generatedMap[i]))
	}

	entryX, entryY := 0, 0
	for x := 0; x < len(CURRENTLEVEL.tiles); x++ {
		for y := 0; y < len(CURRENTLEVEL.tiles[x]); y++ {
			currGeneratedTile := generatedMap[x][y]
			switch currGeneratedTile.Code {
			case layout_tiler.TILE_ENTRYPOINT:
				CURRENTLEVEL.tiles[x][y].data = getTileStatics("ENTRYPOINT")
				entryX, entryY = x, y
			case layout_tiler.TILE_FLOOR:
				CURRENTLEVEL.tiles[x][y].data = getTileStatics("FLOOR")
			case layout_tiler.TILE_WALL:
				CURRENTLEVEL.tiles[x][y].data = getTileStatics("WALL")
			case layout_tiler.TILE_EXITPOINT:
				CURRENTLEVEL.tiles[x][y].data = getTileStatics("EXITPOINT")
			case layout_tiler.TILE_KEY_PLACE:
				CURRENTLEVEL.tiles[x][y].asSwitch = &tileSwitch{
					isActivated: false,
					lockLevel:   currGeneratedTile.LockId,
				}
			case layout_tiler.TILE_DOOR:
				CURRENTLEVEL.tiles[x][y].asDoor = &tileDoor{
					isOpened:  false,
					lockLevel: currGeneratedTile.LockId,
				}
			default: CURRENTLEVEL.tiles[x][y].data = getTileStatics("FLOOR")
			}
		}
	}

	plr := &actor{
		x:  entryX,
		y:  entryY,
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
