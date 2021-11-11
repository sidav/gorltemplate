package main

import "gorltemplate/astar"

func (l *level) getVectorToNextCellForPathTo(a *actor, tox, toy int) (int, int) {
	fromx, fromy := a.getCoords()
	pf := astar.AStarPathfinder{
		DiagonalMoveAllowed:       false,
		ForceGetPath:              true,
		ForceIncludeFinish:        true,
		AutoAdjustDefaultMaxSteps: false,
		MapWidth: len(CURRENTLEVEL.tiles),
		MapHeight: len(CURRENTLEVEL.tiles[0]),
	}
	costFunc := func(x, y int) int {
		if l.isTilePassableFor(x, y, a) {
			return 10
		}
		return -1
	}
	path := pf.FindPath(costFunc, fromx, fromy, tox, toy)
	return path.GetNextStepVector()
}
