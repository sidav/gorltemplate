package main

import (
	"github.com/sidav/cyclicdungeongenerator/generator"
	"github.com/sidav/cyclicdungeongenerator/generator/layout_tiler"
	cw "gorltemplate/console_wrapper"
)

func initLevel() {
	cdg := generator.InitGeneratorsWrapper()
	ptnName := cdg.ListPatternFilenamesInPath("assets/")[0]
	generatedMap := cdg.GenerateTiledMapFromPattern(ptnName,5, 5, 6, 5, "assets/parcels", -1)
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
			case layout_tiler.TILE_WATER:
				CURRENTLEVEL.tiles[x][y].data = getTileStatics("WATER")
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
			name: "Player",
			maxhp: 5,
			char:  '@',
			color: cw.WHITE,
			movementType: MOVEMENT_AMPHIBIOUS,
		},
	}
	plr.init()
	CURRENTLEVEL.actors = append(CURRENTLEVEL.actors, plr)



	emptyCoords := make([][]int, 0)
	for x := 0; x < len(CURRENTLEVEL.tiles); x++ {
		for y := 0; y < len(CURRENTLEVEL.tiles[x]); y++ {
			if CURRENTLEVEL.isTilePassableFor(x, y, nil) {
				emptyCoords = append(emptyCoords, []int{x, y})
			}
		}
	}
	if len(emptyCoords) >= 10 {
		ai_types := []int{AI_TYPE_SIMPLE, AI_TYPE_WAYPOINT_BASED}
		ai_enemies := []string{"TESTENEMY1", "TESTENEMY2"}
		indicesToSelect := rnd.ArrayOfRandomsInRange(2, 0, len(emptyCoords)-1, true)
		for i, index := range indicesToSelect {
			ai_ind := i % len(ai_types)
			newEnemy := &actor{
				data:   getActorStaticByCode(ai_enemies[ai_ind]),
				intent: nil,
				team:   1,
				x:      emptyCoords[index][0],
				y:      emptyCoords[index][1],
				hp:     5,
			}
			newEnemy.init()
			newEnemy.initializeAI(ai_types[ai_ind])
			CURRENTLEVEL.actors = append(CURRENTLEVEL.actors, newEnemy)
		}
	}

	PLAYERCONTROLLER.player = plr
}
