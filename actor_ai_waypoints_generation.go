package main

func (ai *actorAi) initAndGenerateWaypoints(actor *actor) {
	const MAXWAYPOINTS = 10
	const MINWAYPOINTS = 3


	ai.waypoints = [][2]int{{actor.x, actor.y}}

	searchHorizontally := rnd.OneChanceFrom(2)
	for i := 1; i < MAXWAYPOINTS; i++ {
		found, x, y := selectCoordsForNewWaypointOrthogonally(ai.waypoints[i-1][0], ai.waypoints[i-1][1], searchHorizontally)
		if found {
			ai.waypoints = append(ai.waypoints, [2]int{x, y})
			searchHorizontally = !searchHorizontally
		} else {
			if i > MINWAYPOINTS {
				return
			}
			ai.waypoints = ai.waypoints[:i-1]
			i -= 2
			if len(ai.waypoints) == 0 || i <= 0 {
				panic("Waypoints can't be created!")
				return
			}
		}
	}
}

func selectCoordsForNewWaypointOrthogonally(prevx, prevy int, searchHorizontally bool) (bool, int, int) {
	const MINWAYPOINTRADIUS = 2
	const MAXWAYPOINTRADIUS = 100
	candidates := make([][3]int, 0) // [x, y, weight]

	weightMultiplier := 1
	if searchHorizontally {
		for x := prevx; x <= prevx+MAXWAYPOINTRADIUS; x++ {
			if !CURRENTLEVEL.isTilePotentiallyPassable(x, prevy, false) {
				break
			}
			if CURRENTLEVEL.tiles[x][prevy].asDoor != nil {
				weightMultiplier = 3
				continue
			}
			if x < prevx+MINWAYPOINTRADIUS {
				continue
			}
			candidates = append(candidates, [3]int{x, prevy, weightMultiplier*(prevx-x)*(prevx-x)})
		}
		weightMultiplier = 1
		for x := prevx; x >= prevx-MAXWAYPOINTRADIUS; x-- {
			if !CURRENTLEVEL.isTilePotentiallyPassable(x, prevy, false) {
				break
			}
			if CURRENTLEVEL.tiles[x][prevy].asDoor != nil {
				weightMultiplier = 3
				continue
			}
			if x > prevx-MINWAYPOINTRADIUS {
				continue
			}
			candidates = append(candidates, [3]int{x, prevy, weightMultiplier*(prevx-x)*(prevx-x)})
		}
	} else {
		for y := prevy; y <= prevy+MAXWAYPOINTRADIUS; y++ {
			if !CURRENTLEVEL.isTilePotentiallyPassable(prevx, y, false) {
				break
			}
			if CURRENTLEVEL.tiles[prevx][y].asDoor != nil {
				weightMultiplier = 3
				continue
			}
			if y < prevy+MINWAYPOINTRADIUS {
				continue
			}
			candidates = append(candidates, [3]int{prevx, y, weightMultiplier*(prevy-y)*(prevy-y)})
		}
		weightMultiplier = 1
		for y := prevy; y >= prevy-MAXWAYPOINTRADIUS; y-- {
			if !CURRENTLEVEL.isTilePotentiallyPassable(prevx, y, false) {
				break
			}
			if CURRENTLEVEL.tiles[prevx][y].asDoor != nil {
				weightMultiplier = 3
				continue
			}
			if y > prevy-MINWAYPOINTRADIUS {
				continue
			}
			candidates = append(candidates, [3]int{prevx, y, weightMultiplier*(prevy-y)*(prevy-y)})
		}
	}
	if len(candidates) > 0 {
		randIndex := rnd.SelectRandomIndexFromWeighted(len(candidates), func(i int) int {return candidates[i][2]})
		return true, candidates[randIndex][0], candidates[randIndex][1]
	}
	return false, 0, 0
}

// path- and distance-based waypoint generation
func areCoordsGoodForWaypointMethod1() {
	const MAXWAYPOINTPATHLENGTH = 50
}
